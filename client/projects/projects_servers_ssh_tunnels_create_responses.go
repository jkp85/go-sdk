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

// ProjectsServersSSHTunnelsCreateReader is a Reader for the ProjectsServersSSHTunnelsCreate structure.
type ProjectsServersSSHTunnelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersSSHTunnelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersSSHTunnelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersSSHTunnelsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersSSHTunnelsCreateCreated creates a ProjectsServersSSHTunnelsCreateCreated with default headers values
func NewProjectsServersSSHTunnelsCreateCreated() *ProjectsServersSSHTunnelsCreateCreated {
	return &ProjectsServersSSHTunnelsCreateCreated{}
}

/*ProjectsServersSSHTunnelsCreateCreated handles this case with default header values.

SshTunnel created
*/
type ProjectsServersSSHTunnelsCreateCreated struct {
	Payload *models.SSHTunnel
}

func (o *ProjectsServersSSHTunnelsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/ssh-tunnels/][%d] projectsServersSshTunnelsCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsServersSSHTunnelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHTunnel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersSSHTunnelsCreateBadRequest creates a ProjectsServersSSHTunnelsCreateBadRequest with default headers values
func NewProjectsServersSSHTunnelsCreateBadRequest() *ProjectsServersSSHTunnelsCreateBadRequest {
	return &ProjectsServersSSHTunnelsCreateBadRequest{}
}

/*ProjectsServersSSHTunnelsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersSSHTunnelsCreateBadRequest struct {
}

func (o *ProjectsServersSSHTunnelsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/ssh-tunnels/][%d] projectsServersSshTunnelsCreateBadRequest ", 400)
}

func (o *ProjectsServersSSHTunnelsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsServersSSHTunnelsCreateBody projects servers SSH tunnels create body
swagger:model ProjectsServersSSHTunnelsCreateBody
*/
type ProjectsServersSSHTunnelsCreateBody struct {

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// host
	// Required: true
	Host *string `json:"host"`

	// local port
	// Required: true
	LocalPort *int64 `json:"local_port"`

	// name
	// Required: true
	Name *string `json:"name"`

	// remote port
	// Required: true
	RemotePort *int64 `json:"remote_port"`

	// username
	// Required: true
	Username *string `json:"username"`
}