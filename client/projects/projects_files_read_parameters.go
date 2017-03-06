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

// NewProjectsFilesReadParams creates a new ProjectsFilesReadParams object
// with the default values initialized.
func NewProjectsFilesReadParams() *ProjectsFilesReadParams {
	var ()
	return &ProjectsFilesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsFilesReadParamsWithTimeout creates a new ProjectsFilesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsFilesReadParamsWithTimeout(timeout time.Duration) *ProjectsFilesReadParams {
	var ()
	return &ProjectsFilesReadParams{

		timeout: timeout,
	}
}

// NewProjectsFilesReadParamsWithContext creates a new ProjectsFilesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsFilesReadParamsWithContext(ctx context.Context) *ProjectsFilesReadParams {
	var ()
	return &ProjectsFilesReadParams{

		Context: ctx,
	}
}

// NewProjectsFilesReadParamsWithHTTPClient creates a new ProjectsFilesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsFilesReadParamsWithHTTPClient(client *http.Client) *ProjectsFilesReadParams {
	var ()
	return &ProjectsFilesReadParams{
		HTTPClient: client,
	}
}

/*ProjectsFilesReadParams contains all the parameters to send to the API endpoint
for the projects files read operation typically these are written to a http.Request
*/
type ProjectsFilesReadParams struct {

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

// WithTimeout adds the timeout to the projects files read params
func (o *ProjectsFilesReadParams) WithTimeout(timeout time.Duration) *ProjectsFilesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects files read params
func (o *ProjectsFilesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects files read params
func (o *ProjectsFilesReadParams) WithContext(ctx context.Context) *ProjectsFilesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects files read params
func (o *ProjectsFilesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects files read params
func (o *ProjectsFilesReadParams) WithHTTPClient(client *http.Client) *ProjectsFilesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects files read params
func (o *ProjectsFilesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects files read params
func (o *ProjectsFilesReadParams) WithID(id string) *ProjectsFilesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects files read params
func (o *ProjectsFilesReadParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects files read params
func (o *ProjectsFilesReadParams) WithNamespace(namespace string) *ProjectsFilesReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects files read params
func (o *ProjectsFilesReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects files read params
func (o *ProjectsFilesReadParams) WithProjectPk(projectPk string) *ProjectsFilesReadParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects files read params
func (o *ProjectsFilesReadParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsFilesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}