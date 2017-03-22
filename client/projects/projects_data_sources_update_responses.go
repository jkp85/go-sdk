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

// ProjectsDataSourcesUpdateReader is a Reader for the ProjectsDataSourcesUpdate structure.
type ProjectsDataSourcesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDataSourcesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectsDataSourcesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectsDataSourcesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectsDataSourcesUpdateOK creates a ProjectsDataSourcesUpdateOK with default headers values
func NewProjectsDataSourcesUpdateOK() *ProjectsDataSourcesUpdateOK {
	return &ProjectsDataSourcesUpdateOK{}
}

/*ProjectsDataSourcesUpdateOK handles this case with default header values.

DataSource updated
*/
type ProjectsDataSourcesUpdateOK struct {
	Payload *models.DataSource
}

func (o *ProjectsDataSourcesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/data-sources/{server}/][%d] projectsDataSourcesUpdateOK  %+v", 200, o.Payload)
}

func (o *ProjectsDataSourcesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDataSourcesUpdateBadRequest creates a ProjectsDataSourcesUpdateBadRequest with default headers values
func NewProjectsDataSourcesUpdateBadRequest() *ProjectsDataSourcesUpdateBadRequest {
	return &ProjectsDataSourcesUpdateBadRequest{}
}

/*ProjectsDataSourcesUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type ProjectsDataSourcesUpdateBadRequest struct {
}

func (o *ProjectsDataSourcesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v0/{namespace}/projects/{project_pk}/data-sources/{server}/][%d] projectsDataSourcesUpdateBadRequest ", 400)
}

func (o *ProjectsDataSourcesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ProjectsDataSourcesUpdateBody projects data sources update body
swagger:model ProjectsDataSourcesUpdateBody
*/
type ProjectsDataSourcesUpdateBody struct {

	// server
	// Required: true
	Server *models.Server `json:"server"`
}
