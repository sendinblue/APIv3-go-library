// Code generated by go-swagger; DO NOT EDIT.

package smtp

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

// NewCreateSMTPTemplateParams creates a new CreateSMTPTemplateParams object
// with the default values initialized.
func NewCreateSMTPTemplateParams() *CreateSMTPTemplateParams {
	var ()
	return &CreateSMTPTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSMTPTemplateParamsWithTimeout creates a new CreateSMTPTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSMTPTemplateParamsWithTimeout(timeout time.Duration) *CreateSMTPTemplateParams {
	var ()
	return &CreateSMTPTemplateParams{

		timeout: timeout,
	}
}

// NewCreateSMTPTemplateParamsWithContext creates a new CreateSMTPTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSMTPTemplateParamsWithContext(ctx context.Context) *CreateSMTPTemplateParams {
	var ()
	return &CreateSMTPTemplateParams{

		Context: ctx,
	}
}

// NewCreateSMTPTemplateParamsWithHTTPClient creates a new CreateSMTPTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSMTPTemplateParamsWithHTTPClient(client *http.Client) *CreateSMTPTemplateParams {
	var ()
	return &CreateSMTPTemplateParams{
		HTTPClient: client,
	}
}

/*CreateSMTPTemplateParams contains all the parameters to send to the API endpoint
for the create Smtp template operation typically these are written to a http.Request
*/
type CreateSMTPTemplateParams struct {

	/*SMTPTemplate
	  values to update in smtp template

	*/
	SMTPTemplate *models.CreateSMTPTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create Smtp template params
func (o *CreateSMTPTemplateParams) WithTimeout(timeout time.Duration) *CreateSMTPTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Smtp template params
func (o *CreateSMTPTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Smtp template params
func (o *CreateSMTPTemplateParams) WithContext(ctx context.Context) *CreateSMTPTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Smtp template params
func (o *CreateSMTPTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Smtp template params
func (o *CreateSMTPTemplateParams) WithHTTPClient(client *http.Client) *CreateSMTPTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Smtp template params
func (o *CreateSMTPTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSMTPTemplate adds the sMTPTemplate to the create Smtp template params
func (o *CreateSMTPTemplateParams) WithSMTPTemplate(sMTPTemplate *models.CreateSMTPTemplate) *CreateSMTPTemplateParams {
	o.SetSMTPTemplate(sMTPTemplate)
	return o
}

// SetSMTPTemplate adds the smtpTemplate to the create Smtp template params
func (o *CreateSMTPTemplateParams) SetSMTPTemplate(sMTPTemplate *models.CreateSMTPTemplate) {
	o.SMTPTemplate = sMTPTemplate
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSMTPTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SMTPTemplate != nil {
		if err := r.SetBodyParam(o.SMTPTemplate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}