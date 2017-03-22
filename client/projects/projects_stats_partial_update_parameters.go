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

// NewProjectsStatsPartialUpdateParams creates a new ProjectsStatsPartialUpdateParams object
// with the default values initialized.
func NewProjectsStatsPartialUpdateParams() *ProjectsStatsPartialUpdateParams {
	var ()
	return &ProjectsStatsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsStatsPartialUpdateParamsWithTimeout creates a new ProjectsStatsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsStatsPartialUpdateParamsWithTimeout(timeout time.Duration) *ProjectsStatsPartialUpdateParams {
	var ()
	return &ProjectsStatsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsStatsPartialUpdateParamsWithContext creates a new ProjectsStatsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsStatsPartialUpdateParamsWithContext(ctx context.Context) *ProjectsStatsPartialUpdateParams {
	var ()
	return &ProjectsStatsPartialUpdateParams{

		Context: ctx,
	}
}

// NewProjectsStatsPartialUpdateParamsWithHTTPClient creates a new ProjectsStatsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsStatsPartialUpdateParamsWithHTTPClient(client *http.Client) *ProjectsStatsPartialUpdateParams {
	var ()
	return &ProjectsStatsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsStatsPartialUpdateParams contains all the parameters to send to the API endpoint
for the projects stats partial update operation typically these are written to a http.Request
*/
type ProjectsStatsPartialUpdateParams struct {

	/*Data*/
	Data ProjectsStatsPartialUpdateBody
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string
	/*ServerPk*/
	ServerPk string
	/*ServerType*/
	ServerType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithTimeout(timeout time.Duration) *ProjectsStatsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithContext(ctx context.Context) *ProjectsStatsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithHTTPClient(client *http.Client) *ProjectsStatsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithData(data ProjectsStatsPartialUpdateBody) *ProjectsStatsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetData(data ProjectsStatsPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithID(id string) *ProjectsStatsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithNamespace(namespace string) *ProjectsStatsPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithProjectPk(projectPk string) *ProjectsStatsPartialUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithServerPk(serverPk string) *ProjectsStatsPartialUpdateParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WithServerType adds the serverType to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) WithServerType(serverType string) *ProjectsStatsPartialUpdateParams {
	o.SetServerType(serverType)
	return o
}

// SetServerType adds the serverType to the projects stats partial update params
func (o *ProjectsStatsPartialUpdateParams) SetServerType(serverType string) {
	o.ServerType = serverType
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsStatsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param server_type
	if err := r.SetPathParam("server_type", o.ServerType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
