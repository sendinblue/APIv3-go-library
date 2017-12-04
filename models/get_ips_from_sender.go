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

// GetIpsFromSender get ips from sender
// swagger:model getIpsFromSender
type GetIpsFromSender struct {

	// ips
	// Required: true
	Ips GetIpsFromSenderIps `json:"ips"`
}

// Validate validates this get ips from sender
func (m *GetIpsFromSender) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIps(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetIpsFromSender) validateIps(formats strfmt.Registry) error {

	if err := validate.Required("ips", "body", m.Ips); err != nil {
		return err
	}

	if err := m.Ips.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ips")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetIpsFromSender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIpsFromSender) UnmarshalBinary(b []byte) error {
	var res GetIpsFromSender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
