// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsServersAuthParams creates a new ProjectsServersAuthParams object
// with the default values initialized.
func NewProjectsServersAuthParams() *ProjectsServersAuthParams {
	var ()
	return &ProjectsServersAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersAuthParamsWithTimeout creates a new ProjectsServersAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersAuthParamsWithTimeout(timeout time.Duration) *ProjectsServersAuthParams {
	var ()
	return &ProjectsServersAuthParams{

		timeout: timeout,
	}
}

// NewProjectsServersAuthParamsWithContext creates a new ProjectsServersAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersAuthParamsWithContext(ctx context.Context) *ProjectsServersAuthParams {
	var ()
	return &ProjectsServersAuthParams{

		Context: ctx,
	}
}

// NewProjectsServersAuthParamsWithHTTPClient creates a new ProjectsServersAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersAuthParamsWithHTTPClient(client *http.Client) *ProjectsServersAuthParams {
	var ()
	return &ProjectsServersAuthParams{
		HTTPClient: client,
	}
}

/*ProjectsServersAuthParams contains all the parameters to send to the API endpoint
for the projects servers auth operation typically these are written to a http.Request
*/
type ProjectsServersAuthParams struct {

	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Project
	  Project unique identifier expressed as UUID or name.

	*/
	Project string
	/*Server
	  Server unique identifier expressed as UUID or name.

	*/
	Server string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers auth params
func (o *ProjectsServersAuthParams) WithTimeout(timeout time.Duration) *ProjectsServersAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers auth params
func (o *ProjectsServersAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers auth params
func (o *ProjectsServersAuthParams) WithContext(ctx context.Context) *ProjectsServersAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers auth params
func (o *ProjectsServersAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers auth params
func (o *ProjectsServersAuthParams) WithHTTPClient(client *http.Client) *ProjectsServersAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers auth params
func (o *ProjectsServersAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects servers auth params
func (o *ProjectsServersAuthParams) WithNamespace(namespace string) *ProjectsServersAuthParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers auth params
func (o *ProjectsServersAuthParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects servers auth params
func (o *ProjectsServersAuthParams) WithProject(project string) *ProjectsServersAuthParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects servers auth params
func (o *ProjectsServersAuthParams) SetProject(project string) {
	o.Project = project
}

// WithServer adds the server to the projects servers auth params
func (o *ProjectsServersAuthParams) WithServer(server string) *ProjectsServersAuthParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects servers auth params
func (o *ProjectsServersAuthParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param server
	if err := r.SetPathParam("server", o.Server); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
