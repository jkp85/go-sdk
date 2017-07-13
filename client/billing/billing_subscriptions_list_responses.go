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

// BillingSubscriptionsListReader is a Reader for the BillingSubscriptionsList structure.
type BillingSubscriptionsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingSubscriptionsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBillingSubscriptionsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingSubscriptionsListOK creates a BillingSubscriptionsListOK with default headers values
func NewBillingSubscriptionsListOK() *BillingSubscriptionsListOK {
	return &BillingSubscriptionsListOK{}
}

/*BillingSubscriptionsListOK handles this case with default header values.

Subscription list
*/
type BillingSubscriptionsListOK struct {
	Payload []*models.Subscription
}

func (o *BillingSubscriptionsListOK) Error() string {
	return fmt.Sprintf("[GET /{namespace}/billing/subscriptions/][%d] billingSubscriptionsListOK  %+v", 200, o.Payload)
}

func (o *BillingSubscriptionsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
