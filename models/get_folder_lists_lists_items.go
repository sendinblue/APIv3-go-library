// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetFolderListsListsItems get folder lists lists items
// swagger:model getFolderListsListsItems

type GetFolderListsListsItems struct {
	GetList
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetFolderListsListsItems) UnmarshalJSON(raw []byte) error {

	var aO0 GetList
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetList = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetFolderListsListsItems) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.GetList)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get folder lists lists items
func (m *GetFolderListsListsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.GetList.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetFolderListsListsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFolderListsListsItems) UnmarshalBinary(b []byte) error {
	var res GetFolderListsListsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
