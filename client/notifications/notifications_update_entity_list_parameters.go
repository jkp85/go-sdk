// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewNotificationsUpdateEntityListParams creates a new NotificationsUpdateEntityListParams object
// with the default values initialized.
func NewNotificationsUpdateEntityListParams() *NotificationsUpdateEntityListParams {
	var ()
	return &NotificationsUpdateEntityListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationsUpdateEntityListParamsWithTimeout creates a new NotificationsUpdateEntityListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotificationsUpdateEntityListParamsWithTimeout(timeout time.Duration) *NotificationsUpdateEntityListParams {
	var ()
	return &NotificationsUpdateEntityListParams{

		timeout: timeout,
	}
}

// NewNotificationsUpdateEntityListParamsWithContext creates a new NotificationsUpdateEntityListParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotificationsUpdateEntityListParamsWithContext(ctx context.Context) *NotificationsUpdateEntityListParams {
	var ()
	return &NotificationsUpdateEntityListParams{

		Context: ctx,
	}
}

// NewNotificationsUpdateEntityListParamsWithHTTPClient creates a new NotificationsUpdateEntityListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotificationsUpdateEntityListParamsWithHTTPClient(client *http.Client) *NotificationsUpdateEntityListParams {
	var ()
	return &NotificationsUpdateEntityListParams{
		HTTPClient: client,
	}
}

/*NotificationsUpdateEntityListParams contains all the parameters to send to the API endpoint
for the notifications update entity list operation typically these are written to a http.Request
*/
type NotificationsUpdateEntityListParams struct {

	/*Entity
	  Entity to filter notifications by.

	*/
	Entity string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*NotificationData*/
	NotificationData *models.NotificationListUpdateData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) WithTimeout(timeout time.Duration) *NotificationsUpdateEntityListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) WithContext(ctx context.Context) *NotificationsUpdateEntityListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) WithHTTPClient(client *http.Client) *NotificationsUpdateEntityListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) WithEntity(entity string) *NotificationsUpdateEntityListParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithNamespace adds the namespace to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) WithNamespace(namespace string) *NotificationsUpdateEntityListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithNotificationData adds the notificationData to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) WithNotificationData(notificationData *models.NotificationListUpdateData) *NotificationsUpdateEntityListParams {
	o.SetNotificationData(notificationData)
	return o
}

// SetNotificationData adds the notificationData to the notifications update entity list params
func (o *NotificationsUpdateEntityListParams) SetNotificationData(notificationData *models.NotificationListUpdateData) {
	o.NotificationData = notificationData
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationsUpdateEntityListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.NotificationData != nil {
		if err := r.SetBodyParam(o.NotificationData); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}