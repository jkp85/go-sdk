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

// BillingCustomersListReader is a Reader for the BillingCustomersList structure.
type BillingCustomersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingCustomersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBillingCustomersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingCustomersListOK creates a BillingCustomersListOK with default headers values
func NewBillingCustomersListOK() *BillingCustomersListOK {
	return &BillingCustomersListOK{}
}

/*BillingCustomersListOK handles this case with default header values.

Customer list
*/
type BillingCustomersListOK struct {
	Payload []*models.Customer
}

func (o *BillingCustomersListOK) Error() string {
	return fmt.Sprintf("[GET /{namespace}/billing/customers/][%d] billingCustomersListOK  %+v", 200, o.Payload)
}

func (o *BillingCustomersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
