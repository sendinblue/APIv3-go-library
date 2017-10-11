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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSMTPReportParams creates a new GetSMTPReportParams object
// with the default values initialized.
func NewGetSMTPReportParams() *GetSMTPReportParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMTPReportParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSMTPReportParamsWithTimeout creates a new GetSMTPReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSMTPReportParamsWithTimeout(timeout time.Duration) *GetSMTPReportParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMTPReportParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetSMTPReportParamsWithContext creates a new GetSMTPReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSMTPReportParamsWithContext(ctx context.Context) *GetSMTPReportParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMTPReportParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetSMTPReportParamsWithHTTPClient creates a new GetSMTPReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSMTPReportParamsWithHTTPClient(client *http.Client) *GetSMTPReportParams {
	var (
		limitDefault  = int64(50)
		offsetDefault = int64(0)
	)
	return &GetSMTPReportParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetSMTPReportParams contains all the parameters to send to the API endpoint
for the get Smtp report operation typically these are written to a http.Request
*/
type GetSMTPReportParams struct {

	/*Days
	  Number of days in the past including today (positive integer). Not compatible with 'startDate' and 'endDate'

	*/
	Days *int64
	/*EndDate
	  Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD)

	*/
	EndDate *strfmt.Date
	/*Limit
	  Number of documents returned per page

	*/
	Limit *int64
	/*Offset
	  Index of the first document on the page

	*/
	Offset *int64
	/*StartDate
	  Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD)

	*/
	StartDate *strfmt.Date
	/*Tag
	  Tag of the emails

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Smtp report params
func (o *GetSMTPReportParams) WithTimeout(timeout time.Duration) *GetSMTPReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Smtp report params
func (o *GetSMTPReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Smtp report params
func (o *GetSMTPReportParams) WithContext(ctx context.Context) *GetSMTPReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Smtp report params
func (o *GetSMTPReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Smtp report params
func (o *GetSMTPReportParams) WithHTTPClient(client *http.Client) *GetSMTPReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Smtp report params
func (o *GetSMTPReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDays adds the days to the get Smtp report params
func (o *GetSMTPReportParams) WithDays(days *int64) *GetSMTPReportParams {
	o.SetDays(days)
	return o
}

// SetDays adds the days to the get Smtp report params
func (o *GetSMTPReportParams) SetDays(days *int64) {
	o.Days = days
}

// WithEndDate adds the endDate to the get Smtp report params
func (o *GetSMTPReportParams) WithEndDate(endDate *strfmt.Date) *GetSMTPReportParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get Smtp report params
func (o *GetSMTPReportParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithLimit adds the limit to the get Smtp report params
func (o *GetSMTPReportParams) WithLimit(limit *int64) *GetSMTPReportParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get Smtp report params
func (o *GetSMTPReportParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get Smtp report params
func (o *GetSMTPReportParams) WithOffset(offset *int64) *GetSMTPReportParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get Smtp report params
func (o *GetSMTPReportParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithStartDate adds the startDate to the get Smtp report params
func (o *GetSMTPReportParams) WithStartDate(startDate *strfmt.Date) *GetSMTPReportParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get Smtp report params
func (o *GetSMTPReportParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithTag adds the tag to the get Smtp report params
func (o *GetSMTPReportParams) WithTag(tag *string) *GetSMTPReportParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get Smtp report params
func (o *GetSMTPReportParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetSMTPReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Days != nil {

		// query param days
		var qrDays int64
		if o.Days != nil {
			qrDays = *o.Days
		}
		qDays := swag.FormatInt64(qrDays)
		if qDays != "" {
			if err := r.SetQueryParam("days", qDays); err != nil {
				return err
			}
		}

	}

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate strfmt.Date
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate strfmt.Date
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
