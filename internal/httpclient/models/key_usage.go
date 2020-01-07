// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// KeyUsage KeyUsage represents the set of actions that are valid for a given key. It's
// a bitmap of the KeyUsage* constants.
// swagger:model KeyUsage
type KeyUsage int64

// Validate validates this key usage
func (m KeyUsage) Validate(formats strfmt.Registry) error {
	return nil
}
