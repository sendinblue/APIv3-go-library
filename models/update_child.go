// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateChild update child
// swagger:model updateChild
type UpdateChild struct {

	// New Company name to use to update the child account
	CompanyName string `json:"companyName,omitempty"`

	// New Email address to update the child account
	Email strfmt.Email `json:"email,omitempty"`

	// New First name to use to update the child account
	FirstName string `json:"firstName,omitempty"`

	// ips
	Ips []int64 `json:"ips"`

	// New Last name to use to update the child account
	LastName string `json:"lastName,omitempty"`

	// New password for the child account to login
	Password strfmt.Password `json:"password,omitempty"`
}

// Validate validates this update child
func (m *UpdateChild) Validate(formats strfmt.Registry) error {
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

func (m *UpdateChild) validateIps(formats strfmt.Registry) error {

	if swag.IsZero(m.Ips) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateChild) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateChild) UnmarshalBinary(b []byte) error {
	var res UpdateChild
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
