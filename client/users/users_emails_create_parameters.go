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

// NewUsersEmailsCreateParams creates a new UsersEmailsCreateParams object
// with the default values initialized.
func NewUsersEmailsCreateParams() *UsersEmailsCreateParams {
	var ()
	return &UsersEmailsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersEmailsCreateParamsWithTimeout creates a new UsersEmailsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersEmailsCreateParamsWithTimeout(timeout time.Duration) *UsersEmailsCreateParams {
	var ()
	return &UsersEmailsCreateParams{

		timeout: timeout,
	}
}

// NewUsersEmailsCreateParamsWithContext creates a new UsersEmailsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersEmailsCreateParamsWithContext(ctx context.Context) *UsersEmailsCreateParams {
	var ()
	return &UsersEmailsCreateParams{

		Context: ctx,
	}
}

// NewUsersEmailsCreateParamsWithHTTPClient creates a new UsersEmailsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersEmailsCreateParamsWithHTTPClient(client *http.Client) *UsersEmailsCreateParams {
	var ()
	return &UsersEmailsCreateParams{
		HTTPClient: client,
	}
}

/*UsersEmailsCreateParams contains all the parameters to send to the API endpoint
for the users emails create operation typically these are written to a http.Request
*/
type UsersEmailsCreateParams struct {

	/*EmailData*/
	EmailData *models.EmailData
	/*User
	  User unique identifier expressed as UUID or username.

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users emails create params
func (o *UsersEmailsCreateParams) WithTimeout(timeout time.Duration) *UsersEmailsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users emails create params
func (o *UsersEmailsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users emails create params
func (o *UsersEmailsCreateParams) WithContext(ctx context.Context) *UsersEmailsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users emails create params
func (o *UsersEmailsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users emails create params
func (o *UsersEmailsCreateParams) WithHTTPClient(client *http.Client) *UsersEmailsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users emails create params
func (o *UsersEmailsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailData adds the emailData to the users emails create params
func (o *UsersEmailsCreateParams) WithEmailData(emailData *models.EmailData) *UsersEmailsCreateParams {
	o.SetEmailData(emailData)
	return o
}

// SetEmailData adds the emailData to the users emails create params
func (o *UsersEmailsCreateParams) SetEmailData(emailData *models.EmailData) {
	o.EmailData = emailData
}

// WithUser adds the user to the users emails create params
func (o *UsersEmailsCreateParams) WithUser(user string) *UsersEmailsCreateParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the users emails create params
func (o *UsersEmailsCreateParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UsersEmailsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EmailData == nil {
		o.EmailData = new(models.EmailData)
	}

	if err := r.SetBodyParam(o.EmailData); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
