// Code generated by go-swagger; DO NOT EDIT.

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

// NewProjectsServersSSHTunnelsReadParams creates a new ProjectsServersSSHTunnelsReadParams object
// with the default values initialized.
func NewProjectsServersSSHTunnelsReadParams() *ProjectsServersSSHTunnelsReadParams {
	var ()
	return &ProjectsServersSSHTunnelsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersSSHTunnelsReadParamsWithTimeout creates a new ProjectsServersSSHTunnelsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersSSHTunnelsReadParamsWithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsReadParams {
	var ()
	return &ProjectsServersSSHTunnelsReadParams{

		timeout: timeout,
	}
}

// NewProjectsServersSSHTunnelsReadParamsWithContext creates a new ProjectsServersSSHTunnelsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersSSHTunnelsReadParamsWithContext(ctx context.Context) *ProjectsServersSSHTunnelsReadParams {
	var ()
	return &ProjectsServersSSHTunnelsReadParams{

		Context: ctx,
	}
}

// NewProjectsServersSSHTunnelsReadParamsWithHTTPClient creates a new ProjectsServersSSHTunnelsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersSSHTunnelsReadParamsWithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsReadParams {
	var ()
	return &ProjectsServersSSHTunnelsReadParams{
		HTTPClient: client,
	}
}

/*ProjectsServersSSHTunnelsReadParams contains all the parameters to send to the API endpoint
for the projects servers ssh tunnels read operation typically these are written to a http.Request
*/
type ProjectsServersSSHTunnelsReadParams struct {

	/*ID
	  SSH tunnel unique identifier expressed as UUID.

	*/
	ID string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*ProjectID
	  Project unique identifier expressed as UUID.

	*/
	ProjectID string
	/*ServerID
	  Server unique identifier expressed as UUID.

	*/
	ServerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) WithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) WithContext(ctx context.Context) *ProjectsServersSSHTunnelsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) WithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) WithID(id string) *ProjectsServersSSHTunnelsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) WithNamespace(namespace string) *ProjectsServersSSHTunnelsReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) WithProjectID(projectID string) *ProjectsServersSSHTunnelsReadParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServerID adds the serverID to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) WithServerID(serverID string) *ProjectsServersSSHTunnelsReadParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the projects servers ssh tunnels read params
func (o *ProjectsServersSSHTunnelsReadParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersSSHTunnelsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
