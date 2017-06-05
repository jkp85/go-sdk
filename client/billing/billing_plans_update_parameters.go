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

// NewBillingPlansUpdateParams creates a new BillingPlansUpdateParams object
// with the default values initialized.
func NewBillingPlansUpdateParams() *BillingPlansUpdateParams {
	var ()
	return &BillingPlansUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingPlansUpdateParamsWithTimeout creates a new BillingPlansUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingPlansUpdateParamsWithTimeout(timeout time.Duration) *BillingPlansUpdateParams {
	var ()
	return &BillingPlansUpdateParams{

		timeout: timeout,
	}
}

// NewBillingPlansUpdateParamsWithContext creates a new BillingPlansUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingPlansUpdateParamsWithContext(ctx context.Context) *BillingPlansUpdateParams {
	var ()
	return &BillingPlansUpdateParams{

		Context: ctx,
	}
}

// NewBillingPlansUpdateParamsWithHTTPClient creates a new BillingPlansUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingPlansUpdateParamsWithHTTPClient(client *http.Client) *BillingPlansUpdateParams {
	var ()
	return &BillingPlansUpdateParams{
		HTTPClient: client,
	}
}

/*BillingPlansUpdateParams contains all the parameters to send to the API endpoint
for the billing plans update operation typically these are written to a http.Request
*/
type BillingPlansUpdateParams struct {

	/*Data*/
	Data BillingPlansUpdateBody
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billing plans update params
func (o *BillingPlansUpdateParams) WithTimeout(timeout time.Duration) *BillingPlansUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing plans update params
func (o *BillingPlansUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing plans update params
func (o *BillingPlansUpdateParams) WithContext(ctx context.Context) *BillingPlansUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing plans update params
func (o *BillingPlansUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing plans update params
func (o *BillingPlansUpdateParams) WithHTTPClient(client *http.Client) *BillingPlansUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing plans update params
func (o *BillingPlansUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the billing plans update params
func (o *BillingPlansUpdateParams) WithData(data BillingPlansUpdateBody) *BillingPlansUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the billing plans update params
func (o *BillingPlansUpdateParams) SetData(data BillingPlansUpdateBody) {
	o.Data = data
}

// WithID adds the id to the billing plans update params
func (o *BillingPlansUpdateParams) WithID(id string) *BillingPlansUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billing plans update params
func (o *BillingPlansUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the billing plans update params
func (o *BillingPlansUpdateParams) WithNamespace(namespace string) *BillingPlansUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing plans update params
func (o *BillingPlansUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingPlansUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
