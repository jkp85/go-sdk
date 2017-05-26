package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new auth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AuthConvertTokenCreate implements an endpoint to convert a provider token to an access token

Implements an endpoint to convert a provider token to an access token

The endpoint is used in the following flows:

* Authorization code
* Client credentials
*/
func (a *Client) AuthConvertTokenCreate(params *AuthConvertTokenCreateParams) (*AuthConvertTokenCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthConvertTokenCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_convert-token_create",
		Method:             "POST",
		PathPattern:        "/auth/convert-token/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthConvertTokenCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthConvertTokenCreateCreated), nil

}

/*
AuthInvalidateSessionsCreate auth invalidate sessions create API
*/
func (a *Client) AuthInvalidateSessionsCreate(params *AuthInvalidateSessionsCreateParams) (*AuthInvalidateSessionsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthInvalidateSessionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_invalidate-sessions_create",
		Method:             "POST",
		PathPattern:        "/auth/invalidate-sessions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthInvalidateSessionsCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthInvalidateSessionsCreateCreated), nil

}

/*
AuthJwtTokenAuthCreate auth jwt token auth create API
*/
func (a *Client) AuthJwtTokenAuthCreate(params *AuthJwtTokenAuthCreateParams) (*AuthJwtTokenAuthCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthJwtTokenAuthCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_jwt-token-auth_create",
		Method:             "POST",
		PathPattern:        "/auth/jwt-token-auth/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthJwtTokenAuthCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthJwtTokenAuthCreateCreated), nil

}

/*
AuthJwtTokenRefreshCreate APIs view that returns a refreshed token with new expiration based on

API View that returns a refreshed token (with new expiration) based on
existing token

If 'orig_iat' field (original issued-at-time) is found, will first check
if it's within expiration window, then copy it to the new token
*/
func (a *Client) AuthJwtTokenRefreshCreate(params *AuthJwtTokenRefreshCreateParams) (*AuthJwtTokenRefreshCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthJwtTokenRefreshCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_jwt-token-refresh_create",
		Method:             "POST",
		PathPattern:        "/auth/jwt-token-refresh/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthJwtTokenRefreshCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthJwtTokenRefreshCreateCreated), nil

}

/*
AuthJwtTokenVerifyCreate APIs view that checks the veracity of a token returning the token if it

API View that checks the veracity of a token, returning the token if it
is valid.
*/
func (a *Client) AuthJwtTokenVerifyCreate(params *AuthJwtTokenVerifyCreateParams) (*AuthJwtTokenVerifyCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthJwtTokenVerifyCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_jwt-token-verify_create",
		Method:             "POST",
		PathPattern:        "/auth/jwt-token-verify/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthJwtTokenVerifyCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthJwtTokenVerifyCreateCreated), nil

}

/*
AuthRevokeTokenCreate implements an endpoint to revoke access or refresh tokens

Implements an endpoint to revoke access or refresh tokens
*/
func (a *Client) AuthRevokeTokenCreate(params *AuthRevokeTokenCreateParams) (*AuthRevokeTokenCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthRevokeTokenCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_revoke-token_create",
		Method:             "POST",
		PathPattern:        "/auth/revoke-token/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthRevokeTokenCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthRevokeTokenCreateCreated), nil

}

/*
AuthSimpleTokenAuthCreate auth simple token auth create API
*/
func (a *Client) AuthSimpleTokenAuthCreate(params *AuthSimpleTokenAuthCreateParams) (*AuthSimpleTokenAuthCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthSimpleTokenAuthCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_simple-token-auth_create",
		Method:             "POST",
		PathPattern:        "/auth/simple-token-auth/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthSimpleTokenAuthCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthSimpleTokenAuthCreateCreated), nil

}

/*
AuthTokenCreate implements an endpoint to provide access tokens

Implements an endpoint to provide access tokens

The endpoint is used in the following flows:

* Authorization code
* Password
* Client credentials
*/
func (a *Client) AuthTokenCreate(params *AuthTokenCreateParams) (*AuthTokenCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthTokenCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_token_create",
		Method:             "POST",
		PathPattern:        "/auth/token/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AuthTokenCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthTokenCreateCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
