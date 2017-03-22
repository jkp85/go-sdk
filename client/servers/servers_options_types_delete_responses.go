package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ServersOptionsTypesDeleteReader is a Reader for the ServersOptionsTypesDelete structure.
type ServersOptionsTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewServersOptionsTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewServersOptionsTypesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsTypesDeleteNoContent creates a ServersOptionsTypesDeleteNoContent with default headers values
func NewServersOptionsTypesDeleteNoContent() *ServersOptionsTypesDeleteNoContent {
	return &ServersOptionsTypesDeleteNoContent{}
}

/*ServersOptionsTypesDeleteNoContent handles this case with default header values.

EnvironmentType deleted
*/
type ServersOptionsTypesDeleteNoContent struct {
}

func (o *ServersOptionsTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/servers/options/types/{id}/][%d] serversOptionsTypesDeleteNoContent ", 204)
}

func (o *ServersOptionsTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServersOptionsTypesDeleteNotFound creates a ServersOptionsTypesDeleteNotFound with default headers values
func NewServersOptionsTypesDeleteNotFound() *ServersOptionsTypesDeleteNotFound {
	return &ServersOptionsTypesDeleteNotFound{}
}

/*ServersOptionsTypesDeleteNotFound handles this case with default header values.

EnvironmentType not found
*/
type ServersOptionsTypesDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *ServersOptionsTypesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/servers/options/types/{id}/][%d] serversOptionsTypesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ServersOptionsTypesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
