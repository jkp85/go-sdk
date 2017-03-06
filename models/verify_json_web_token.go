package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VerifyJSONWebToken verify JSON web token
// swagger:model VerifyJSONWebToken
type VerifyJSONWebToken struct {

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this verify JSON web token
func (m *VerifyJSONWebToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VerifyJSONWebToken) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}
