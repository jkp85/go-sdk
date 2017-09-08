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
)

// NewUsersEmailsReadParams creates a new UsersEmailsReadParams object
// with the default values initialized.
func NewUsersEmailsReadParams() *UsersEmailsReadParams {
	var ()
	return &UsersEmailsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersEmailsReadParamsWithTimeout creates a new UsersEmailsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersEmailsReadParamsWithTimeout(timeout time.Duration) *UsersEmailsReadParams {
	var ()
	return &UsersEmailsReadParams{

		timeout: timeout,
	}
}

// NewUsersEmailsReadParamsWithContext creates a new UsersEmailsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersEmailsReadParamsWithContext(ctx context.Context) *UsersEmailsReadParams {
	var ()
	return &UsersEmailsReadParams{

		Context: ctx,
	}
}

// NewUsersEmailsReadParamsWithHTTPClient creates a new UsersEmailsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersEmailsReadParamsWithHTTPClient(client *http.Client) *UsersEmailsReadParams {
	var ()
	return &UsersEmailsReadParams{
		HTTPClient: client,
	}
}

/*UsersEmailsReadParams contains all the parameters to send to the API endpoint
for the users emails read operation typically these are written to a http.Request
*/
type UsersEmailsReadParams struct {

	/*EmailID
	  Email unique identifier expressed as UUID.

	*/
	EmailID string
	/*User
	  User unique identifier expressed as UUID or username.

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users emails read params
func (o *UsersEmailsReadParams) WithTimeout(timeout time.Duration) *UsersEmailsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users emails read params
func (o *UsersEmailsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users emails read params
func (o *UsersEmailsReadParams) WithContext(ctx context.Context) *UsersEmailsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users emails read params
func (o *UsersEmailsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users emails read params
func (o *UsersEmailsReadParams) WithHTTPClient(client *http.Client) *UsersEmailsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users emails read params
func (o *UsersEmailsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailID adds the emailID to the users emails read params
func (o *UsersEmailsReadParams) WithEmailID(emailID string) *UsersEmailsReadParams {
	o.SetEmailID(emailID)
	return o
}

// SetEmailID adds the emailId to the users emails read params
func (o *UsersEmailsReadParams) SetEmailID(emailID string) {
	o.EmailID = emailID
}

// WithUser adds the user to the users emails read params
func (o *UsersEmailsReadParams) WithUser(user string) *UsersEmailsReadParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the users emails read params
func (o *UsersEmailsReadParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UsersEmailsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email_id
	if err := r.SetPathParam("email_id", o.EmailID); err != nil {
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
