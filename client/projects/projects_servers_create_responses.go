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

// ProjectsServersCreateReader is a Reader for the ProjectsServersCreate structure.
type ProjectsServersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersCreateCreated creates a ProjectsServersCreateCreated with default headers values
func NewProjectsServersCreateCreated() *ProjectsServersCreateCreated {
	return &ProjectsServersCreateCreated{}
}

/*ProjectsServersCreateCreated handles this case with default header values.

Server created
*/
type ProjectsServersCreateCreated struct {
	Payload *models.Server
}

func (o *ProjectsServersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/][%d] projectsServersCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsServersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersCreateBadRequest creates a ProjectsServersCreateBadRequest with default headers values
func NewProjectsServersCreateBadRequest() *ProjectsServersCreateBadRequest {
	return &ProjectsServersCreateBadRequest{}
}

/*ProjectsServersCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersCreateBadRequest struct {
}

func (o *ProjectsServersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/][%d] projectsServersCreateBadRequest ", 400)
}

func (o *ProjectsServersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsServersCreateBody projects servers create body
swagger:model ProjectsServersCreateBody
*/
type ProjectsServersCreateBody struct {

	// config
	Config interface{} `json:"config,omitempty"`

	// connected
	// Required: true
	Connected []string `json:"connected"`

	// environment resources
	// Required: true
	EnvironmentResources *string `json:"environment_resources"`

	// environment type
	// Required: true
	EnvironmentType *string `json:"environment_type"`

	// name
	// Required: true
	Name *string `json:"name"`

	// startup script
	StartupScript string `json:"startup_script,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}