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

	models "github.com/IllumiDesk/go-sdk/models"
)

// NewProjectsServersSSHTunnelsReplaceParams creates a new ProjectsServersSSHTunnelsReplaceParams object
// with the default values initialized.
func NewProjectsServersSSHTunnelsReplaceParams() *ProjectsServersSSHTunnelsReplaceParams {
	var ()
	return &ProjectsServersSSHTunnelsReplaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersSSHTunnelsReplaceParamsWithTimeout creates a new ProjectsServersSSHTunnelsReplaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersSSHTunnelsReplaceParamsWithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsReplaceParams {
	var ()
	return &ProjectsServersSSHTunnelsReplaceParams{

		timeout: timeout,
	}
}

// NewProjectsServersSSHTunnelsReplaceParamsWithContext creates a new ProjectsServersSSHTunnelsReplaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersSSHTunnelsReplaceParamsWithContext(ctx context.Context) *ProjectsServersSSHTunnelsReplaceParams {
	var ()
	return &ProjectsServersSSHTunnelsReplaceParams{

		Context: ctx,
	}
}

// NewProjectsServersSSHTunnelsReplaceParamsWithHTTPClient creates a new ProjectsServersSSHTunnelsReplaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersSSHTunnelsReplaceParamsWithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsReplaceParams {
	var ()
	return &ProjectsServersSSHTunnelsReplaceParams{
		HTTPClient: client,
	}
}

/*ProjectsServersSSHTunnelsReplaceParams contains all the parameters to send to the API endpoint
for the projects servers ssh tunnels replace operation typically these are written to a http.Request
*/
type ProjectsServersSSHTunnelsReplaceParams struct {

	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Project
	  Project unique identifier expressed as UUID or name.

	*/
	Project string
	/*Server
	  Server unique identifier expressed as UUID or name.

	*/
	Server string
	/*SshtunnelData*/
	SshtunnelData *models.SSHTunnelData
	/*Tunnel
	  SSH tunnel unique identifier expressed as UUID or name.

	*/
	Tunnel string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) WithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) WithContext(ctx context.Context) *ProjectsServersSSHTunnelsReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) WithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) WithNamespace(namespace string) *ProjectsServersSSHTunnelsReplaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) WithProject(project string) *ProjectsServersSSHTunnelsReplaceParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) SetProject(project string) {
	o.Project = project
}

// WithServer adds the server to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) WithServer(server string) *ProjectsServersSSHTunnelsReplaceParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) SetServer(server string) {
	o.Server = server
}

// WithSshtunnelData adds the sshtunnelData to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) WithSshtunnelData(sshtunnelData *models.SSHTunnelData) *ProjectsServersSSHTunnelsReplaceParams {
	o.SetSshtunnelData(sshtunnelData)
	return o
}

// SetSshtunnelData adds the sshtunnelData to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) SetSshtunnelData(sshtunnelData *models.SSHTunnelData) {
	o.SshtunnelData = sshtunnelData
}

// WithTunnel adds the tunnel to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) WithTunnel(tunnel string) *ProjectsServersSSHTunnelsReplaceParams {
	o.SetTunnel(tunnel)
	return o
}

// SetTunnel adds the tunnel to the projects servers ssh tunnels replace params
func (o *ProjectsServersSSHTunnelsReplaceParams) SetTunnel(tunnel string) {
	o.Tunnel = tunnel
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersSSHTunnelsReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param server
	if err := r.SetPathParam("server", o.Server); err != nil {
		return err
	}

	if o.SshtunnelData != nil {
		if err := r.SetBodyParam(o.SshtunnelData); err != nil {
			return err
		}
	}

	// path param tunnel
	if err := r.SetPathParam("tunnel", o.Tunnel); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
