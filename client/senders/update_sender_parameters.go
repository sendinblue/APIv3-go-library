// Code generated by go-swagger; DO NOT EDIT.

package senders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// NewUpdateSenderParams creates a new UpdateSenderParams object
// with the default values initialized.
func NewUpdateSenderParams() *UpdateSenderParams {
	var ()
	return &UpdateSenderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSenderParamsWithTimeout creates a new UpdateSenderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSenderParamsWithTimeout(timeout time.Duration) *UpdateSenderParams {
	var ()
	return &UpdateSenderParams{

		timeout: timeout,
	}
}

// NewUpdateSenderParamsWithContext creates a new UpdateSenderParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSenderParamsWithContext(ctx context.Context) *UpdateSenderParams {
	var ()
	return &UpdateSenderParams{

		Context: ctx,
	}
}

// NewUpdateSenderParamsWithHTTPClient creates a new UpdateSenderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSenderParamsWithHTTPClient(client *http.Client) *UpdateSenderParams {
	var ()
	return &UpdateSenderParams{
		HTTPClient: client,
	}
}

/*UpdateSenderParams contains all the parameters to send to the API endpoint
for the update sender operation typically these are written to a http.Request
*/
type UpdateSenderParams struct {

	/*Sender
	  sender's name

	*/
	Sender *models.UpdateSender
	/*SenderID
	  Id of the sender

	*/
	SenderID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update sender params
func (o *UpdateSenderParams) WithTimeout(timeout time.Duration) *UpdateSenderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update sender params
func (o *UpdateSenderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update sender params
func (o *UpdateSenderParams) WithContext(ctx context.Context) *UpdateSenderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update sender params
func (o *UpdateSenderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update sender params
func (o *UpdateSenderParams) WithHTTPClient(client *http.Client) *UpdateSenderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update sender params
func (o *UpdateSenderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSender adds the sender to the update sender params
func (o *UpdateSenderParams) WithSender(sender *models.UpdateSender) *UpdateSenderParams {
	o.SetSender(sender)
	return o
}

// SetSender adds the sender to the update sender params
func (o *UpdateSenderParams) SetSender(sender *models.UpdateSender) {
	o.Sender = sender
}

// WithSenderID adds the senderID to the update sender params
func (o *UpdateSenderParams) WithSenderID(senderID int64) *UpdateSenderParams {
	o.SetSenderID(senderID)
	return o
}

// SetSenderID adds the senderId to the update sender params
func (o *UpdateSenderParams) SetSenderID(senderID int64) {
	o.SenderID = senderID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSenderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Sender != nil {
		if err := r.SetBodyParam(o.Sender); err != nil {
			return err
		}
	}

	// path param senderId
	if err := r.SetPathParam("senderId", swag.FormatInt64(o.SenderID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
