// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// UserAvatarUpdateReader is a Reader for the UserAvatarUpdate structure.
type UserAvatarUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAvatarUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserAvatarUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserAvatarUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserAvatarUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserAvatarUpdateOK creates a UserAvatarUpdateOK with default headers values
func NewUserAvatarUpdateOK() *UserAvatarUpdateOK {
	return &UserAvatarUpdateOK{}
}

/*UserAvatarUpdateOK handles this case with default header values.

Avatar updated.
*/
type UserAvatarUpdateOK struct {
	Payload *models.User
}

func (o *UserAvatarUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/{user}/avatar/][%d] userAvatarUpdateOK  %+v", 200, o.Payload)
}

func (o *UserAvatarUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAvatarUpdateBadRequest creates a UserAvatarUpdateBadRequest with default headers values
func NewUserAvatarUpdateBadRequest() *UserAvatarUpdateBadRequest {
	return &UserAvatarUpdateBadRequest{}
}

/*UserAvatarUpdateBadRequest handles this case with default header values.

Invalid data supplied.
*/
type UserAvatarUpdateBadRequest struct {
	Payload *models.UserError
}

func (o *UserAvatarUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/{user}/avatar/][%d] userAvatarUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UserAvatarUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAvatarUpdateNotFound creates a UserAvatarUpdateNotFound with default headers values
func NewUserAvatarUpdateNotFound() *UserAvatarUpdateNotFound {
	return &UserAvatarUpdateNotFound{}
}

/*UserAvatarUpdateNotFound handles this case with default header values.

Avatar not found.
*/
type UserAvatarUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *UserAvatarUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/{user}/avatar/][%d] userAvatarUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UserAvatarUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
