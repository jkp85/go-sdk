// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlanData plan data
// swagger:model PlanData
type PlanData struct {

	// Plan amount in currency.
	// Required: true
	Amount *int64 `json:"amount"`

	// Currency for plan.
	Currency string `json:"currency,omitempty"`

	// Plan interval.
	// Required: true
	Interval *string `json:"interval"`

	// Number of intervals.
	// Required: true
	IntervalCount *int64 `json:"interval_count"`

	// Is plan live, or not.
	Livemode bool `json:"livemode,omitempty"`

	// Plan meta data.
	Metadata interface{} `json:"metadata,omitempty"`

	// Plan name.
	// Required: true
	Name *string `json:"name"`

	// Plan description.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// Trial days for try and buy campaigns.
	TrialPeriodDays int64 `json:"trial_period_days,omitempty"`
}

// Validate validates this plan data
func (m *PlanData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIntervalCount(formats); err != nil {
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

func (m *PlanData) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

var planDataTypeIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["day","week","month","year"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planDataTypeIntervalPropEnum = append(planDataTypeIntervalPropEnum, v)
	}
}

const (
	// PlanDataIntervalDay captures enum value "day"
	PlanDataIntervalDay string = "day"
	// PlanDataIntervalWeek captures enum value "week"
	PlanDataIntervalWeek string = "week"
	// PlanDataIntervalMonth captures enum value "month"
	PlanDataIntervalMonth string = "month"
	// PlanDataIntervalYear captures enum value "year"
	PlanDataIntervalYear string = "year"
)

// prop value enum
func (m *PlanData) validateIntervalEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, planDataTypeIntervalPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PlanData) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	// value enum
	if err := m.validateIntervalEnum("interval", "body", *m.Interval); err != nil {
		return err
	}

	return nil
}

func (m *PlanData) validateIntervalCount(formats strfmt.Registry) error {

	if err := validate.Required("interval_count", "body", m.IntervalCount); err != nil {
		return err
	}

	return nil
}

func (m *PlanData) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanData) UnmarshalBinary(b []byte) error {
	var res PlanData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}