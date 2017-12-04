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

// SendSMTPEmailReplyTo Email on which transactional mail recipients will be able to reply to
// swagger:model sendSmtpEmailReplyTo
type SendSMTPEmailReplyTo struct {

	// Email address in reply to
	// Required: true
	Email *strfmt.Email `json:"email"`

	// Name in reply to
	Name string `json:"name,omitempty"`
}

// Validate validates this send Smtp email reply to
func (m *SendSMTPEmailReplyTo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmailReplyTo) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmailReplyTo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmailReplyTo) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmailReplyTo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
