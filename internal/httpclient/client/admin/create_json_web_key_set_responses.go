// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// CreateJSONWebKeySetReader is a Reader for the CreateJSONWebKeySet structure.
type CreateJSONWebKeySetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateJSONWebKeySetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateJSONWebKeySetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateJSONWebKeySetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateJSONWebKeySetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateJSONWebKeySetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateJSONWebKeySetCreated creates a CreateJSONWebKeySetCreated with default headers values
func NewCreateJSONWebKeySetCreated() *CreateJSONWebKeySetCreated {
	return &CreateJSONWebKeySetCreated{}
}

/* CreateJSONWebKeySetCreated describes a response with status code 201, with default header values.

JSONWebKeySet
*/
type CreateJSONWebKeySetCreated struct {
	Payload *models.JSONWebKeySet
}

func (o *CreateJSONWebKeySetCreated) Error() string {
	return fmt.Sprintf("[POST /keys/{set}][%d] createJsonWebKeySetCreated  %+v", 201, o.Payload)
}
func (o *CreateJSONWebKeySetCreated) GetPayload() *models.JSONWebKeySet {
	return o.Payload
}

func (o *CreateJSONWebKeySetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONWebKeySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJSONWebKeySetUnauthorized creates a CreateJSONWebKeySetUnauthorized with default headers values
func NewCreateJSONWebKeySetUnauthorized() *CreateJSONWebKeySetUnauthorized {
	return &CreateJSONWebKeySetUnauthorized{}
}

/* CreateJSONWebKeySetUnauthorized describes a response with status code 401, with default header values.

oAuth2ApiError
*/
type CreateJSONWebKeySetUnauthorized struct {
	Payload *models.OAuth2APIError
}

func (o *CreateJSONWebKeySetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /keys/{set}][%d] createJsonWebKeySetUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateJSONWebKeySetUnauthorized) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *CreateJSONWebKeySetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJSONWebKeySetForbidden creates a CreateJSONWebKeySetForbidden with default headers values
func NewCreateJSONWebKeySetForbidden() *CreateJSONWebKeySetForbidden {
	return &CreateJSONWebKeySetForbidden{}
}

/* CreateJSONWebKeySetForbidden describes a response with status code 403, with default header values.

oAuth2ApiError
*/
type CreateJSONWebKeySetForbidden struct {
	Payload *models.OAuth2APIError
}

func (o *CreateJSONWebKeySetForbidden) Error() string {
	return fmt.Sprintf("[POST /keys/{set}][%d] createJsonWebKeySetForbidden  %+v", 403, o.Payload)
}
func (o *CreateJSONWebKeySetForbidden) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *CreateJSONWebKeySetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJSONWebKeySetInternalServerError creates a CreateJSONWebKeySetInternalServerError with default headers values
func NewCreateJSONWebKeySetInternalServerError() *CreateJSONWebKeySetInternalServerError {
	return &CreateJSONWebKeySetInternalServerError{}
}

/* CreateJSONWebKeySetInternalServerError describes a response with status code 500, with default header values.

oAuth2ApiError
*/
type CreateJSONWebKeySetInternalServerError struct {
	Payload *models.OAuth2APIError
}

func (o *CreateJSONWebKeySetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /keys/{set}][%d] createJsonWebKeySetInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateJSONWebKeySetInternalServerError) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *CreateJSONWebKeySetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
