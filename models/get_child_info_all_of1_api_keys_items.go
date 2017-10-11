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

// GetChildInfoAllOf1APIKeysItems get child info all of1 Api keys items
// swagger:model getChildInfoAllOf1ApiKeysItems

type GetChildInfoAllOf1APIKeysItems struct {

	// API Key
	// Required: true
	Key *string `json:"key"`

	// Name of the key
	// Required: true
	Name *string `json:"name"`

	// Secret Key associated to the API Key (in case v1 Key is used only)
	Secret string `json:"secret,omitempty"`
}

/* polymorph getChildInfoAllOf1ApiKeysItems key false */

/* polymorph getChildInfoAllOf1ApiKeysItems name false */

/* polymorph getChildInfoAllOf1ApiKeysItems secret false */

// Validate validates this get child info all of1 Api keys items
func (m *GetChildInfoAllOf1APIKeysItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetChildInfoAllOf1APIKeysItems) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *GetChildInfoAllOf1APIKeysItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetChildInfoAllOf1APIKeysItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetChildInfoAllOf1APIKeysItems) UnmarshalBinary(b []byte) error {
	var res GetChildInfoAllOf1APIKeysItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}