// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewProjectsProjectFilesReplaceParams creates a new ProjectsProjectFilesReplaceParams object
// with the default values initialized.
func NewProjectsProjectFilesReplaceParams() *ProjectsProjectFilesReplaceParams {
	var ()
	return &ProjectsProjectFilesReplaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectFilesReplaceParamsWithTimeout creates a new ProjectsProjectFilesReplaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsProjectFilesReplaceParamsWithTimeout(timeout time.Duration) *ProjectsProjectFilesReplaceParams {
	var ()
	return &ProjectsProjectFilesReplaceParams{

		timeout: timeout,
	}
}

// NewProjectsProjectFilesReplaceParamsWithContext creates a new ProjectsProjectFilesReplaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsProjectFilesReplaceParamsWithContext(ctx context.Context) *ProjectsProjectFilesReplaceParams {
	var ()
	return &ProjectsProjectFilesReplaceParams{

		Context: ctx,
	}
}

// NewProjectsProjectFilesReplaceParamsWithHTTPClient creates a new ProjectsProjectFilesReplaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsProjectFilesReplaceParamsWithHTTPClient(client *http.Client) *ProjectsProjectFilesReplaceParams {
	var ()
	return &ProjectsProjectFilesReplaceParams{
		HTTPClient: client,
	}
}

/*ProjectsProjectFilesReplaceParams contains all the parameters to send to the API endpoint
for the projects project files replace operation typically these are written to a http.Request
*/
type ProjectsProjectFilesReplaceParams struct {

	/*Base64Data
	  File contents as base64.

	*/
	Base64Data *string
	/*File
	  File to update.

	*/
	File os.File
	/*ID
	  Project file identifier expressed as UUID.

	*/
	ID string
	/*Name
	  File name.

	*/
	Name *string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*ProjectID
	  Project unique identifer expressed as UUID.

	*/
	ProjectID string
	/*Public
	  Is project public or private.

	*/
	Public *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithTimeout(timeout time.Duration) *ProjectsProjectFilesReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithContext(ctx context.Context) *ProjectsProjectFilesReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithHTTPClient(client *http.Client) *ProjectsProjectFilesReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBase64Data adds the base64Data to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithBase64Data(base64Data *string) *ProjectsProjectFilesReplaceParams {
	o.SetBase64Data(base64Data)
	return o
}

// SetBase64Data adds the base64Data to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetBase64Data(base64Data *string) {
	o.Base64Data = base64Data
}

// WithFile adds the file to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithFile(file os.File) *ProjectsProjectFilesReplaceParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetFile(file os.File) {
	o.File = file
}

// WithID adds the id to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithID(id string) *ProjectsProjectFilesReplaceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithName(name *string) *ProjectsProjectFilesReplaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetName(name *string) {
	o.Name = name
}

// WithNamespace adds the namespace to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithNamespace(namespace string) *ProjectsProjectFilesReplaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithProjectID(projectID string) *ProjectsProjectFilesReplaceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithPublic adds the public to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) WithPublic(public *bool) *ProjectsProjectFilesReplaceParams {
	o.SetPublic(public)
	return o
}

// SetPublic adds the public to the projects project files replace params
func (o *ProjectsProjectFilesReplaceParams) SetPublic(public *bool) {
	o.Public = public
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectFilesReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Base64Data != nil {

		// form param base64_data
		var frBase64Data string
		if o.Base64Data != nil {
			frBase64Data = *o.Base64Data
		}
		fBase64Data := frBase64Data
		if fBase64Data != "" {
			if err := r.SetFormParam("base64_data", fBase64Data); err != nil {
				return err
			}
		}

	}

	// form file param file
	if err := r.SetFileParam("file", &o.File); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if o.Public != nil {

		// form param public
		var frPublic bool
		if o.Public != nil {
			frPublic = *o.Public
		}
		fPublic := swag.FormatBool(frPublic)
		if fPublic != "" {
			if err := r.SetFormParam("public", fPublic); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}