// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// BillingPlansUpdateReader is a Reader for the BillingPlansUpdate structure.
type BillingPlansUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingPlansUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBillingPlansUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBillingPlansUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBillingPlansUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingPlansUpdateOK creates a BillingPlansUpdateOK with default headers values
func NewBillingPlansUpdateOK() *BillingPlansUpdateOK {
	return &BillingPlansUpdateOK{}
}

/*BillingPlansUpdateOK handles this case with default header values.

Plan updated
*/
type BillingPlansUpdateOK struct {
	Payload *models.Plan
}

func (o *BillingPlansUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/billing/plans/{id}/][%d] billingPlansUpdateOK  %+v", 200, o.Payload)
}

func (o *BillingPlansUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingPlansUpdateBadRequest creates a BillingPlansUpdateBadRequest with default headers values
func NewBillingPlansUpdateBadRequest() *BillingPlansUpdateBadRequest {
	return &BillingPlansUpdateBadRequest{}
}

/*BillingPlansUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type BillingPlansUpdateBadRequest struct {
	Payload *models.PlanError
}

func (o *BillingPlansUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/billing/plans/{id}/][%d] billingPlansUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *BillingPlansUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingPlansUpdateNotFound creates a BillingPlansUpdateNotFound with default headers values
func NewBillingPlansUpdateNotFound() *BillingPlansUpdateNotFound {
	return &BillingPlansUpdateNotFound{}
}

/*BillingPlansUpdateNotFound handles this case with default header values.

Plan not found
*/
type BillingPlansUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *BillingPlansUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/billing/plans/{id}/][%d] billingPlansUpdateNotFound  %+v", 404, o.Payload)
}

func (o *BillingPlansUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
