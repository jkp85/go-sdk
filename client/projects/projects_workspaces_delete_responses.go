package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsWorkspacesDeleteReader is a Reader for the ProjectsWorkspacesDelete structure.
type ProjectsWorkspacesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsWorkspacesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsWorkspacesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsWorkspacesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsWorkspacesDeleteNoContent creates a ProjectsWorkspacesDeleteNoContent with default headers values
func NewProjectsWorkspacesDeleteNoContent() *ProjectsWorkspacesDeleteNoContent {
	return &ProjectsWorkspacesDeleteNoContent{}
}

/*ProjectsWorkspacesDeleteNoContent handles this case with default header values.

Workspace deleted
*/
type ProjectsWorkspacesDeleteNoContent struct {
}

func (o *ProjectsWorkspacesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/workspaces/{server}/][%d] projectsWorkspacesDeleteNoContent ", 204)
}

func (o *ProjectsWorkspacesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsWorkspacesDeleteNotFound creates a ProjectsWorkspacesDeleteNotFound with default headers values
func NewProjectsWorkspacesDeleteNotFound() *ProjectsWorkspacesDeleteNotFound {
	return &ProjectsWorkspacesDeleteNotFound{}
}

/*ProjectsWorkspacesDeleteNotFound handles this case with default header values.

Workspace not found
*/
type ProjectsWorkspacesDeleteNotFound struct {
}

func (o *ProjectsWorkspacesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/workspaces/{server}/][%d] projectsWorkspacesDeleteNotFound ", 404)
}

func (o *ProjectsWorkspacesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
