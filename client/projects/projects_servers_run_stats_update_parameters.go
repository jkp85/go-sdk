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

// NewProjectsServersRunStatsUpdateParams creates a new ProjectsServersRunStatsUpdateParams object
// with the default values initialized.
func NewProjectsServersRunStatsUpdateParams() *ProjectsServersRunStatsUpdateParams {
	var ()
	return &ProjectsServersRunStatsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersRunStatsUpdateParamsWithTimeout creates a new ProjectsServersRunStatsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersRunStatsUpdateParamsWithTimeout(timeout time.Duration) *ProjectsServersRunStatsUpdateParams {
	var ()
	return &ProjectsServersRunStatsUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsServersRunStatsUpdateParamsWithContext creates a new ProjectsServersRunStatsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersRunStatsUpdateParamsWithContext(ctx context.Context) *ProjectsServersRunStatsUpdateParams {
	var ()
	return &ProjectsServersRunStatsUpdateParams{

		Context: ctx,
	}
}

// NewProjectsServersRunStatsUpdateParamsWithHTTPClient creates a new ProjectsServersRunStatsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersRunStatsUpdateParamsWithHTTPClient(client *http.Client) *ProjectsServersRunStatsUpdateParams {
	var ()
	return &ProjectsServersRunStatsUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersRunStatsUpdateParams contains all the parameters to send to the API endpoint
for the projects servers run stats update operation typically these are written to a http.Request
*/
type ProjectsServersRunStatsUpdateParams struct {

	/*Data*/
	Data ProjectsServersRunStatsUpdateBody
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

// WithTimeout adds the timeout to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) WithTimeout(timeout time.Duration) *ProjectsServersRunStatsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) WithContext(ctx context.Context) *ProjectsServersRunStatsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) WithHTTPClient(client *http.Client) *ProjectsServersRunStatsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) WithData(data ProjectsServersRunStatsUpdateBody) *ProjectsServersRunStatsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) SetData(data ProjectsServersRunStatsUpdateBody) {
	o.Data = data
}

// WithID adds the id to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) WithID(id string) *ProjectsServersRunStatsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) WithNamespace(namespace string) *ProjectsServersRunStatsUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) WithProjectPk(projectPk string) *ProjectsServersRunStatsUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) WithServerPk(serverPk string) *ProjectsServersRunStatsUpdateParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects servers run stats update params
func (o *ProjectsServersRunStatsUpdateParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersRunStatsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

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
