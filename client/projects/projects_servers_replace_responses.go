// Code generated by go-swagger; DO NOT EDIT.

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

// ProjectsServersReplaceReader is a Reader for the ProjectsServersReplace structure.
type ProjectsServersReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersReplaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersReplaceOK creates a ProjectsServersReplaceOK with default headers values
func NewProjectsServersReplaceOK() *ProjectsServersReplaceOK {
	return &ProjectsServersReplaceOK{}
}

/*ProjectsServersReplaceOK handles this case with default header values.

Server updated.
*/
type ProjectsServersReplaceOK struct {
	Payload *models.Server
}

func (o *ProjectsServersReplaceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/projects/{project_id}/servers/{id}/][%d] projectsServersReplaceOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersReplaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersReplaceBadRequest creates a ProjectsServersReplaceBadRequest with default headers values
func NewProjectsServersReplaceBadRequest() *ProjectsServersReplaceBadRequest {
	return &ProjectsServersReplaceBadRequest{}
}

/*ProjectsServersReplaceBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ProjectsServersReplaceBadRequest struct {
	Payload *models.ServerError
}

func (o *ProjectsServersReplaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/projects/{project_id}/servers/{id}/][%d] projectsServersReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}