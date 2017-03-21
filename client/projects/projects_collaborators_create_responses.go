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

	"github.com/jkp85/go-sdk/models"
)

// ProjectsCollaboratorsCreateReader is a Reader for the ProjectsCollaboratorsCreate structure.
type ProjectsCollaboratorsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsCollaboratorsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProjectsCollaboratorsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsCollaboratorsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsCollaboratorsCreateCreated creates a ProjectsCollaboratorsCreateCreated with default headers values
func NewProjectsCollaboratorsCreateCreated() *ProjectsCollaboratorsCreateCreated {
	return &ProjectsCollaboratorsCreateCreated{}
}

/*ProjectsCollaboratorsCreateCreated handles this case with default header values.

Collaborator created
*/
type ProjectsCollaboratorsCreateCreated struct {
	Payload *models.Collaborator
}

func (o *ProjectsCollaboratorsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/collaborators/][%d] projectsCollaboratorsCreateCreated  %+v", 201, o.Payload)
}

func (o *ProjectsCollaboratorsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collaborator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCollaboratorsCreateBadRequest creates a ProjectsCollaboratorsCreateBadRequest with default headers values
func NewProjectsCollaboratorsCreateBadRequest() *ProjectsCollaboratorsCreateBadRequest {
	return &ProjectsCollaboratorsCreateBadRequest{}
}

/*ProjectsCollaboratorsCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsCollaboratorsCreateBadRequest struct {
	Payload ProjectsCollaboratorsCreateBadRequestBody
}

func (o *ProjectsCollaboratorsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/{namespace}/projects/{project_pk}/collaborators/][%d] projectsCollaboratorsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsCollaboratorsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsCollaboratorsCreateBadRequestBody projects collaborators create bad request body
swagger:model ProjectsCollaboratorsCreateBadRequestBody
*/
type ProjectsCollaboratorsCreateBadRequestBody struct {

	// email field errors
	// Required: true
	Email []string `json:"email"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// joined field errors
	// Required: true
	Joined []string `json:"joined"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// owner field errors
	// Required: true
	Owner []string `json:"owner"`
}

// Validate validates this projects collaborators create bad request body
func (o *ProjectsCollaboratorsCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateJoined(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOwner(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsCollaboratorsCreateBadRequestBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsCreateBadRequest"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCollaboratorsCreateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsCreateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCollaboratorsCreateBadRequestBody) validateJoined(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsCreateBadRequest"+"."+"joined", "body", o.Joined); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCollaboratorsCreateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsCreateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCollaboratorsCreateBadRequestBody) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsCreateBadRequest"+"."+"owner", "body", o.Owner); err != nil {
		return err
	}

	return nil
}

/*ProjectsCollaboratorsCreateBody projects collaborators create body
swagger:model ProjectsCollaboratorsCreateBody
*/
type ProjectsCollaboratorsCreateBody struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// owner
	Owner bool `json:"owner,omitempty"`
}
