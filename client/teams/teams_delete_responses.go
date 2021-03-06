package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// TeamsDeleteReader is a Reader for the TeamsDelete structure.
type TeamsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewTeamsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewTeamsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsDeleteNoContent creates a TeamsDeleteNoContent with default headers values
func NewTeamsDeleteNoContent() *TeamsDeleteNoContent {
	return &TeamsDeleteNoContent{}
}

/*TeamsDeleteNoContent handles this case with default header values.

Team deleted.
*/
type TeamsDeleteNoContent struct {
}

func (o *TeamsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/teams/{team}/][%d] teamsDeleteNoContent ", 204)
}

func (o *TeamsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeamsDeleteNotFound creates a TeamsDeleteNotFound with default headers values
func NewTeamsDeleteNotFound() *TeamsDeleteNotFound {
	return &TeamsDeleteNotFound{}
}

/*TeamsDeleteNotFound handles this case with default header values.

Team not found.
*/
type TeamsDeleteNotFound struct {
	Payload *models.NotFound
}

func (o *TeamsDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/teams/{team}/][%d] teamsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *TeamsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
