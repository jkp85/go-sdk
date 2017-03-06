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

// NewProjectsServersCreateParams creates a new ProjectsServersCreateParams object
// with the default values initialized.
func NewProjectsServersCreateParams() *ProjectsServersCreateParams {
	var ()
	return &ProjectsServersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersCreateParamsWithTimeout creates a new ProjectsServersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersCreateParamsWithTimeout(timeout time.Duration) *ProjectsServersCreateParams {
	var ()
	return &ProjectsServersCreateParams{

		timeout: timeout,
	}
}

// NewProjectsServersCreateParamsWithContext creates a new ProjectsServersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersCreateParamsWithContext(ctx context.Context) *ProjectsServersCreateParams {
	var ()
	return &ProjectsServersCreateParams{

		Context: ctx,
	}
}

// NewProjectsServersCreateParamsWithHTTPClient creates a new ProjectsServersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersCreateParamsWithHTTPClient(client *http.Client) *ProjectsServersCreateParams {
	var ()
	return &ProjectsServersCreateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersCreateParams contains all the parameters to send to the API endpoint
for the projects servers create operation typically these are written to a http.Request
*/
type ProjectsServersCreateParams struct {

	/*Data*/
	Data ProjectsServersCreateBody
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers create params
func (o *ProjectsServersCreateParams) WithTimeout(timeout time.Duration) *ProjectsServersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers create params
func (o *ProjectsServersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers create params
func (o *ProjectsServersCreateParams) WithContext(ctx context.Context) *ProjectsServersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers create params
func (o *ProjectsServersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers create params
func (o *ProjectsServersCreateParams) WithHTTPClient(client *http.Client) *ProjectsServersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers create params
func (o *ProjectsServersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects servers create params
func (o *ProjectsServersCreateParams) WithData(data ProjectsServersCreateBody) *ProjectsServersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects servers create params
func (o *ProjectsServersCreateParams) SetData(data ProjectsServersCreateBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects servers create params
func (o *ProjectsServersCreateParams) WithNamespace(namespace string) *ProjectsServersCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers create params
func (o *ProjectsServersCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects servers create params
func (o *ProjectsServersCreateParams) WithProjectPk(projectPk string) *ProjectsServersCreateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers create params
func (o *ProjectsServersCreateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
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