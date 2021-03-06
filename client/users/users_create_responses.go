package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// UsersCreateReader is a Reader for the UsersCreate structure.
type UsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUsersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUsersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersCreateCreated creates a UsersCreateCreated with default headers values
func NewUsersCreateCreated() *UsersCreateCreated {
	return &UsersCreateCreated{}
}

/*UsersCreateCreated handles this case with default header values.

User created
*/
type UsersCreateCreated struct {
	Payload *models.User
}

func (o *UsersCreateCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/users/profiles/][%d] usersCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersCreateBadRequest creates a UsersCreateBadRequest with default headers values
func NewUsersCreateBadRequest() *UsersCreateBadRequest {
	return &UsersCreateBadRequest{}
}

/*UsersCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type UsersCreateBadRequest struct {
	Payload *models.UserError
}

func (o *UsersCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/users/profiles/][%d] usersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
