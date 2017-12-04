// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetSMTPTemplateOverviewSender get Smtp template overview sender
// swagger:model getSmtpTemplateOverviewSender
type GetSMTPTemplateOverviewSender struct {

	// From email for the template
	Email strfmt.Email `json:"email,omitempty"`

	// From email for the template
	Name string `json:"name,omitempty"`
}

// Validate validates this get Smtp template overview sender
func (m *GetSMTPTemplateOverviewSender) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetSMTPTemplateOverviewSender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSMTPTemplateOverviewSender) UnmarshalBinary(b []byte) error {
	var res GetSMTPTemplateOverviewSender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
