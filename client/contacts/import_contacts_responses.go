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

// ImportContactsReader is a Reader for the ImportContacts structure.
type ImportContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewImportContactsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewImportContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportContactsAccepted creates a ImportContactsAccepted with default headers values
func NewImportContactsAccepted() *ImportContactsAccepted {
	return &ImportContactsAccepted{}
}

/*ImportContactsAccepted handles this case with default header values.

Contact import request has been accepted
*/
type ImportContactsAccepted struct {
	Payload *models.CreatedProcessID
}

func (o *ImportContactsAccepted) Error() string {
	return fmt.Sprintf("[POST /contacts/import][%d] importContactsAccepted  %+v", 202, o.Payload)
}

func (o *ImportContactsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatedProcessID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportContactsBadRequest creates a ImportContactsBadRequest with default headers values
func NewImportContactsBadRequest() *ImportContactsBadRequest {
	return &ImportContactsBadRequest{}
}

/*ImportContactsBadRequest handles this case with default header values.

bad request
*/
type ImportContactsBadRequest struct {
	Payload *models.ErrorModel
}

func (o *ImportContactsBadRequest) Error() string {
	return fmt.Sprintf("[POST /contacts/import][%d] importContactsBadRequest  %+v", 400, o.Payload)
}

func (o *ImportContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
