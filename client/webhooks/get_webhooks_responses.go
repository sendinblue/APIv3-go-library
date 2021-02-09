// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// GetWebhooksReader is a Reader for the GetWebhooks structure.
type GetWebhooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebhooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWebhooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetWebhooksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWebhooksOK creates a GetWebhooksOK with default headers values
func NewGetWebhooksOK() *GetWebhooksOK {
	return &GetWebhooksOK{}
}

/*GetWebhooksOK handles this case with default header values.

Webhooks informations
*/
type GetWebhooksOK struct {
	Payload *models.GetWebhooks
}

func (o *GetWebhooksOK) Error() string {
	return fmt.Sprintf("[GET /webhooks][%d] getWebhooksOK  %+v", 200, o.Payload)
}

func (o *GetWebhooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetWebhooks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhooksBadRequest creates a GetWebhooksBadRequest with default headers values
func NewGetWebhooksBadRequest() *GetWebhooksBadRequest {
	return &GetWebhooksBadRequest{}
}

/*GetWebhooksBadRequest handles this case with default header values.

bad request
*/
type GetWebhooksBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetWebhooksBadRequest) Error() string {
	return fmt.Sprintf("[GET /webhooks][%d] getWebhooksBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebhooksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}