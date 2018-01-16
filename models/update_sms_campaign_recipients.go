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

// UpdateSMSCampaignRecipients update Sms campaign recipients
// swagger:model updateSmsCampaignRecipients
type UpdateSMSCampaignRecipients struct {

	// List ids which have to be excluded from a campaign
	// Required: true
	ExclusionListIds []int64 `json:"exclusionListIds"`

	// Lists Ids to send the campaign to. REQUIRED if scheduledAt is not empty
	// Required: true
	ListIds []int64 `json:"listIds"`
}

// Validate validates this update Sms campaign recipients
func (m *UpdateSMSCampaignRecipients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusionListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSMSCampaignRecipients) validateExclusionListIds(formats strfmt.Registry) error {

	if err := validate.Required("exclusionListIds", "body", m.ExclusionListIds); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSMSCampaignRecipients) validateListIds(formats strfmt.Registry) error {

	if err := validate.Required("listIds", "body", m.ListIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSMSCampaignRecipients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSMSCampaignRecipients) UnmarshalBinary(b []byte) error {
	var res UpdateSMSCampaignRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
