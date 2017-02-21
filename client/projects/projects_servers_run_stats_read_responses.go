package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkp85/go-sdk/models"
)

// ProjectsServersRunStatsReadReader is a Reader for the ProjectsServersRunStatsRead structure.
type ProjectsServersRunStatsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersRunStatsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersRunStatsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsServersRunStatsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersRunStatsReadOK creates a ProjectsServersRunStatsReadOK with default headers values
func NewProjectsServersRunStatsReadOK() *ProjectsServersRunStatsReadOK {
	return &ProjectsServersRunStatsReadOK{}
}

/*ProjectsServersRunStatsReadOK handles this case with default header values.

ServerRunStatistics retrieved
*/
type ProjectsServersRunStatsReadOK struct {
	Payload *models.ServerRunStatistics
}

func (o *ProjectsServersRunStatsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/run-stats/{id}/][%d] projectsServersRunStatsReadOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersRunStatsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRunStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersRunStatsReadNotFound creates a ProjectsServersRunStatsReadNotFound with default headers values
func NewProjectsServersRunStatsReadNotFound() *ProjectsServersRunStatsReadNotFound {
	return &ProjectsServersRunStatsReadNotFound{}
}

/*ProjectsServersRunStatsReadNotFound handles this case with default header values.

ServerRunStatistics not found
*/
type ProjectsServersRunStatsReadNotFound struct {
}

func (o *ProjectsServersRunStatsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/run-stats/{id}/][%d] projectsServersRunStatsReadNotFound ", 404)
}

func (o *ProjectsServersRunStatsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
