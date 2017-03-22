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

// ProjectsServersDeleteReader is a Reader for the ProjectsServersDelete structure.
type ProjectsServersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsServersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsServersDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersDeleteNoContent creates a ProjectsServersDeleteNoContent with default headers values
func NewProjectsServersDeleteNoContent() *ProjectsServersDeleteNoContent {
	return &ProjectsServersDeleteNoContent{}
}

/*ProjectsServersDeleteNoContent handles this case with default header values.

Server deleted
*/
type ProjectsServersDeleteNoContent struct {
}

func (o *ProjectsServersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/servers/{id}/][%d] projectsServersDeleteNoContent ", 204)
}

func (o *ProjectsServersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsServersDeleteNotFound creates a ProjectsServersDeleteNotFound with default headers values
func NewProjectsServersDeleteNotFound() *ProjectsServersDeleteNotFound {
	return &ProjectsServersDeleteNotFound{}
}

/*ProjectsServersDeleteNotFound handles this case with default header values.

Server not found
*/
type ProjectsServersDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsServersDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/servers/{id}/][%d] projectsServersDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsServersDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
