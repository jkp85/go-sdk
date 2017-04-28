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

// NewProjectsStatsListParams creates a new ProjectsStatsListParams object
// with the default values initialized.
func NewProjectsStatsListParams() *ProjectsStatsListParams {
	var ()
	return &ProjectsStatsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsStatsListParamsWithTimeout creates a new ProjectsStatsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsStatsListParamsWithTimeout(timeout time.Duration) *ProjectsStatsListParams {
	var ()
	return &ProjectsStatsListParams{

		timeout: timeout,
	}
}

// NewProjectsStatsListParamsWithContext creates a new ProjectsStatsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsStatsListParamsWithContext(ctx context.Context) *ProjectsStatsListParams {
	var ()
	return &ProjectsStatsListParams{

		Context: ctx,
	}
}

// NewProjectsStatsListParamsWithHTTPClient creates a new ProjectsStatsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsStatsListParamsWithHTTPClient(client *http.Client) *ProjectsStatsListParams {
	var ()
	return &ProjectsStatsListParams{
		HTTPClient: client,
	}
}

/*ProjectsStatsListParams contains all the parameters to send to the API endpoint
for the projects stats list operation typically these are written to a http.Request
*/
type ProjectsStatsListParams struct {

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

// WithTimeout adds the timeout to the projects stats list params
func (o *ProjectsStatsListParams) WithTimeout(timeout time.Duration) *ProjectsStatsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects stats list params
func (o *ProjectsStatsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects stats list params
func (o *ProjectsStatsListParams) WithContext(ctx context.Context) *ProjectsStatsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects stats list params
func (o *ProjectsStatsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects stats list params
func (o *ProjectsStatsListParams) WithHTTPClient(client *http.Client) *ProjectsStatsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects stats list params
func (o *ProjectsStatsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the projects stats list params
func (o *ProjectsStatsListParams) WithLimit(limit *string) *ProjectsStatsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the projects stats list params
func (o *ProjectsStatsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the projects stats list params
func (o *ProjectsStatsListParams) WithNamespace(namespace string) *ProjectsStatsListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects stats list params
func (o *ProjectsStatsListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the projects stats list params
func (o *ProjectsStatsListParams) WithOffset(offset *string) *ProjectsStatsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the projects stats list params
func (o *ProjectsStatsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithProjectPk adds the projectPk to the projects stats list params
func (o *ProjectsStatsListParams) WithProjectPk(projectPk string) *ProjectsStatsListParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects stats list params
func (o *ProjectsStatsListParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects stats list params
func (o *ProjectsStatsListParams) WithServerPk(serverPk string) *ProjectsStatsListParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects stats list params
func (o *ProjectsStatsListParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WithServerType adds the serverType to the projects stats list params
func (o *ProjectsStatsListParams) WithServerType(serverType string) *ProjectsStatsListParams {
	o.SetServerType(serverType)
	return o
}

// SetServerType adds the serverType to the projects stats list params
func (o *ProjectsStatsListParams) SetServerType(serverType string) {
	o.ServerType = serverType
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsStatsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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