package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsCollaboratorsDeleteReader is a Reader for the ProjectsCollaboratorsDelete structure.
type ProjectsCollaboratorsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsCollaboratorsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsCollaboratorsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsCollaboratorsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsCollaboratorsDeleteNoContent creates a ProjectsCollaboratorsDeleteNoContent with default headers values
func NewProjectsCollaboratorsDeleteNoContent() *ProjectsCollaboratorsDeleteNoContent {
	return &ProjectsCollaboratorsDeleteNoContent{}
}

/*ProjectsCollaboratorsDeleteNoContent handles this case with default header values.

Collaborator deleted
*/
type ProjectsCollaboratorsDeleteNoContent struct {
}

func (o *ProjectsCollaboratorsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsDeleteNoContent ", 204)
}

func (o *ProjectsCollaboratorsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsCollaboratorsDeleteNotFound creates a ProjectsCollaboratorsDeleteNotFound with default headers values
func NewProjectsCollaboratorsDeleteNotFound() *ProjectsCollaboratorsDeleteNotFound {
	return &ProjectsCollaboratorsDeleteNotFound{}
}

/*ProjectsCollaboratorsDeleteNotFound handles this case with default header values.

Collaborator not found
*/
type ProjectsCollaboratorsDeleteNotFound struct {
}

func (o *ProjectsCollaboratorsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsDeleteNotFound ", 404)
}

func (o *ProjectsCollaboratorsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}