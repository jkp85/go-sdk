package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// ProjectsServersUpdateReader is a Reader for the ProjectsServersUpdate structure.
type ProjectsServersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersUpdateOK creates a ProjectsServersUpdateOK with default headers values
func NewProjectsServersUpdateOK() *ProjectsServersUpdateOK {
	return &ProjectsServersUpdateOK{}
}

/*ProjectsServersUpdateOK handles this case with default header values.

Server updated
*/
type ProjectsServersUpdateOK struct {
	Payload *models.Server
}

func (o *ProjectsServersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/projects/{project_pk}/servers/{id}/][%d] projectsServersUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersUpdateBadRequest creates a ProjectsServersUpdateBadRequest with default headers values
func NewProjectsServersUpdateBadRequest() *ProjectsServersUpdateBadRequest {
	return &ProjectsServersUpdateBadRequest{}
}

/*ProjectsServersUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersUpdateBadRequest struct {
	Payload ProjectsServersUpdateBadRequestBody
}

func (o *ProjectsServersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/projects/{project_pk}/servers/{id}/][%d] projectsServersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsServersUpdateBadRequestBody projects servers update bad request body
swagger:model ProjectsServersUpdateBadRequestBody
*/
type ProjectsServersUpdateBadRequestBody struct {

	// config field errors
	// Required: true
	Config []string `json:"config"`

	// connected field errors
	// Required: true
	Connected []string `json:"connected"`

	// created_at field errors
	// Required: true
	CreatedAt []string `json:"created_at"`

	// endpoint field errors
	// Required: true
	Endpoint []string `json:"endpoint"`

	// environment_resources field errors
	// Required: true
	EnvironmentResources []string `json:"environment_resources"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// image_name field errors
	// Required: true
	ImageName []string `json:"image_name"`

	// name field errors
	// Required: true
	Name []string `json:"name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// startup_script field errors
	// Required: true
	StartupScript []string `json:"startup_script"`

	// status field errors
	// Required: true
	Status []string `json:"status"`
}

// Validate validates this projects servers update bad request body
func (o *ProjectsServersUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateConnected(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEndpoint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEnvironmentResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateImageName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStartupScript(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"config", "body", o.Config); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateConnected(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"connected", "body", o.Connected); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"endpoint", "body", o.Endpoint); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateEnvironmentResources(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"environment_resources", "body", o.EnvironmentResources); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateImageName(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"image_name", "body", o.ImageName); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateStartupScript(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"startup_script", "body", o.StartupScript); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersUpdateBadRequestBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersUpdateBadRequest"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

/*ProjectsServersUpdateBody projects servers update body
swagger:model ProjectsServersUpdateBody
*/
type ProjectsServersUpdateBody struct {

	// config
	Config interface{} `json:"config,omitempty"`

	// connected
	Connected []string `json:"connected"`

	// environment resources
	EnvironmentResources string `json:"environment_resources,omitempty"`

	// image name
	ImageName string `json:"image_name,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// startup script
	StartupScript string `json:"startup_script,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}
