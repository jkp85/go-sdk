// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupUserError group user error
// swagger:model GroupUserError
type GroupUserError struct {

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`

	// User field errors.
	User []string `json:"user"`
}

// Validate validates this group user error
func (m *GroupUserError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupUserError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *GroupUserError) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupUserError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupUserError) UnmarshalBinary(b []byte) error {
	var res GroupUserError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}