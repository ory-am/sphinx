/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package oauth2_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/ory/x/snapshotx"

	"github.com/ory/hydra/internal/testhelpers"
	"github.com/ory/x/contextx"

	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/client/admin"
	"github.com/ory/hydra/internal/httpclient/models"

	"github.com/ory/hydra/jwk"
	"github.com/ory/hydra/x"

	"github.com/ory/hydra/driver/config"
	"github.com/ory/hydra/internal"
	"github.com/ory/x/urlx"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	jwt2 "github.com/ory/fosite/token/jwt"

	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"
	"github.com/ory/hydra/client"
	hydra "github.com/ory/hydra/internal/httpclient/client"
	"github.com/ory/hydra/oauth2"
)

var lifespan = time.Hour

func TestHandlerDeleteHandler(t *testing.T) {
	ctx := context.Background()
	conf := internal.NewConfigurationWithDefaults()
	conf.MustSet(ctx, config.KeyIssuerURL, "http://hydra.localhost")
	reg := internal.NewRegistryMemory(t, conf, &contextx.Default{})

	cm := reg.ClientManager()
	store := reg.OAuth2Storage()

	h := oauth2.NewHandler(reg, conf)

	deleteRequest := &fosite.Request{
		ID:             "del-1",
		RequestedAt:    time.Now().Round(time.Second),
		Client:         &client.Client{LegacyClientID: "foobar"},
		RequestedScope: fosite.Arguments{"fa", "ba"},
		GrantedScope:   fosite.Arguments{"fa", "ba"},
		Form:           url.Values{"foo": []string{"bar", "baz"}},
		Session:        &oauth2.Session{DefaultSession: &openid.DefaultSession{Subject: "bar"}},
	}
	require.NoError(t, cm.CreateClient(context.Background(), deleteRequest.Client.(*client.Client)))
	require.NoError(t, store.CreateAccessTokenSession(context.Background(), deleteRequest.ID, deleteRequest))

	r := x.NewRouterAdmin()
	h.SetRoutes(r, r.RouterPublic(), func(h http.Handler) http.Handler {
		return h
	})
	ts := httptest.NewServer(r)
	defer ts.Close()

	c := hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{Schemes: []string{"http"}, Host: urlx.ParseOrPanic(ts.URL).Host})
	_, err := c.Admin.DeleteOAuth2Token(admin.NewDeleteOAuth2TokenParams().WithClientID("foobar"))
	require.NoError(t, err)

	ds := new(oauth2.Session)
	_, err = store.GetAccessTokenSession(ctx, "del-1", ds)
	require.Error(t, err, "not_found")
}

func TestHandlerFlushHandler(t *testing.T) {

	t.Run("case=not-after", func(t *testing.T) {
		ctx := context.Background()

		// Setup database and tests
		jt := testhelpers.NewConsentJanitorTestHelper(t.Name())
		reg, err := jt.GetRegistry(ctx, t.Name())
		require.NoError(t, err)

		// Setup rest server
		h := oauth2.NewHandler(reg, jt.GetConfig())

		r := x.NewRouterAdmin()
		h.SetRoutes(r, r.RouterPublic(), func(h http.Handler) http.Handler {
			return h
		})

		ts := httptest.NewServer(r)
		defer ts.Close()
		c := hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{Schemes: []string{"http"}, Host: urlx.ParseOrPanic(ts.URL).Host})

		// Get the notAfter test cycles
		notAfterTests := testhelpers.NewConsentJanitorTestHelper(t.Name()).GetNotAfterTestCycles()

		for k, v := range notAfterTests {
			t.Run(fmt.Sprintf("case=%s", k), func(t *testing.T) {
				ctx := context.Background()
				jt := testhelpers.NewConsentJanitorTestHelper(t.Name())

				notAfter := time.Now().Round(time.Second).Add(-v)

				t.Run("step=setup-access-token", jt.AccessTokenNotAfterSetup(ctx, reg.ClientManager(), reg.OAuth2Storage()))
				t.Run("step=setup-refresh-token", jt.RefreshTokenNotAfterSetup(ctx, reg.ClientManager(), reg.OAuth2Storage()))

				t.Run("step=cleanup", func(t *testing.T) {
					_, err := c.Admin.FlushInactiveOAuth2Tokens(admin.NewFlushInactiveOAuth2TokensParams().WithBody(&models.FlushInactiveOAuth2TokensRequest{NotAfter: strfmt.DateTime(notAfter)}))
					require.NoError(t, err)
				})

				t.Run("step=validate-access-token", jt.AccessTokenNotAfterValidate(ctx, notAfter, reg.OAuth2Storage()))
				t.Run("step=validate-refresh-token", jt.RefreshTokenNotAfterValidate(ctx, notAfter, reg.OAuth2Storage()))
			})
		}

	})
}

