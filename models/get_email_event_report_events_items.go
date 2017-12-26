// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetEmailEventReportEventsItems get email event report events items
// swagger:model getEmailEventReportEventsItems
type GetEmailEventReportEventsItems struct {

	// UTC date-time on which the event has been generated
	// Required: true
	Date *strfmt.DateTime `json:"date"`

	// Email address which generates the event
	// Required: true
	Email *strfmt.Email `json:"email"`

	// Event which occurred
	// Required: true
	Event *string `json:"event"`

	// Sender email from which the emails are sent
	// Required: true
	From *strfmt.Email `json:"from"`

	// IP from which the user has opened the email or clicked on the link (only available if the event is opened or clicks)
	IP string `json:"ip,omitempty"`

	// The link which is sent to the user (only available if the event is requests or opened or clicks)
	Link string `json:"link,omitempty"`

	// Message ID which generated the event
	// Required: true
	MessageID *string `json:"messageId"`

	// Reason of bounce (only available if the event is hardbounce or softbounce)
	Reason string `json:"reason,omitempty"`

	// Subject of the event
	Subject string `json:"subject,omitempty"`

	// Tag of the email which generated the event
	// Required: true
	Tag *string `json:"tag"`
}

// Validate validates this get email event report events items
func (m *GetEmailEventReportEventsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessageID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetEmailEventReportEventsItems) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetEmailEventReportEventsItems) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

var getEmailEventReportEventsItemsTypeEventPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bounces","hardBounces","softBounces","delivered","spam","requests","opened","clicks","invalid","deferred","blocked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getEmailEventReportEventsItemsTypeEventPropEnum = append(getEmailEventReportEventsItemsTypeEventPropEnum, v)
	}
}

const (
	// GetEmailEventReportEventsItemsEventBounces captures enum value "bounces"
	GetEmailEventReportEventsItemsEventBounces string = "bounces"
	// GetEmailEventReportEventsItemsEventHardBounces captures enum value "hardBounces"
	GetEmailEventReportEventsItemsEventHardBounces string = "hardBounces"
	// GetEmailEventReportEventsItemsEventSoftBounces captures enum value "softBounces"
	GetEmailEventReportEventsItemsEventSoftBounces string = "softBounces"
	// GetEmailEventReportEventsItemsEventDelivered captures enum value "delivered"
	GetEmailEventReportEventsItemsEventDelivered string = "delivered"
	// GetEmailEventReportEventsItemsEventSpam captures enum value "spam"
	GetEmailEventReportEventsItemsEventSpam string = "spam"
	// GetEmailEventReportEventsItemsEventRequests captures enum value "requests"
	GetEmailEventReportEventsItemsEventRequests string = "requests"
	// GetEmailEventReportEventsItemsEventOpened captures enum value "opened"
	GetEmailEventReportEventsItemsEventOpened string = "opened"
	// GetEmailEventReportEventsItemsEventClicks captures enum value "clicks"
	GetEmailEventReportEventsItemsEventClicks string = "clicks"
	// GetEmailEventReportEventsItemsEventInvalid captures enum value "invalid"
	GetEmailEventReportEventsItemsEventInvalid string = "invalid"
	// GetEmailEventReportEventsItemsEventDeferred captures enum value "deferred"
	GetEmailEventReportEventsItemsEventDeferred string = "deferred"
	// GetEmailEventReportEventsItemsEventBlocked captures enum value "blocked"
	GetEmailEventReportEventsItemsEventBlocked string = "blocked"
)

// prop value enum
func (m *GetEmailEventReportEventsItems) validateEventEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getEmailEventReportEventsItemsTypeEventPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetEmailEventReportEventsItems) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event); err != nil {
		return err
	}

	// value enum
	if err := m.validateEventEnum("event", "body", *m.Event); err != nil {
		return err
	}

	return nil
}

func (m *GetEmailEventReportEventsItems) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if err := validate.FormatOf("from", "body", "email", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetEmailEventReportEventsItems) validateMessageID(formats strfmt.Registry) error {

	if err := validate.Required("messageId", "body", m.MessageID); err != nil {
		return err
	}

	return nil
}

func (m *GetEmailEventReportEventsItems) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetEmailEventReportEventsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEmailEventReportEventsItems) UnmarshalBinary(b []byte) error {
	var res GetEmailEventReportEventsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
