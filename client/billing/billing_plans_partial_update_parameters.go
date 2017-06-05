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

// NewBillingPlansPartialUpdateParams creates a new BillingPlansPartialUpdateParams object
// with the default values initialized.
func NewBillingPlansPartialUpdateParams() *BillingPlansPartialUpdateParams {
	var ()
	return &BillingPlansPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingPlansPartialUpdateParamsWithTimeout creates a new BillingPlansPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingPlansPartialUpdateParamsWithTimeout(timeout time.Duration) *BillingPlansPartialUpdateParams {
	var ()
	return &BillingPlansPartialUpdateParams{

		timeout: timeout,
	}
}

// NewBillingPlansPartialUpdateParamsWithContext creates a new BillingPlansPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingPlansPartialUpdateParamsWithContext(ctx context.Context) *BillingPlansPartialUpdateParams {
	var ()
	return &BillingPlansPartialUpdateParams{

		Context: ctx,
	}
}

// NewBillingPlansPartialUpdateParamsWithHTTPClient creates a new BillingPlansPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingPlansPartialUpdateParamsWithHTTPClient(client *http.Client) *BillingPlansPartialUpdateParams {
	var ()
	return &BillingPlansPartialUpdateParams{
		HTTPClient: client,
	}
}

/*BillingPlansPartialUpdateParams contains all the parameters to send to the API endpoint
for the billing plans partial update operation typically these are written to a http.Request
*/
type BillingPlansPartialUpdateParams struct {

	/*Data*/
	Data BillingPlansPartialUpdateBody
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) WithTimeout(timeout time.Duration) *BillingPlansPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) WithContext(ctx context.Context) *BillingPlansPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) WithHTTPClient(client *http.Client) *BillingPlansPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) WithData(data BillingPlansPartialUpdateBody) *BillingPlansPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) SetData(data BillingPlansPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) WithID(id string) *BillingPlansPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) WithNamespace(namespace string) *BillingPlansPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing plans partial update params
func (o *BillingPlansPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingPlansPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
