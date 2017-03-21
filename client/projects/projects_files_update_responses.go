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

// ProjectsFilesUpdateReader is a Reader for the ProjectsFilesUpdate structure.
type ProjectsFilesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsFilesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsFilesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsFilesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsFilesUpdateOK creates a ProjectsFilesUpdateOK with default headers values
func NewProjectsFilesUpdateOK() *ProjectsFilesUpdateOK {
	return &ProjectsFilesUpdateOK{}
}

/*ProjectsFilesUpdateOK handles this case with default header values.

File updated
*/
type ProjectsFilesUpdateOK struct {
	Payload *models.File
}

func (o *ProjectsFilesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/files/{id}/][%d] projectsFilesUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsFilesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsFilesUpdateBadRequest creates a ProjectsFilesUpdateBadRequest with default headers values
func NewProjectsFilesUpdateBadRequest() *ProjectsFilesUpdateBadRequest {
	return &ProjectsFilesUpdateBadRequest{}
}

/*ProjectsFilesUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsFilesUpdateBadRequest struct {
	Payload ProjectsFilesUpdateBadRequestBody
}

func (o *ProjectsFilesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/files/{id}/][%d] projectsFilesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsFilesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProjectsFilesUpdateBadRequestBody projects files update bad request body
swagger:model ProjectsFilesUpdateBadRequestBody
*/
type ProjectsFilesUpdateBadRequestBody struct {

	// author field errors
	// Required: true
	Author []string `json:"author"`

	// content field errors
	// Required: true
	Content []string `json:"content"`

	// encoding field errors
	// Required: true
	Encoding []string `json:"encoding"`

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// path field errors
	// Required: true
	Path []string `json:"path"`

	// project field errors
	// Required: true
	Project []string `json:"project"`

	// public field errors
	// Required: true
	Public []string `json:"public"`

	// size field errors
	// Required: true
	Size []string `json:"size"`
}

// Validate validates this projects files update bad request body
func (o *ProjectsFilesUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEncoding(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNonFieldErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePath(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateProject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePublic(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"author", "body", o.Author); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"content", "body", o.Content); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validateEncoding(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"encoding", "body", o.Encoding); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"path", "body", o.Path); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validatePublic(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"public", "body", o.Public); err != nil {
		return err
	}

	return nil
}

func (o *ProjectsFilesUpdateBadRequestBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("projectsFilesUpdateBadRequest"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

/*ProjectsFilesUpdateBody projects files update body
swagger:model ProjectsFilesUpdateBody
*/
type ProjectsFilesUpdateBody struct {

	// author
	// Required: true
	Author *string `json:"author"`

	// content
	// Required: true
	Content *string `json:"content"`

	// encoding
	// Required: true
	Encoding *string `json:"encoding"`

	// path
	// Required: true
	Path *string `json:"path"`

	// project
	// Required: true
	Project *string `json:"project"`

	// public
	Public bool `json:"public,omitempty"`
}
