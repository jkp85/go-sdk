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

// NewUsersEmailsListParams creates a new UsersEmailsListParams object
// with the default values initialized.
func NewUsersEmailsListParams() *UsersEmailsListParams {
	var ()
	return &UsersEmailsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersEmailsListParamsWithTimeout creates a new UsersEmailsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersEmailsListParamsWithTimeout(timeout time.Duration) *UsersEmailsListParams {
	var ()
	return &UsersEmailsListParams{

		timeout: timeout,
	}
}

// NewUsersEmailsListParamsWithContext creates a new UsersEmailsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersEmailsListParamsWithContext(ctx context.Context) *UsersEmailsListParams {
	var ()
	return &UsersEmailsListParams{

		Context: ctx,
	}
}

// NewUsersEmailsListParamsWithHTTPClient creates a new UsersEmailsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersEmailsListParamsWithHTTPClient(client *http.Client) *UsersEmailsListParams {
	var ()
	return &UsersEmailsListParams{
		HTTPClient: client,
	}
}

/*UsersEmailsListParams contains all the parameters to send to the API endpoint
for the users emails list operation typically these are written to a http.Request
*/
type UsersEmailsListParams struct {

	/*Limit*/
	Limit *string
	/*Offset*/
	Offset *string
	/*Ordering*/
	Ordering *string
	/*UserPk*/
	UserPk string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users emails list params
func (o *UsersEmailsListParams) WithTimeout(timeout time.Duration) *UsersEmailsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users emails list params
func (o *UsersEmailsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users emails list params
func (o *UsersEmailsListParams) WithContext(ctx context.Context) *UsersEmailsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users emails list params
func (o *UsersEmailsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users emails list params
func (o *UsersEmailsListParams) WithHTTPClient(client *http.Client) *UsersEmailsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users emails list params
func (o *UsersEmailsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the users emails list params
func (o *UsersEmailsListParams) WithLimit(limit *string) *UsersEmailsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the users emails list params
func (o *UsersEmailsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the users emails list params
func (o *UsersEmailsListParams) WithOffset(offset *string) *UsersEmailsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the users emails list params
func (o *UsersEmailsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the users emails list params
func (o *UsersEmailsListParams) WithOrdering(ordering *string) *UsersEmailsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the users emails list params
func (o *UsersEmailsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithUserPk adds the userPk to the users emails list params
func (o *UsersEmailsListParams) WithUserPk(userPk string) *UsersEmailsListParams {
	o.SetUserPk(userPk)
	return o
}

// SetUserPk adds the userPk to the users emails list params
func (o *UsersEmailsListParams) SetUserPk(userPk string) {
	o.UserPk = userPk
}

// WriteToRequest writes these params to a swagger request
func (o *UsersEmailsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user_pk
	if err := r.SetPathParam("user_pk", o.UserPk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
