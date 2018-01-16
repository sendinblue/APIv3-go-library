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

// GetTransacAggregatedSMSReport get transac aggregated Sms report
// swagger:model getTransacAggregatedSmsReport
type GetTransacAggregatedSMSReport struct {

	// Number of accepted for the timeframe
	// Required: true
	Accepted *int64 `json:"accepted"`

	// Number of blocked contact for the timeframe
	// Required: true
	Blocked *int64 `json:"blocked"`

	// Number of delivered SMS for the timeframe
	// Required: true
	Delivered *int64 `json:"delivered"`

	// Number of hardbounces for the timeframe
	// Required: true
	HardBounces *int64 `json:"hardBounces"`

	// Time frame of the report
	// Required: true
	Range *string `json:"range"`

	// Number of rejected for the timeframe
	// Required: true
	Rejected *int64 `json:"rejected"`

	// Number of answered SMS for the timeframe
	// Required: true
	Replied *int64 `json:"replied"`

	// Number of requests for the timeframe
	// Required: true
	Requests *int64 `json:"requests"`

	// Number of softbounces for the timeframe
	// Required: true
	SoftBounces *int64 `json:"softBounces"`

	// Number of unsubscription for the timeframe
	// Required: true
	Unsubscribed *int64 `json:"unsubscribed"`
}

// Validate validates this get transac aggregated Sms report
func (m *GetTransacAggregatedSMSReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccepted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBlocked(formats); err != nil {
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

	if err := m.validateRange(formats); err != nil {
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

	if err := m.validateUnsubscribed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTransacAggregatedSMSReport) validateAccepted(formats strfmt.Registry) error {

	if err := validate.Required("accepted", "body", m.Accepted); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateBlocked(formats strfmt.Registry) error {

	if err := validate.Required("blocked", "body", m.Blocked); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateDelivered(formats strfmt.Registry) error {

	if err := validate.Required("delivered", "body", m.Delivered); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateHardBounces(formats strfmt.Registry) error {

	if err := validate.Required("hardBounces", "body", m.HardBounces); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateRejected(formats strfmt.Registry) error {

	if err := validate.Required("rejected", "body", m.Rejected); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateReplied(formats strfmt.Registry) error {

	if err := validate.Required("replied", "body", m.Replied); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateRequests(formats strfmt.Registry) error {

	if err := validate.Required("requests", "body", m.Requests); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateSoftBounces(formats strfmt.Registry) error {

	if err := validate.Required("softBounces", "body", m.SoftBounces); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacAggregatedSMSReport) validateUnsubscribed(formats strfmt.Registry) error {

	if err := validate.Required("unsubscribed", "body", m.Unsubscribed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacAggregatedSMSReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacAggregatedSMSReport) UnmarshalBinary(b []byte) error {
	var res GetTransacAggregatedSMSReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
