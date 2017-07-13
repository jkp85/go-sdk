package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ActionsReadReader is a Reader for the ActionsRead structure.
type ActionsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewActionsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewActionsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsReadOK creates a ActionsReadOK with default headers values
func NewActionsReadOK() *ActionsReadOK {
	return &ActionsReadOK{}
}

/*ActionsReadOK handles this case with default header values.

Action retrieved
*/
type ActionsReadOK struct {
	Payload *models.Action
}

func (o *ActionsReadOK) Error() string {
	return fmt.Sprintf("[GET /actions/{id}/][%d] actionsReadOK  %+v", 200, o.Payload)
}

func (o *ActionsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsReadNotFound creates a ActionsReadNotFound with default headers values
func NewActionsReadNotFound() *ActionsReadNotFound {
	return &ActionsReadNotFound{}
}

/*ActionsReadNotFound handles this case with default header values.

Action not found
*/
type ActionsReadNotFound struct {
	Payload *models.NotFound
}

func (o *ActionsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /actions/{id}/][%d] actionsReadNotFound  %+v", 404, o.Payload)
}

func (o *ActionsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
