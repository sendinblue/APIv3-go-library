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

// GetExtendedContactDetailsAllOf1StatisticsComplaintsItems get extended contact details all of1 statistics complaints items
// swagger:model getExtendedContactDetailsAllOf1StatisticsComplaintsItems
type GetExtendedContactDetailsAllOf1StatisticsComplaintsItems struct {

	// ID of the campaign which generated the event
	// Required: true
	CampaignID *int64 `json:"campaignId"`

	// UTC date-time of the event
	// Required: true
	EventTime *strfmt.DateTime `json:"eventTime"`
}

// Validate validates this get extended contact details all of1 statistics complaints items
func (m *GetExtendedContactDetailsAllOf1StatisticsComplaintsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaignID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetExtendedContactDetailsAllOf1StatisticsComplaintsItems) validateCampaignID(formats strfmt.Registry) error {

	if err := validate.Required("campaignId", "body", m.CampaignID); err != nil {
		return err
	}

	return nil
}

func (m *GetExtendedContactDetailsAllOf1StatisticsComplaintsItems) validateEventTime(formats strfmt.Registry) error {

	if err := validate.Required("eventTime", "body", m.EventTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetExtendedContactDetailsAllOf1StatisticsComplaintsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetExtendedContactDetailsAllOf1StatisticsComplaintsItems) UnmarshalBinary(b []byte) error {
	var res GetExtendedContactDetailsAllOf1StatisticsComplaintsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
