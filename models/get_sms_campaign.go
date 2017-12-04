// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetSmsCampaign get sms campaign
// swagger:model getSmsCampaign
type GetSmsCampaign struct {
	GetSmsCampaignOverview

	GetSmsCampaignAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetSmsCampaign) UnmarshalJSON(raw []byte) error {

	var aO0 GetSmsCampaignOverview
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetSmsCampaignOverview = aO0

	var aO1 GetSmsCampaignAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetSmsCampaignAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetSmsCampaign) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.GetSmsCampaignOverview)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetSmsCampaignAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get sms campaign
func (m *GetSmsCampaign) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.GetSmsCampaignOverview.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.GetSmsCampaignAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetSmsCampaign) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSmsCampaign) UnmarshalBinary(b []byte) error {
	var res GetSmsCampaign
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
