// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteHardbounces delete hardbounces
// swagger:model deleteHardbounces
type DeleteHardbounces struct {

	// Target a specific email address
	ContactEmail strfmt.Email `json:"contactEmail,omitempty"`

	// Ending date (YYYY-MM-DD) of the period from which the hardbounces will be deleted. Must be greater than equal to startDate
	EndDate strfmt.Date `json:"endDate,omitempty"`

	// Starting date (YYYY-MM-DD) of the period from which the hardbounces will be deleted. Must be lower than equal to endDate
	StartDate strfmt.Date `json:"startDate,omitempty"`
}

// Validate validates this delete hardbounces
func (m *DeleteHardbounces) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteHardbounces) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteHardbounces) UnmarshalBinary(b []byte) error {
	var res DeleteHardbounces
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
