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

// NewProjectsServersStatsListParams creates a new ProjectsServersStatsListParams object
// with the default values initialized.
func NewProjectsServersStatsListParams() *ProjectsServersStatsListParams {
	var ()
	return &ProjectsServersStatsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersStatsListParamsWithTimeout creates a new ProjectsServersStatsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersStatsListParamsWithTimeout(timeout time.Duration) *ProjectsServersStatsListParams {
	var ()
	return &ProjectsServersStatsListParams{

		timeout: timeout,
	}
}

// NewProjectsServersStatsListParamsWithContext creates a new ProjectsServersStatsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersStatsListParamsWithContext(ctx context.Context) *ProjectsServersStatsListParams {
	var ()
	return &ProjectsServersStatsListParams{

		Context: ctx,
	}
}

// NewProjectsServersStatsListParamsWithHTTPClient creates a new ProjectsServersStatsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersStatsListParamsWithHTTPClient(client *http.Client) *ProjectsServersStatsListParams {
	var ()
	return &ProjectsServersStatsListParams{
		HTTPClient: client,
	}
}

/*ProjectsServersStatsListParams contains all the parameters to send to the API endpoint
for the projects servers stats list operation typically these are written to a http.Request
*/
type ProjectsServersStatsListParams struct {

	/*Limit
	  Limit items when retrieving data.

	*/
	Limit *string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Offset
	  Offset items when retrieving items.

	*/
	Offset *string
	/*Ordering
	  Order when retrieving items.

	*/
	Ordering *string
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

// WithTimeout adds the timeout to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithTimeout(timeout time.Duration) *ProjectsServersStatsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithContext(ctx context.Context) *ProjectsServersStatsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithHTTPClient(client *http.Client) *ProjectsServersStatsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithLimit(limit *string) *ProjectsServersStatsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithNamespace(namespace string) *ProjectsServersStatsListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithOffset(offset *string) *ProjectsServersStatsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithOrdering(ordering *string) *ProjectsServersStatsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithProjectID adds the projectID to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithProjectID(projectID string) *ProjectsServersStatsListParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServerID adds the serverID to the projects servers stats list params
func (o *ProjectsServersStatsListParams) WithServerID(serverID string) *ProjectsServersStatsListParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the projects servers stats list params
func (o *ProjectsServersStatsListParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersStatsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string
		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {
			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}

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
