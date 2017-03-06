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

// NewProjectsCollaboratorsCreateParams creates a new ProjectsCollaboratorsCreateParams object
// with the default values initialized.
func NewProjectsCollaboratorsCreateParams() *ProjectsCollaboratorsCreateParams {
	var ()
	return &ProjectsCollaboratorsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsCollaboratorsCreateParamsWithTimeout creates a new ProjectsCollaboratorsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsCollaboratorsCreateParamsWithTimeout(timeout time.Duration) *ProjectsCollaboratorsCreateParams {
	var ()
	return &ProjectsCollaboratorsCreateParams{

		timeout: timeout,
	}
}

// NewProjectsCollaboratorsCreateParamsWithContext creates a new ProjectsCollaboratorsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsCollaboratorsCreateParamsWithContext(ctx context.Context) *ProjectsCollaboratorsCreateParams {
	var ()
	return &ProjectsCollaboratorsCreateParams{

		Context: ctx,
	}
}

// NewProjectsCollaboratorsCreateParamsWithHTTPClient creates a new ProjectsCollaboratorsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsCollaboratorsCreateParamsWithHTTPClient(client *http.Client) *ProjectsCollaboratorsCreateParams {
	var ()
	return &ProjectsCollaboratorsCreateParams{
		HTTPClient: client,
	}
}

/*ProjectsCollaboratorsCreateParams contains all the parameters to send to the API endpoint
for the projects collaborators create operation typically these are written to a http.Request
*/
type ProjectsCollaboratorsCreateParams struct {

	/*Data*/
	Data ProjectsCollaboratorsCreateBody
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) WithTimeout(timeout time.Duration) *ProjectsCollaboratorsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) WithContext(ctx context.Context) *ProjectsCollaboratorsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) WithHTTPClient(client *http.Client) *ProjectsCollaboratorsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) WithData(data ProjectsCollaboratorsCreateBody) *ProjectsCollaboratorsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) SetData(data ProjectsCollaboratorsCreateBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) WithNamespace(namespace string) *ProjectsCollaboratorsCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) WithProjectPk(projectPk string) *ProjectsCollaboratorsCreateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects collaborators create params
func (o *ProjectsCollaboratorsCreateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsCollaboratorsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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