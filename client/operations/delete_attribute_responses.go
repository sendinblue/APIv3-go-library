// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// DeleteAttributeReader is a Reader for the DeleteAttribute structure.
type DeleteAttributeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAttributeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAttributeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAttributeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteAttributeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAttributeNoContent creates a DeleteAttributeNoContent with default headers values
func NewDeleteAttributeNoContent() *DeleteAttributeNoContent {
	return &DeleteAttributeNoContent{}
}

/*DeleteAttributeNoContent handles this case with default header values.

Attribute deleted
*/
type DeleteAttributeNoContent struct {
}

func (o *DeleteAttributeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /contacts/attributes/{attributeCategory}/{attributeName}][%d] deleteAttributeNoContent ", 204)
}

func (o *DeleteAttributeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAttributeBadRequest creates a DeleteAttributeBadRequest with default headers values
func NewDeleteAttributeBadRequest() *DeleteAttributeBadRequest {
	return &DeleteAttributeBadRequest{}
}

/*DeleteAttributeBadRequest handles this case with default header values.

bad request
*/
type DeleteAttributeBadRequest struct {
	Payload *models.ErrorModel
}

func (o *DeleteAttributeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /contacts/attributes/{attributeCategory}/{attributeName}][%d] deleteAttributeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAttributeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAttributeNotFound creates a DeleteAttributeNotFound with default headers values
func NewDeleteAttributeNotFound() *DeleteAttributeNotFound {
	return &DeleteAttributeNotFound{}
}

/*DeleteAttributeNotFound handles this case with default header values.

Attribute not found
*/
type DeleteAttributeNotFound struct {
	Payload *models.ErrorModel
}

func (o *DeleteAttributeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /contacts/attributes/{attributeCategory}/{attributeName}][%d] deleteAttributeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAttributeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
