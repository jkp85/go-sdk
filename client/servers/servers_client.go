package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new servers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for servers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServersOptionsResourcesCreate servers options resources create API
*/
func (a *Client) ServersOptionsResourcesCreate(params *ServersOptionsResourcesCreateParams) (*ServersOptionsResourcesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsResourcesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_resources_create",
		Method:             "POST",
		PathPattern:        "/api/v0/{namespace}/servers/options/resources/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsResourcesCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsResourcesCreateCreated), nil

}

/*
ServersOptionsResourcesDelete servers options resources delete API
*/
func (a *Client) ServersOptionsResourcesDelete(params *ServersOptionsResourcesDeleteParams) (*ServersOptionsResourcesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsResourcesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_resources_delete",
		Method:             "DELETE",
		PathPattern:        "/api/v0/{namespace}/servers/options/resources/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsResourcesDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsResourcesDeleteNoContent), nil

}

/*
ServersOptionsResourcesList servers options resources list API
*/
func (a *Client) ServersOptionsResourcesList(params *ServersOptionsResourcesListParams) (*ServersOptionsResourcesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsResourcesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_resources_list",
		Method:             "GET",
		PathPattern:        "/api/v0/{namespace}/servers/options/resources/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsResourcesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsResourcesListOK), nil

}

/*
ServersOptionsResourcesPartialUpdate servers options resources partial update API
*/
func (a *Client) ServersOptionsResourcesPartialUpdate(params *ServersOptionsResourcesPartialUpdateParams) (*ServersOptionsResourcesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsResourcesPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_resources_partial_update",
		Method:             "PATCH",
		PathPattern:        "/api/v0/{namespace}/servers/options/resources/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsResourcesPartialUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsResourcesPartialUpdateOK), nil

}

/*
ServersOptionsResourcesRead servers options resources read API
*/
func (a *Client) ServersOptionsResourcesRead(params *ServersOptionsResourcesReadParams) (*ServersOptionsResourcesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsResourcesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_resources_read",
		Method:             "GET",
		PathPattern:        "/api/v0/{namespace}/servers/options/resources/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsResourcesReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsResourcesReadOK), nil

}

/*
ServersOptionsResourcesUpdate servers options resources update API
*/
func (a *Client) ServersOptionsResourcesUpdate(params *ServersOptionsResourcesUpdateParams) (*ServersOptionsResourcesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsResourcesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_resources_update",
		Method:             "PUT",
		PathPattern:        "/api/v0/{namespace}/servers/options/resources/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsResourcesUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsResourcesUpdateOK), nil

}

/*
ServersOptionsTypesCreate servers options types create API
*/
func (a *Client) ServersOptionsTypesCreate(params *ServersOptionsTypesCreateParams) (*ServersOptionsTypesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsTypesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_types_create",
		Method:             "POST",
		PathPattern:        "/api/v0/{namespace}/servers/options/types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsTypesCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsTypesCreateCreated), nil

}

/*
ServersOptionsTypesDelete servers options types delete API
*/
func (a *Client) ServersOptionsTypesDelete(params *ServersOptionsTypesDeleteParams) (*ServersOptionsTypesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsTypesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_types_delete",
		Method:             "DELETE",
		PathPattern:        "/api/v0/{namespace}/servers/options/types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsTypesDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsTypesDeleteNoContent), nil

}

/*
ServersOptionsTypesList servers options types list API
*/
func (a *Client) ServersOptionsTypesList(params *ServersOptionsTypesListParams) (*ServersOptionsTypesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsTypesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_types_list",
		Method:             "GET",
		PathPattern:        "/api/v0/{namespace}/servers/options/types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsTypesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsTypesListOK), nil

}

/*
ServersOptionsTypesPartialUpdate servers options types partial update API
*/
func (a *Client) ServersOptionsTypesPartialUpdate(params *ServersOptionsTypesPartialUpdateParams) (*ServersOptionsTypesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsTypesPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_types_partial_update",
		Method:             "PATCH",
		PathPattern:        "/api/v0/{namespace}/servers/options/types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsTypesPartialUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsTypesPartialUpdateOK), nil

}

/*
ServersOptionsTypesRead servers options types read API
*/
func (a *Client) ServersOptionsTypesRead(params *ServersOptionsTypesReadParams) (*ServersOptionsTypesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsTypesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_types_read",
		Method:             "GET",
		PathPattern:        "/api/v0/{namespace}/servers/options/types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsTypesReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsTypesReadOK), nil

}

/*
ServersOptionsTypesUpdate servers options types update API
*/
func (a *Client) ServersOptionsTypesUpdate(params *ServersOptionsTypesUpdateParams) (*ServersOptionsTypesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServersOptionsTypesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servers_options_types_update",
		Method:             "PUT",
		PathPattern:        "/api/v0/{namespace}/servers/options/types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServersOptionsTypesUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServersOptionsTypesUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}