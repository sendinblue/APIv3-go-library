// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetTransacSmsReportReportsItems get transac sms report reports items
// swagger:model getTransacSmsReportReportsItems

type GetTransacSmsReportReportsItems struct {

	// Number of accepted for the date
	// Required: true
	Accepted *int64 `json:"accepted"`

	// Number of blocked contact for the date
	// Required: true
	Blocked *int64 `json:"blocked"`

	// Date for which statistics are retrieved
	// Required: true
	Date *strfmt.Date `json:"date"`

	// Number of delivered SMS for the date
	// Required: true
	Delivered *int64 `json:"delivered"`

	// Number of hardbounces for the date
	// Required: true
	HardBounces *int64 `json:"hardBounces"`

	// Number of rejected for the date
	// Required: true
	Rejected *int64 `json:"rejected"`

	// Number of answered SMS for the date
	// Required: true
	Replied *int64 `json:"replied"`

	// Number of requests for the date
	// Required: true
	Requests *int64 `json:"requests"`

	// Number of softbounces for the date
	// Required: true
	SoftBounces *int64 `json:"softBounces"`

	// Tag specified in request
	// Required: true
	Tag *string `json:"tag"`

	// Number of unsubscription for the date
	// Required: true
	Unsubscribed *int64 `json:"unsubscribed"`
}

/* polymorph getTransacSmsReportReportsItems accepted false */

/* polymorph getTransacSmsReportReportsItems blocked false */

/* polymorph getTransacSmsReportReportsItems date false */

/* polymorph getTransacSmsReportReportsItems delivered false */

/* polymorph getTransacSmsReportReportsItems hardBounces false */

/* polymorph getTransacSmsReportReportsItems rejected false */

/* polymorph getTransacSmsReportReportsItems replied false */

/* polymorph getTransacSmsReportReportsItems requests false */

/* polymorph getTransacSmsReportReportsItems softBounces false */

/* polymorph getTransacSmsReportReportsItems tag false */

/* polymorph getTransacSmsReportReportsItems unsubscribed false */

// Validate validates this get transac sms report reports items
func (m *GetTransacSmsReportReportsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccepted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBlocked(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDelivered(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHardBounces(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRejected(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplied(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRequests(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSoftBounces(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUnsubscribed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTransacSmsReportReportsItems) validateAccepted(formats strfmt.Registry) error {

	if err := validate.Required("accepted", "body", m.Accepted); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateBlocked(formats strfmt.Registry) error {

	if err := validate.Required("blocked", "body", m.Blocked); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateDelivered(formats strfmt.Registry) error {

	if err := validate.Required("delivered", "body", m.Delivered); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateHardBounces(formats strfmt.Registry) error {

	if err := validate.Required("hardBounces", "body", m.HardBounces); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateRejected(formats strfmt.Registry) error {

	if err := validate.Required("rejected", "body", m.Rejected); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateReplied(formats strfmt.Registry) error {

	if err := validate.Required("replied", "body", m.Replied); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateRequests(formats strfmt.Registry) error {

	if err := validate.Required("requests", "body", m.Requests); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateSoftBounces(formats strfmt.Registry) error {

	if err := validate.Required("softBounces", "body", m.SoftBounces); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems) validateUnsubscribed(formats strfmt.Registry) error {

	if err := validate.Required("unsubscribed", "body", m.Unsubscribed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacSmsReportReportsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacSmsReportReportsItems) UnmarshalBinary(b []byte) error {
	var res GetTransacSmsReportReportsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
