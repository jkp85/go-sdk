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

// NewProjectsServersStopCreateParams creates a new ProjectsServersStopCreateParams object
// with the default values initialized.
func NewProjectsServersStopCreateParams() *ProjectsServersStopCreateParams {
	var ()
	return &ProjectsServersStopCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersStopCreateParamsWithTimeout creates a new ProjectsServersStopCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersStopCreateParamsWithTimeout(timeout time.Duration) *ProjectsServersStopCreateParams {
	var ()
	return &ProjectsServersStopCreateParams{

		timeout: timeout,
	}
}

// NewProjectsServersStopCreateParamsWithContext creates a new ProjectsServersStopCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersStopCreateParamsWithContext(ctx context.Context) *ProjectsServersStopCreateParams {
	var ()
	return &ProjectsServersStopCreateParams{

		Context: ctx,
	}
}

// NewProjectsServersStopCreateParamsWithHTTPClient creates a new ProjectsServersStopCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersStopCreateParamsWithHTTPClient(client *http.Client) *ProjectsServersStopCreateParams {
	var ()
	return &ProjectsServersStopCreateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersStopCreateParams contains all the parameters to send to the API endpoint
for the projects servers stop create operation typically these are written to a http.Request
*/
type ProjectsServersStopCreateParams struct {

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

// WithTimeout adds the timeout to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) WithTimeout(timeout time.Duration) *ProjectsServersStopCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) WithContext(ctx context.Context) *ProjectsServersStopCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) WithHTTPClient(client *http.Client) *ProjectsServersStopCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) WithNamespace(namespace string) *ProjectsServersStopCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) WithProjectPk(projectPk string) *ProjectsServersStopCreateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) WithServerPk(serverPk string) *ProjectsServersStopCreateParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects servers stop create params
func (o *ProjectsServersStopCreateParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersStopCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

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