// Code generated by go-swagger; DO NOT EDIT.

package sms_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// SendSMSReportReader is a Reader for the SendSMSReport structure.
type SendSMSReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendSMSReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewSendSMSReportNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendSMSReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendSMSReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendSMSReportNoContent creates a SendSMSReportNoContent with default headers values
func NewSendSMSReportNoContent() *SendSMSReportNoContent {
	return &SendSMSReportNoContent{}
}

/*SendSMSReportNoContent handles this case with default header values.

Report has been successfully sent to the defined recipients
*/
type SendSMSReportNoContent struct {
}

func (o *SendSMSReportNoContent) Error() string {
	return fmt.Sprintf("[POST /smsCampaigns/{campaignId}/sendReport][%d] sendSMSReportNoContent ", 204)
}

func (o *SendSMSReportNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendSMSReportBadRequest creates a SendSMSReportBadRequest with default headers values
func NewSendSMSReportBadRequest() *SendSMSReportBadRequest {
	return &SendSMSReportBadRequest{}
}

/*SendSMSReportBadRequest handles this case with default header values.

bad request
*/
type SendSMSReportBadRequest struct {
	Payload *models.ErrorModel
}

func (o *SendSMSReportBadRequest) Error() string {
	return fmt.Sprintf("[POST /smsCampaigns/{campaignId}/sendReport][%d] sendSMSReportBadRequest  %+v", 400, o.Payload)
}

func (o *SendSMSReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendSMSReportNotFound creates a SendSMSReportNotFound with default headers values
func NewSendSMSReportNotFound() *SendSMSReportNotFound {
	return &SendSMSReportNotFound{}
}

/*SendSMSReportNotFound handles this case with default header values.

Campaign ID not found
*/
type SendSMSReportNotFound struct {
	Payload *models.ErrorModel
}

func (o *SendSMSReportNotFound) Error() string {
	return fmt.Sprintf("[POST /smsCampaigns/{campaignId}/sendReport][%d] sendSMSReportNotFound  %+v", 404, o.Payload)
}

func (o *SendSMSReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
