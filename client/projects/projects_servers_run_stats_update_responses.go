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

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersRunStatsUpdateOK creates a ProjectsServersRunStatsUpdateOK with default headers values
func NewProjectsServersRunStatsUpdateOK() *ProjectsServersRunStatsUpdateOK {
	return &ProjectsServersRunStatsUpdateOK{}
}

/*ProjectsServersRunStatsUpdateOK handles this case with default header values.

ServerRunStatistics updated
*/
type ProjectsServersRunStatsUpdateOK struct {
	Payload *models.ServerRunStatistics
}

func (o *ProjectsServersRunStatsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/run-stats/{id}/][%d] projectsServersRunStatsUpdateOK  %+v", 200, o.Payload)
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

Invalid data supplied
*/
type ProjectsServersRunStatsUpdateBadRequest struct {
}

func (o *ProjectsServersRunStatsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/run-stats/{id}/][%d] projectsServersRunStatsUpdateBadRequest ", 400)
}

func (o *ProjectsServersRunStatsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsServersRunStatsUpdateBody projects servers run stats update body
swagger:model ProjectsServersRunStatsUpdateBody
*/
type ProjectsServersRunStatsUpdateBody struct {

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