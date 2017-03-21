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

// NewProjectsServersTerminateCreateParams creates a new ProjectsServersTerminateCreateParams object
// with the default values initialized.
func NewProjectsServersTerminateCreateParams() *ProjectsServersTerminateCreateParams {
	var ()
	return &ProjectsServersTerminateCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersTerminateCreateParamsWithTimeout creates a new ProjectsServersTerminateCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersTerminateCreateParamsWithTimeout(timeout time.Duration) *ProjectsServersTerminateCreateParams {
	var ()
	return &ProjectsServersTerminateCreateParams{

		timeout: timeout,
	}
}

// NewProjectsServersTerminateCreateParamsWithContext creates a new ProjectsServersTerminateCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersTerminateCreateParamsWithContext(ctx context.Context) *ProjectsServersTerminateCreateParams {
	var ()
	return &ProjectsServersTerminateCreateParams{

		Context: ctx,
	}
}

// NewProjectsServersTerminateCreateParamsWithHTTPClient creates a new ProjectsServersTerminateCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersTerminateCreateParamsWithHTTPClient(client *http.Client) *ProjectsServersTerminateCreateParams {
	var ()
	return &ProjectsServersTerminateCreateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersTerminateCreateParams contains all the parameters to send to the API endpoint
for the projects servers terminate create operation typically these are written to a http.Request
*/
type ProjectsServersTerminateCreateParams struct {

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

// WithTimeout adds the timeout to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) WithTimeout(timeout time.Duration) *ProjectsServersTerminateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) WithContext(ctx context.Context) *ProjectsServersTerminateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) WithHTTPClient(client *http.Client) *ProjectsServersTerminateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) WithNamespace(namespace string) *ProjectsServersTerminateCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) WithProjectPk(projectPk string) *ProjectsServersTerminateCreateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) WithServerPk(serverPk string) *ProjectsServersTerminateCreateParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects servers terminate create params
func (o *ProjectsServersTerminateCreateParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersTerminateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
