package servers

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

// NewServersOptionsTypesCreateParams creates a new ServersOptionsTypesCreateParams object
// with the default values initialized.
func NewServersOptionsTypesCreateParams() *ServersOptionsTypesCreateParams {
	var ()
	return &ServersOptionsTypesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServersOptionsTypesCreateParamsWithTimeout creates a new ServersOptionsTypesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServersOptionsTypesCreateParamsWithTimeout(timeout time.Duration) *ServersOptionsTypesCreateParams {
	var ()
	return &ServersOptionsTypesCreateParams{

		timeout: timeout,
	}
}

// NewServersOptionsTypesCreateParamsWithContext creates a new ServersOptionsTypesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewServersOptionsTypesCreateParamsWithContext(ctx context.Context) *ServersOptionsTypesCreateParams {
	var ()
	return &ServersOptionsTypesCreateParams{

		Context: ctx,
	}
}

// NewServersOptionsTypesCreateParamsWithHTTPClient creates a new ServersOptionsTypesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServersOptionsTypesCreateParamsWithHTTPClient(client *http.Client) *ServersOptionsTypesCreateParams {
	var ()
	return &ServersOptionsTypesCreateParams{
		HTTPClient: client,
	}
}

/*ServersOptionsTypesCreateParams contains all the parameters to send to the API endpoint
for the servers options types create operation typically these are written to a http.Request
*/
type ServersOptionsTypesCreateParams struct {

	/*Data*/
	Data ServersOptionsTypesCreateBody
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the servers options types create params
func (o *ServersOptionsTypesCreateParams) WithTimeout(timeout time.Duration) *ServersOptionsTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the servers options types create params
func (o *ServersOptionsTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the servers options types create params
func (o *ServersOptionsTypesCreateParams) WithContext(ctx context.Context) *ServersOptionsTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the servers options types create params
func (o *ServersOptionsTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the servers options types create params
func (o *ServersOptionsTypesCreateParams) WithHTTPClient(client *http.Client) *ServersOptionsTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the servers options types create params
func (o *ServersOptionsTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the servers options types create params
func (o *ServersOptionsTypesCreateParams) WithData(data ServersOptionsTypesCreateBody) *ServersOptionsTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the servers options types create params
func (o *ServersOptionsTypesCreateParams) SetData(data ServersOptionsTypesCreateBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the servers options types create params
func (o *ServersOptionsTypesCreateParams) WithNamespace(namespace string) *ServersOptionsTypesCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the servers options types create params
func (o *ServersOptionsTypesCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *ServersOptionsTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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