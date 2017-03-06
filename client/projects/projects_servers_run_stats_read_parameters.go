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

// NewProjectsServersRunStatsReadParams creates a new ProjectsServersRunStatsReadParams object
// with the default values initialized.
func NewProjectsServersRunStatsReadParams() *ProjectsServersRunStatsReadParams {
	var ()
	return &ProjectsServersRunStatsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersRunStatsReadParamsWithTimeout creates a new ProjectsServersRunStatsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersRunStatsReadParamsWithTimeout(timeout time.Duration) *ProjectsServersRunStatsReadParams {
	var ()
	return &ProjectsServersRunStatsReadParams{

		timeout: timeout,
	}
}

// NewProjectsServersRunStatsReadParamsWithContext creates a new ProjectsServersRunStatsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersRunStatsReadParamsWithContext(ctx context.Context) *ProjectsServersRunStatsReadParams {
	var ()
	return &ProjectsServersRunStatsReadParams{

		Context: ctx,
	}
}

// NewProjectsServersRunStatsReadParamsWithHTTPClient creates a new ProjectsServersRunStatsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersRunStatsReadParamsWithHTTPClient(client *http.Client) *ProjectsServersRunStatsReadParams {
	var ()
	return &ProjectsServersRunStatsReadParams{
		HTTPClient: client,
	}
}

/*ProjectsServersRunStatsReadParams contains all the parameters to send to the API endpoint
for the projects servers run stats read operation typically these are written to a http.Request
*/
type ProjectsServersRunStatsReadParams struct {

	/*ID*/
	ID string
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string
	/*ServerPk*/
	ServerPk string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) WithTimeout(timeout time.Duration) *ProjectsServersRunStatsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) WithContext(ctx context.Context) *ProjectsServersRunStatsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) WithHTTPClient(client *http.Client) *ProjectsServersRunStatsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) WithID(id string) *ProjectsServersRunStatsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) WithNamespace(namespace string) *ProjectsServersRunStatsReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) WithProjectPk(projectPk string) *ProjectsServersRunStatsReadParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) WithServerPk(serverPk string) *ProjectsServersRunStatsReadParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects servers run stats read params
func (o *ProjectsServersRunStatsReadParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersRunStatsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_pk
	if err := r.SetPathParam("project_pk", o.ProjectPk); err != nil {
		return err
	}

	// path param server_pk
	if err := r.SetPathParam("server_pk", o.ServerPk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}