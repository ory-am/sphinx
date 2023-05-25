// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/dgraph-io/ristretto"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/randx"
	"github.com/gofrs/uuid"

	"github.com/ory/hydra/v2/client"
	"github.com/ory/hydra/v2/flow"

	"github.com/pkg/errors"

	"github.com/ory/fosite"
	"github.com/ory/fosite/storage"
	"github.com/ory/hydra/v2/driver/config"
	"github.com/ory/hydra/v2/jwk"
	"github.com/ory/hydra/v2/persistence"
	"github.com/ory/hydra/v2/x"
	"github.com/ory/x/contextx"
	"github.com/ory/x/errorsx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/networkx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/popx"
)

var (
	_ persistence.Persister = new(Persister)
	_ storage.Transactional = new(Persister)

	ptrCost   = int64(unsafe.Sizeof(new(int)))
	clientTTL = 5 * time.Minute
)

var (
	ErrTransactionOpen   = errors.New("There is already a transaction in this context.")
	ErrNoTransactionOpen = errors.New("There is no transaction in this context.")
)

type (
	getFlowArgs struct {
		NID            uuid.UUID
		LoginChallenge string
	}
	Persister struct {
		conn        *pop.Connection
		mb          *popx.MigrationBox
		mbs         popx.MigrationStatuses
		r           Dependencies
		config      *config.DefaultProvider
		l           *logrusx.Logger
		fallbackNID uuid.UUID
		p           *networkx.Manager
		queries     struct {
			getFlow string
		}

		cache *ristretto.Cache
	}
	Dependencies interface {
		ClientHasher() fosite.Hasher
		KeyCipher() *jwk.AEAD
		contextx.Provider
		x.RegistryLogger
		x.TracingProvider
	}
)

func (p *Persister) BeginTX(ctx context.Context) (_ context.Context, err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.BeginTX")
	defer otelx.End(span, &err)

	fallback := &pop.Connection{TX: &pop.Tx{}}
	if popx.GetConnection(ctx, fallback).TX != fallback.TX {
		return ctx, errorsx.WithStack(ErrTransactionOpen)
	}

	tx, err := p.conn.Store.TransactionContextOptions(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	c := &pop.Connection{
		TX:      tx,
		Store:   tx,
		ID:      randx.String(30),
		Dialect: p.conn.Dialect,
	}
	return popx.WithTransaction(ctx, c), err
}

func (p *Persister) Commit(ctx context.Context) (err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.Commit")
	defer otelx.End(span, &err)

	fallback := &pop.Connection{TX: &pop.Tx{}}
	tx := popx.GetConnection(ctx, fallback)
	if tx.TX == fallback.TX || tx.TX == nil {
		return errorsx.WithStack(ErrNoTransactionOpen)
	}

	return errorsx.WithStack(tx.TX.Commit())
}

func (p *Persister) Rollback(ctx context.Context) (err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.Rollback")
	defer otelx.End(span, &err)

	fallback := &pop.Connection{TX: &pop.Tx{}}
	tx := popx.GetConnection(ctx, fallback)
	if tx.TX == fallback.TX || tx.TX == nil {
		return errorsx.WithStack(ErrNoTransactionOpen)
	}

	return errorsx.WithStack(tx.TX.Rollback())
}

func NewPersister(ctx context.Context, c *pop.Connection, r Dependencies, config *config.DefaultProvider, l *logrusx.Logger) (*Persister, error) {
	mb, err := popx.NewMigrationBox(migrations, popx.NewMigrator(c, r.Logger(), r.Tracer(ctx), 0))
	if err != nil {
		return nil, errorsx.WithStack(err)
	}

	maxItems := int64(1000)
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: maxItems * 10,
		MaxCost:     maxItems * ptrCost, // approx 1000 items with 8 bytes each
		BufferItems: 64,
		Metrics:     true,
	})
	if err != nil {
		return nil, errorsx.WithStack(err)
	}

	p := &Persister{
		conn:    c,
		mb:      mb,
		r:       r,
		config:  config,
		l:       l,
		p:       networkx.NewManager(c, r.Logger(), r.Tracer(ctx)),
		cache:   cache,
		queries: struct{ getFlow string }{getFlow: ""},
	}

	p.queries.getFlow = fmt.Sprintf(`
SELECT %s, %s
FROM hydra_oauth2_flow AS flow
    JOIN hydra_client AS client ON
        client.id = flow.client_id and
        client.nid = flow.nid
WHERE flow.login_challenge = :loginchallenge AND
      flow.nid = :nid
LIMIT 1
`,
		p.selectColumns(ctx, new(flow.Flow), func(s string) string { return fmt.Sprintf(`flow.%s AS "flow.%s"`, s, s) }),
		p.selectColumns(ctx, new(client.Client), func(s string) string { return fmt.Sprintf(`client.%s AS "client.%s"`, s, s) }),
	)

	return p, nil
}