func TestUserinfo(t *testing.T) {
	ctx := context.Background()
	conf := internal.NewConfigurationWithDefaults()
	conf.MustSet(ctx, config.KeyScopeStrategy, "")
	conf.MustSet(ctx, config.KeyAuthCodeLifespan, lifespan)
	conf.MustSet(ctx, config.KeyIssuerURL, "http://hydra.localhost")
	reg := internal.NewRegistryMemory(t, conf, &contextx.Default{})
	internal.MustEnsureRegistryKeys(reg, x.OpenIDConnectKeyName)

	ctrl := gomock.NewController(t)
	op := NewMockOAuth2Provider(ctrl)
	defer ctrl.Finish()
	reg.WithOAuth2Provider(op)

	h := reg.OAuth2Handler()

	router := x.NewRouterAdmin()
	h.SetRoutes(router, router.RouterPublic(), func(h http.Handler) http.Handler {
		return h
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	for k, tc := range []struct {
		setup                func(t *testing.T)
		checkForSuccess      func(t *testing.T, body []byte)
		checkForUnauthorized func(t *testing.T, body []byte, header http.Header)
		expectStatusCode     int
	}{
		{
			setup: func(t *testing.T) {
				op.EXPECT().IntrospectToken(gomock.Any(), gomock.Eq("access-token"), gomock.Eq(fosite.AccessToken), gomock.Any()).Return(fosite.AccessToken, nil, errors.New("asdf"))
			},
			expectStatusCode: http.StatusInternalServerError,
		},
		{
			setup: func(t *testing.T) {
				op.EXPECT().
					IntrospectToken(gomock.Any(), gomock.Eq("access-token"), gomock.Eq(fosite.AccessToken), gomock.Any()).
					Return(fosite.RefreshToken, nil, nil)
			},
			checkForUnauthorized: func(t *testing.T, body []byte, headers http.Header) {
				assert.True(t, headers.Get("WWW-Authenticate") == `Bearer error="invalid_token",error_description="Only access tokens are allowed in the authorization header."`, "%s", headers)
			},
			expectStatusCode: http.StatusUnauthorized,
		},
		{
			setup: func(t *testing.T) {
				op.EXPECT().
					IntrospectToken(gomock.Any(), gomock.Eq("access-token"), gomock.Eq(fosite.AccessToken), gomock.Any()).
					Return(fosite.AccessToken, nil, fosite.ErrRequestUnauthorized)
			},
			checkForUnauthorized: func(t *testing.T, body []byte, headers http.Header) {
				assert.True(t, headers.Get("WWW-Authenticate") == `Bearer error="request_unauthorized",error_description="The request could not be authorized. Check that you provided valid credentials in the right format."`, "%s", headers)
			},
			expectStatusCode: http.StatusUnauthorized,
		},
		{
			setup: func(t *testing.T) {
				op.EXPECT().
					IntrospectToken(gomock.Any(), gomock.Eq("access-token"), gomock.Eq(fosite.AccessToken), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, _ fosite.TokenType, session fosite.Session, _ ...string) (fosite.TokenType, fosite.AccessRequester, error) {
						session = &oauth2.Session{
							DefaultSession: &openid.DefaultSession{
								Claims: &jwt.IDTokenClaims{
									Subject: "alice",
								},
								Headers: new(jwt.Headers),
								Subject: "alice",
							},
							Extra: map[string]interface{}{},
						}

						return fosite.AccessToken, &fosite.AccessRequest{
							Request: fosite.Request{
								Client: &client.Client{
									LegacyClientID: "foobar",
								},
								Session: session,
							},
						}, nil
					})
			},
			expectStatusCode: http.StatusOK,
			checkForSuccess: func(t *testing.T, body []byte) {
				bodyString := string(body)
				assert.True(t, strings.Contains(bodyString, `"sub":"alice"`), "%s", body)
				assert.True(t, strings.Contains(bodyString, `"aud":["foobar"]`), "%s", body)
			},
		},
		{
			setup: func(t *testing.T) {
				op.EXPECT().
					IntrospectToken(gomock.Any(), gomock.Eq("access-token"), gomock.Eq(fosite.AccessToken), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, _ fosite.TokenType, session fosite.Session, _ ...string) (fosite.TokenType, fosite.AccessRequester, error) {
						session = &oauth2.Session{
							DefaultSession: &openid.DefaultSession{
								Claims: &jwt.IDTokenClaims{
									Subject:  "another-alice",
									Audience: []string{"something-else"},
								},
								Headers: new(jwt.Headers),
								Subject: "alice",
							},
							Extra: map[string]interface{}{},
						}

						return fosite.AccessToken, &fosite.AccessRequest{
							Request: fosite.Request{
								Client: &client.Client{
									LegacyClientID: "foobar",
								},
								Session: session,
							},
						}, nil
					})
			},
			expectStatusCode: http.StatusOK,
			checkForSuccess: func(t *testing.T, body []byte) {
				bodyString := string(body)
				assert.False(t, strings.Contains(bodyString, `"sub":"alice"`), "%s", body)
				assert.True(t, strings.Contains(bodyString, `"sub":"another-alice"`), "%s", body)
				assert.True(t, strings.Contains(bodyString, `"aud":["something-else","foobar"]`), "%s", body)
			},
		},
		{
			setup: func(t *testing.T) {
				op.EXPECT().
					IntrospectToken(gomock.Any(), gomock.Eq("access-token"), gomock.Eq(fosite.AccessToken), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, _ fosite.TokenType, session fosite.Session, _ ...string) (fosite.TokenType, fosite.AccessRequester, error) {
						session = &oauth2.Session{
							DefaultSession: &openid.DefaultSession{
								Claims: &jwt.IDTokenClaims{
									Subject:  "alice",
									Audience: []string{"foobar"},
								},
								Headers: new(jwt.Headers),
								Subject: "alice",
							},
							Extra: map[string]interface{}{},
						}

						return fosite.AccessToken, &fosite.AccessRequest{
							Request: fosite.Request{
								Client: &client.Client{
									LegacyClientID:            "foobar",
									UserinfoSignedResponseAlg: "none",
								},
								Session: session,
							},
						}, nil
					})
			},
			expectStatusCode: http.StatusOK,
			checkForSuccess: func(t *testing.T, body []byte) {
				bodyString := string(body)
				assert.True(t, strings.Contains(bodyString, `"sub":"alice"`), "%s", body)
				assert.True(t, strings.Contains(bodyString, `"aud":["foobar"]`), "%s", body)
			},
		},
		{
			setup: func(t *testing.T) {
				op.EXPECT().
					IntrospectToken(gomock.Any(), gomock.Eq("access-token"), gomock.Eq(fosite.AccessToken), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, _ fosite.TokenType, session fosite.Session, _ ...string) (fosite.TokenType, fosite.AccessRequester, error) {
						session = &oauth2.Session{
							DefaultSession: &openid.DefaultSession{
								Claims: &jwt.IDTokenClaims{
									Subject: "alice",
								},
								Headers: new(jwt.Headers),
								Subject: "alice",
							},
							Extra: map[string]interface{}{},
						}

						return fosite.AccessToken, &fosite.AccessRequest{
							Request: fosite.Request{
								Client: &client.Client{
									UserinfoSignedResponseAlg: "asdfasdf",
								},
								Session: session,
							},
						}, nil
					})
			},
			expectStatusCode: http.StatusInternalServerError,
		},
		{
			setup: func(t *testing.T) {
				op.EXPECT().
					IntrospectToken(gomock.Any(), gomock.Eq("access-token"), gomock.Eq(fosite.AccessToken), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, _ fosite.TokenType, session fosite.Session, _ ...string) (fosite.TokenType, fosite.AccessRequester, error) {
						session = &oauth2.Session{
							DefaultSession: &openid.DefaultSession{
								Claims: &jwt.IDTokenClaims{
									Subject: "alice",
								},
								Headers: new(jwt.Headers),
								Subject: "alice",
							},
							Extra: map[string]interface{}{},
						}

						return fosite.AccessToken, &fosite.AccessRequest{
							Request: fosite.Request{
								Client: &client.Client{
									LegacyClientID:            "foobar-client",
									UserinfoSignedResponseAlg: "RS256",
								},
								Session: session,
							},
						}, nil
					})
			},
			expectStatusCode: http.StatusOK,
			checkForSuccess: func(t *testing.T, body []byte) {
				claims, err := jwt2.Parse(string(body), func(token *jwt2.Token) (interface{}, error) {
					keys, err := reg.KeyManager().GetKeySet(context.Background(), x.OpenIDConnectKeyName)
					require.NoError(t, err)
					t.Logf("%+v", keys)
					key, err := jwk.FindPublicKey(keys)
					return key.Key, nil
				})
				require.NoError(t, err)
				assert.EqualValues(t, "alice", claims.Claims["sub"])
				assert.EqualValues(t, []interface{}{"foobar-client"}, claims.Claims["aud"], "%#v", claims.Claims)
				assert.NotEmpty(t, claims.Claims["jti"])
			},
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			tc.setup(t)

			req, err := http.NewRequest("GET", ts.URL+"/userinfo", nil)
			require.NoError(t, err)
			req.Header.Set("Authorization", "Bearer access-token")
			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()
			require.EqualValues(t, tc.expectStatusCode, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			if tc.expectStatusCode == http.StatusOK {
				tc.checkForSuccess(t, body)
			} else if tc.expectStatusCode == http.StatusUnauthorized {
				tc.checkForUnauthorized(t, body, resp.Header)
			}
		})
	}
}

