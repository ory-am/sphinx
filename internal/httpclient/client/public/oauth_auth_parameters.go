// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewOauthAuthParams creates a new OauthAuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOauthAuthParams() *OauthAuthParams {
	return &OauthAuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOauthAuthParamsWithTimeout creates a new OauthAuthParams object
// with the ability to set a timeout on a request.
func NewOauthAuthParamsWithTimeout(timeout time.Duration) *OauthAuthParams {
	return &OauthAuthParams{
		timeout: timeout,
	}
}

// NewOauthAuthParamsWithContext creates a new OauthAuthParams object
// with the ability to set a context for a request.
func NewOauthAuthParamsWithContext(ctx context.Context) *OauthAuthParams {
	return &OauthAuthParams{
		Context: ctx,
	}
}

// NewOauthAuthParamsWithHTTPClient creates a new OauthAuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewOauthAuthParamsWithHTTPClient(client *http.Client) *OauthAuthParams {
	return &OauthAuthParams{
		HTTPClient: client,
	}
}

/* OauthAuthParams contains all the parameters to send to the API endpoint
   for the oauth auth operation.

   Typically these are written to a http.Request.
*/
type OauthAuthParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the oauth auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OauthAuthParams) WithDefaults() *OauthAuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the oauth auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OauthAuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the oauth auth params
func (o *OauthAuthParams) WithTimeout(timeout time.Duration) *OauthAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oauth auth params
func (o *OauthAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oauth auth params
func (o *OauthAuthParams) WithContext(ctx context.Context) *OauthAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oauth auth params
func (o *OauthAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oauth auth params
func (o *OauthAuthParams) WithHTTPClient(client *http.Client) *OauthAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oauth auth params
func (o *OauthAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OauthAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
