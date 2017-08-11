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

// ServiceTriggerUpdateReader is a Reader for the ServiceTriggerUpdate structure.
type ServiceTriggerUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceTriggerUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceTriggerUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceTriggerUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewServiceTriggerUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceTriggerUpdateOK creates a ServiceTriggerUpdateOK with default headers values
func NewServiceTriggerUpdateOK() *ServiceTriggerUpdateOK {
	return &ServiceTriggerUpdateOK{}
}

/*ServiceTriggerUpdateOK handles this case with default header values.

Server action updated.
*/
type ServiceTriggerUpdateOK struct {
	Payload *models.ServerAction
}

func (o *ServiceTriggerUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/servers/{server_id}/triggers/{id}/][%d] serviceTriggerUpdateOK  %+v", 200, o.Payload)
}

func (o *ServiceTriggerUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerAction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceTriggerUpdateBadRequest creates a ServiceTriggerUpdateBadRequest with default headers values
func NewServiceTriggerUpdateBadRequest() *ServiceTriggerUpdateBadRequest {
	return &ServiceTriggerUpdateBadRequest{}
}

/*ServiceTriggerUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ServiceTriggerUpdateBadRequest struct {
	Payload *models.ServerActionError
}

func (o *ServiceTriggerUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/servers/{server_id}/triggers/{id}/][%d] serviceTriggerUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceTriggerUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerActionError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceTriggerUpdateNotFound creates a ServiceTriggerUpdateNotFound with default headers values
func NewServiceTriggerUpdateNotFound() *ServiceTriggerUpdateNotFound {
	return &ServiceTriggerUpdateNotFound{}
}

/*ServiceTriggerUpdateNotFound handles this case with default header values.

Server action not found.
*/
type ServiceTriggerUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ServiceTriggerUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project_id}/servers/{server_id}/triggers/{id}/][%d] serviceTriggerUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ServiceTriggerUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
