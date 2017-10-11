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

// GetIpsReader is a Reader for the GetIps structure.
type GetIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIpsOK creates a GetIpsOK with default headers values
func NewGetIpsOK() *GetIpsOK {
	return &GetIpsOK{}
}

/*GetIpsOK handles this case with default header values.

list of dedicated IPs
*/
type GetIpsOK struct {
	Payload *models.GetIps
}

func (o *GetIpsOK) Error() string {
	return fmt.Sprintf("[GET /senders/ips][%d] getIpsOK  %+v", 200, o.Payload)
}

func (o *GetIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
