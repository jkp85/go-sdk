package triggers

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

// NewTriggersCreateParams creates a new TriggersCreateParams object
// with the default values initialized.
func NewTriggersCreateParams() *TriggersCreateParams {
	var ()
	return &TriggersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTriggersCreateParamsWithTimeout creates a new TriggersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTriggersCreateParamsWithTimeout(timeout time.Duration) *TriggersCreateParams {
	var ()
	return &TriggersCreateParams{

		timeout: timeout,
	}
}

// NewTriggersCreateParamsWithContext creates a new TriggersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTriggersCreateParamsWithContext(ctx context.Context) *TriggersCreateParams {
	var ()
	return &TriggersCreateParams{

		Context: ctx,
	}
}

// NewTriggersCreateParamsWithHTTPClient creates a new TriggersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTriggersCreateParamsWithHTTPClient(client *http.Client) *TriggersCreateParams {
	var ()
	return &TriggersCreateParams{
		HTTPClient: client,
	}
}

/*TriggersCreateParams contains all the parameters to send to the API endpoint
for the triggers create operation typically these are written to a http.Request
*/
type TriggersCreateParams struct {

	/*Data*/
	Data TriggersCreateBody
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the triggers create params
func (o *TriggersCreateParams) WithTimeout(timeout time.Duration) *TriggersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the triggers create params
func (o *TriggersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the triggers create params
func (o *TriggersCreateParams) WithContext(ctx context.Context) *TriggersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the triggers create params
func (o *TriggersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the triggers create params
func (o *TriggersCreateParams) WithHTTPClient(client *http.Client) *TriggersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the triggers create params
func (o *TriggersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the triggers create params
func (o *TriggersCreateParams) WithData(data TriggersCreateBody) *TriggersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the triggers create params
func (o *TriggersCreateParams) SetData(data TriggersCreateBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the triggers create params
func (o *TriggersCreateParams) WithNamespace(namespace string) *TriggersCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the triggers create params
func (o *TriggersCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *TriggersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
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
