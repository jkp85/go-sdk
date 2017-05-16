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

// NewTriggersPartialUpdateParams creates a new TriggersPartialUpdateParams object
// with the default values initialized.
func NewTriggersPartialUpdateParams() *TriggersPartialUpdateParams {
	var ()
	return &TriggersPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTriggersPartialUpdateParamsWithTimeout creates a new TriggersPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTriggersPartialUpdateParamsWithTimeout(timeout time.Duration) *TriggersPartialUpdateParams {
	var ()
	return &TriggersPartialUpdateParams{

		timeout: timeout,
	}
}

// NewTriggersPartialUpdateParamsWithContext creates a new TriggersPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTriggersPartialUpdateParamsWithContext(ctx context.Context) *TriggersPartialUpdateParams {
	var ()
	return &TriggersPartialUpdateParams{

		Context: ctx,
	}
}

// NewTriggersPartialUpdateParamsWithHTTPClient creates a new TriggersPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTriggersPartialUpdateParamsWithHTTPClient(client *http.Client) *TriggersPartialUpdateParams {
	var ()
	return &TriggersPartialUpdateParams{
		HTTPClient: client,
	}
}

/*TriggersPartialUpdateParams contains all the parameters to send to the API endpoint
for the triggers partial update operation typically these are written to a http.Request
*/
type TriggersPartialUpdateParams struct {

	/*Data*/
	Data TriggersPartialUpdateBody
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the triggers partial update params
func (o *TriggersPartialUpdateParams) WithTimeout(timeout time.Duration) *TriggersPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the triggers partial update params
func (o *TriggersPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the triggers partial update params
func (o *TriggersPartialUpdateParams) WithContext(ctx context.Context) *TriggersPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the triggers partial update params
func (o *TriggersPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the triggers partial update params
func (o *TriggersPartialUpdateParams) WithHTTPClient(client *http.Client) *TriggersPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the triggers partial update params
func (o *TriggersPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the triggers partial update params
func (o *TriggersPartialUpdateParams) WithData(data TriggersPartialUpdateBody) *TriggersPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the triggers partial update params
func (o *TriggersPartialUpdateParams) SetData(data TriggersPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the triggers partial update params
func (o *TriggersPartialUpdateParams) WithID(id string) *TriggersPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the triggers partial update params
func (o *TriggersPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the triggers partial update params
func (o *TriggersPartialUpdateParams) WithNamespace(namespace string) *TriggersPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the triggers partial update params
func (o *TriggersPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *TriggersPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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