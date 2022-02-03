// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FlushInactiveOAuth2TokensRequest FlushInactiveOAuth2TokensRequest flush inactive o auth2 tokens request
//
// swagger:model flushInactiveOAuth2TokensRequest
type FlushInactiveOAuth2TokensRequest struct {

	// NotAfter sets after which point tokens should not be flushed. This is useful when you want to keep a history
	// of recently issued tokens for auditing.
	// Format: date-time
	// Format: date-time
	NotAfter strfmt.DateTime `json:"notAfter,omitempty"`
}

// Validate validates this flush inactive o auth2 tokens request
func (m *FlushInactiveOAuth2TokensRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotAfter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlushInactiveOAuth2TokensRequest) validateNotAfter(formats strfmt.Registry) error {
	if swag.IsZero(m.NotAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("notAfter", "body", "date-time", m.NotAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this flush inactive o auth2 tokens request based on context it is used
func (m *FlushInactiveOAuth2TokensRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlushInactiveOAuth2TokensRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlushInactiveOAuth2TokensRequest) UnmarshalBinary(b []byte) error {
	var res FlushInactiveOAuth2TokensRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
