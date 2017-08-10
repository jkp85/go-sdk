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

	/*ID*/
	ID string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*ProjectID*/
	ProjectID string
	/*ServerID*/
	ServerID string
	/*SshtunnelData*/
	SshtunnelData *models.SSHTunnelData

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

// WithProjectID adds the projectID to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithProjectID(projectID string) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServerID adds the serverID to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithServerID(serverID string) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithSshtunnelData adds the sshtunnelData to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) WithSshtunnelData(sshtunnelData *models.SSHTunnelData) *ProjectsServersSSHTunnelsUpdateParams {
	o.SetSshtunnelData(sshtunnelData)
	return o
}

// SetSshtunnelData adds the sshtunnelData to the projects servers ssh tunnels update params
func (o *ProjectsServersSSHTunnelsUpdateParams) SetSshtunnelData(sshtunnelData *models.SSHTunnelData) {
	o.SshtunnelData = sshtunnelData
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersSSHTunnelsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SshtunnelData == nil {
		o.SshtunnelData = new(models.SSHTunnelData)
	}

	if err := r.SetBodyParam(o.SshtunnelData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
