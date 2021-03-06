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

// NewProjectsServersSSHTunnelsListParams creates a new ProjectsServersSSHTunnelsListParams object
// with the default values initialized.
func NewProjectsServersSSHTunnelsListParams() *ProjectsServersSSHTunnelsListParams {
	var ()
	return &ProjectsServersSSHTunnelsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersSSHTunnelsListParamsWithTimeout creates a new ProjectsServersSSHTunnelsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersSSHTunnelsListParamsWithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsListParams {
	var ()
	return &ProjectsServersSSHTunnelsListParams{

		timeout: timeout,
	}
}

// NewProjectsServersSSHTunnelsListParamsWithContext creates a new ProjectsServersSSHTunnelsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersSSHTunnelsListParamsWithContext(ctx context.Context) *ProjectsServersSSHTunnelsListParams {
	var ()
	return &ProjectsServersSSHTunnelsListParams{

		Context: ctx,
	}
}

// NewProjectsServersSSHTunnelsListParamsWithHTTPClient creates a new ProjectsServersSSHTunnelsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersSSHTunnelsListParamsWithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsListParams {
	var ()
	return &ProjectsServersSSHTunnelsListParams{
		HTTPClient: client,
	}
}

/*ProjectsServersSSHTunnelsListParams contains all the parameters to send to the API endpoint
for the projects servers ssh tunnels list operation typically these are written to a http.Request
*/
type ProjectsServersSSHTunnelsListParams struct {

	/*Limit
	  Limit retrieved items.

	*/
	Limit *string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Offset
	  Offset retrieved items.

	*/
	Offset *string
	/*Ordering
	  Order retrieved items.

	*/
	Ordering *string
	/*Project
	  Project unique identifier expressed as UUID or name.

	*/
	Project string
	/*Server
	  Server unique identifier expressed as UUID or name.

	*/
	Server string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithTimeout(timeout time.Duration) *ProjectsServersSSHTunnelsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithContext(ctx context.Context) *ProjectsServersSSHTunnelsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithHTTPClient(client *http.Client) *ProjectsServersSSHTunnelsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithLimit(limit *string) *ProjectsServersSSHTunnelsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithNamespace(namespace string) *ProjectsServersSSHTunnelsListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithOffset(offset *string) *ProjectsServersSSHTunnelsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithOrdering(ordering *string) *ProjectsServersSSHTunnelsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithProject adds the project to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithProject(project string) *ProjectsServersSSHTunnelsListParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetProject(project string) {
	o.Project = project
}

// WithServer adds the server to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) WithServer(server string) *ProjectsServersSSHTunnelsListParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects servers ssh tunnels list params
func (o *ProjectsServersSSHTunnelsListParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersSSHTunnelsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param server
	if err := r.SetPathParam("server", o.Server); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
