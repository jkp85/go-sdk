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

// NewBillingSubscriptionRequiredListParams creates a new BillingSubscriptionRequiredListParams object
// with the default values initialized.
func NewBillingSubscriptionRequiredListParams() *BillingSubscriptionRequiredListParams {
	var ()
	return &BillingSubscriptionRequiredListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingSubscriptionRequiredListParamsWithTimeout creates a new BillingSubscriptionRequiredListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingSubscriptionRequiredListParamsWithTimeout(timeout time.Duration) *BillingSubscriptionRequiredListParams {
	var ()
	return &BillingSubscriptionRequiredListParams{

		timeout: timeout,
	}
}

// NewBillingSubscriptionRequiredListParamsWithContext creates a new BillingSubscriptionRequiredListParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingSubscriptionRequiredListParamsWithContext(ctx context.Context) *BillingSubscriptionRequiredListParams {
	var ()
	return &BillingSubscriptionRequiredListParams{

		Context: ctx,
	}
}

// NewBillingSubscriptionRequiredListParamsWithHTTPClient creates a new BillingSubscriptionRequiredListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingSubscriptionRequiredListParamsWithHTTPClient(client *http.Client) *BillingSubscriptionRequiredListParams {
	var ()
	return &BillingSubscriptionRequiredListParams{
		HTTPClient: client,
	}
}

/*BillingSubscriptionRequiredListParams contains all the parameters to send to the API endpoint
for the billing subscription required list operation typically these are written to a http.Request
*/
type BillingSubscriptionRequiredListParams struct {

	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billing subscription required list params
func (o *BillingSubscriptionRequiredListParams) WithTimeout(timeout time.Duration) *BillingSubscriptionRequiredListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing subscription required list params
func (o *BillingSubscriptionRequiredListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing subscription required list params
func (o *BillingSubscriptionRequiredListParams) WithContext(ctx context.Context) *BillingSubscriptionRequiredListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing subscription required list params
func (o *BillingSubscriptionRequiredListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing subscription required list params
func (o *BillingSubscriptionRequiredListParams) WithHTTPClient(client *http.Client) *BillingSubscriptionRequiredListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing subscription required list params
func (o *BillingSubscriptionRequiredListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the billing subscription required list params
func (o *BillingSubscriptionRequiredListParams) WithNamespace(namespace string) *BillingSubscriptionRequiredListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing subscription required list params
func (o *BillingSubscriptionRequiredListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingSubscriptionRequiredListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
