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

// GetExtendedCampaignStats get extended campaign stats
// swagger:model getExtendedCampaignStats
type GetExtendedCampaignStats struct {

	// campaign stats
	// Required: true
	CampaignStats GetExtendedCampaignStatsCampaignStats `json:"campaignStats"`

	// links stats
	// Required: true
	LinksStats GetExtendedCampaignStatsLinksStats `json:"linksStats"`

	// Number of clicks on mirror link
	// Required: true
	MirrorClick *int64 `json:"mirrorClick"`

	// Number of remaning emails to send
	// Required: true
	Remaining *int64 `json:"remaining"`

	// stats by domain
	// Required: true
	StatsByDomain GetStatsByDomain `json:"statsByDomain"`
}

// Validate validates this get extended campaign stats
func (m *GetExtendedCampaignStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaignStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinksStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMirrorClick(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemaining(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatsByDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetExtendedCampaignStats) validateCampaignStats(formats strfmt.Registry) error {

	if err := validate.Required("campaignStats", "body", m.CampaignStats); err != nil {
		return err
	}

	if err := m.CampaignStats.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("campaignStats")
		}
		return err
	}

	return nil
}

func (m *GetExtendedCampaignStats) validateLinksStats(formats strfmt.Registry) error {

	return nil
}

func (m *GetExtendedCampaignStats) validateMirrorClick(formats strfmt.Registry) error {

	if err := validate.Required("mirrorClick", "body", m.MirrorClick); err != nil {
		return err
	}

	return nil
}

func (m *GetExtendedCampaignStats) validateRemaining(formats strfmt.Registry) error {

	if err := validate.Required("remaining", "body", m.Remaining); err != nil {
		return err
	}

	return nil
}

func (m *GetExtendedCampaignStats) validateStatsByDomain(formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetExtendedCampaignStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetExtendedCampaignStats) UnmarshalBinary(b []byte) error {
	var res GetExtendedCampaignStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}