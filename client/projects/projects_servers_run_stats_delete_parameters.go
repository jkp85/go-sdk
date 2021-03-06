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

// NewProjectsServersRunStatsDeleteParams creates a new ProjectsServersRunStatsDeleteParams object
// with the default values initialized.
func NewProjectsServersRunStatsDeleteParams() *ProjectsServersRunStatsDeleteParams {
	var ()
	return &ProjectsServersRunStatsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersRunStatsDeleteParamsWithTimeout creates a new ProjectsServersRunStatsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersRunStatsDeleteParamsWithTimeout(timeout time.Duration) *ProjectsServersRunStatsDeleteParams {
	var ()
	return &ProjectsServersRunStatsDeleteParams{

		timeout: timeout,
	}
}

// NewProjectsServersRunStatsDeleteParamsWithContext creates a new ProjectsServersRunStatsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersRunStatsDeleteParamsWithContext(ctx context.Context) *ProjectsServersRunStatsDeleteParams {
	var ()
	return &ProjectsServersRunStatsDeleteParams{

		Context: ctx,
	}
}

// NewProjectsServersRunStatsDeleteParamsWithHTTPClient creates a new ProjectsServersRunStatsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersRunStatsDeleteParamsWithHTTPClient(client *http.Client) *ProjectsServersRunStatsDeleteParams {
	var ()
	return &ProjectsServersRunStatsDeleteParams{
		HTTPClient: client,
	}
}

/*ProjectsServersRunStatsDeleteParams contains all the parameters to send to the API endpoint
for the projects servers run stats delete operation typically these are written to a http.Request
*/
type ProjectsServersRunStatsDeleteParams struct {

	/*ID
	  Server run statistics unique identifier expressed as UUID.

	*/
	ID string
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

// WithTimeout adds the timeout to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) WithTimeout(timeout time.Duration) *ProjectsServersRunStatsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) WithContext(ctx context.Context) *ProjectsServersRunStatsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) WithHTTPClient(client *http.Client) *ProjectsServersRunStatsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) WithID(id string) *ProjectsServersRunStatsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) WithNamespace(namespace string) *ProjectsServersRunStatsDeleteParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) WithProject(project string) *ProjectsServersRunStatsDeleteParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) SetProject(project string) {
	o.Project = project
}

// WithServer adds the server to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) WithServer(server string) *ProjectsServersRunStatsDeleteParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects servers run stats delete params
func (o *ProjectsServersRunStatsDeleteParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersRunStatsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

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
