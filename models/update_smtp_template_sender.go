// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateSMTPTemplateSender update Smtp template sender
// swagger:model updateSmtpTemplateSender

type UpdateSMTPTemplateSender struct {

	// Email of the sender
	Email strfmt.Email `json:"email,omitempty"`

	// Name of the sender
	Name string `json:"name,omitempty"`
}

/* polymorph updateSmtpTemplateSender email false */

/* polymorph updateSmtpTemplateSender name false */

// Validate validates this update Smtp template sender
func (m *UpdateSMTPTemplateSender) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSMTPTemplateSender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSMTPTemplateSender) UnmarshalBinary(b []byte) error {
	var res UpdateSMTPTemplateSender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
