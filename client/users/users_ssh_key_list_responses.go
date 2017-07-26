// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UsersSSHKeyListReader is a Reader for the UsersSSHKeyList structure.
type UsersSSHKeyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersSSHKeyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUsersSSHKeyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersSSHKeyListOK creates a UsersSSHKeyListOK with default headers values
func NewUsersSSHKeyListOK() *UsersSSHKeyListOK {
	return &UsersSSHKeyListOK{}
}

/*UsersSSHKeyListOK handles this case with default header values.

SSH Key.
*/
type UsersSSHKeyListOK struct {
}

func (o *UsersSSHKeyListOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/{user_id}/ssh-key/][%d] usersSshKeyListOK ", 200)
}

func (o *UsersSSHKeyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
