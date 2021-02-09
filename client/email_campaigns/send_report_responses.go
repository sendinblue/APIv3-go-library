// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// SendReportReader is a Reader for the SendReport structure.
type SendReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewSendReportNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendReportNoContent creates a SendReportNoContent with default headers values
func NewSendReportNoContent() *SendReportNoContent {
	return &SendReportNoContent{}
}

/*SendReportNoContent handles this case with default header values.

Report has been successfully sent to the defined recipients
*/
type SendReportNoContent struct {
}

func (o *SendReportNoContent) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/sendReport][%d] sendReportNoContent ", 204)
}

func (o *SendReportNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendReportBadRequest creates a SendReportBadRequest with default headers values
func NewSendReportBadRequest() *SendReportBadRequest {
	return &SendReportBadRequest{}
}

/*SendReportBadRequest handles this case with default header values.

bad request
*/
type SendReportBadRequest struct {
	Payload *models.ErrorModel
}

func (o *SendReportBadRequest) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/sendReport][%d] sendReportBadRequest  %+v", 400, o.Payload)
}

func (o *SendReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendReportNotFound creates a SendReportNotFound with default headers values
func NewSendReportNotFound() *SendReportNotFound {
	return &SendReportNotFound{}
}

/*SendReportNotFound handles this case with default header values.

Campaign ID not found
*/
type SendReportNotFound struct {
	Payload *models.ErrorModel
}

func (o *SendReportNotFound) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/sendReport][%d] sendReportNotFound  %+v", 404, o.Payload)
}

func (o *SendReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}