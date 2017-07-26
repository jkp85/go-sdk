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

// BillingPlansReplaceReader is a Reader for the BillingPlansReplace structure.
type BillingPlansReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingPlansReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBillingPlansReplaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBillingPlansReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingPlansReplaceOK creates a BillingPlansReplaceOK with default headers values
func NewBillingPlansReplaceOK() *BillingPlansReplaceOK {
	return &BillingPlansReplaceOK{}
}

/*BillingPlansReplaceOK handles this case with default header values.

Plan updated.
*/
type BillingPlansReplaceOK struct {
	Payload *models.Plan
}

func (o *BillingPlansReplaceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/billing/plans/{id}/][%d] billingPlansReplaceOK  %+v", 200, o.Payload)
}

func (o *BillingPlansReplaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingPlansReplaceBadRequest creates a BillingPlansReplaceBadRequest with default headers values
func NewBillingPlansReplaceBadRequest() *BillingPlansReplaceBadRequest {
	return &BillingPlansReplaceBadRequest{}
}

/*BillingPlansReplaceBadRequest handles this case with default header values.

Invalid data supplied.
*/
type BillingPlansReplaceBadRequest struct {
	Payload *models.PlanError
}

func (o *BillingPlansReplaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/{namespace}/billing/plans/{id}/][%d] billingPlansReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *BillingPlansReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}