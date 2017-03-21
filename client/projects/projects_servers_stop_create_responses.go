package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsServersStopCreateReader is a Reader for the ProjectsServersStopCreate structure.
type ProjectsServersStopCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersStopCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersStopCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersStopCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersStopCreateCreated creates a ProjectsServersStopCreateCreated with default headers values
func NewProjectsServersStopCreateCreated() *ProjectsServersStopCreateCreated {
	return &ProjectsServersStopCreateCreated{}
}

/*ProjectsServersStopCreateCreated handles this case with default header values.

Stop
*/
type ProjectsServersStopCreateCreated struct {
}

func (o *ProjectsServersStopCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/stop/][%d] projectsServersStopCreateCreated ", 201)
}

func (o *ProjectsServersStopCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsServersStopCreateBadRequest creates a ProjectsServersStopCreateBadRequest with default headers values
func NewProjectsServersStopCreateBadRequest() *ProjectsServersStopCreateBadRequest {
	return &ProjectsServersStopCreateBadRequest{}
}

/*ProjectsServersStopCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersStopCreateBadRequest struct {
}

func (o *ProjectsServersStopCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/stop/][%d] projectsServersStopCreateBadRequest ", 400)
}

func (o *ProjectsServersStopCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
