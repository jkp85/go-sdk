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

// ProjectsServersAPIKeyReader is a Reader for the ProjectsServersAPIKey structure.
type ProjectsServersAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersAPIKeyOK creates a ProjectsServersAPIKeyOK with default headers values
func NewProjectsServersAPIKeyOK() *ProjectsServersAPIKeyOK {
	return &ProjectsServersAPIKeyOK{}
}

/*ProjectsServersAPIKeyOK handles this case with default header values.

Server API key
*/
type ProjectsServersAPIKeyOK struct {
	Payload *models.JWT
}

func (o *ProjectsServersAPIKeyOK) Error() string {
	return fmt.Sprintf("[GET /v1/{namespace}/projects/{project}/servers/{server}/api-key/][%d] projectsServersApiKeyOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JWT)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
