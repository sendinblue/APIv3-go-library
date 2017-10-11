// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateUpdateFolder create update folder
// swagger:model createUpdateFolder

type CreateUpdateFolder struct {

	// Name of the folder
	Name string `json:"name,omitempty"`
}

/* polymorph createUpdateFolder name false */

// Validate validates this create update folder
func (m *CreateUpdateFolder) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreateUpdateFolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUpdateFolder) UnmarshalBinary(b []byte) error {
	var res CreateUpdateFolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
