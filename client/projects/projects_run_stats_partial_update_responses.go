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

// ProjectsRunStatsPartialUpdateReader is a Reader for the ProjectsRunStatsPartialUpdate structure.
type ProjectsRunStatsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsRunStatsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsRunStatsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsRunStatsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectsRunStatsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsRunStatsPartialUpdateOK creates a ProjectsRunStatsPartialUpdateOK with default headers values
func NewProjectsRunStatsPartialUpdateOK() *ProjectsRunStatsPartialUpdateOK {
	return &ProjectsRunStatsPartialUpdateOK{}
}

/*ProjectsRunStatsPartialUpdateOK handles this case with default header values.

ServerRunStatistics updated
*/
type ProjectsRunStatsPartialUpdateOK struct {
	Payload *models.ServerRunStatistics
}

func (o *ProjectsRunStatsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/run-stats/{id}/][%d] projectsRunStatsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsRunStatsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerRunStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsRunStatsPartialUpdateBadRequest creates a ProjectsRunStatsPartialUpdateBadRequest with default headers values
func NewProjectsRunStatsPartialUpdateBadRequest() *ProjectsRunStatsPartialUpdateBadRequest {
	return &ProjectsRunStatsPartialUpdateBadRequest{}
}

/*ProjectsRunStatsPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsRunStatsPartialUpdateBadRequest struct {
}

func (o *ProjectsRunStatsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/run-stats/{id}/][%d] projectsRunStatsPartialUpdateBadRequest ", 400)
}

func (o *ProjectsRunStatsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsRunStatsPartialUpdateNotFound creates a ProjectsRunStatsPartialUpdateNotFound with default headers values
func NewProjectsRunStatsPartialUpdateNotFound() *ProjectsRunStatsPartialUpdateNotFound {
	return &ProjectsRunStatsPartialUpdateNotFound{}
}

/*ProjectsRunStatsPartialUpdateNotFound handles this case with default header values.

ServerRunStatistics not found
*/
type ProjectsRunStatsPartialUpdateNotFound struct {
}

func (o *ProjectsRunStatsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/{server_type}/{server_pk}/run-stats/{id}/][%d] projectsRunStatsPartialUpdateNotFound ", 404)
}

func (o *ProjectsRunStatsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsRunStatsPartialUpdateBody projects run stats partial update body
swagger:model ProjectsRunStatsPartialUpdateBody
*/
type ProjectsRunStatsPartialUpdateBody struct {

	// exit code
	ExitCode int64 `json:"exit_code,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// stacktrace
	Stacktrace string `json:"stacktrace,omitempty"`

	// start
	Start string `json:"start,omitempty"`

	// stop
	Stop string `json:"stop,omitempty"`
}
