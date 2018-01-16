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

// UpdateEmailCampaignSender update email campaign sender
// swagger:model updateEmailCampaignSender
type UpdateEmailCampaignSender struct {

	// Sender email from which the campaign emails are sent
	Email strfmt.Email `json:"email,omitempty"`

	// Sender Name from which the campaign emails are sent
	Name string `json:"name,omitempty"`
}

// Validate validates this update email campaign sender
func (m *UpdateEmailCampaignSender) Validate(formats strfmt.Registry) error {
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

func (m *UpdateEmailCampaignSender) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateEmailCampaignSender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateEmailCampaignSender) UnmarshalBinary(b []byte) error {
	var res UpdateEmailCampaignSender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
