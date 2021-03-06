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

	models "github.com/IllumiDesk/go-sdk/models"
)

// NewUsersEmailsUpdateParams creates a new UsersEmailsUpdateParams object
// with the default values initialized.
func NewUsersEmailsUpdateParams() *UsersEmailsUpdateParams {
	var ()
	return &UsersEmailsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersEmailsUpdateParamsWithTimeout creates a new UsersEmailsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersEmailsUpdateParamsWithTimeout(timeout time.Duration) *UsersEmailsUpdateParams {
	var ()
	return &UsersEmailsUpdateParams{

		timeout: timeout,
	}
}

// NewUsersEmailsUpdateParamsWithContext creates a new UsersEmailsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersEmailsUpdateParamsWithContext(ctx context.Context) *UsersEmailsUpdateParams {
	var ()
	return &UsersEmailsUpdateParams{

		Context: ctx,
	}
}

// NewUsersEmailsUpdateParamsWithHTTPClient creates a new UsersEmailsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersEmailsUpdateParamsWithHTTPClient(client *http.Client) *UsersEmailsUpdateParams {
	var ()
	return &UsersEmailsUpdateParams{
		HTTPClient: client,
	}
}

/*UsersEmailsUpdateParams contains all the parameters to send to the API endpoint
for the users emails update operation typically these are written to a http.Request
*/
type UsersEmailsUpdateParams struct {

	/*EmailData*/
	EmailData *models.EmailData
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

// WithTimeout adds the timeout to the users emails update params
func (o *UsersEmailsUpdateParams) WithTimeout(timeout time.Duration) *UsersEmailsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users emails update params
func (o *UsersEmailsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users emails update params
func (o *UsersEmailsUpdateParams) WithContext(ctx context.Context) *UsersEmailsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users emails update params
func (o *UsersEmailsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users emails update params
func (o *UsersEmailsUpdateParams) WithHTTPClient(client *http.Client) *UsersEmailsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users emails update params
func (o *UsersEmailsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailData adds the emailData to the users emails update params
func (o *UsersEmailsUpdateParams) WithEmailData(emailData *models.EmailData) *UsersEmailsUpdateParams {
	o.SetEmailData(emailData)
	return o
}

// SetEmailData adds the emailData to the users emails update params
func (o *UsersEmailsUpdateParams) SetEmailData(emailData *models.EmailData) {
	o.EmailData = emailData
}

// WithEmailID adds the emailID to the users emails update params
func (o *UsersEmailsUpdateParams) WithEmailID(emailID string) *UsersEmailsUpdateParams {
	o.SetEmailID(emailID)
	return o
}

// SetEmailID adds the emailId to the users emails update params
func (o *UsersEmailsUpdateParams) SetEmailID(emailID string) {
	o.EmailID = emailID
}

// WithUser adds the user to the users emails update params
func (o *UsersEmailsUpdateParams) WithUser(user string) *UsersEmailsUpdateParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the users emails update params
func (o *UsersEmailsUpdateParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UsersEmailsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EmailData != nil {
		if err := r.SetBodyParam(o.EmailData); err != nil {
			return err
		}
	}

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
