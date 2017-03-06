package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsServersSSHTunnelsDeleteReader is a Reader for the ProjectsServersSSHTunnelsDelete structure.
type ProjectsServersSSHTunnelsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersSSHTunnelsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewProjectsServersSSHTunnelsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectsServersSSHTunnelsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersSSHTunnelsDeleteNoContent creates a ProjectsServersSSHTunnelsDeleteNoContent with default headers values
func NewProjectsServersSSHTunnelsDeleteNoContent() *ProjectsServersSSHTunnelsDeleteNoContent {
	return &ProjectsServersSSHTunnelsDeleteNoContent{}
}

/*ProjectsServersSSHTunnelsDeleteNoContent handles this case with default header values.

SshTunnel deleted
*/
type ProjectsServersSSHTunnelsDeleteNoContent struct {
}

func (o *ProjectsServersSSHTunnelsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/ssh-tunnels/{id}/][%d] projectsServersSshTunnelsDeleteNoContent ", 204)
}

func (o *ProjectsServersSSHTunnelsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsServersSSHTunnelsDeleteNotFound creates a ProjectsServersSSHTunnelsDeleteNotFound with default headers values
func NewProjectsServersSSHTunnelsDeleteNotFound() *ProjectsServersSSHTunnelsDeleteNotFound {
	return &ProjectsServersSSHTunnelsDeleteNotFound{}
}

/*ProjectsServersSSHTunnelsDeleteNotFound handles this case with default header values.

SshTunnel not found
*/
type ProjectsServersSSHTunnelsDeleteNotFound struct {
}

func (o *ProjectsServersSSHTunnelsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v0/{namespace}/projects/{project_pk}/servers/{server_pk}/ssh-tunnels/{id}/][%d] projectsServersSshTunnelsDeleteNotFound ", 404)
}

func (o *ProjectsServersSSHTunnelsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}