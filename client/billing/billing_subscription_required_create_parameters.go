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

// NewBillingSubscriptionRequiredCreateParams creates a new BillingSubscriptionRequiredCreateParams object
// with the default values initialized.
func NewBillingSubscriptionRequiredCreateParams() *BillingSubscriptionRequiredCreateParams {
	var ()
	return &BillingSubscriptionRequiredCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingSubscriptionRequiredCreateParamsWithTimeout creates a new BillingSubscriptionRequiredCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingSubscriptionRequiredCreateParamsWithTimeout(timeout time.Duration) *BillingSubscriptionRequiredCreateParams {
	var ()
	return &BillingSubscriptionRequiredCreateParams{

		timeout: timeout,
	}
}

// NewBillingSubscriptionRequiredCreateParamsWithContext creates a new BillingSubscriptionRequiredCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingSubscriptionRequiredCreateParamsWithContext(ctx context.Context) *BillingSubscriptionRequiredCreateParams {
	var ()
	return &BillingSubscriptionRequiredCreateParams{

		Context: ctx,
	}
}

// NewBillingSubscriptionRequiredCreateParamsWithHTTPClient creates a new BillingSubscriptionRequiredCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingSubscriptionRequiredCreateParamsWithHTTPClient(client *http.Client) *BillingSubscriptionRequiredCreateParams {
	var ()
	return &BillingSubscriptionRequiredCreateParams{
		HTTPClient: client,
	}
}

/*BillingSubscriptionRequiredCreateParams contains all the parameters to send to the API endpoint
for the billing subscription required create operation typically these are written to a http.Request
*/
type BillingSubscriptionRequiredCreateParams struct {

	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billing subscription required create params
func (o *BillingSubscriptionRequiredCreateParams) WithTimeout(timeout time.Duration) *BillingSubscriptionRequiredCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing subscription required create params
func (o *BillingSubscriptionRequiredCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing subscription required create params
func (o *BillingSubscriptionRequiredCreateParams) WithContext(ctx context.Context) *BillingSubscriptionRequiredCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing subscription required create params
func (o *BillingSubscriptionRequiredCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing subscription required create params
func (o *BillingSubscriptionRequiredCreateParams) WithHTTPClient(client *http.Client) *BillingSubscriptionRequiredCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing subscription required create params
func (o *BillingSubscriptionRequiredCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the billing subscription required create params
func (o *BillingSubscriptionRequiredCreateParams) WithNamespace(namespace string) *BillingSubscriptionRequiredCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing subscription required create params
func (o *BillingSubscriptionRequiredCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingSubscriptionRequiredCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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