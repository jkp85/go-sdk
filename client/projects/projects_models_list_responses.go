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

// ProjectsModelsListReader is a Reader for the ProjectsModelsList structure.
type ProjectsModelsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsModelsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsModelsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsModelsListOK creates a ProjectsModelsListOK with default headers values
func NewProjectsModelsListOK() *ProjectsModelsListOK {
	return &ProjectsModelsListOK{}
}

/*ProjectsModelsListOK handles this case with default header values.

Model list
*/
type ProjectsModelsListOK struct {
	Payload []*models.Model
}

func (o *ProjectsModelsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v0/{namespace}/projects/{project_pk}/models/][%d] projectsModelsListOK  %+v", 200, o.Payload)
}

func (o *ProjectsModelsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
