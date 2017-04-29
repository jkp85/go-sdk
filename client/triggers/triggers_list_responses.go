package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// TriggersListReader is a Reader for the TriggersList structure.
type TriggersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTriggersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTriggersListOK creates a TriggersListOK with default headers values
func NewTriggersListOK() *TriggersListOK {
	return &TriggersListOK{}
}

/*TriggersListOK handles this case with default header values.

Trigger list
*/
type TriggersListOK struct {
	Payload []*models.Trigger
}

func (o *TriggersListOK) Error() string {
	return fmt.Sprintf("[GET /{namespace}/triggers/][%d] triggersListOK  %+v", 200, o.Payload)
}

func (o *TriggersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
