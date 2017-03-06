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

// ProjectsServersTerminateReader is a Reader for the ProjectsServersTerminate structure.
type ProjectsServersTerminateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersTerminateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersTerminateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersTerminateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersTerminateCreated creates a ProjectsServersTerminateCreated with default headers values
func NewProjectsServersTerminateCreated() *ProjectsServersTerminateCreated {
	return &ProjectsServersTerminateCreated{}
}

/*ProjectsServersTerminateCreated handles this case with default header values.

Server created
*/
type ProjectsServersTerminateCreated struct {
	Payload *models.Server
}

func (o *ProjectsServersTerminateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/{id}/terminate/][%d] projectsServersTerminateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsServersTerminateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersTerminateBadRequest creates a ProjectsServersTerminateBadRequest with default headers values
func NewProjectsServersTerminateBadRequest() *ProjectsServersTerminateBadRequest {
	return &ProjectsServersTerminateBadRequest{}
}

/*ProjectsServersTerminateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersTerminateBadRequest struct {
}

func (o *ProjectsServersTerminateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/{id}/terminate/][%d] projectsServersTerminateBadRequest ", 400)
}

func (o *ProjectsServersTerminateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsServersTerminateBody projects servers terminate body
swagger:model ProjectsServersTerminateBody
*/
type ProjectsServersTerminateBody struct {

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