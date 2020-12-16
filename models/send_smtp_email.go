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

// SendSMTPEmail send Smtp email
// swagger:model sendSmtpEmail
type SendSMTPEmail struct {

	// Pass the absolute URL (no local file) or the base64 content of the attachment along with the attachment name (Mandatory if attachment content is passed). For example, `[{"url":"https://attachment.domain.com/myAttachmentFromUrl.jpg", "name":"My attachment 1"}, {"content":"base64 exmaple content", "name":"My attachment 2"}]`. Allowed extensions for attachment file: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff, rtf, bmp, cgm, css, shtml, html, htm, zip, xml, ppt, pptx, tar, ez, ics, mobi, msg, pub, eps, odt, mp3, m4a, m4v, wma, ogg, flac, wav, aif, aifc, aiff, mp4, mov, avi, mkv, mpeg, mpg and wmv ( If 'templateId' is passed and is in New Template Language format then both attachment url and content are accepted. If template is in Old template Language format, then 'attachment' is ignored )
	Attachment []*SendSMTPEmailAttachmentItems0 `json:"attachment"`

	// List of email addresses and names (optional) of the recipients in bcc
	Bcc []*SendSMTPEmailBccItems0 `json:"bcc,omitempty"`

	// List of email addresses and names (optional) of the recipients in cc
	Cc []*SendSMTPEmailCcItems0 `json:"cc,omitempty"`

	// Pass the set of custom headers (not the standard headers) that shall be sent along the mail headers in the original email. 'sender.ip' header can be set (only for dedicated ip users) to mention the IP to be used for sending transactional emails. For example, `{"sender.ip":"1.2.3.4", "X-Mailin-custom":"some_custom_header"}`.
	Headers interface{} `json:"headers,omitempty"`

	// HTML body of the message ( Mandatory if 'templateId' is not passed, ignored if 'templateId' is passed )
	HTMLContent string `json:"htmlContent,omitempty"`

	// Pass the set of attributes to customize the template. For example, {'FNAME':'Joe', 'LNAME':'Doe'}. It's considered only if template is in New Template Language format.
	Params interface{} `json:"params,omitempty"`

	// reply to
	ReplyTo *SendSMTPEmailReplyTo `json:"replyTo,omitempty"`

	// sender
	Sender *SendSMTPEmailSender `json:"sender,omitempty"`

	// Subject of the message. Mandatory if 'templateId' is not passed
	Subject string `json:"subject,omitempty"`

	// Tag your emails to find them more easily
	Tags []string `json:"tags"`

	// Id of the template
	TemplateID int64 `json:"templateId,omitempty"`

	// Plain Text body of the message ( Ignored if 'templateId' is passed )
	TextContent string `json:"textContent,omitempty"`

	// List of email addresses and names (optional) of the recipients. For example, [{'name':'Jimmy', 'email':'jimmy98@example.com'}, {'name':'Joe', 'email':'joe@example.com'}]
	// Required: true
	To []*SendSMTPEmailToItems0 `json:"to"`
}

