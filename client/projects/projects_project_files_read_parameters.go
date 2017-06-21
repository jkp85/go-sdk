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

// NewProjectsProjectFilesReadParams creates a new ProjectsProjectFilesReadParams object
// with the default values initialized.
func NewProjectsProjectFilesReadParams() *ProjectsProjectFilesReadParams {
	var ()
	return &ProjectsProjectFilesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectFilesReadParamsWithTimeout creates a new ProjectsProjectFilesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsProjectFilesReadParamsWithTimeout(timeout time.Duration) *ProjectsProjectFilesReadParams {
	var ()
	return &ProjectsProjectFilesReadParams{

		timeout: timeout,
	}
}

// NewProjectsProjectFilesReadParamsWithContext creates a new ProjectsProjectFilesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsProjectFilesReadParamsWithContext(ctx context.Context) *ProjectsProjectFilesReadParams {
	var ()
	return &ProjectsProjectFilesReadParams{

		Context: ctx,
	}
}

// NewProjectsProjectFilesReadParamsWithHTTPClient creates a new ProjectsProjectFilesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsProjectFilesReadParamsWithHTTPClient(client *http.Client) *ProjectsProjectFilesReadParams {
	var ()
	return &ProjectsProjectFilesReadParams{
		HTTPClient: client,
	}
}

/*ProjectsProjectFilesReadParams contains all the parameters to send to the API endpoint
for the projects project files read operation typically these are written to a http.Request
*/
type ProjectsProjectFilesReadParams struct {

	/*ID*/
	ID string
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithTimeout(timeout time.Duration) *ProjectsProjectFilesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithContext(ctx context.Context) *ProjectsProjectFilesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithHTTPClient(client *http.Client) *ProjectsProjectFilesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithID(id string) *ProjectsProjectFilesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithNamespace(namespace string) *ProjectsProjectFilesReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithProjectPk(projectPk string) *ProjectsProjectFilesReadParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectFilesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_pk
	if err := r.SetPathParam("project_pk", o.ProjectPk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}