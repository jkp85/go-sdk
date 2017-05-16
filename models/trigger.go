package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Trigger trigger
// swagger:model Trigger
type Trigger struct {

	// cause
	Cause *TriggerAction `json:"cause,omitempty"`

	// effect
	Effect *TriggerAction `json:"effect,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Cron schedule
	Schedule string `json:"schedule,omitempty"`

	// webhook
	Webhook *Webhook `json:"webhook,omitempty"`
}

// Validate validates this trigger
func (m *Trigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEffect(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWebhook(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Trigger) validateCause(formats strfmt.Registry) error {

	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	if m.Cause != nil {

		if err := m.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *Trigger) validateEffect(formats strfmt.Registry) error {

	if swag.IsZero(m.Effect) { // not required
		return nil
	}

	if m.Effect != nil {

		if err := m.Effect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

func (m *Trigger) validateWebhook(formats strfmt.Registry) error {

	if swag.IsZero(m.Webhook) { // not required
		return nil
	}

	if m.Webhook != nil {

		if err := m.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook")
			}
			return err
		}
	}

	return nil
}