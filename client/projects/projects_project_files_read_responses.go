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

// ProjectsProjectFilesReadReader is a Reader for the ProjectsProjectFilesRead structure.
type ProjectsProjectFilesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectFilesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsProjectFilesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsProjectFilesReadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsProjectFilesReadOK creates a ProjectsProjectFilesReadOK with default headers values
func NewProjectsProjectFilesReadOK() *ProjectsProjectFilesReadOK {
	return &ProjectsProjectFilesReadOK{}
}

/*ProjectsProjectFilesReadOK handles this case with default header values.

ProjectFile retrieved
*/
type ProjectsProjectFilesReadOK struct {
	Payload *models.ProjectFile
}

func (o *ProjectsProjectFilesReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsProjectFilesReadOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/projects/{project}/project_files/{id}/][%d] projectsProjectFilesReadOK  %+v", 200, o.Payload)
}

func (o *ProjectsProjectFilesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsProjectFilesReadNotFound creates a ProjectsProjectFilesReadNotFound with default headers values
func NewProjectsProjectFilesReadNotFound() *ProjectsProjectFilesReadNotFound {
	return &ProjectsProjectFilesReadNotFound{}
}

/*ProjectsProjectFilesReadNotFound handles this case with default header values.

ProjectFile not found
*/
type ProjectsProjectFilesReadNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsProjectFilesReadNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsProjectFilesReadNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/projects/{project}/project_files/{id}/][%d] projectsProjectFilesReadNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsProjectFilesReadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
