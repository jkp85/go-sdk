package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkp85/go-sdk/models"
)

// AuthJwtTokenAuthCreateReader is a Reader for the AuthJwtTokenAuthCreate structure.
type AuthJwtTokenAuthCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthJwtTokenAuthCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAuthJwtTokenAuthCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAuthJwtTokenAuthCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthJwtTokenAuthCreateCreated creates a AuthJwtTokenAuthCreateCreated with default headers values
func NewAuthJwtTokenAuthCreateCreated() *AuthJwtTokenAuthCreateCreated {
	return &AuthJwtTokenAuthCreateCreated{}
}

/*AuthJwtTokenAuthCreateCreated handles this case with default header values.

JSONWebToken created
*/
type AuthJwtTokenAuthCreateCreated struct {
	Payload *models.JSONWebToken
}

func (o *AuthJwtTokenAuthCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/auth/jwt-token-auth/][%d] authJwtTokenAuthCreateCreated  %+v", 201, o.Payload)
}

func (o *AuthJwtTokenAuthCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONWebToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthJwtTokenAuthCreateBadRequest creates a AuthJwtTokenAuthCreateBadRequest with default headers values
func NewAuthJwtTokenAuthCreateBadRequest() *AuthJwtTokenAuthCreateBadRequest {
	return &AuthJwtTokenAuthCreateBadRequest{}
}

/*AuthJwtTokenAuthCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type AuthJwtTokenAuthCreateBadRequest struct {
	Payload AuthJwtTokenAuthCreateBadRequestBody
}

func (o *AuthJwtTokenAuthCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/auth/jwt-token-auth/][%d] authJwtTokenAuthCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AuthJwtTokenAuthCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AuthJwtTokenAuthCreateBadRequestBody auth jwt token auth create bad request body
swagger:model AuthJwtTokenAuthCreateBadRequestBody
*/
type AuthJwtTokenAuthCreateBadRequestBody struct {

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// password field errors
	// Required: true
	Password []string `json:"password"`

	// username field errors
	// Required: true
	Username []string `json:"username"`
}

// Validate validates this auth jwt token auth create bad request body
func (o *AuthJwtTokenAuthCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AuthJwtTokenAuthCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("authJwtTokenAuthCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *AuthJwtTokenAuthCreateBadRequestBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("authJwtTokenAuthCreateBadRequest"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	return nil
}

func (o *AuthJwtTokenAuthCreateBadRequestBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("authJwtTokenAuthCreateBadRequest"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

/*AuthJwtTokenAuthCreateBody auth jwt token auth create body
swagger:model AuthJwtTokenAuthCreateBody
*/
type AuthJwtTokenAuthCreateBody struct {

	// password
	// Required: true
	Password *string `json:"password"`

	// username
	// Required: true
	Username *string `json:"username"`
}