// Validate validates this send Smtp email
func (m *SendSMTPEmail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplyTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmail) validateAttachment(formats strfmt.Registry) error {

	if swag.IsZero(m.Attachment) { // not required
		return nil
	}

	for i := 0; i < len(m.Attachment); i++ {
		if swag.IsZero(m.Attachment[i]) { // not required
			continue
		}

		if m.Attachment[i] != nil {
			if err := m.Attachment[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachment" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SendSMTPEmail) validateBcc(formats strfmt.Registry) error {

	if swag.IsZero(m.Bcc) { // not required
		return nil
	}

	for i := 0; i < len(m.Bcc); i++ {
		if swag.IsZero(m.Bcc[i]) { // not required
			continue
		}

		if m.Bcc[i] != nil {
			if err := m.Bcc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bcc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SendSMTPEmail) validateCc(formats strfmt.Registry) error {

	if swag.IsZero(m.Cc) { // not required
		return nil
	}

	for i := 0; i < len(m.Cc); i++ {
		if swag.IsZero(m.Cc[i]) { // not required
			continue
		}

		if m.Cc[i] != nil {
			if err := m.Cc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SendSMTPEmail) validateReplyTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyTo) { // not required
		return nil
	}

	if m.ReplyTo != nil {
		if err := m.ReplyTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replyTo")
			}
			return err
		}
	}

	return nil
}

func (m *SendSMTPEmail) validateSender(formats strfmt.Registry) error {

	if swag.IsZero(m.Sender) { // not required
		return nil
	}

	if m.Sender != nil {
		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

func (m *SendSMTPEmail) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	for i := 0; i < len(m.To); i++ {
		if swag.IsZero(m.To[i]) { // not required
			continue
		}

		if m.To[i] != nil {
			if err := m.To[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmail) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SendSMTPEmailAttachmentItems0 send SMTP email attachment items0
// swagger:model SendSMTPEmailAttachmentItems0
type SendSMTPEmailAttachmentItems0 struct {

	// Base64 encoded chunk data of the attachment generated on the fly
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// Required if content is passed. Name of the attachment
	Name string `json:"name,omitempty"`

	// Absolute url of the attachment (no local file).
	URL string `json:"url,omitempty"`
}

// Validate validates this send SMTP email attachment items0
func (m *SendSMTPEmailAttachmentItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmailAttachmentItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmailAttachmentItems0) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmailAttachmentItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SendSMTPEmailBccItems0 send SMTP email bcc items0
// swagger:model SendSMTPEmailBccItems0
type SendSMTPEmailBccItems0 struct {

	// Email address of the recipient in bcc
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// Name of the recipient in bcc. Maximum allowed characters are 70.
	Name string `json:"name,omitempty"`
}

// Validate validates this send SMTP email bcc items0
func (m *SendSMTPEmailBccItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmailBccItems0) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmailBccItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmailBccItems0) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmailBccItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SendSMTPEmailCcItems0 send SMTP email cc items0
// swagger:model SendSMTPEmailCcItems0
type SendSMTPEmailCcItems0 struct {

	// Email address of the recipient in cc
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// Name of the recipient in cc. Maximum allowed characters are 70.
	Name string `json:"name,omitempty"`
}

// Validate validates this send SMTP email cc items0
func (m *SendSMTPEmailCcItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmailCcItems0) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmailCcItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmailCcItems0) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmailCcItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SendSMTPEmailReplyTo Email (required), along with name (optional), on which transactional mail recipients will be able to reply back. For example, {'email':'ann6533@example.com', 'name':'Ann'}.
// swagger:model SendSMTPEmailReplyTo
type SendSMTPEmailReplyTo struct {

	// Email address in reply to
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// Name in reply to. Maximum allowed characters are 70.
	Name string `json:"name,omitempty"`
}

// Validate validates this send SMTP email reply to
func (m *SendSMTPEmailReplyTo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmailReplyTo) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("replyTo"+"."+"email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("replyTo"+"."+"email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmailReplyTo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmailReplyTo) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmailReplyTo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SendSMTPEmailSender Mandatory if 'templateId' is not passed. Pass name (optional) and email of sender from which emails will be sent. For example, {'name':'Mary from MyShop', 'email':'no-reply@myshop.com'}
// swagger:model SendSMTPEmailSender
type SendSMTPEmailSender struct {

	// Email of the sender from which the emails will be sent
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// Name of the sender from which the emails will be sent. Maximum allowed characters are 70.
	Name string `json:"name,omitempty"`
}

// Validate validates this send SMTP email sender
func (m *SendSMTPEmailSender) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmailSender) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("sender"+"."+"email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("sender"+"."+"email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmailSender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmailSender) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmailSender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SendSMTPEmailToItems0 send SMTP email to items0
// swagger:model SendSMTPEmailToItems0
type SendSMTPEmailToItems0 struct {

	// Email address of the recipient
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// Name of the recipient. Maximum allowed characters are 70.
	Name string `json:"name,omitempty"`
}

// Validate validates this send SMTP email to items0
func (m *SendSMTPEmailToItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmailToItems0) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmailToItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmailToItems0) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmailToItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
