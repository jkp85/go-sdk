package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// TriggersUpdateReader is a Reader for the TriggersUpdate structure.
type TriggersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTriggersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTriggersUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTriggersUpdateOK creates a TriggersUpdateOK with default headers values
func NewTriggersUpdateOK() *TriggersUpdateOK {
	return &TriggersUpdateOK{}
}

/*TriggersUpdateOK handles this case with default header values.

Trigger updated
*/
type TriggersUpdateOK struct {
	Payload *models.Trigger
}

func (o *TriggersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/triggers/{id}/][%d] triggersUpdateOK  %+v", 200, o.Payload)
}

func (o *TriggersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Trigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggersUpdateBadRequest creates a TriggersUpdateBadRequest with default headers values
func NewTriggersUpdateBadRequest() *TriggersUpdateBadRequest {
	return &TriggersUpdateBadRequest{}
}

/*TriggersUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type TriggersUpdateBadRequest struct {
	Payload TriggersUpdateBadRequestBody
}

func (o *TriggersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/triggers/{id}/][%d] triggersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TriggersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TriggersUpdateBadRequestBody triggers update bad request body
swagger:model TriggersUpdateBadRequestBody
*/
type TriggersUpdateBadRequestBody struct {

	// cause
	// Required: true
	Cause *TriggersUpdateBadRequestBodyCause `json:"cause"`

	// effect
	// Required: true
	Effect *TriggersUpdateBadRequestBodyEffect `json:"effect"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// schedule field errors
	// Required: true
	Schedule []string `json:"schedule"`

	// webhook
	// Required: true
	Webhook *TriggersUpdateBadRequestBodyWebhook `json:"webhook"`
}

// Validate validates this triggers update bad request body
func (o *TriggersUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCause(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEffect(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSchedule(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateWebhook(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TriggersUpdateBadRequestBody) validateCause(formats strfmt.Registry) error {

	if err := validate.Required("triggersUpdateBadRequest"+"."+"cause", "body", o.Cause); err != nil {
		return err
	}

	if o.Cause != nil {

		if err := o.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggersUpdateBadRequest" + "." + "cause")
			}
			return err
		}
	}

	return nil
}

func (o *TriggersUpdateBadRequestBody) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("triggersUpdateBadRequest"+"."+"effect", "body", o.Effect); err != nil {
		return err
	}

	if o.Effect != nil {

		if err := o.Effect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggersUpdateBadRequest" + "." + "effect")
			}
			return err
		}
	}

	return nil
}

func (o *TriggersUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("triggersUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *TriggersUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("triggersUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *TriggersUpdateBadRequestBody) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("triggersUpdateBadRequest"+"."+"schedule", "body", o.Schedule); err != nil {
		return err
	}

	return nil
}

func (o *TriggersUpdateBadRequestBody) validateWebhook(formats strfmt.Registry) error {

	if err := validate.Required("triggersUpdateBadRequest"+"."+"webhook", "body", o.Webhook); err != nil {
		return err
	}

	if o.Webhook != nil {

		if err := o.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggersUpdateBadRequest" + "." + "webhook")
			}
			return err
		}
	}

	return nil
}

/*TriggersUpdateBadRequestBodyCause triggers update bad request body cause
swagger:model TriggersUpdateBadRequestBodyCause
*/
type TriggersUpdateBadRequestBodyCause struct {

	// action_name field errors
	ActionName []string `json:"action_name"`

	// id field errors
	ID []string `json:"id"`

	// method field errors
	Method []string `json:"method"`

	// model field errors
	Model []string `json:"model"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// object_id field errors
	ObjectID []string `json:"object_id"`

	// payload field errors
	Payload []string `json:"payload"`
}

// Validate validates this triggers update bad request body cause
func (o *TriggersUpdateBadRequestBodyCause) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActionName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateModel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateObjectID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePayload(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TriggersUpdateBadRequestBodyCause) validateActionName(formats strfmt.Registry) error {

	if swag.IsZero(o.ActionName) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyCause) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyCause) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(o.Method) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyCause) validateModel(formats strfmt.Registry) error {

	if swag.IsZero(o.Model) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyCause) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyCause) validateObjectID(formats strfmt.Registry) error {

	if swag.IsZero(o.ObjectID) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyCause) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(o.Payload) { // not required
		return nil
	}

	return nil
}

/*TriggersUpdateBadRequestBodyEffect triggers update bad request body effect
swagger:model TriggersUpdateBadRequestBodyEffect
*/
type TriggersUpdateBadRequestBodyEffect struct {

	// action_name field errors
	ActionName []string `json:"action_name"`

	// id field errors
	ID []string `json:"id"`

	// method field errors
	Method []string `json:"method"`

	// model field errors
	Model []string `json:"model"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// object_id field errors
	ObjectID []string `json:"object_id"`

	// payload field errors
	Payload []string `json:"payload"`
}

// Validate validates this triggers update bad request body effect
func (o *TriggersUpdateBadRequestBodyEffect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActionName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateModel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateObjectID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePayload(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TriggersUpdateBadRequestBodyEffect) validateActionName(formats strfmt.Registry) error {

	if swag.IsZero(o.ActionName) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyEffect) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyEffect) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(o.Method) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyEffect) validateModel(formats strfmt.Registry) error {

	if swag.IsZero(o.Model) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyEffect) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyEffect) validateObjectID(formats strfmt.Registry) error {

	if swag.IsZero(o.ObjectID) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyEffect) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(o.Payload) { // not required
		return nil
	}

	return nil
}

/*TriggersUpdateBadRequestBodyWebhook triggers update bad request body webhook
swagger:model TriggersUpdateBadRequestBodyWebhook
*/
type TriggersUpdateBadRequestBodyWebhook struct {

	// config field errors
	Config []string `json:"config"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// url field errors
	URL []string `json:"url"`
}

// Validate validates this triggers update bad request body webhook
func (o *TriggersUpdateBadRequestBodyWebhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TriggersUpdateBadRequestBodyWebhook) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.Config) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyWebhook) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (o *TriggersUpdateBadRequestBodyWebhook) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(o.URL) { // not required
		return nil
	}

	return nil
}

/*TriggersUpdateBody triggers update body
swagger:model TriggersUpdateBody
*/
type TriggersUpdateBody struct {

	// cause
	Cause *models.TriggerAction `json:"cause,omitempty"`

	// effect
	Effect *models.TriggerAction `json:"effect,omitempty"`

	// Cron schedule
	Schedule string `json:"schedule,omitempty"`

	// webhook
	Webhook *models.Webhook `json:"webhook,omitempty"`
}
