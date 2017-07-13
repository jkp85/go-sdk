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

// ProjectsModelsReadReader is a Reader for the ProjectsModelsRead structure.
type ProjectsModelsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsModelsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsModelsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsModelsReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsModelsReadOK creates a ProjectsModelsReadOK with default headers values
func NewProjectsModelsReadOK() *ProjectsModelsReadOK {
	return &ProjectsModelsReadOK{}
}

/*ProjectsModelsReadOK handles this case with default header values.

Model retrieved
*/
type ProjectsModelsReadOK struct {
	Payload *models.Model
}

func (o *ProjectsModelsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/models/{server}/][%d] projectsModelsReadOK  %+v", 200, o.Payload)
}

func (o *ProjectsModelsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsModelsReadNotFound creates a ProjectsModelsReadNotFound with default headers values
func NewProjectsModelsReadNotFound() *ProjectsModelsReadNotFound {
	return &ProjectsModelsReadNotFound{}
}

/*ProjectsModelsReadNotFound handles this case with default header values.

Model not found
*/
type ProjectsModelsReadNotFound struct {
}

func (o *ProjectsModelsReadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/models/{server}/][%d] projectsModelsReadNotFound ", 404)
}

func (o *ProjectsModelsReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
