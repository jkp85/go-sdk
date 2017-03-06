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

// ProjectsCollaboratorsUpdateReader is a Reader for the ProjectsCollaboratorsUpdate structure.
type ProjectsCollaboratorsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsCollaboratorsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsCollaboratorsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsCollaboratorsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsCollaboratorsUpdateOK creates a ProjectsCollaboratorsUpdateOK with default headers values
func NewProjectsCollaboratorsUpdateOK() *ProjectsCollaboratorsUpdateOK {
	return &ProjectsCollaboratorsUpdateOK{}
}

/*ProjectsCollaboratorsUpdateOK handles this case with default header values.

Collaborator updated
*/
type ProjectsCollaboratorsUpdateOK struct {
	Payload *models.Collaborator
}

func (o *ProjectsCollaboratorsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsCollaboratorsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collaborator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCollaboratorsUpdateBadRequest creates a ProjectsCollaboratorsUpdateBadRequest with default headers values
func NewProjectsCollaboratorsUpdateBadRequest() *ProjectsCollaboratorsUpdateBadRequest {
	return &ProjectsCollaboratorsUpdateBadRequest{}
}

/*ProjectsCollaboratorsUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsCollaboratorsUpdateBadRequest struct {
}

func (o *ProjectsCollaboratorsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsUpdateBadRequest ", 400)
}

func (o *ProjectsCollaboratorsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsCollaboratorsUpdateBody projects collaborators update body
swagger:model ProjectsCollaboratorsUpdateBody
*/
type ProjectsCollaboratorsUpdateBody struct {

	// owner
	Owner bool `json:"owner,omitempty"`

	// user
	// Required: true
	User *models.User `json:"user"`
}