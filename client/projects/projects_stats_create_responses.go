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

// ProjectsStatsCreateReader is a Reader for the ProjectsStatsCreate structure.
type ProjectsStatsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsStatsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsStatsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsStatsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsStatsCreateCreated creates a ProjectsStatsCreateCreated with default headers values
func NewProjectsStatsCreateCreated() *ProjectsStatsCreateCreated {
	return &ProjectsStatsCreateCreated{}
}

/*ProjectsStatsCreateCreated handles this case with default header values.

ServerStatistics created
*/
type ProjectsStatsCreateCreated struct {
	Payload *models.ServerStatistics
}

func (o *ProjectsStatsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/stats/][%d] projectsStatsCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsStatsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsStatsCreateBadRequest creates a ProjectsStatsCreateBadRequest with default headers values
func NewProjectsStatsCreateBadRequest() *ProjectsStatsCreateBadRequest {
	return &ProjectsStatsCreateBadRequest{}
}

/*ProjectsStatsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsStatsCreateBadRequest struct {
}

func (o *ProjectsStatsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/stats/][%d] projectsStatsCreateBadRequest ", 400)
}

func (o *ProjectsStatsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsStatsCreateBody projects stats create body
swagger:model ProjectsStatsCreateBody
*/
type ProjectsStatsCreateBody struct {

	// size
	Size int64 `json:"size,omitempty"`

	// start
	Start string `json:"start,omitempty"`

	// stop
	Stop string `json:"stop,omitempty"`
}
