// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JSONWebKeySetGeneratorRequest JSONWebKeySetGeneratorRequest JSONWebKeySetGeneratorRequest JSONWebKeySetGeneratorRequest JSONWebKeySetGeneratorRequest JSONWebKeySetGeneratorRequest JSONWebKeySetGeneratorRequest JSONWebKeySetGeneratorRequest JSONWebKeySetGeneratorRequest json web key set generator request
//
// swagger:model jsonWebKeySetGeneratorRequest
type JSONWebKeySetGeneratorRequest struct {

	// The algorithm to be used for creating the key. Supports "RS256", "ES512", "HS512", and "HS256"
	// Required: true
	Alg *string `json:"alg"`

	// The kid of the key to be created
	// Required: true
	Kid *string `json:"kid"`

	// The "use" (public key use) parameter identifies the intended use of
	// the public key. The "use" parameter is employed to indicate whether
	// a public key is used for encrypting data or verifying the signature
	// on data. Valid values are "enc" and "sig".
	// Required: true
	Use *string `json:"use"`
}

// Validate validates this json web key set generator request
func (m *JSONWebKeySetGeneratorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONWebKeySetGeneratorRequest) validateAlg(formats strfmt.Registry) error {

	if err := validate.Required("alg", "body", m.Alg); err != nil {
		return err
	}

	return nil
}

func (m *JSONWebKeySetGeneratorRequest) validateKid(formats strfmt.Registry) error {

	if err := validate.Required("kid", "body", m.Kid); err != nil {
		return err
	}

	return nil
}

func (m *JSONWebKeySetGeneratorRequest) validateUse(formats strfmt.Registry) error {

	if err := validate.Required("use", "body", m.Use); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSONWebKeySetGeneratorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONWebKeySetGeneratorRequest) UnmarshalBinary(b []byte) error {
	var res JSONWebKeySetGeneratorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