func (p *Persister) DetermineNetwork(ctx context.Context) (*networkx.Network, error) {
	return p.p.Determine(ctx)
}

func (p Persister) WithFallbackNetworkID(nid uuid.UUID) persistence.Persister {
	p.fallbackNID = nid
	return &p
}

func (p *Persister) CreateWithNetwork(ctx context.Context, v interface{}) error {
	n := p.NetworkID(ctx)
	return p.Connection(ctx).Create(p.mustSetNetwork(n, v))
}

func (p *Persister) UpdateWithNetwork(ctx context.Context, v interface{}) (int64, error) {
	n := p.NetworkID(ctx)
	v = p.mustSetNetwork(n, v)

	m := pop.NewModel(v, ctx)
	var cs []string
	for _, t := range m.Columns().Cols {
		cs = append(cs, t.Name)
	}

	return p.Connection(ctx).Where(m.IDField()+" = ? AND nid = ?", m.ID(), n).UpdateQuery(v, cs...)
}

func (p *Persister) NetworkID(ctx context.Context) uuid.UUID {
	return p.r.Contextualizer().Network(ctx, p.fallbackNID)
}

func (p *Persister) QueryWithNetwork(ctx context.Context) *pop.Query {
	return p.Connection(ctx).Where("nid = ?", p.NetworkID(ctx))
}

func (p *Persister) Connection(ctx context.Context) *pop.Connection {
	return popx.GetConnection(ctx, p.conn)
}

func (p *Persister) Ping() error {
	type pinger interface{ Ping() error }
	return p.conn.Store.(pinger).Ping()
}

func (p *Persister) mustSetNetwork(nid uuid.UUID, v interface{}) interface{} {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || (rv.Kind() == reflect.Ptr && rv.Elem().Kind() != reflect.Struct) {
		panic("v must be a pointer to a struct")
	}
	nf := rv.Elem().FieldByName("NID")
	if !nf.IsValid() || !nf.CanSet() {
		panic("v must have settable a field 'NID uuid.UUID'")
	}
	nf.Set(reflect.ValueOf(nid))
	return v
}

func (p *Persister) transaction(ctx context.Context, f func(ctx context.Context, c *pop.Connection) error) error {
	return popx.Transaction(ctx, p.conn, f)
}

// cacheKey returns a network specific cache key for the given method and arguments.
func (p *Persister) cacheKey(ctx context.Context, method string, args ...string) string {
	parts := []string{p.NetworkID(ctx).String(), method}
	parts = append(parts, args...)

	return strings.Join(parts, "/")
}

func (p *Persister) selectColumns(ctx context.Context, model any, transform func(string) string) string {
	m := pop.NewModel(model, ctx)
	columns := make([]string, 0)
	for _, col := range m.Columns().Cols {
		columns = append(columns, transform(col.Name))
	}

	return strings.Join(columns, ", ")
}
