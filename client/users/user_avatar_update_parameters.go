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

// NewUserAvatarUpdateParams creates a new UserAvatarUpdateParams object
// with the default values initialized.
func NewUserAvatarUpdateParams() *UserAvatarUpdateParams {
	var ()
	return &UserAvatarUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserAvatarUpdateParamsWithTimeout creates a new UserAvatarUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserAvatarUpdateParamsWithTimeout(timeout time.Duration) *UserAvatarUpdateParams {
	var ()
	return &UserAvatarUpdateParams{

		timeout: timeout,
	}
}

// NewUserAvatarUpdateParamsWithContext creates a new UserAvatarUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserAvatarUpdateParamsWithContext(ctx context.Context) *UserAvatarUpdateParams {
	var ()
	return &UserAvatarUpdateParams{

		Context: ctx,
	}
}

// NewUserAvatarUpdateParamsWithHTTPClient creates a new UserAvatarUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserAvatarUpdateParamsWithHTTPClient(client *http.Client) *UserAvatarUpdateParams {
	var ()
	return &UserAvatarUpdateParams{
		HTTPClient: client,
	}
}

/*UserAvatarUpdateParams contains all the parameters to send to the API endpoint
for the user avatar update operation typically these are written to a http.Request
*/
type UserAvatarUpdateParams struct {

	/*User
	  User unique identifier expressed as UUID or username.

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user avatar update params
func (o *UserAvatarUpdateParams) WithTimeout(timeout time.Duration) *UserAvatarUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user avatar update params
func (o *UserAvatarUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user avatar update params
func (o *UserAvatarUpdateParams) WithContext(ctx context.Context) *UserAvatarUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user avatar update params
func (o *UserAvatarUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user avatar update params
func (o *UserAvatarUpdateParams) WithHTTPClient(client *http.Client) *UserAvatarUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user avatar update params
func (o *UserAvatarUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the user avatar update params
func (o *UserAvatarUpdateParams) WithUser(user string) *UserAvatarUpdateParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the user avatar update params
func (o *UserAvatarUpdateParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UserAvatarUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
