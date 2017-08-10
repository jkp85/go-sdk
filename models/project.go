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

// Project project
// swagger:model Project
type Project struct {

	// Array of project collaborators.
	Collaborators []string `json:"collaborators"`

	// Project description.
	Description string `json:"description,omitempty"`

	// Unique identifier for project as UUID.
	ID string `json:"id,omitempty"`

	// Date and time when project was last updated.
	LastUpdated string `json:"last_updated,omitempty"`

	// Project name.
	// Required: true
	Name *string `json:"name"`

	// Username of project owner.
	Owner string `json:"owner,omitempty"`

	// Value that states whether project is private or public.
	Private bool `json:"private,omitempty"`
}

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollaborators(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Project) validateCollaborators(formats strfmt.Registry) error {

	if swag.IsZero(m.Collaborators) { // not required
		return nil
	}

	return nil
}

func (m *Project) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
