package hosts

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

// HostsUpdateReader is a Reader for the HostsUpdate structure.
type HostsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHostsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewHostsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHostsUpdateOK creates a HostsUpdateOK with default headers values
func NewHostsUpdateOK() *HostsUpdateOK {
	return &HostsUpdateOK{}
}

/*HostsUpdateOK handles this case with default header values.

DockerHost updated
*/
type HostsUpdateOK struct {
	Payload *models.DockerHost
}

func (o *HostsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/hosts/{id}/][%d] hostsUpdateOK  %+v", 200, o.Payload)
}

func (o *HostsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DockerHost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHostsUpdateBadRequest creates a HostsUpdateBadRequest with default headers values
func NewHostsUpdateBadRequest() *HostsUpdateBadRequest {
	return &HostsUpdateBadRequest{}
}

/*HostsUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type HostsUpdateBadRequest struct {
	Payload HostsUpdateBadRequestBody
}

func (o *HostsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{namespace}/hosts/{id}/][%d] hostsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *HostsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*HostsUpdateBadRequestBody hosts update bad request body
swagger:model HostsUpdateBadRequestBody
*/
type HostsUpdateBadRequestBody struct {

	// id field errors
	// Required: true
	ID []string `json:"id"`

	// ip field errors
	// Required: true
	IP []string `json:"ip"`

	// name field errors
	// Required: true
	Name []string `json:"name"`

	// Errors not connected to any field
	// Required: true
	NonFieldErrors []string `json:"non_field_errors"`

	// port field errors
	// Required: true
	Port []string `json:"port"`

	// status field errors
	// Required: true
	Status []string `json:"status"`
}

// Validate validates this hosts update bad request body
func (o *HostsUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIP(formats); err != nil {
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

	if err := o.validatePort(formats); err != nil {
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

func (o *HostsUpdateBadRequestBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("hostsUpdateBadRequest"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *HostsUpdateBadRequestBody) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("hostsUpdateBadRequest"+"."+"ip", "body", o.IP); err != nil {
		return err
	}

	return nil
}

func (o *HostsUpdateBadRequestBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("hostsUpdateBadRequest"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *HostsUpdateBadRequestBody) validateNonFieldErrors(formats strfmt.Registry) error {

	if err := validate.Required("hostsUpdateBadRequest"+"."+"non_field_errors", "body", o.NonFieldErrors); err != nil {
		return err
	}

	return nil
}

func (o *HostsUpdateBadRequestBody) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("hostsUpdateBadRequest"+"."+"port", "body", o.Port); err != nil {
		return err
	}

	return nil
}

func (o *HostsUpdateBadRequestBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("hostsUpdateBadRequest"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

/*HostsUpdateBody hosts update body
swagger:model HostsUpdateBody
*/
type HostsUpdateBody struct {

	// ip
	// Required: true
	IP *string `json:"ip"`

	// name
	// Required: true
	Name *string `json:"name"`

	// port
	Port int64 `json:"port,omitempty"`
}
