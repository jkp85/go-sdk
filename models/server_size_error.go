// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerSizeError server size error
// swagger:model ServerSizeError
type ServerSizeError struct {

	// Active field errors.
	Active []string `json:"active"`

	// CPU field errors.
	CPU []string `json:"cpu"`

	// Id field errors.
	ID []string `json:"id"`

	// Memory field errors.
	Memory []string `json:"memory"`

	// Name field errors.
	Name []string `json:"name"`

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`
}

// Validate validates this server size error
func (m *ServerSizeError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerSizeError) validateActive(formats strfmt.Registry) error {

	if swag.IsZero(m.Active) { // not required
		return nil
	}

	return nil
}

func (m *ServerSizeError) validateCPU(formats strfmt.Registry) error {

	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	return nil
}

func (m *ServerSizeError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *ServerSizeError) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	return nil
}

func (m *ServerSizeError) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	return nil
}

func (m *ServerSizeError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerSizeError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerSizeError) UnmarshalBinary(b []byte) error {
	var res ServerSizeError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}