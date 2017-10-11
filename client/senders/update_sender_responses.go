// Code generated by go-swagger; DO NOT EDIT.

package senders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// UpdateSenderReader is a Reader for the UpdateSender structure.
type UpdateSenderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSenderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSenderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSenderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSenderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSenderNoContent creates a UpdateSenderNoContent with default headers values
func NewUpdateSenderNoContent() *UpdateSenderNoContent {
	return &UpdateSenderNoContent{}
}

/*UpdateSenderNoContent handles this case with default header values.

sender updated
*/
type UpdateSenderNoContent struct {
}

func (o *UpdateSenderNoContent) Error() string {
	return fmt.Sprintf("[PUT /senders/{senderId}][%d] updateSenderNoContent ", 204)
}

func (o *UpdateSenderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSenderBadRequest creates a UpdateSenderBadRequest with default headers values
func NewUpdateSenderBadRequest() *UpdateSenderBadRequest {
	return &UpdateSenderBadRequest{}
}

/*UpdateSenderBadRequest handles this case with default header values.

bad request
*/
type UpdateSenderBadRequest struct {
	Payload *models.ErrorModel
}

func (o *UpdateSenderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /senders/{senderId}][%d] updateSenderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSenderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSenderNotFound creates a UpdateSenderNotFound with default headers values
func NewUpdateSenderNotFound() *UpdateSenderNotFound {
	return &UpdateSenderNotFound{}
}

/*UpdateSenderNotFound handles this case with default header values.

Sender ID not found
*/
type UpdateSenderNotFound struct {
	Payload *models.ErrorModel
}

func (o *UpdateSenderNotFound) Error() string {
	return fmt.Sprintf("[PUT /senders/{senderId}][%d] updateSenderNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSenderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
