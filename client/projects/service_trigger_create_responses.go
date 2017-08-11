// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ServiceTriggerCreateReader is a Reader for the ServiceTriggerCreate structure.
type ServiceTriggerCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceTriggerCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewServiceTriggerCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceTriggerCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceTriggerCreateCreated creates a ServiceTriggerCreateCreated with default headers values
func NewServiceTriggerCreateCreated() *ServiceTriggerCreateCreated {
	return &ServiceTriggerCreateCreated{}
}

/*ServiceTriggerCreateCreated handles this case with default header values.

Server action created.
*/
type ServiceTriggerCreateCreated struct {
	Payload *models.ServerAction
}

func (o *ServiceTriggerCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/projects/{project_id}/servers/{server_id}/triggers/][%d] serviceTriggerCreateCreated  %+v", 201, o.Payload)
}

func (o *ServiceTriggerCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerAction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceTriggerCreateBadRequest creates a ServiceTriggerCreateBadRequest with default headers values
func NewServiceTriggerCreateBadRequest() *ServiceTriggerCreateBadRequest {
	return &ServiceTriggerCreateBadRequest{}
}

/*ServiceTriggerCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ServiceTriggerCreateBadRequest struct {
	Payload *models.ServerActionError
}

func (o *ServiceTriggerCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/projects/{project_id}/servers/{server_id}/triggers/][%d] serviceTriggerCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceTriggerCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerActionError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
