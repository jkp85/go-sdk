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

// ProjectsProjectFilesListReader is a Reader for the ProjectsProjectFilesList structure.
type ProjectsProjectFilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectFilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsProjectFilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsProjectFilesListOK creates a ProjectsProjectFilesListOK with default headers values
func NewProjectsProjectFilesListOK() *ProjectsProjectFilesListOK {
	return &ProjectsProjectFilesListOK{}
}

/*ProjectsProjectFilesListOK handles this case with default header values.

ProjectFile list
*/
type ProjectsProjectFilesListOK struct {
	Payload []*models.ProjectFile
}

func (o *ProjectsProjectFilesListOK) Error() string {
	return fmt.Sprintf("[GET /{namespace}/projects/{project_pk}/project_files/][%d] projectsProjectFilesListOK  %+v", 200, o.Payload)
}

func (o *ProjectsProjectFilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}