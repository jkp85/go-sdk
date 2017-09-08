// Code generated by go-swagger; DO NOT EDIT.

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
AuthJwtTokenAuth creates JSON web token j w t
*/
func (a *Client) AuthJwtTokenAuth(params *AuthJwtTokenAuthParams) (*AuthJwtTokenAuthCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthJwtTokenAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_jwt-token-auth",
		Method:             "POST",
		PathPattern:        "/auth/jwt-token-auth/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthJwtTokenAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthJwtTokenAuthCreated), nil

}

/*
AuthJwtTokenRefresh refreshes a JSON web token j w t

Obtains a new JSON Web Token using existing user credentials.
*/
func (a *Client) AuthJwtTokenRefresh(params *AuthJwtTokenRefreshParams) (*AuthJwtTokenRefreshCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthJwtTokenRefreshParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_jwt-token-refresh",
		Method:             "POST",
		PathPattern:        "/auth/jwt-token-refresh/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthJwtTokenRefreshReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthJwtTokenRefreshCreated), nil

}

/*
AuthJwtTokenVerify validates JSON web token j w t

Checks veraciy of token.
*/
func (a *Client) AuthJwtTokenVerify(params *AuthJwtTokenVerifyParams) (*AuthJwtTokenVerifyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthJwtTokenVerifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_jwt-token-verify",
		Method:             "POST",
		PathPattern:        "/auth/jwt-token-verify/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthJwtTokenVerifyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthJwtTokenVerifyCreated), nil

}

/*
AuthRegister registers a user

User registration requires confirming email address to activate user.
*/
func (a *Client) AuthRegister(params *AuthRegisterParams) (*AuthRegisterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthRegisterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "auth_register",
		Method:             "POST",
		PathPattern:        "/auth/register/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthRegisterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthRegisterCreated), nil

}

/*
OauthLogin oauth login API
*/
func (a *Client) OauthLogin(params *OauthLoginParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOauthLoginParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "oauth_login",
		Method:             "GET",
		PathPattern:        "/auth/login/{provider}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OauthLoginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
