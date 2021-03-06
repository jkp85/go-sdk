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

// ProjectsDeploymentsCreateReader is a Reader for the ProjectsDeploymentsCreate structure.
type ProjectsDeploymentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDeploymentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsDeploymentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsDeploymentsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsDeploymentsCreateCreated creates a ProjectsDeploymentsCreateCreated with default headers values
func NewProjectsDeploymentsCreateCreated() *ProjectsDeploymentsCreateCreated {
	return &ProjectsDeploymentsCreateCreated{}
}

/*ProjectsDeploymentsCreateCreated handles this case with default header values.

Deployment created.
*/
type ProjectsDeploymentsCreateCreated struct {
	Payload *models.Deployment
}

func (o *ProjectsDeploymentsCreateCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDeploymentsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/projects/{project}/deployments/][%d] projectsDeploymentsCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsDeploymentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeploymentsCreateBadRequest creates a ProjectsDeploymentsCreateBadRequest with default headers values
func NewProjectsDeploymentsCreateBadRequest() *ProjectsDeploymentsCreateBadRequest {
	return &ProjectsDeploymentsCreateBadRequest{}
}

/*ProjectsDeploymentsCreateBadRequest handles this case with default header values.

Invalid data supplied.
*/
type ProjectsDeploymentsCreateBadRequest struct {
	Payload *models.DeploymentError
}

func (o *ProjectsDeploymentsCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDeploymentsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/{namespace}/projects/{project}/deployments/][%d] projectsDeploymentsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsDeploymentsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
