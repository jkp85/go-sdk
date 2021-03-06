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

// ProjectsDeploymentsUpdateReader is a Reader for the ProjectsDeploymentsUpdate structure.
type ProjectsDeploymentsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDeploymentsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsDeploymentsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsDeploymentsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectsDeploymentsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsDeploymentsUpdateOK creates a ProjectsDeploymentsUpdateOK with default headers values
func NewProjectsDeploymentsUpdateOK() *ProjectsDeploymentsUpdateOK {
	return &ProjectsDeploymentsUpdateOK{}
}

/*ProjectsDeploymentsUpdateOK handles this case with default header values.

Deployment updated
*/
type ProjectsDeploymentsUpdateOK struct {
	Payload *models.Deployment
}

func (o *ProjectsDeploymentsUpdateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDeploymentsUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/deployments/{deployment}/][%d] projectsDeploymentsUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsDeploymentsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeploymentsUpdateBadRequest creates a ProjectsDeploymentsUpdateBadRequest with default headers values
func NewProjectsDeploymentsUpdateBadRequest() *ProjectsDeploymentsUpdateBadRequest {
	return &ProjectsDeploymentsUpdateBadRequest{}
}

/*ProjectsDeploymentsUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsDeploymentsUpdateBadRequest struct {
	Payload *models.DeploymentData
}

func (o *ProjectsDeploymentsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDeploymentsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/deployments/{deployment}/][%d] projectsDeploymentsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsDeploymentsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeploymentsUpdateNotFound creates a ProjectsDeploymentsUpdateNotFound with default headers values
func NewProjectsDeploymentsUpdateNotFound() *ProjectsDeploymentsUpdateNotFound {
	return &ProjectsDeploymentsUpdateNotFound{}
}

/*ProjectsDeploymentsUpdateNotFound handles this case with default header values.

Deployment not found
*/
type ProjectsDeploymentsUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsDeploymentsUpdateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDeploymentsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/projects/{project}/deployments/{deployment}/][%d] projectsDeploymentsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDeploymentsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
