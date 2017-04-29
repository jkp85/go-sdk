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

// NewTriggersDeleteParams creates a new TriggersDeleteParams object
// with the default values initialized.
func NewTriggersDeleteParams() *TriggersDeleteParams {
	var ()
	return &TriggersDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTriggersDeleteParamsWithTimeout creates a new TriggersDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTriggersDeleteParamsWithTimeout(timeout time.Duration) *TriggersDeleteParams {
	var ()
	return &TriggersDeleteParams{

		timeout: timeout,
	}
}

// NewTriggersDeleteParamsWithContext creates a new TriggersDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTriggersDeleteParamsWithContext(ctx context.Context) *TriggersDeleteParams {
	var ()
	return &TriggersDeleteParams{

		Context: ctx,
	}
}

// NewTriggersDeleteParamsWithHTTPClient creates a new TriggersDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTriggersDeleteParamsWithHTTPClient(client *http.Client) *TriggersDeleteParams {
	var ()
	return &TriggersDeleteParams{
		HTTPClient: client,
	}
}

/*TriggersDeleteParams contains all the parameters to send to the API endpoint
for the triggers delete operation typically these are written to a http.Request
*/
type TriggersDeleteParams struct {

	/*ID*/
	ID string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the triggers delete params
func (o *TriggersDeleteParams) WithTimeout(timeout time.Duration) *TriggersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the triggers delete params
func (o *TriggersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the triggers delete params
func (o *TriggersDeleteParams) WithContext(ctx context.Context) *TriggersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the triggers delete params
func (o *TriggersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the triggers delete params
func (o *TriggersDeleteParams) WithHTTPClient(client *http.Client) *TriggersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the triggers delete params
func (o *TriggersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the triggers delete params
func (o *TriggersDeleteParams) WithID(id string) *TriggersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the triggers delete params
func (o *TriggersDeleteParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the triggers delete params
func (o *TriggersDeleteParams) WithNamespace(namespace string) *TriggersDeleteParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the triggers delete params
func (o *TriggersDeleteParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *TriggersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

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
