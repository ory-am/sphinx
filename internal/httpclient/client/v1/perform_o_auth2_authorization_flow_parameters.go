// Code generated by go-swagger; DO NOT EDIT.

package v1

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

	"github.com/ory/hydra/internal/httpclient/client/v1"
)

// NewPerformOAuth2AuthorizationFlowParams creates a new PerformOAuth2AuthorizationFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformOAuth2AuthorizationFlowParams() *PerformOAuth2AuthorizationFlowParams {
	return &PerformOAuth2AuthorizationFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformOAuth2AuthorizationFlowParamsWithTimeout creates a new PerformOAuth2AuthorizationFlowParams object
// with the ability to set a timeout on a request.
func NewPerformOAuth2AuthorizationFlowParamsWithTimeout(timeout time.Duration) *PerformOAuth2AuthorizationFlowParams {
	return &PerformOAuth2AuthorizationFlowParams{
		timeout: timeout,
	}
}

// NewPerformOAuth2AuthorizationFlowParamsWithContext creates a new PerformOAuth2AuthorizationFlowParams object
// with the ability to set a context for a request.
func NewPerformOAuth2AuthorizationFlowParamsWithContext(ctx context.Context) *PerformOAuth2AuthorizationFlowParams {
	return &PerformOAuth2AuthorizationFlowParams{
		Context: ctx,
	}
}

// NewPerformOAuth2AuthorizationFlowParamsWithHTTPClient creates a new PerformOAuth2AuthorizationFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformOAuth2AuthorizationFlowParamsWithHTTPClient(client *http.Client) *PerformOAuth2AuthorizationFlowParams {
	return &PerformOAuth2AuthorizationFlowParams{
		HTTPClient: client,
	}
}

/* PerformOAuth2AuthorizationFlowParams contains all the parameters to send to the API endpoint
   for the perform o auth2 authorization flow operation.

   Typically these are written to a http.Request.
*/
type PerformOAuth2AuthorizationFlowParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform o auth2 authorization flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformOAuth2AuthorizationFlowParams) WithDefaults() *PerformOAuth2AuthorizationFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform o auth2 authorization flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformOAuth2AuthorizationFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform o auth2 authorization flow params
func (o *PerformOAuth2AuthorizationFlowParams) WithTimeout(timeout time.Duration) *PerformOAuth2AuthorizationFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform o auth2 authorization flow params
func (o *PerformOAuth2AuthorizationFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform o auth2 authorization flow params
func (o *PerformOAuth2AuthorizationFlowParams) WithContext(ctx context.Context) *PerformOAuth2AuthorizationFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform o auth2 authorization flow params
func (o *PerformOAuth2AuthorizationFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform o auth2 authorization flow params
func (o *PerformOAuth2AuthorizationFlowParams) WithHTTPClient(client *http.Client) *PerformOAuth2AuthorizationFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform o auth2 authorization flow params
func (o *PerformOAuth2AuthorizationFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PerformOAuth2AuthorizationFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
