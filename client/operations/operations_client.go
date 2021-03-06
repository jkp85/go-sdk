// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TeamsBillingInvoiceItemsList gets team invoice items for a given invoice
*/
func (a *Client) TeamsBillingInvoiceItemsList(params *TeamsBillingInvoiceItemsListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsBillingInvoiceItemsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsBillingInvoiceItemsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_billing_invoice_items_list",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/billing/invoices/{invoice_id}/invoice-items/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsBillingInvoiceItemsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsBillingInvoiceItemsListOK), nil

}

/*
TeamsBillingInvoiceItemsRead gets a specific team invoice item
*/
func (a *Client) TeamsBillingInvoiceItemsRead(params *TeamsBillingInvoiceItemsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsBillingInvoiceItemsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsBillingInvoiceItemsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_billing_invoice_items_read",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/billing/invoices/{invoice_id}/invoice-items/{id}",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsBillingInvoiceItemsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsBillingInvoiceItemsReadOK), nil

}

/*
TeamsBillingInvoicesList gets team invoices
*/
func (a *Client) TeamsBillingInvoicesList(params *TeamsBillingInvoicesListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsBillingInvoicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsBillingInvoicesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_billing_invoices_list",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/billing/invoices/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsBillingInvoicesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsBillingInvoicesListOK), nil

}

/*
TeamsBillingInvoicesRead gets an invoice
*/
func (a *Client) TeamsBillingInvoicesRead(params *TeamsBillingInvoicesReadParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsBillingInvoicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsBillingInvoicesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_billing_invoices_read",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/billing/invoices/{id}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsBillingInvoicesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsBillingInvoicesReadOK), nil

}

/*
TeamsBillingSubscriptionsCreate creates a new team subscription
*/
func (a *Client) TeamsBillingSubscriptionsCreate(params *TeamsBillingSubscriptionsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsBillingSubscriptionsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsBillingSubscriptionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_billing_subscriptions_create",
		Method:             "POST",
		PathPattern:        "/v1/teams/{team}/billing/subscriptions/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsBillingSubscriptionsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsBillingSubscriptionsCreateCreated), nil

}

/*
TeamsBillingSubscriptionsDelete deletes a subscription
*/
func (a *Client) TeamsBillingSubscriptionsDelete(params *TeamsBillingSubscriptionsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsBillingSubscriptionsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsBillingSubscriptionsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_billing_subscriptions_delete",
		Method:             "DELETE",
		PathPattern:        "/v1/teams/{team}/billing/subscriptions/{id}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsBillingSubscriptionsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsBillingSubscriptionsDeleteNoContent), nil

}

/*
TeamsBillingSubscriptionsList gets active team subscriptons
*/
func (a *Client) TeamsBillingSubscriptionsList(params *TeamsBillingSubscriptionsListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsBillingSubscriptionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsBillingSubscriptionsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_billing_subscriptions_list",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/billing/subscriptions/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsBillingSubscriptionsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsBillingSubscriptionsListOK), nil

}

/*
TeamsBillingSubscriptionsRead gets team subscriptions
*/
func (a *Client) TeamsBillingSubscriptionsRead(params *TeamsBillingSubscriptionsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsBillingSubscriptionsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsBillingSubscriptionsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_billing_subscriptions_read",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/billing/subscriptions/{id}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsBillingSubscriptionsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsBillingSubscriptionsReadOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
