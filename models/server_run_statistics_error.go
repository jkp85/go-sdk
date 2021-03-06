// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerRunStatisticsError server run statistics error
// swagger:model ServerRunStatisticsError
type ServerRunStatisticsError struct {

	// exit_code field errors.
	ExitCode []string `json:"exit_code"`

	// id field errors.
	ID []string `json:"id"`

	// Errors not connected to any field.
	NonFieldErrors []string `json:"non_field_errors"`

	// size field errors.
	Size []string `json:"size"`

	// stacktrace field errors.
	Stacktrace []string `json:"stacktrace"`

	// start field errors.
	Start []string `json:"start"`

	// stop field errors.
	Stop []string `json:"stop"`
}

// Validate validates this server run statistics error
func (m *ServerRunStatisticsError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExitCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStacktrace(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStop(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerRunStatisticsError) validateExitCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ExitCode) { // not required
		return nil
	}

	return nil
}

func (m *ServerRunStatisticsError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *ServerRunStatisticsError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *ServerRunStatisticsError) validateSize(formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	return nil
}

func (m *ServerRunStatisticsError) validateStacktrace(formats strfmt.Registry) error {

	if swag.IsZero(m.Stacktrace) { // not required
		return nil
	}

	return nil
}

func (m *ServerRunStatisticsError) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	return nil
}

func (m *ServerRunStatisticsError) validateStop(formats strfmt.Registry) error {

	if swag.IsZero(m.Stop) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerRunStatisticsError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerRunStatisticsError) UnmarshalBinary(b []byte) error {
	var res ServerRunStatisticsError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
