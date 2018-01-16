// Code generated by go-swagger; DO NOT EDIT.

package transactional_sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// NewSendTransacSMSParams creates a new SendTransacSMSParams object
// with the default values initialized.
func NewSendTransacSMSParams() *SendTransacSMSParams {
	var ()
	return &SendTransacSMSParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendTransacSMSParamsWithTimeout creates a new SendTransacSMSParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendTransacSMSParamsWithTimeout(timeout time.Duration) *SendTransacSMSParams {
	var ()
	return &SendTransacSMSParams{

		timeout: timeout,
	}
}

// NewSendTransacSMSParamsWithContext creates a new SendTransacSMSParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendTransacSMSParamsWithContext(ctx context.Context) *SendTransacSMSParams {
	var ()
	return &SendTransacSMSParams{

		Context: ctx,
	}
}

// NewSendTransacSMSParamsWithHTTPClient creates a new SendTransacSMSParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendTransacSMSParamsWithHTTPClient(client *http.Client) *SendTransacSMSParams {
	var ()
	return &SendTransacSMSParams{
		HTTPClient: client,
	}
}

/*SendTransacSMSParams contains all the parameters to send to the API endpoint
for the send transac Sms operation typically these are written to a http.Request
*/
type SendTransacSMSParams struct {

	/*SendTransacSMS
	  Values to send a transactional SMS

	*/
	SendTransacSMS *models.SendTransacSMS

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send transac Sms params
func (o *SendTransacSMSParams) WithTimeout(timeout time.Duration) *SendTransacSMSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send transac Sms params
func (o *SendTransacSMSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send transac Sms params
func (o *SendTransacSMSParams) WithContext(ctx context.Context) *SendTransacSMSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send transac Sms params
func (o *SendTransacSMSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send transac Sms params
func (o *SendTransacSMSParams) WithHTTPClient(client *http.Client) *SendTransacSMSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send transac Sms params
func (o *SendTransacSMSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSendTransacSMS adds the sendTransacSMS to the send transac Sms params
func (o *SendTransacSMSParams) WithSendTransacSMS(sendTransacSMS *models.SendTransacSMS) *SendTransacSMSParams {
	o.SetSendTransacSMS(sendTransacSMS)
	return o
}

// SetSendTransacSMS adds the sendTransacSms to the send transac Sms params
func (o *SendTransacSMSParams) SetSendTransacSMS(sendTransacSMS *models.SendTransacSMS) {
	o.SendTransacSMS = sendTransacSMS
}

// WriteToRequest writes these params to a swagger request
func (o *SendTransacSMSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SendTransacSMS != nil {
		if err := r.SetBodyParam(o.SendTransacSMS); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
