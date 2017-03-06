package auth

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

// NewConvertTokenCreateParams creates a new ConvertTokenCreateParams object
// with the default values initialized.
func NewConvertTokenCreateParams() *ConvertTokenCreateParams {

	return &ConvertTokenCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConvertTokenCreateParamsWithTimeout creates a new ConvertTokenCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConvertTokenCreateParamsWithTimeout(timeout time.Duration) *ConvertTokenCreateParams {

	return &ConvertTokenCreateParams{

		timeout: timeout,
	}
}

// NewConvertTokenCreateParamsWithContext creates a new ConvertTokenCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewConvertTokenCreateParamsWithContext(ctx context.Context) *ConvertTokenCreateParams {

	return &ConvertTokenCreateParams{

		Context: ctx,
	}
}

// NewConvertTokenCreateParamsWithHTTPClient creates a new ConvertTokenCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConvertTokenCreateParamsWithHTTPClient(client *http.Client) *ConvertTokenCreateParams {

	return &ConvertTokenCreateParams{
		HTTPClient: client,
	}
}

/*ConvertTokenCreateParams contains all the parameters to send to the API endpoint
for the convert token create operation typically these are written to a http.Request
*/
type ConvertTokenCreateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the convert token create params
func (o *ConvertTokenCreateParams) WithTimeout(timeout time.Duration) *ConvertTokenCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the convert token create params
func (o *ConvertTokenCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the convert token create params
func (o *ConvertTokenCreateParams) WithContext(ctx context.Context) *ConvertTokenCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the convert token create params
func (o *ConvertTokenCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the convert token create params
func (o *ConvertTokenCreateParams) WithHTTPClient(client *http.Client) *ConvertTokenCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the convert token create params
func (o *ConvertTokenCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ConvertTokenCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
