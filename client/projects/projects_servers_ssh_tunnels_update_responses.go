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

// ProjectsServersSSHTunnelsUpdateReader is a Reader for the ProjectsServersSSHTunnelsUpdate structure.
type ProjectsServersSSHTunnelsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersSSHTunnelsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersSSHTunnelsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersSSHTunnelsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectsServersSSHTunnelsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersSSHTunnelsUpdateOK creates a ProjectsServersSSHTunnelsUpdateOK with default headers values
func NewProjectsServersSSHTunnelsUpdateOK() *ProjectsServersSSHTunnelsUpdateOK {
	return &ProjectsServersSSHTunnelsUpdateOK{}
}

/*ProjectsServersSSHTunnelsUpdateOK handles this case with default header values.

SSH Tunnel updated.
*/
type ProjectsServersSSHTunnelsUpdateOK struct {
	Payload *models.SSHTunnel
}

func (o *ProjectsServersSSHTunnelsUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/servers/{server}/ssh-tunnels/{tunnel}/][%d] projectsServersSshTunnelsUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersSSHTunnelsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHTunnel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersSSHTunnelsUpdateBadRequest creates a ProjectsServersSSHTunnelsUpdateBadRequest with default headers values
func NewProjectsServersSSHTunnelsUpdateBadRequest() *ProjectsServersSSHTunnelsUpdateBadRequest {
	return &ProjectsServersSSHTunnelsUpdateBadRequest{}
}

/*ProjectsServersSSHTunnelsUpdateBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ProjectsServersSSHTunnelsUpdateBadRequest struct {
	Payload *models.SSHTunnelError
}

func (o *ProjectsServersSSHTunnelsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/servers/{server}/ssh-tunnels/{tunnel}/][%d] projectsServersSshTunnelsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersSSHTunnelsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHTunnelError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersSSHTunnelsUpdateNotFound creates a ProjectsServersSSHTunnelsUpdateNotFound with default headers values
func NewProjectsServersSSHTunnelsUpdateNotFound() *ProjectsServersSSHTunnelsUpdateNotFound {
	return &ProjectsServersSSHTunnelsUpdateNotFound{}
}

/*ProjectsServersSSHTunnelsUpdateNotFound handles this case with default header values.

SSH tunnel not found.
*/
type ProjectsServersSSHTunnelsUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsServersSSHTunnelsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/servers/{server}/ssh-tunnels/{tunnel}/][%d] projectsServersSshTunnelsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsServersSSHTunnelsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
