// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RefreshJSONWebTokenError refresh JSON web token error
// swagger:model RefreshJSONWebTokenError
type RefreshJSONWebTokenError struct {

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`

	// Token field errors.
	Token []string `json:"token"`
}

// Validate validates this refresh JSON web token error
func (m *RefreshJSONWebTokenError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RefreshJSONWebTokenError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *RefreshJSONWebTokenError) validateToken(formats strfmt.Registry) error {

	if swag.IsZero(m.Token) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RefreshJSONWebTokenError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RefreshJSONWebTokenError) UnmarshalBinary(b []byte) error {
	var res RefreshJSONWebTokenError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}