// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ProjectsServersStartReader is a Reader for the ProjectsServersStart structure.
type ProjectsServersStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsServersStartCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersStartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersStartCreated creates a ProjectsServersStartCreated with default headers values
func NewProjectsServersStartCreated() *ProjectsServersStartCreated {
	return &ProjectsServersStartCreated{}
}

/*ProjectsServersStartCreated handles this case with default header values.

Server started.
*/
type ProjectsServersStartCreated struct {
}

func (o *ProjectsServersStartCreated) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/projects/{project}/servers/{server}/start/][%d] projectsServersStartCreated ", 201)
}

func (o *ProjectsServersStartCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsServersStartBadRequest creates a ProjectsServersStartBadRequest with default headers values
func NewProjectsServersStartBadRequest() *ProjectsServersStartBadRequest {
	return &ProjectsServersStartBadRequest{}
}

/*ProjectsServersStartBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ProjectsServersStartBadRequest struct {
}

func (o *ProjectsServersStartBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/projects/{project}/servers/{server}/start/][%d] projectsServersStartBadRequest ", 400)
}

func (o *ProjectsServersStartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
