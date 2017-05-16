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

// TriggersPartialUpdateReader is a Reader for the TriggersPartialUpdate structure.
type TriggersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTriggersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTriggersPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTriggersPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTriggersPartialUpdateOK creates a TriggersPartialUpdateOK with default headers values
func NewTriggersPartialUpdateOK() *TriggersPartialUpdateOK {
	return &TriggersPartialUpdateOK{}
}

/*TriggersPartialUpdateOK handles this case with default header values.

Trigger updated
*/
type TriggersPartialUpdateOK struct {
	Payload *models.Trigger
}

func (o *TriggersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/triggers/{id}/][%d] triggersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TriggersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Trigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggersPartialUpdateBadRequest creates a TriggersPartialUpdateBadRequest with default headers values
func NewTriggersPartialUpdateBadRequest() *TriggersPartialUpdateBadRequest {
	return &TriggersPartialUpdateBadRequest{}
}

/*TriggersPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type TriggersPartialUpdateBadRequest struct {
	Payload TriggersPartialUpdateBadRequestBody
}

func (o *TriggersPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/triggers/{id}/][%d] triggersPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TriggersPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggersPartialUpdateNotFound creates a TriggersPartialUpdateNotFound with default headers values
func NewTriggersPartialUpdateNotFound() *TriggersPartialUpdateNotFound {
	return &TriggersPartialUpdateNotFound{}
}

/*TriggersPartialUpdateNotFound handles this case with default header values.

Trigger not found
*/
type TriggersPartialUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *TriggersPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/triggers/{id}/][%d] triggersPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *TriggersPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TriggersPartialUpdateBadRequestBody triggers partial update bad request body
swagger:model TriggersPartialUpdateBadRequestBody
*/
type TriggersPartialUpdateBadRequestBody struct {

	// cause
	// Required: true
	Cause *TriggersPartialUpdateBadRequestBodyCause `json:"cause"`

	// effect
	// Required: true
	Effect *TriggersPartialUpdateBadRequestBodyEffect `json:"effect"`

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
	Webhook *TriggersPartialUpdateBadRequestBodyWebhook `json:"webhook"`
}

// Validate validates this triggers partial update bad request body
func (o *TriggersPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *TriggersPartialUpdateBadRequestBody) validateCause(formats strfmt.Registry) error {

	if err := validate.Required("triggersPartialUpdateBadRequest"+"."+"cause", "body", o.Cause); err != nil {
		return err
	}

	if o.Cause != nil {

		if err := o.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggersPartialUpdateBadRequest" + "." + "cause")
			}
			return err
		}
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBody) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("triggersPartialUpdateBadRequest"+"."+"effect", "body", o.Effect); err != nil {
		return err
	}

	if o.Effect != nil {

		if err := o.Effect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggersPartialUpdateBadRequest" + "." + "effect")
			}
			return err
		}
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("triggersPartialUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("triggersPartialUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBody) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("triggersPartialUpdateBadRequest"+"."+"schedule", "body", o.Schedule); err != nil {
		return err
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBody) validateWebhook(formats strfmt.Registry) error {

	if err := validate.Required("triggersPartialUpdateBadRequest"+"."+"webhook", "body", o.Webhook); err != nil {
		return err
	}

	if o.Webhook != nil {

		if err := o.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggersPartialUpdateBadRequest" + "." + "webhook")
			}
			return err
		}
	}

	return nil
}

/*TriggersPartialUpdateBadRequestBodyCause triggers partial update bad request body cause
swagger:model TriggersPartialUpdateBadRequestBodyCause
*/
type TriggersPartialUpdateBadRequestBodyCause struct {

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

// Validate validates this triggers partial update bad request body cause
func (o *TriggersPartialUpdateBadRequestBodyCause) Validate(formats strfmt.Registry) error {
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

func (o *TriggersPartialUpdateBadRequestBodyCause) validateActionName(formats strfmt.Registry) error {

	if swag.IsZero(o.ActionName) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyCause) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyCause) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(o.Method) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyCause) validateModel(formats strfmt.Registry) error {

	if swag.IsZero(o.Model) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyCause) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyCause) validateObjectID(formats strfmt.Registry) error {

	if swag.IsZero(o.ObjectID) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyCause) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(o.Payload) { // not required
		return nil
	}

	return nil
}

/*TriggersPartialUpdateBadRequestBodyEffect triggers partial update bad request body effect
swagger:model TriggersPartialUpdateBadRequestBodyEffect
*/
type TriggersPartialUpdateBadRequestBodyEffect struct {

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

// Validate validates this triggers partial update bad request body effect
func (o *TriggersPartialUpdateBadRequestBodyEffect) Validate(formats strfmt.Registry) error {
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

func (o *TriggersPartialUpdateBadRequestBodyEffect) validateActionName(formats strfmt.Registry) error {

	if swag.IsZero(o.ActionName) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyEffect) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyEffect) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(o.Method) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyEffect) validateModel(formats strfmt.Registry) error {

	if swag.IsZero(o.Model) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyEffect) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyEffect) validateObjectID(formats strfmt.Registry) error {

	if swag.IsZero(o.ObjectID) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyEffect) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(o.Payload) { // not required
		return nil
	}

	return nil
}

/*TriggersPartialUpdateBadRequestBodyWebhook triggers partial update bad request body webhook
swagger:model TriggersPartialUpdateBadRequestBodyWebhook
*/
type TriggersPartialUpdateBadRequestBodyWebhook struct {

	// config field errors
	Config []string `json:"config"`

	// Errors not connected to any field
	NonFieldErrors []string `json:"non_field_errors"`

	// url field errors
	URL []string `json:"url"`
}

// Validate validates this triggers partial update bad request body webhook
func (o *TriggersPartialUpdateBadRequestBodyWebhook) Validate(formats strfmt.Registry) error {
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

func (o *TriggersPartialUpdateBadRequestBodyWebhook) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.Config) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyWebhook) validateNonFieldErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.NonFieldErrors) { // not required
		return nil
	}

	return nil
}

func (o *TriggersPartialUpdateBadRequestBodyWebhook) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(o.URL) { // not required
		return nil
	}

	return nil
}

/*TriggersPartialUpdateBody triggers partial update body
swagger:model TriggersPartialUpdateBody
*/
type TriggersPartialUpdateBody struct {

	// cause
	Cause *models.TriggerAction `json:"cause,omitempty"`

	// effect
	Effect *models.TriggerAction `json:"effect,omitempty"`

	// Cron schedule
	Schedule string `json:"schedule,omitempty"`

	// webhook
	Webhook *models.Webhook `json:"webhook,omitempty"`
}
