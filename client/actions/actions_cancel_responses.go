// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ActionsCancelReader is a Reader for the ActionsCancel structure.
type ActionsCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewActionsCancelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewActionsCancelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsCancelCreated creates a ActionsCancelCreated with default headers values
func NewActionsCancelCreated() *ActionsCancelCreated {
	return &ActionsCancelCreated{}
}

/*ActionsCancelCreated handles this case with default header values.

Cancel action.
*/
type ActionsCancelCreated struct {
}

func (o *ActionsCancelCreated) Error() string {
	return fmt.Sprintf("[POST /v1/actions/{id}/cancel/][%d] actionsCancelCreated ", 201)
}

func (o *ActionsCancelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsCancelBadRequest creates a ActionsCancelBadRequest with default headers values
func NewActionsCancelBadRequest() *ActionsCancelBadRequest {
	return &ActionsCancelBadRequest{}
}

/*ActionsCancelBadRequest handles this case with default header values.

Invalid action id supplied.
*/
type ActionsCancelBadRequest struct {
}

func (o *ActionsCancelBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/actions/{id}/cancel/][%d] actionsCancelBadRequest ", 400)
}

func (o *ActionsCancelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
