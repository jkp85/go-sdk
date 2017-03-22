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

// ProjectsCollaboratorsPartialUpdateReader is a Reader for the ProjectsCollaboratorsPartialUpdate structure.
type ProjectsCollaboratorsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsCollaboratorsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsCollaboratorsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsCollaboratorsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectsCollaboratorsPartialUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsCollaboratorsPartialUpdateOK creates a ProjectsCollaboratorsPartialUpdateOK with default headers values
func NewProjectsCollaboratorsPartialUpdateOK() *ProjectsCollaboratorsPartialUpdateOK {
	return &ProjectsCollaboratorsPartialUpdateOK{}
}

/*ProjectsCollaboratorsPartialUpdateOK handles this case with default header values.

Collaborator updated
*/
type ProjectsCollaboratorsPartialUpdateOK struct {
	Payload *models.Collaborator
}

func (o *ProjectsCollaboratorsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsCollaboratorsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collaborator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCollaboratorsPartialUpdateBadRequest creates a ProjectsCollaboratorsPartialUpdateBadRequest with default headers values
func NewProjectsCollaboratorsPartialUpdateBadRequest() *ProjectsCollaboratorsPartialUpdateBadRequest {
	return &ProjectsCollaboratorsPartialUpdateBadRequest{}
}

/*ProjectsCollaboratorsPartialUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsCollaboratorsPartialUpdateBadRequest struct {
	Payload ProjectsCollaboratorsPartialUpdateBadRequestBody
}

func (o *ProjectsCollaboratorsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsCollaboratorsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCollaboratorsPartialUpdateNotFound creates a ProjectsCollaboratorsPartialUpdateNotFound with default headers values
func NewProjectsCollaboratorsPartialUpdateNotFound() *ProjectsCollaboratorsPartialUpdateNotFound {
	return &ProjectsCollaboratorsPartialUpdateNotFound{}
}

/*ProjectsCollaboratorsPartialUpdateNotFound handles this case with default header values.

Collaborator not found
*/
type ProjectsCollaboratorsPartialUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *ProjectsCollaboratorsPartialUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v0/{namespace}/projects/{project_pk}/collaborators/{id}/][%d] projectsCollaboratorsPartialUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsCollaboratorsPartialUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsCollaboratorsPartialUpdateBadRequestBody projects collaborators partial update bad request body
swagger:model ProjectsCollaboratorsPartialUpdateBadRequestBody
*/
type ProjectsCollaboratorsPartialUpdateBadRequestBody struct {

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

// Validate validates this projects collaborators partial update bad request body
func (o *ProjectsCollaboratorsPartialUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *ProjectsCollaboratorsPartialUpdateBadRequestBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsPartialUpdateBadRequest"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCollaboratorsPartialUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsPartialUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCollaboratorsPartialUpdateBadRequestBody) validateJoined(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsPartialUpdateBadRequest"+"."+"joined", "body", o.Joined); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCollaboratorsPartialUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsPartialUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsCollaboratorsPartialUpdateBadRequestBody) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("projectsCollaboratorsPartialUpdateBadRequest"+"."+"owner", "body", o.Owner); err != nil {
		return err
	}

	return nil
}

/*ProjectsCollaboratorsPartialUpdateBody projects collaborators partial update body
swagger:model ProjectsCollaboratorsPartialUpdateBody
*/
type ProjectsCollaboratorsPartialUpdateBody struct {

	// email
	Email string `json:"email,omitempty"`

	// owner
	Owner bool `json:"owner,omitempty"`
}
