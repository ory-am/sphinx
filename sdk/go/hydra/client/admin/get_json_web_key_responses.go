// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra/sdk/go/hydra/models"
)

// GetJSONWebKeyReader is a Reader for the GetJSONWebKey structure.
type GetJSONWebKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJSONWebKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJSONWebKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetJSONWebKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJSONWebKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetJSONWebKeyOK creates a GetJSONWebKeyOK with default headers values
func NewGetJSONWebKeyOK() *GetJSONWebKeyOK {
	return &GetJSONWebKeyOK{}
}

/*GetJSONWebKeyOK handles this case with default header values.

JSONWebKeySet
*/
type GetJSONWebKeyOK struct {
	Payload *models.SwaggerJSONWebKeySet
}

func (o *GetJSONWebKeyOK) Error() string {
	return fmt.Sprintf("[GET /keys/{set}/{kid}][%d] getJsonWebKeyOK  %+v", 200, o.Payload)
}

func (o *GetJSONWebKeyOK) GetPayload() *models.SwaggerJSONWebKeySet {
	return o.Payload
}

func (o *GetJSONWebKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwaggerJSONWebKeySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJSONWebKeyNotFound creates a GetJSONWebKeyNotFound with default headers values
func NewGetJSONWebKeyNotFound() *GetJSONWebKeyNotFound {
	return &GetJSONWebKeyNotFound{}
}

/*GetJSONWebKeyNotFound handles this case with default header values.

genericError
*/
type GetJSONWebKeyNotFound struct {
	Payload *models.GenericError
}

func (o *GetJSONWebKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /keys/{set}/{kid}][%d] getJsonWebKeyNotFound  %+v", 404, o.Payload)
}

func (o *GetJSONWebKeyNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetJSONWebKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJSONWebKeyInternalServerError creates a GetJSONWebKeyInternalServerError with default headers values
func NewGetJSONWebKeyInternalServerError() *GetJSONWebKeyInternalServerError {
	return &GetJSONWebKeyInternalServerError{}
}

/*GetJSONWebKeyInternalServerError handles this case with default header values.

genericError
*/
type GetJSONWebKeyInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetJSONWebKeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /keys/{set}/{kid}][%d] getJsonWebKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJSONWebKeyInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetJSONWebKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
