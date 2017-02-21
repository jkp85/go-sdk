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

// NewProjectsServersSSHTunnelsUpdateParams creates a new ProjectsServersSSHTunnelsUpdateParams object
// with the default values initialized.
func NewProjectsServersSSHTunnelsUpdateParams() *ProjectsServersSSHTunnelsUpdateParams {
	var ()
	return &ProjectsServersSSHTunnelsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersSSHTunnelsUpdateParamsWithTimeout creates a new ProjectsServersSSHTunnelsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersSSHTunnelsUpdateParamsWithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsUpdateParams {
	var ()
	return &ProjectsServersSSHTunnelsUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsServersSSHTunnelsUpdateParamsWithContext creates a new ProjectsServersSSHTunnelsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersSSHTunnelsUpdateParamsWithContext(ctx context.Context) *ProjectsServersSSHTunnelsUpdateParams {
	var ()
	return &ProjectsServersSSHTunnelsUpdateParams{

		Context: ctx,
	}
}

// NewProjectsServersSSHTunnelsUpdateParamsWithHTTPClient creates a new ProjectsServersSSHTunnelsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersSSHTunnelsUpdateParamsWithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsUpdateParams {
	var ()
	return &ProjectsServersSSHTunnelsUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersSSHTunnelsUpdateParams contains all the parameters to send to the API endpoint
for the projects servers ssh tunnels update operation typically these are written to a http.Request
*/
type ProjectsServersSSHTunnelsUpdateParams struct {

	/*Data*/
	Data ProjectsServersSSHTunnelsUpdateBody
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

// WithTimeout adds the timeout to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithContext(ctx context.Context) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithData(data ProjectsServersSSHTunnelsUpdateBody) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetData(data ProjectsServersSSHTunnelsUpdateBody) {
	o.Data = data
}

// WithID adds the id to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithID(id string) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithNamespace(namespace string) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithProjectPk(projectPk string) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithServerPk(serverPk string) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersSSHTunnelsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
