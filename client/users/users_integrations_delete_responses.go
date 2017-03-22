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

// UsersIntegrationsDeleteReader is a Reader for the UsersIntegrationsDelete structure.
type UsersIntegrationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersIntegrationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUsersIntegrationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUsersIntegrationsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersIntegrationsDeleteNoContent creates a UsersIntegrationsDeleteNoContent with default headers values
func NewUsersIntegrationsDeleteNoContent() *UsersIntegrationsDeleteNoContent {
	return &UsersIntegrationsDeleteNoContent{}
}

/*UsersIntegrationsDeleteNoContent handles this case with default header values.

Integration deleted
*/
type UsersIntegrationsDeleteNoContent struct {
}

func (o *UsersIntegrationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsDeleteNoContent ", 204)
}

func (o *UsersIntegrationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersIntegrationsDeleteNotFound creates a UsersIntegrationsDeleteNotFound with default headers values
func NewUsersIntegrationsDeleteNotFound() *UsersIntegrationsDeleteNotFound {
	return &UsersIntegrationsDeleteNotFound{}
}

/*UsersIntegrationsDeleteNotFound handles this case with default header values.

Integration not found
*/
type UsersIntegrationsDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *UsersIntegrationsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/users/{user_pk}/integrations/{id}/][%d] usersIntegrationsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *UsersIntegrationsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
