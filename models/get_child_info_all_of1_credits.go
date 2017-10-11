// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetChildInfoAllOf1Credits Credits available for your child
// swagger:model getChildInfoAllOf1Credits

type GetChildInfoAllOf1Credits struct {

	// Email credits available for your child
	EmailCredits int64 `json:"emailCredits,omitempty"`

	// SMS credits available for your child
	SmsCredits int64 `json:"smsCredits,omitempty"`
}

/* polymorph getChildInfoAllOf1Credits emailCredits false */

/* polymorph getChildInfoAllOf1Credits smsCredits false */

// Validate validates this get child info all of1 credits
func (m *GetChildInfoAllOf1Credits) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetChildInfoAllOf1Credits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetChildInfoAllOf1Credits) UnmarshalBinary(b []byte) error {
	var res GetChildInfoAllOf1Credits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
