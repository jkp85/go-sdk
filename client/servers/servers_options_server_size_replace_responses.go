// Code generated by go-swagger; DO NOT EDIT.

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

// ServersOptionsServerSizeReplaceReader is a Reader for the ServersOptionsServerSizeReplace structure.
type ServersOptionsServerSizeReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersOptionsServerSizeReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServersOptionsServerSizeReplaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServersOptionsServerSizeReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewServersOptionsServerSizeReplaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServersOptionsServerSizeReplaceOK creates a ServersOptionsServerSizeReplaceOK with default headers values
func NewServersOptionsServerSizeReplaceOK() *ServersOptionsServerSizeReplaceOK {
	return &ServersOptionsServerSizeReplaceOK{}
}

/*ServersOptionsServerSizeReplaceOK handles this case with default header values.

Server size replaced. This operation is available only to super users.
*/
type ServersOptionsServerSizeReplaceOK struct {
	Payload *models.ServerSize
}

func (o *ServersOptionsServerSizeReplaceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/servers/options/server-size/{size}/][%d] serversOptionsServerSizeReplaceOK  %+v", 200, o.Payload)
}

func (o *ServersOptionsServerSizeReplaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerSize)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsServerSizeReplaceBadRequest creates a ServersOptionsServerSizeReplaceBadRequest with default headers values
func NewServersOptionsServerSizeReplaceBadRequest() *ServersOptionsServerSizeReplaceBadRequest {
	return &ServersOptionsServerSizeReplaceBadRequest{}
}

/*ServersOptionsServerSizeReplaceBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ServersOptionsServerSizeReplaceBadRequest struct {
	Payload *models.ServerSizeError
}

func (o *ServersOptionsServerSizeReplaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/servers/options/server-size/{size}/][%d] serversOptionsServerSizeReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *ServersOptionsServerSizeReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerSizeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersOptionsServerSizeReplaceNotFound creates a ServersOptionsServerSizeReplaceNotFound with default headers values
func NewServersOptionsServerSizeReplaceNotFound() *ServersOptionsServerSizeReplaceNotFound {
	return &ServersOptionsServerSizeReplaceNotFound{}
}

/*ServersOptionsServerSizeReplaceNotFound handles this case with default header values.

ServerSize not found
*/
type ServersOptionsServerSizeReplaceNotFound struct {
	Payload *models.NotFound
}

func (o *ServersOptionsServerSizeReplaceNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/servers/options/server-size/{size}/][%d] serversOptionsServerSizeReplaceNotFound  %+v", 404, o.Payload)
}

func (o *ServersOptionsServerSizeReplaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
