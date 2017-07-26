// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriptionError subscription error
// swagger:model SubscriptionError
type SubscriptionError struct {

	// application_fee_percent field errors
	ApplicationFeePercent []string `json:"application_fee_percent"`

	// cancel_at_period_end field errors
	CancelAtPeriodEnd []string `json:"cancel_at_period_end"`

	// canceled_at field errors
	CanceledAt []string `json:"canceled_at"`

	// created field errors
	Created []string `json:"created"`

	// current_period_end field errors
	CurrentPeriodEnd []string `json:"current_period_end"`

	// current_period_start field errors
	CurrentPeriodStart []string `json:"current_period_start"`

	// ended_at field errors
	EndedAt []string `json:"ended_at"`

	// id field errors
	ID []string `json:"id"`

	// livemode field errors
	Livemode []string `json:"livemode"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// plan field errors
	Plan []string `json:"plan"`

	// quantity field errors
	Quantity []string `json:"quantity"`

	// start field errors
	Start []string `json:"start"`

	// status field errors
	Status []string `json:"status"`

	// stripe_id field errors
	StripeID []string `json:"stripe_id"`

	// trial_end field errors
	TrialEnd []string `json:"trial_end"`

	// trial_start field errors
	TrialStart []string `json:"trial_start"`
}

// Validate validates this subscription error
func (m *SubscriptionError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationFeePercent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCancelAtPeriodEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCanceledAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrentPeriodEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrentPeriodStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEndedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLivemode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStripeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrialEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrialStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionError) validateApplicationFeePercent(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicationFeePercent) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateCancelAtPeriodEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.CancelAtPeriodEnd) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateCanceledAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CanceledAt) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateCurrentPeriodEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentPeriodEnd) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateCurrentPeriodStart(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentPeriodStart) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateEndedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EndedAt) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateLivemode(formats strfmt.Registry) error {

	if swag.IsZero(m.Livemode) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateQuantity(formats strfmt.Registry) error {

	if swag.IsZero(m.Quantity) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateStripeID(formats strfmt.Registry) error {

	if swag.IsZero(m.StripeID) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateTrialEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.TrialEnd) { // not required
		return nil
	}

	return nil
}

func (m *SubscriptionError) validateTrialStart(formats strfmt.Registry) error {

	if swag.IsZero(m.TrialStart) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionError) UnmarshalBinary(b []byte) error {
	var res SubscriptionError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}