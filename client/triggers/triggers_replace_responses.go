// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// TriggersReplaceReader is a Reader for the TriggersReplace structure.
type TriggersReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggersReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTriggersReplaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTriggersReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTriggersReplaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTriggersReplaceOK creates a TriggersReplaceOK with default headers values
func NewTriggersReplaceOK() *TriggersReplaceOK {
	return &TriggersReplaceOK{}
}

/*TriggersReplaceOK handles this case with default header values.

Trigger updated.
*/
type TriggersReplaceOK struct {
	Payload *models.Trigger
}

func (o *TriggersReplaceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/triggers/{id}/][%d] triggersReplaceOK  %+v", 200, o.Payload)
}

func (o *TriggersReplaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Trigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggersReplaceBadRequest creates a TriggersReplaceBadRequest with default headers values
func NewTriggersReplaceBadRequest() *TriggersReplaceBadRequest {
	return &TriggersReplaceBadRequest{}
}

/*TriggersReplaceBadRequest handles this case with default header values.

Invalid data supplied.
*/
type TriggersReplaceBadRequest struct {
	Payload *models.TriggerError
}

func (o *TriggersReplaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/triggers/{id}/][%d] triggersReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *TriggersReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TriggerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggersReplaceNotFound creates a TriggersReplaceNotFound with default headers values
func NewTriggersReplaceNotFound() *TriggersReplaceNotFound {
	return &TriggersReplaceNotFound{}
}

/*TriggersReplaceNotFound handles this case with default header values.

Trigger not found
*/
type TriggersReplaceNotFound struct {
	Payload *models.NotFound
}

func (o *TriggersReplaceNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/triggers/{id}/][%d] triggersReplaceNotFound  %+v", 404, o.Payload)
}

func (o *TriggersReplaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}