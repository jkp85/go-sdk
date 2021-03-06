// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewTeamsBillingSubscriptionsReadParams creates a new TeamsBillingSubscriptionsReadParams object
// with the default values initialized.
func NewTeamsBillingSubscriptionsReadParams() *TeamsBillingSubscriptionsReadParams {
	var ()
	return &TeamsBillingSubscriptionsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsBillingSubscriptionsReadParamsWithTimeout creates a new TeamsBillingSubscriptionsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsBillingSubscriptionsReadParamsWithTimeout(timeout time.Duration) *TeamsBillingSubscriptionsReadParams {
	var ()
	return &TeamsBillingSubscriptionsReadParams{

		timeout: timeout,
	}
}

// NewTeamsBillingSubscriptionsReadParamsWithContext creates a new TeamsBillingSubscriptionsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsBillingSubscriptionsReadParamsWithContext(ctx context.Context) *TeamsBillingSubscriptionsReadParams {
	var ()
	return &TeamsBillingSubscriptionsReadParams{

		Context: ctx,
	}
}

// NewTeamsBillingSubscriptionsReadParamsWithHTTPClient creates a new TeamsBillingSubscriptionsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsBillingSubscriptionsReadParamsWithHTTPClient(client *http.Client) *TeamsBillingSubscriptionsReadParams {
	var ()
	return &TeamsBillingSubscriptionsReadParams{
		HTTPClient: client,
	}
}

/*TeamsBillingSubscriptionsReadParams contains all the parameters to send to the API endpoint
for the teams billing subscriptions read operation typically these are written to a http.Request
*/
type TeamsBillingSubscriptionsReadParams struct {

	/*ID
	  Unique identifier expressed as UUID.

	*/
	ID string
	/*Team
	  Team unique identifier expressed as UUID or name.

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) WithTimeout(timeout time.Duration) *TeamsBillingSubscriptionsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) WithContext(ctx context.Context) *TeamsBillingSubscriptionsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) WithHTTPClient(client *http.Client) *TeamsBillingSubscriptionsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) WithID(id string) *TeamsBillingSubscriptionsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) SetID(id string) {
	o.ID = id
}

// WithTeam adds the team to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) WithTeam(team string) *TeamsBillingSubscriptionsReadParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the teams billing subscriptions read params
func (o *TeamsBillingSubscriptionsReadParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsBillingSubscriptionsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