func TestHandlerWellKnown(t *testing.T) {
	ctx := context.Background()
	conf := internal.NewConfigurationWithDefaults()
	t.Run(fmt.Sprintf("hsm_enabled=%v", conf.HSMEnabled()), func(t *testing.T) {
		conf.MustSet(ctx, config.KeyScopeStrategy, "DEPRECATED_HIERARCHICAL_SCOPE_STRATEGY")
		conf.MustSet(ctx, config.KeyIssuerURL, "http://hydra.localhost")
		conf.MustSet(ctx, config.KeySubjectTypesSupported, []string{"pairwise", "public"})
		conf.MustSet(ctx, config.KeyOIDCDiscoverySupportedClaims, []string{"sub"})
		conf.MustSet(ctx, config.KeyOAuth2ClientRegistrationURL, "http://client-register/registration")
		conf.MustSet(ctx, config.KeyOIDCDiscoveryUserinfoEndpoint, "/userinfo")
		reg := internal.NewRegistryMemory(t, conf, &contextx.Default{})

		h := oauth2.NewHandler(reg, conf)

		r := x.NewRouterAdmin()
		h.SetRoutes(r, r.RouterPublic(), func(h http.Handler) http.Handler {
			return h
		})
		ts := httptest.NewServer(r)
		defer ts.Close()

		res, err := http.Get(ts.URL + "/.well-known/openid-configuration")
		require.NoError(t, err)
		defer res.Body.Close()

		var wellKnownResp oauth2.WellKnown
		err = json.NewDecoder(res.Body).Decode(&wellKnownResp)
		require.NoError(t, err, "problem decoding wellknown json response: %+v", err)
		snapshotx.SnapshotT(t, wellKnownResp)
	})
}
