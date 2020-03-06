// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetTransacEmailsList get transac emails list
// swagger:model getTransacEmailsList
type GetTransacEmailsList struct {

	// transactional emails
	TransactionalEmails []*GetTransacEmailsListTransactionalEmailsItems0 `json:"transactionalEmails"`
}

// Validate validates this get transac emails list
func (m *GetTransacEmailsList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransactionalEmails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTransacEmailsList) validateTransactionalEmails(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionalEmails) { // not required
		return nil
	}

	for i := 0; i < len(m.TransactionalEmails); i++ {
		if swag.IsZero(m.TransactionalEmails[i]) { // not required
			continue
		}

		if m.TransactionalEmails[i] != nil {
			if err := m.TransactionalEmails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transactionalEmails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacEmailsList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacEmailsList) UnmarshalBinary(b []byte) error {
	var res GetTransacEmailsList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetTransacEmailsListTransactionalEmailsItems0 get transac emails list transactional emails items0
// swagger:model GetTransacEmailsListTransactionalEmailsItems0
type GetTransacEmailsListTransactionalEmailsItems0 struct {

	// Date on which transactional email was sent
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// Email address to which transactional email has been sent
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// Message Id of the sent email
	// Required: true
	MessageID *string `json:"messageId"`

	// Subject of the sent email
	// Required: true
	Subject *string `json:"subject"`

	// Id of the template
	TemplateID int64 `json:"templateId,omitempty"`

	// Unique id of the email sent to a particular contact
	// Required: true
	UUID *string `json:"uuid"`
}

// Validate validates this get transac emails list transactional emails items0
func (m *GetTransacEmailsListTransactionalEmailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTransacEmailsListTransactionalEmailsItems0) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacEmailsListTransactionalEmailsItems0) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacEmailsListTransactionalEmailsItems0) validateMessageID(formats strfmt.Registry) error {

	if err := validate.Required("messageId", "body", m.MessageID); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacEmailsListTransactionalEmailsItems0) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacEmailsListTransactionalEmailsItems0) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacEmailsListTransactionalEmailsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacEmailsListTransactionalEmailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransacEmailsListTransactionalEmailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}