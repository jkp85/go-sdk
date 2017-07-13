package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TriggerAction trigger action
// swagger:model TriggerAction
type TriggerAction struct {

	// action name
	// Required: true
	ActionName *string `json:"action_name"`

	// id
	ID string `json:"id,omitempty"`

	// method
	// Required: true
	Method *string `json:"method"`

	// model
	Model string `json:"model,omitempty"`

	// object id
	ObjectID string `json:"object_id,omitempty"`

	// payload
	Payload interface{} `json:"payload,omitempty"`
}

// Validate validates this trigger action
func (m *TriggerAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TriggerAction) validateActionName(formats strfmt.Registry) error {

	if err := validate.Required("action_name", "body", m.ActionName); err != nil {
		return err
	}

	return nil
}

func (m *TriggerAction) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}
