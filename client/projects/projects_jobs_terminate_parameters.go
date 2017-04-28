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

// NewProjectsJobsTerminateParams creates a new ProjectsJobsTerminateParams object
// with the default values initialized.
func NewProjectsJobsTerminateParams() *ProjectsJobsTerminateParams {
	var ()
	return &ProjectsJobsTerminateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsJobsTerminateParamsWithTimeout creates a new ProjectsJobsTerminateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsJobsTerminateParamsWithTimeout(timeout time.Duration) *ProjectsJobsTerminateParams {
	var ()
	return &ProjectsJobsTerminateParams{

		timeout: timeout,
	}
}

// NewProjectsJobsTerminateParamsWithContext creates a new ProjectsJobsTerminateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsJobsTerminateParamsWithContext(ctx context.Context) *ProjectsJobsTerminateParams {
	var ()
	return &ProjectsJobsTerminateParams{

		Context: ctx,
	}
}

// NewProjectsJobsTerminateParamsWithHTTPClient creates a new ProjectsJobsTerminateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsJobsTerminateParamsWithHTTPClient(client *http.Client) *ProjectsJobsTerminateParams {
	var ()
	return &ProjectsJobsTerminateParams{
		HTTPClient: client,
	}
}

/*ProjectsJobsTerminateParams contains all the parameters to send to the API endpoint
for the projects jobs terminate operation typically these are written to a http.Request
*/
type ProjectsJobsTerminateParams struct {

	/*Data*/
	Data ProjectsJobsTerminateBody
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string
	/*Server*/
	Server string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) WithTimeout(timeout time.Duration) *ProjectsJobsTerminateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) WithContext(ctx context.Context) *ProjectsJobsTerminateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) WithHTTPClient(client *http.Client) *ProjectsJobsTerminateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) WithData(data ProjectsJobsTerminateBody) *ProjectsJobsTerminateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) SetData(data ProjectsJobsTerminateBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) WithNamespace(namespace string) *ProjectsJobsTerminateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) WithProjectPk(projectPk string) *ProjectsJobsTerminateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServer adds the server to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) WithServer(server string) *ProjectsJobsTerminateParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects jobs terminate params
func (o *ProjectsJobsTerminateParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsJobsTerminateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param server
	if err := r.SetPathParam("server", o.Server); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}