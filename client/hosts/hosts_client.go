// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new hosts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hosts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
HostsCreate creates a new host
*/
func (a *Client) HostsCreate(params *HostsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*HostsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hosts_create",
		Method:             "POST",
		PathPattern:        "/v1/{namespace}/hosts/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HostsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HostsCreateCreated), nil

}

/*
HostsDelete deletes a host
*/
func (a *Client) HostsDelete(params *HostsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*HostsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hosts_delete",
		Method:             "DELETE",
		PathPattern:        "/v1/{namespace}/hosts/{host}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HostsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HostsDeleteNoContent), nil

}

/*
HostsList gets available hosts
*/
func (a *Client) HostsList(params *HostsListParams, authInfo runtime.ClientAuthInfoWriter) (*HostsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hosts_list",
		Method:             "GET",
		PathPattern:        "/v1/{namespace}/hosts/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HostsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HostsListOK), nil

}

/*
HostsRead gets a host
*/
func (a *Client) HostsRead(params *HostsReadParams, authInfo runtime.ClientAuthInfoWriter) (*HostsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hosts_read",
		Method:             "GET",
		PathPattern:        "/v1/{namespace}/hosts/{host}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HostsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HostsReadOK), nil

}

/*
HostsReplace replaces a host
*/
func (a *Client) HostsReplace(params *HostsReplaceParams, authInfo runtime.ClientAuthInfoWriter) (*HostsReplaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsReplaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hosts_replace",
		Method:             "PUT",
		PathPattern:        "/v1/{namespace}/hosts/{host}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HostsReplaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HostsReplaceOK), nil

}

/*
HostsUpdate updates a host
*/
func (a *Client) HostsUpdate(params *HostsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*HostsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hosts_update",
		Method:             "PATCH",
		PathPattern:        "/v1/{namespace}/hosts/{host}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HostsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HostsUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
