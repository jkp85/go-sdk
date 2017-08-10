// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewAuthSimpleTokenAuthParams creates a new AuthSimpleTokenAuthParams object
// with the default values initialized.
func NewAuthSimpleTokenAuthParams() *AuthSimpleTokenAuthParams {
	var ()
	return &AuthSimpleTokenAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthSimpleTokenAuthParamsWithTimeout creates a new AuthSimpleTokenAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthSimpleTokenAuthParamsWithTimeout(timeout time.Duration) *AuthSimpleTokenAuthParams {
	var ()
	return &AuthSimpleTokenAuthParams{

		timeout: timeout,
	}
}

// NewAuthSimpleTokenAuthParamsWithContext creates a new AuthSimpleTokenAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthSimpleTokenAuthParamsWithContext(ctx context.Context) *AuthSimpleTokenAuthParams {
	var ()
	return &AuthSimpleTokenAuthParams{

		Context: ctx,
	}
}

// NewAuthSimpleTokenAuthParamsWithHTTPClient creates a new AuthSimpleTokenAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthSimpleTokenAuthParamsWithHTTPClient(client *http.Client) *AuthSimpleTokenAuthParams {
	var ()
	return &AuthSimpleTokenAuthParams{
		HTTPClient: client,
	}
}

/*AuthSimpleTokenAuthParams contains all the parameters to send to the API endpoint
for the auth simple token auth operation typically these are written to a http.Request
*/
type AuthSimpleTokenAuthParams struct {

	/*AuthTokenData*/
	AuthTokenData *models.AuthTokenData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth simple token auth params
func (o *AuthSimpleTokenAuthParams) WithTimeout(timeout time.Duration) *AuthSimpleTokenAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth simple token auth params
func (o *AuthSimpleTokenAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth simple token auth params
func (o *AuthSimpleTokenAuthParams) WithContext(ctx context.Context) *AuthSimpleTokenAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth simple token auth params
func (o *AuthSimpleTokenAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth simple token auth params
func (o *AuthSimpleTokenAuthParams) WithHTTPClient(client *http.Client) *AuthSimpleTokenAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth simple token auth params
func (o *AuthSimpleTokenAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthTokenData adds the authTokenData to the auth simple token auth params
func (o *AuthSimpleTokenAuthParams) WithAuthTokenData(authTokenData *models.AuthTokenData) *AuthSimpleTokenAuthParams {
	o.SetAuthTokenData(authTokenData)
	return o
}

// SetAuthTokenData adds the authTokenData to the auth simple token auth params
func (o *AuthSimpleTokenAuthParams) SetAuthTokenData(authTokenData *models.AuthTokenData) {
	o.AuthTokenData = authTokenData
}

// WriteToRequest writes these params to a swagger request
func (o *AuthSimpleTokenAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthTokenData == nil {
		o.AuthTokenData = new(models.AuthTokenData)
	}

	if err := r.SetBodyParam(o.AuthTokenData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
