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
	"github.com/ory/hydra/internal/httpclient/models"
)

// NewAdminTrustOAuth2JwtGrantIssuerParams creates a new AdminTrustOAuth2JwtGrantIssuerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminTrustOAuth2JwtGrantIssuerParams() *AdminTrustOAuth2JwtGrantIssuerParams {
	return &AdminTrustOAuth2JwtGrantIssuerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminTrustOAuth2JwtGrantIssuerParamsWithTimeout creates a new AdminTrustOAuth2JwtGrantIssuerParams object
// with the ability to set a timeout on a request.
func NewAdminTrustOAuth2JwtGrantIssuerParamsWithTimeout(timeout time.Duration) *AdminTrustOAuth2JwtGrantIssuerParams {
	return &AdminTrustOAuth2JwtGrantIssuerParams{
		timeout: timeout,
	}
}

// NewAdminTrustOAuth2JwtGrantIssuerParamsWithContext creates a new AdminTrustOAuth2JwtGrantIssuerParams object
// with the ability to set a context for a request.
func NewAdminTrustOAuth2JwtGrantIssuerParamsWithContext(ctx context.Context) *AdminTrustOAuth2JwtGrantIssuerParams {
	return &AdminTrustOAuth2JwtGrantIssuerParams{
		Context: ctx,
	}
}

// NewAdminTrustOAuth2JwtGrantIssuerParamsWithHTTPClient creates a new AdminTrustOAuth2JwtGrantIssuerParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminTrustOAuth2JwtGrantIssuerParamsWithHTTPClient(client *http.Client) *AdminTrustOAuth2JwtGrantIssuerParams {
	return &AdminTrustOAuth2JwtGrantIssuerParams{
		HTTPClient: client,
	}
}

/* AdminTrustOAuth2JwtGrantIssuerParams contains all the parameters to send to the API endpoint
   for the admin trust o auth2 jwt grant issuer operation.

   Typically these are written to a http.Request.
*/
type AdminTrustOAuth2JwtGrantIssuerParams struct {

	// Body.
	Body *models.AdminTrustOAuth2JwtGrantIssuerBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin trust o auth2 jwt grant issuer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminTrustOAuth2JwtGrantIssuerParams) WithDefaults() *AdminTrustOAuth2JwtGrantIssuerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin trust o auth2 jwt grant issuer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminTrustOAuth2JwtGrantIssuerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin trust o auth2 jwt grant issuer params
func (o *AdminTrustOAuth2JwtGrantIssuerParams) WithTimeout(timeout time.Duration) *AdminTrustOAuth2JwtGrantIssuerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin trust o auth2 jwt grant issuer params
func (o *AdminTrustOAuth2JwtGrantIssuerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin trust o auth2 jwt grant issuer params
func (o *AdminTrustOAuth2JwtGrantIssuerParams) WithContext(ctx context.Context) *AdminTrustOAuth2JwtGrantIssuerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin trust o auth2 jwt grant issuer params
func (o *AdminTrustOAuth2JwtGrantIssuerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin trust o auth2 jwt grant issuer params
func (o *AdminTrustOAuth2JwtGrantIssuerParams) WithHTTPClient(client *http.Client) *AdminTrustOAuth2JwtGrantIssuerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin trust o auth2 jwt grant issuer params
func (o *AdminTrustOAuth2JwtGrantIssuerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin trust o auth2 jwt grant issuer params
func (o *AdminTrustOAuth2JwtGrantIssuerParams) WithBody(body *models.AdminTrustOAuth2JwtGrantIssuerBody) *AdminTrustOAuth2JwtGrantIssuerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin trust o auth2 jwt grant issuer params
func (o *AdminTrustOAuth2JwtGrantIssuerParams) SetBody(body *models.AdminTrustOAuth2JwtGrantIssuerBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AdminTrustOAuth2JwtGrantIssuerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}