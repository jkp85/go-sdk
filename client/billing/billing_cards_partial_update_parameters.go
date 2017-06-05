package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBillingCardsPartialUpdateParams creates a new BillingCardsPartialUpdateParams object
// with the default values initialized.
func NewBillingCardsPartialUpdateParams() *BillingCardsPartialUpdateParams {
	var ()
	return &BillingCardsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingCardsPartialUpdateParamsWithTimeout creates a new BillingCardsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingCardsPartialUpdateParamsWithTimeout(timeout time.Duration) *BillingCardsPartialUpdateParams {
	var ()
	return &BillingCardsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewBillingCardsPartialUpdateParamsWithContext creates a new BillingCardsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingCardsPartialUpdateParamsWithContext(ctx context.Context) *BillingCardsPartialUpdateParams {
	var ()
	return &BillingCardsPartialUpdateParams{

		Context: ctx,
	}
}

// NewBillingCardsPartialUpdateParamsWithHTTPClient creates a new BillingCardsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingCardsPartialUpdateParamsWithHTTPClient(client *http.Client) *BillingCardsPartialUpdateParams {
	var ()
	return &BillingCardsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*BillingCardsPartialUpdateParams contains all the parameters to send to the API endpoint
for the billing cards partial update operation typically these are written to a http.Request
*/
type BillingCardsPartialUpdateParams struct {

	/*Data*/
	Data BillingCardsPartialUpdateBody
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) WithTimeout(timeout time.Duration) *BillingCardsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) WithContext(ctx context.Context) *BillingCardsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) WithHTTPClient(client *http.Client) *BillingCardsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) WithData(data BillingCardsPartialUpdateBody) *BillingCardsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) SetData(data BillingCardsPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) WithID(id string) *BillingCardsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) WithNamespace(namespace string) *BillingCardsPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing cards partial update params
func (o *BillingCardsPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingCardsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}