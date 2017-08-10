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

// ServerData server data
// swagger:model ServerData
type ServerData struct {

	// Server configuration option. Values are jupyter, restful and cron.
	Config *ServerConfig `json:"config,omitempty"`

	// Array of other servers the server is connected to.
	Connected []string `json:"connected"`

	// External host IPv4 address or hostname.
	Host string `json:"host,omitempty"`

	// Image name.
	ImageName string `json:"image_name,omitempty"`

	// Server name.
	// Required: true
	Name *string `json:"name"`

	// Server size unique identifier.
	ServerSize string `json:"server_size,omitempty"`

	// Startup script to run when launching server.
	StartupScript string `json:"startup_script,omitempty"`
}

// Validate validates this server data
func (m *ServerData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnected(formats); err != nil {
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

func (m *ServerData) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {

		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *ServerData) validateConnected(formats strfmt.Registry) error {

	if swag.IsZero(m.Connected) { // not required
		return nil
	}

	return nil
}

func (m *ServerData) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerData) UnmarshalBinary(b []byte) error {
	var res ServerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
