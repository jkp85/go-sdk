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

// NewProjectsSSHTunnelsListParams creates a new ProjectsSSHTunnelsListParams object
// with the default values initialized.
func NewProjectsSSHTunnelsListParams() *ProjectsSSHTunnelsListParams {
	var ()
	return &ProjectsSSHTunnelsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsSSHTunnelsListParamsWithTimeout creates a new ProjectsSSHTunnelsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsSSHTunnelsListParamsWithTimeout(timeout time.Duration) *ProjectsSSHTunnelsListParams {
	var ()
	return &ProjectsSSHTunnelsListParams{

		timeout: timeout,
	}
}

// NewProjectsSSHTunnelsListParamsWithContext creates a new ProjectsSSHTunnelsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsSSHTunnelsListParamsWithContext(ctx context.Context) *ProjectsSSHTunnelsListParams {
	var ()
	return &ProjectsSSHTunnelsListParams{

		Context: ctx,
	}
}

// NewProjectsSSHTunnelsListParamsWithHTTPClient creates a new ProjectsSSHTunnelsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsSSHTunnelsListParamsWithHTTPClient(client *http.Client) *ProjectsSSHTunnelsListParams {
	var ()
	return &ProjectsSSHTunnelsListParams{
		HTTPClient: client,
	}
}

/*ProjectsSSHTunnelsListParams contains all the parameters to send to the API endpoint
for the projects ssh tunnels list operation typically these are written to a http.Request
*/
type ProjectsSSHTunnelsListParams struct {

	/*Limit*/
	Limit *string
	/*Namespace*/
	Namespace string
	/*Offset*/
	Offset *string
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

// WithTimeout adds the timeout to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithTimeout(timeout time.Duration) *ProjectsSSHTunnelsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithContext(ctx context.Context) *ProjectsSSHTunnelsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithHTTPClient(client *http.Client) *ProjectsSSHTunnelsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithLimit(limit *string) *ProjectsSSHTunnelsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithNamespace(namespace string) *ProjectsSSHTunnelsListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithOffset(offset *string) *ProjectsSSHTunnelsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithProjectPk adds the projectPk to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithProjectPk(projectPk string) *ProjectsSSHTunnelsListParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithServerPk(serverPk string) *ProjectsSSHTunnelsListParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WithServerType adds the serverType to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) WithServerType(serverType string) *ProjectsSSHTunnelsListParams {
	o.SetServerType(serverType)
	return o
}

// SetServerType adds the serverType to the projects ssh tunnels list params
func (o *ProjectsSSHTunnelsListParams) SetServerType(serverType string) {
	o.ServerType = serverType
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsSSHTunnelsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

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
