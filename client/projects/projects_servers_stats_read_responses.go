package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// ProjectsServersStatsReadReader is a Reader for the ProjectsServersStatsRead structure.
type ProjectsServersStatsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersStatsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersStatsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsServersStatsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersStatsReadOK creates a ProjectsServersStatsReadOK with default headers values
func NewProjectsServersStatsReadOK() *ProjectsServersStatsReadOK {
	return &ProjectsServersStatsReadOK{}
}

/*ProjectsServersStatsReadOK handles this case with default header values.

ServerStatistics retrieved
*/
type ProjectsServersStatsReadOK struct {
	Payload *models.ServerStatistics
}

func (o *ProjectsServersStatsReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersStatsReadOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/projects/{project}/servers/{server}/stats/{id}/][%d] projectsServersStatsReadOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersStatsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersStatsReadNotFound creates a ProjectsServersStatsReadNotFound with default headers values
func NewProjectsServersStatsReadNotFound() *ProjectsServersStatsReadNotFound {
	return &ProjectsServersStatsReadNotFound{}
}

/*ProjectsServersStatsReadNotFound handles this case with default header values.

ServerStatistics not found
*/
type ProjectsServersStatsReadNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsServersStatsReadNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersStatsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/projects/{project}/servers/{server}/stats/{id}/][%d] projectsServersStatsReadNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsServersStatsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
