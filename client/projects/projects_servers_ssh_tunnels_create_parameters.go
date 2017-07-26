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

	"github.com/3Blades/go-sdk/models"
)

// NewProjectsServersSSHTunnelsCreateParams creates a new ProjectsServersSSHTunnelsCreateParams object
// with the default values initialized.
func NewProjectsServersSSHTunnelsCreateParams() *ProjectsServersSSHTunnelsCreateParams {
	var ()
	return &ProjectsServersSSHTunnelsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersSSHTunnelsCreateParamsWithTimeout creates a new ProjectsServersSSHTunnelsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersSSHTunnelsCreateParamsWithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsCreateParams {
	var ()
	return &ProjectsServersSSHTunnelsCreateParams{

		timeout: timeout,
	}
}

// NewProjectsServersSSHTunnelsCreateParamsWithContext creates a new ProjectsServersSSHTunnelsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersSSHTunnelsCreateParamsWithContext(ctx context.Context) *ProjectsServersSSHTunnelsCreateParams {
	var ()
	return &ProjectsServersSSHTunnelsCreateParams{

		Context: ctx,
	}
}

// NewProjectsServersSSHTunnelsCreateParamsWithHTTPClient creates a new ProjectsServersSSHTunnelsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersSSHTunnelsCreateParamsWithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsCreateParams {
	var ()
	return &ProjectsServersSSHTunnelsCreateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersSSHTunnelsCreateParams contains all the parameters to send to the API endpoint
for the projects servers ssh tunnels create operation typically these are written to a http.Request
*/
type ProjectsServersSSHTunnelsCreateParams struct {

	/*Data*/
	Data *models.SSHTunnelData
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

// WithTimeout adds the timeout to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) WithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) WithContext(ctx context.Context) *ProjectsServersSSHTunnelsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) WithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) WithData(data *models.SSHTunnelData) *ProjectsServersSSHTunnelsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) SetData(data *models.SSHTunnelData) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) WithNamespace(namespace string) *ProjectsServersSSHTunnelsCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) WithProjectID(projectID string) *ProjectsServersSSHTunnelsCreateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServerID adds the serverID to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) WithServerID(serverID string) *ProjectsServersSSHTunnelsCreateParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the projects servers ssh tunnels create params
func (o *ProjectsServersSSHTunnelsCreateParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersSSHTunnelsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data == nil {
		o.Data = new(models.SSHTunnelData)
	}

	if err := r.SetBodyParam(o.Data); err != nil {
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
