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

// ProjectsServersRunStatsDeleteReader is a Reader for the ProjectsServersRunStatsDelete structure.
type ProjectsServersRunStatsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersRunStatsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsServersRunStatsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsServersRunStatsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersRunStatsDeleteNoContent creates a ProjectsServersRunStatsDeleteNoContent with default headers values
func NewProjectsServersRunStatsDeleteNoContent() *ProjectsServersRunStatsDeleteNoContent {
	return &ProjectsServersRunStatsDeleteNoContent{}
}

/*ProjectsServersRunStatsDeleteNoContent handles this case with default header values.

ServerRunStatistics deleted
*/
type ProjectsServersRunStatsDeleteNoContent struct {
}

func (o *ProjectsServersRunStatsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/run-stats/{id}/][%d] projectsServersRunStatsDeleteNoContent ", 204)
}

func (o *ProjectsServersRunStatsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsServersRunStatsDeleteNotFound creates a ProjectsServersRunStatsDeleteNotFound with default headers values
func NewProjectsServersRunStatsDeleteNotFound() *ProjectsServersRunStatsDeleteNotFound {
	return &ProjectsServersRunStatsDeleteNotFound{}
}

/*ProjectsServersRunStatsDeleteNotFound handles this case with default header values.

ServerRunStatistics not found
*/
type ProjectsServersRunStatsDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsServersRunStatsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/run-stats/{id}/][%d] projectsServersRunStatsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsServersRunStatsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
