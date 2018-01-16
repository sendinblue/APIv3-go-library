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

// PostSendSMSTestFailed post send Sms test failed
// swagger:model postSendSmsTestFailed
type PostSendSMSTestFailed struct {

	// Response code
	// Required: true
	Code *int64 `json:"code"`

	// Response message
	// Required: true
	Message *string `json:"message"`

	// unexisting Sms
	UnexistingSMS []strfmt.Email `json:"unexistingSms"`

	// without list Sms
	WithoutListSMS []strfmt.Email `json:"withoutListSms"`
}

// Validate validates this post send Sms test failed
func (m *PostSendSMSTestFailed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUnexistingSMS(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWithoutListSMS(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostSendSMSTestFailed) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *PostSendSMSTestFailed) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *PostSendSMSTestFailed) validateUnexistingSMS(formats strfmt.Registry) error {

	if swag.IsZero(m.UnexistingSMS) { // not required
		return nil
	}

	return nil
}

func (m *PostSendSMSTestFailed) validateWithoutListSMS(formats strfmt.Registry) error {

	if swag.IsZero(m.WithoutListSMS) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostSendSMSTestFailed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSendSMSTestFailed) UnmarshalBinary(b []byte) error {
	var res PostSendSMSTestFailed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
