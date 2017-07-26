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

// BillingCustomersUpdateReader is a Reader for the BillingCustomersUpdate structure.
type BillingCustomersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingCustomersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBillingCustomersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBillingCustomersUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBillingCustomersUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingCustomersUpdateOK creates a BillingCustomersUpdateOK with default headers values
func NewBillingCustomersUpdateOK() *BillingCustomersUpdateOK {
	return &BillingCustomersUpdateOK{}
}

/*BillingCustomersUpdateOK handles this case with default header values.

Customer updated
*/
type BillingCustomersUpdateOK struct {
	Payload *models.Customer
}

func (o *BillingCustomersUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/billing/customers/{id}/][%d] billingCustomersUpdateOK  %+v", 200, o.Payload)
}

func (o *BillingCustomersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCustomersUpdateBadRequest creates a BillingCustomersUpdateBadRequest with default headers values
func NewBillingCustomersUpdateBadRequest() *BillingCustomersUpdateBadRequest {
	return &BillingCustomersUpdateBadRequest{}
}

/*BillingCustomersUpdateBadRequest handles this case with default header values.

Invalid data supplied
*/
type BillingCustomersUpdateBadRequest struct {
	Payload *models.CustomerError
}

func (o *BillingCustomersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/billing/customers/{id}/][%d] billingCustomersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *BillingCustomersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCustomersUpdateNotFound creates a BillingCustomersUpdateNotFound with default headers values
func NewBillingCustomersUpdateNotFound() *BillingCustomersUpdateNotFound {
	return &BillingCustomersUpdateNotFound{}
}

/*BillingCustomersUpdateNotFound handles this case with default header values.

Customer not found
*/
type BillingCustomersUpdateNotFound struct {
	Payload *models.NotFound
}

func (o *BillingCustomersUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/{namespace}/billing/customers/{id}/][%d] billingCustomersUpdateNotFound  %+v", 404, o.Payload)
}

func (o *BillingCustomersUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
