// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersCreateParams creates a new UsersCreateParams object
// with the default values initialized.
func NewUsersCreateParams() *UsersCreateParams {
	var ()
	return &UsersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersCreateParamsWithTimeout creates a new UsersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersCreateParamsWithTimeout(timeout time.Duration) *UsersCreateParams {
	var ()
	return &UsersCreateParams{

		timeout: timeout,
	}
}

// NewUsersCreateParamsWithContext creates a new UsersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersCreateParamsWithContext(ctx context.Context) *UsersCreateParams {
	var ()
	return &UsersCreateParams{

		Context: ctx,
	}
}

// NewUsersCreateParamsWithHTTPClient creates a new UsersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersCreateParamsWithHTTPClient(client *http.Client) *UsersCreateParams {
	var ()
	return &UsersCreateParams{
		HTTPClient: client,
	}
}

/*UsersCreateParams contains all the parameters to send to the API endpoint
for the users create operation typically these are written to a http.Request
*/
type UsersCreateParams struct {

	/*UserData*/
	UserData *models.UserData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users create params
func (o *UsersCreateParams) WithTimeout(timeout time.Duration) *UsersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users create params
func (o *UsersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users create params
func (o *UsersCreateParams) WithContext(ctx context.Context) *UsersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users create params
func (o *UsersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users create params
func (o *UsersCreateParams) WithHTTPClient(client *http.Client) *UsersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users create params
func (o *UsersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserData adds the userData to the users create params
func (o *UsersCreateParams) WithUserData(userData *models.UserData) *UsersCreateParams {
	o.SetUserData(userData)
	return o
}

// SetUserData adds the userData to the users create params
func (o *UsersCreateParams) SetUserData(userData *models.UserData) {
	o.UserData = userData
}

// WriteToRequest writes these params to a swagger request
func (o *UsersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserData == nil {
		o.UserData = new(models.UserData)
	}

	if err := r.SetBodyParam(o.UserData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
