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

// ProjectsServersRunStatsUpdateReader is a Reader for the ProjectsServersRunStatsUpdate structure.
type ProjectsServersRunStatsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersRunStatsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersRunStatsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersRunStatsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectsServersRunStatsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersRunStatsUpdateOK creates a ProjectsServersRunStatsUpdateOK with default headers values
func NewProjectsServersRunStatsUpdateOK() *ProjectsServersRunStatsUpdateOK {
	return &ProjectsServersRunStatsUpdateOK{}
}

/*ProjectsServersRunStatsUpdateOK handles this case with default header values.

ServerRunStatistics updated.
*/
type ProjectsServersRunStatsUpdateOK struct {
	Payload *models.ServerRunStatistics
}

func (o *ProjectsServersRunStatsUpdateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersRunStatsUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/servers/{server}/run-stats/{id}/][%d] projectsServersRunStatsUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersRunStatsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRunStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersRunStatsUpdateBadRequest creates a ProjectsServersRunStatsUpdateBadRequest with default headers values
func NewProjectsServersRunStatsUpdateBadRequest() *ProjectsServersRunStatsUpdateBadRequest {
	return &ProjectsServersRunStatsUpdateBadRequest{}
}

/*ProjectsServersRunStatsUpdateBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ProjectsServersRunStatsUpdateBadRequest struct {
	Payload *models.ServerRunStatisticsError
}

func (o *ProjectsServersRunStatsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersRunStatsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/servers/{server}/run-stats/{id}/][%d] projectsServersRunStatsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersRunStatsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRunStatisticsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersRunStatsUpdateNotFound creates a ProjectsServersRunStatsUpdateNotFound with default headers values
func NewProjectsServersRunStatsUpdateNotFound() *ProjectsServersRunStatsUpdateNotFound {
	return &ProjectsServersRunStatsUpdateNotFound{}
}

/*ProjectsServersRunStatsUpdateNotFound handles this case with default header values.

ServerRunStatistics not found.
*/
type ProjectsServersRunStatsUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsServersRunStatsUpdateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsServersRunStatsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/servers/{server}/run-stats/{id}/][%d] projectsServersRunStatsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsServersRunStatsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
