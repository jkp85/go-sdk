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

// ProjectsServersStatsPartialUpdateReader is a Reader for the ProjectsServersStatsPartialUpdate structure.
type ProjectsServersStatsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsServersStatsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsServersStatsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsServersStatsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectsServersStatsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsServersStatsPartialUpdateOK creates a ProjectsServersStatsPartialUpdateOK with default headers values
func NewProjectsServersStatsPartialUpdateOK() *ProjectsServersStatsPartialUpdateOK {
	return &ProjectsServersStatsPartialUpdateOK{}
}

/*ProjectsServersStatsPartialUpdateOK handles this case with default header values.

ServerStatistics updated
*/
type ProjectsServersStatsPartialUpdateOK struct {
	Payload *models.ServerStatistics
}

func (o *ProjectsServersStatsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/projects/{project_pk}/servers/{server_pk}/stats/{id}/][%d] projectsServersStatsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsServersStatsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersStatsPartialUpdateBadRequest creates a ProjectsServersStatsPartialUpdateBadRequest with default headers values
func NewProjectsServersStatsPartialUpdateBadRequest() *ProjectsServersStatsPartialUpdateBadRequest {
	return &ProjectsServersStatsPartialUpdateBadRequest{}
}

/*ProjectsServersStatsPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsServersStatsPartialUpdateBadRequest struct {
	Payload ProjectsServersStatsPartialUpdateBadRequestBody
}

func (o *ProjectsServersStatsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/projects/{project_pk}/servers/{server_pk}/stats/{id}/][%d] projectsServersStatsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsServersStatsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsServersStatsPartialUpdateNotFound creates a ProjectsServersStatsPartialUpdateNotFound with default headers values
func NewProjectsServersStatsPartialUpdateNotFound() *ProjectsServersStatsPartialUpdateNotFound {
	return &ProjectsServersStatsPartialUpdateNotFound{}
}

/*ProjectsServersStatsPartialUpdateNotFound handles this case with default header values.

ServerStatistics not found
*/
type ProjectsServersStatsPartialUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsServersStatsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /{namespace}/projects/{project_pk}/servers/{server_pk}/stats/{id}/][%d] projectsServersStatsPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsServersStatsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsServersStatsPartialUpdateBadRequestBody projects servers stats partial update bad request body
swagger:model ProjectsServersStatsPartialUpdateBadRequestBody
*/
type ProjectsServersStatsPartialUpdateBadRequestBody struct {

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// size field errors
	// Required: true
	Size []string `json:"size"`

	// start field errors
	// Required: true
	Start []string `json:"start"`

	// stop field errors
	// Required: true
	Stop []string `json:"stop"`
}

// Validate validates this projects servers stats partial update bad request body
func (o *ProjectsServersStatsPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStop(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsServersStatsPartialUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsPartialUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersStatsPartialUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsPartialUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersStatsPartialUpdateBadRequestBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsPartialUpdateBadRequest"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersStatsPartialUpdateBadRequestBody) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsPartialUpdateBadRequest"+"."+"start", "body", o.Start); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsServersStatsPartialUpdateBadRequestBody) validateStop(formats strfmt.Registry) error {

	if err := validate.Required("projectsServersStatsPartialUpdateBadRequest"+"."+"stop", "body", o.Stop); err != nil {
		return err
	}

	return nil
}

/*ProjectsServersStatsPartialUpdateBody projects servers stats partial update body
swagger:model ProjectsServersStatsPartialUpdateBody
*/
type ProjectsServersStatsPartialUpdateBody struct {

	// size
	Size int64 `json:"size,omitempty"`

	// start
	Start string `json:"start,omitempty"`

	// stop
	Stop string `json:"stop,omitempty"`
}
