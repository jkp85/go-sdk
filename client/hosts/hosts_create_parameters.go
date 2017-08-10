// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

	"github.com/3Blades/go-sdk/models"
)

// NewHostsCreateParams creates a new HostsCreateParams object
// with the default values initialized.
func NewHostsCreateParams() *HostsCreateParams {
	var ()
	return &HostsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHostsCreateParamsWithTimeout creates a new HostsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHostsCreateParamsWithTimeout(timeout time.Duration) *HostsCreateParams {
	var ()
	return &HostsCreateParams{

		timeout: timeout,
	}
}

// NewHostsCreateParamsWithContext creates a new HostsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewHostsCreateParamsWithContext(ctx context.Context) *HostsCreateParams {
	var ()
	return &HostsCreateParams{

		Context: ctx,
	}
}

// NewHostsCreateParamsWithHTTPClient creates a new HostsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHostsCreateParamsWithHTTPClient(client *http.Client) *HostsCreateParams {
	var ()
	return &HostsCreateParams{
		HTTPClient: client,
	}
}

/*HostsCreateParams contains all the parameters to send to the API endpoint
for the hosts create operation typically these are written to a http.Request
*/
type HostsCreateParams struct {

	/*DockerhostData*/
	DockerhostData *models.DockerHostData
	/*Namespace
	  User or team name.

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hosts create params
func (o *HostsCreateParams) WithTimeout(timeout time.Duration) *HostsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts create params
func (o *HostsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts create params
func (o *HostsCreateParams) WithContext(ctx context.Context) *HostsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts create params
func (o *HostsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts create params
func (o *HostsCreateParams) WithHTTPClient(client *http.Client) *HostsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts create params
func (o *HostsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDockerhostData adds the dockerhostData to the hosts create params
func (o *HostsCreateParams) WithDockerhostData(dockerhostData *models.DockerHostData) *HostsCreateParams {
	o.SetDockerhostData(dockerhostData)
	return o
}

// SetDockerhostData adds the dockerhostData to the hosts create params
func (o *HostsCreateParams) SetDockerhostData(dockerhostData *models.DockerHostData) {
	o.DockerhostData = dockerhostData
}

// WithNamespace adds the namespace to the hosts create params
func (o *HostsCreateParams) WithNamespace(namespace string) *HostsCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the hosts create params
func (o *HostsCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *HostsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DockerhostData == nil {
		o.DockerhostData = new(models.DockerHostData)
	}

	if err := r.SetBodyParam(o.DockerhostData); err != nil {
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
