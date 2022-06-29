// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AdminCreateOAuth2Client(params *AdminCreateOAuth2ClientParams, opts ...ClientOption) (*AdminCreateOAuth2ClientCreated, error)

	AdminDeleteOAuth2Client(params *AdminDeleteOAuth2ClientParams, opts ...ClientOption) (*AdminDeleteOAuth2ClientNoContent, error)

	AdminGetOAuth2Client(params *AdminGetOAuth2ClientParams, opts ...ClientOption) (*AdminGetOAuth2ClientOK, error)

	AdminListOAuth2Clients(params *AdminListOAuth2ClientsParams, opts ...ClientOption) (*AdminListOAuth2ClientsOK, error)

	AdminPatchOAuth2Client(params *AdminPatchOAuth2ClientParams, opts ...ClientOption) (*AdminPatchOAuth2ClientOK, error)

	AdminUpdateOAuth2Client(params *AdminUpdateOAuth2ClientParams, opts ...ClientOption) (*AdminUpdateOAuth2ClientOK, error)

	DynamicClientRegistrationCreateOAuth2Client(params *DynamicClientRegistrationCreateOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationCreateOAuth2ClientCreated, error)

	DynamicClientRegistrationDeleteOAuth2Client(params *DynamicClientRegistrationDeleteOAuth2ClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DynamicClientRegistrationDeleteOAuth2ClientNoContent, error)

	DynamicClientRegistrationGetOAuth2Client(params *DynamicClientRegistrationGetOAuth2ClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DynamicClientRegistrationGetOAuth2ClientOK, error)

	DynamicClientRegistrationUpdateOAuth2Client(params *DynamicClientRegistrationUpdateOAuth2ClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DynamicClientRegistrationUpdateOAuth2ClientOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AdminCreateOAuth2Client creates an o auth 2 0 client

  Create a new OAuth 2.0 client. If you pass `client_secret` the secret is used, otherwise a random secret
is generated. The secret is echoed in the response. It is not possible to retrieve it later on.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) AdminCreateOAuth2Client(params *AdminCreateOAuth2ClientParams, opts ...ClientOption) (*AdminCreateOAuth2ClientCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminCreateOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminCreateOAuth2Client",
		Method:             "POST",
		PathPattern:        "/admin/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminCreateOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminCreateOAuth2ClientCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AdminCreateOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AdminDeleteOAuth2Client deletes an o auth 2 0 client

  Delete an existing OAuth 2.0 Client by its ID.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.

Make sure that this endpoint is well protected and only callable by first-party components.
*/
func (a *Client) AdminDeleteOAuth2Client(params *AdminDeleteOAuth2ClientParams, opts ...ClientOption) (*AdminDeleteOAuth2ClientNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminDeleteOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminDeleteOAuth2Client",
		Method:             "DELETE",
		PathPattern:        "/clients/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminDeleteOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminDeleteOAuth2ClientNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AdminDeleteOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AdminGetOAuth2Client gets an o auth 2 0 client

  Get an OAuth 2.0 client by its ID. This endpoint never returns the client secret.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) AdminGetOAuth2Client(params *AdminGetOAuth2ClientParams, opts ...ClientOption) (*AdminGetOAuth2ClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminGetOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminGetOAuth2Client",
		Method:             "GET",
		PathPattern:        "/clients/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminGetOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminGetOAuth2ClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AdminGetOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AdminListOAuth2Clients lists o auth 2 0 clients

  This endpoint lists all clients in the database, and never returns client secrets.
As a default it lists the first 100 clients. The `limit` parameter can be used to retrieve more clients,
but it has an upper bound at 500 objects. Pagination should be used to retrieve more than 500 objects.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.

The "Link" header is also included in successful responses, which contains one or more links for pagination, formatted like so: '<https://hydra-url/admin/clients?limit={limit}&offset={offset}>; rel="{page}"', where page is one of the following applicable pages: 'first', 'next', 'last', and 'previous'.
Multiple links can be included in this header, and will be separated by a comma.
*/
func (a *Client) AdminListOAuth2Clients(params *AdminListOAuth2ClientsParams, opts ...ClientOption) (*AdminListOAuth2ClientsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminListOAuth2ClientsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminListOAuth2Clients",
		Method:             "GET",
		PathPattern:        "/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminListOAuth2ClientsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminListOAuth2ClientsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AdminListOAuth2ClientsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AdminPatchOAuth2Client patches an o auth 2 0 client

  Patch an existing OAuth 2.0 Client. If you pass `client_secret`
the secret will be updated and returned via the API. This is the
only time you will be able to retrieve the client secret, so write it down and keep it safe.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) AdminPatchOAuth2Client(params *AdminPatchOAuth2ClientParams, opts ...ClientOption) (*AdminPatchOAuth2ClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminPatchOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminPatchOAuth2Client",
		Method:             "PATCH",
		PathPattern:        "/clients/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminPatchOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminPatchOAuth2ClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AdminPatchOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AdminUpdateOAuth2Client updates an o auth 2 0 client

  Update an existing OAuth 2.0 Client. If you pass `client_secret` the secret is used, otherwise a random secret
is generated. The secret is echoed in the response. It is not possible to retrieve it later on.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) AdminUpdateOAuth2Client(params *AdminUpdateOAuth2ClientParams, opts ...ClientOption) (*AdminUpdateOAuth2ClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminUpdateOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminUpdateOAuth2Client",
		Method:             "PUT",
		PathPattern:        "/admin/clients/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminUpdateOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminUpdateOAuth2ClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AdminUpdateOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DynamicClientRegistrationCreateOAuth2Client registers an o auth 2 0 client using the open ID o auth2 dynamic client registration management protocol

  This endpoint behaves like the administrative counterpart (`createOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

Please note that using this endpoint you are not able to choose the `client_secret` nor the `client_id` as those
values will be server generated when specifying `token_endpoint_auth_method` as `client_secret_basic` or
`client_secret_post`.

The `client_secret` will be returned in the response and you will not be able to retrieve it later on.
Write the secret down and keep it somewhere safe.
*/
func (a *Client) DynamicClientRegistrationCreateOAuth2Client(params *DynamicClientRegistrationCreateOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationCreateOAuth2ClientCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationCreateOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationCreateOAuth2Client",
		Method:             "POST",
		PathPattern:        "/oauth2/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationCreateOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationCreateOAuth2ClientCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DynamicClientRegistrationCreateOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DynamicClientRegistrationDeleteOAuth2Client deletes an o auth 2 0 client using the open ID o auth2 dynamic client registration management protocol

  This endpoint behaves like the administrative counterpart (`deleteOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) DynamicClientRegistrationDeleteOAuth2Client(params *DynamicClientRegistrationDeleteOAuth2ClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DynamicClientRegistrationDeleteOAuth2ClientNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationDeleteOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationDeleteOAuth2Client",
		Method:             "DELETE",
		PathPattern:        "/oauth2/register/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationDeleteOAuth2ClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationDeleteOAuth2ClientNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DynamicClientRegistrationDeleteOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DynamicClientRegistrationGetOAuth2Client gets an o auth 2 0 client using the open ID o auth2 dynamic client registration management protocol

  This endpoint behaves like the administrative counterpart (`getOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) DynamicClientRegistrationGetOAuth2Client(params *DynamicClientRegistrationGetOAuth2ClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DynamicClientRegistrationGetOAuth2ClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationGetOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationGetOAuth2Client",
		Method:             "GET",
		PathPattern:        "/oauth2/register/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationGetOAuth2ClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationGetOAuth2ClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DynamicClientRegistrationGetOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DynamicClientRegistrationUpdateOAuth2Client updates an o auth 2 0 client using the open ID o auth2 dynamic client registration management protocol

  This endpoint behaves like the administrative counterpart (`updateOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

If you pass `client_secret` the secret is used, otherwise a random secret
is generated. The secret is echoed in the response. It is not possible to retrieve it later on.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) DynamicClientRegistrationUpdateOAuth2Client(params *DynamicClientRegistrationUpdateOAuth2ClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DynamicClientRegistrationUpdateOAuth2ClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationUpdateOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationUpdateOAuth2Client",
		Method:             "PUT",
		PathPattern:        "/oauth2/register/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationUpdateOAuth2ClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationUpdateOAuth2ClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DynamicClientRegistrationUpdateOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
