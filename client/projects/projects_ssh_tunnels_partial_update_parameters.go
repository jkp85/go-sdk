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

// NewProjectsSSHTunnelsPartialUpdateParams creates a new ProjectsSSHTunnelsPartialUpdateParams object
// with the default values initialized.
func NewProjectsSSHTunnelsPartialUpdateParams() *ProjectsSSHTunnelsPartialUpdateParams {
	var ()
	return &ProjectsSSHTunnelsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsSSHTunnelsPartialUpdateParamsWithTimeout creates a new ProjectsSSHTunnelsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsSSHTunnelsPartialUpdateParamsWithTimeout(timeout time.Duration) *ProjectsSSHTunnelsPartialUpdateParams {
	var ()
	return &ProjectsSSHTunnelsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsSSHTunnelsPartialUpdateParamsWithContext creates a new ProjectsSSHTunnelsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsSSHTunnelsPartialUpdateParamsWithContext(ctx context.Context) *ProjectsSSHTunnelsPartialUpdateParams {
	var ()
	return &ProjectsSSHTunnelsPartialUpdateParams{

		Context: ctx,
	}
}

// NewProjectsSSHTunnelsPartialUpdateParamsWithHTTPClient creates a new ProjectsSSHTunnelsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsSSHTunnelsPartialUpdateParamsWithHTTPClient(client *http.Client) *ProjectsSSHTunnelsPartialUpdateParams {
	var ()
	return &ProjectsSSHTunnelsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsSSHTunnelsPartialUpdateParams contains all the parameters to send to the API endpoint
for the projects ssh tunnels partial update operation typically these are written to a http.Request
*/
type ProjectsSSHTunnelsPartialUpdateParams struct {

	/*Data*/
	Data ProjectsSSHTunnelsPartialUpdateBody
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

// WithTimeout adds the timeout to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithTimeout(timeout time.Duration) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithContext(ctx context.Context) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithHTTPClient(client *http.Client) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithData(data ProjectsSSHTunnelsPartialUpdateBody) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetData(data ProjectsSSHTunnelsPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithID(id string) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithNamespace(namespace string) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithProjectPk(projectPk string) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithServerPk(serverPk string) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WithServerType adds the serverType to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) WithServerType(serverType string) *ProjectsSSHTunnelsPartialUpdateParams {
	o.SetServerType(serverType)
	return o
}

// SetServerType adds the serverType to the projects ssh tunnels partial update params
func (o *ProjectsSSHTunnelsPartialUpdateParams) SetServerType(serverType string) {
	o.ServerType = serverType
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsSSHTunnelsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
