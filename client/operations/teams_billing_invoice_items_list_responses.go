package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IllumiDesk/go-sdk/models"
)

// TeamsBillingInvoiceItemsListReader is a Reader for the TeamsBillingInvoiceItemsList structure.
type TeamsBillingInvoiceItemsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsBillingInvoiceItemsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTeamsBillingInvoiceItemsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsBillingInvoiceItemsListOK creates a TeamsBillingInvoiceItemsListOK with default headers values
func NewTeamsBillingInvoiceItemsListOK() *TeamsBillingInvoiceItemsListOK {
	return &TeamsBillingInvoiceItemsListOK{}
}

/*TeamsBillingInvoiceItemsListOK handles this case with default header values.

Team invoiceItem list.
*/
type TeamsBillingInvoiceItemsListOK struct {
	Payload models.TeamsBillingInvoiceItemsListOKBody
}

func (o *TeamsBillingInvoiceItemsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsBillingInvoiceItemsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{team}/billing/invoices/{invoice_id}/invoice-items/][%d] teamsBillingInvoiceItemsListOK  %+v", 200, o.Payload)
}

func (o *TeamsBillingInvoiceItemsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
