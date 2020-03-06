// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// RemoveContactFromListReader is a Reader for the RemoveContactFromList structure.
type RemoveContactFromListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveContactFromListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRemoveContactFromListCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveContactFromListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveContactFromListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveContactFromListCreated creates a RemoveContactFromListCreated with default headers values
func NewRemoveContactFromListCreated() *RemoveContactFromListCreated {
	return &RemoveContactFromListCreated{}
}

/*RemoveContactFromListCreated handles this case with default header values.

All contacts have been removed successfully from the list with details of failed ones
*/
type RemoveContactFromListCreated struct {
	Payload *models.PostContactInfo
}

func (o *RemoveContactFromListCreated) Error() string {
	return fmt.Sprintf("[POST /contacts/lists/{listId}/contacts/remove][%d] removeContactFromListCreated  %+v", 201, o.Payload)
}

func (o *RemoveContactFromListCreated) GetPayload() *models.PostContactInfo {
	return o.Payload
}

func (o *RemoveContactFromListCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostContactInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveContactFromListBadRequest creates a RemoveContactFromListBadRequest with default headers values
func NewRemoveContactFromListBadRequest() *RemoveContactFromListBadRequest {
	return &RemoveContactFromListBadRequest{}
}

/*RemoveContactFromListBadRequest handles this case with default header values.

bad request
*/
type RemoveContactFromListBadRequest struct {
	Payload *models.ErrorModel
}

func (o *RemoveContactFromListBadRequest) Error() string {
	return fmt.Sprintf("[POST /contacts/lists/{listId}/contacts/remove][%d] removeContactFromListBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveContactFromListBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *RemoveContactFromListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveContactFromListNotFound creates a RemoveContactFromListNotFound with default headers values
func NewRemoveContactFromListNotFound() *RemoveContactFromListNotFound {
	return &RemoveContactFromListNotFound{}
}

/*RemoveContactFromListNotFound handles this case with default header values.

List ID not found
*/
type RemoveContactFromListNotFound struct {
	Payload *models.ErrorModel
}

func (o *RemoveContactFromListNotFound) Error() string {
	return fmt.Sprintf("[POST /contacts/lists/{listId}/contacts/remove][%d] removeContactFromListNotFound  %+v", 404, o.Payload)
}

func (o *RemoveContactFromListNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *RemoveContactFromListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}