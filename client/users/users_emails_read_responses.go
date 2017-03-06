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

// UsersEmailsReadReader is a Reader for the UsersEmailsRead structure.
type UsersEmailsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersEmailsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersEmailsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUsersEmailsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersEmailsReadOK creates a UsersEmailsReadOK with default headers values
func NewUsersEmailsReadOK() *UsersEmailsReadOK {
	return &UsersEmailsReadOK{}
}

/*UsersEmailsReadOK handles this case with default header values.

Email retrieved
*/
type UsersEmailsReadOK struct {
	Payload *models.Email
}

func (o *UsersEmailsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/users/{user_pk}/emails/{address}/][%d] usersEmailsReadOK  %+v", 200, o.Payload)
}

func (o *UsersEmailsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Email)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersEmailsReadNotFound creates a UsersEmailsReadNotFound with default headers values
func NewUsersEmailsReadNotFound() *UsersEmailsReadNotFound {
	return &UsersEmailsReadNotFound{}
}

/*UsersEmailsReadNotFound handles this case with default header values.

Email not found
*/
type UsersEmailsReadNotFound struct {
}

func (o *UsersEmailsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/users/{user_pk}/emails/{address}/][%d] usersEmailsReadNotFound ", 404)
}

func (o *UsersEmailsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}