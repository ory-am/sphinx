package jwk

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/go-errors/errors"
	"github.com/square/go-jose"
	"crypto/x509"
)

type ECDSA256Generator struct{}

func (g *ECDSA256Generator) Generate(id string) (*jose.JsonWebKeySet, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, errors.Errorf("Could not generate key because %s", err)
	}

	return &jose.JsonWebKeySet{
		Keys: []jose.JsonWebKey{
			{
				Key:   key,
				KeyID: ider("private", id),
				Certificates: []*x509.Certificate{},
			},
			{
				Key:   &key.PublicKey,
				KeyID: ider("public", id),
				Certificates: []*x509.Certificate{},
			},
		},
	}, nil
}
