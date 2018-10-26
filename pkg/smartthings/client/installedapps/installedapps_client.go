// Code generated by go-swagger; DO NOT EDIT.

package installedapps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new installedapps API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for installedapps API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteInstallation deletes an installed app

Delete an Installed App.
*/
func (a *Client) DeleteInstallation(params *DeleteInstallationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteInstallationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstallationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInstallation",
		Method:             "DELETE",
		PathPattern:        "/installedapps/{installedAppId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteInstallationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteInstallationOK), nil

}

/*
GetInstallation gets an installed app

Fetch a single installed application.
*/
func (a *Client) GetInstallation(params *GetInstallationParams, authInfo runtime.ClientAuthInfoWriter) (*GetInstallationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstallationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInstallation",
		Method:             "GET",
		PathPattern:        "/installedapps/{installedAppId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInstallationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInstallationOK), nil

}

/*
GetInstallationConfig gets an installed app configuration

Fetch a detailed install configuration model containing actual config entries / values.
*/
func (a *Client) GetInstallationConfig(params *GetInstallationConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetInstallationConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstallationConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInstallationConfig",
		Method:             "GET",
		PathPattern:        "/installedapps/{installedAppId}/configs/{configurationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInstallationConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInstallationConfigOK), nil

}

/*
ListInstallationConfig lists an installed app s configurations

List all configurations potentially filtered by status for an installed app.
*/
func (a *Client) ListInstallationConfig(params *ListInstallationConfigParams, authInfo runtime.ClientAuthInfoWriter) (*ListInstallationConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListInstallationConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listInstallationConfig",
		Method:             "GET",
		PathPattern:        "/installedapps/{installedAppId}/configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListInstallationConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListInstallationConfigOK), nil

}

/*
ListInstallations lists installed apps

List all installed applications within a location.
*/
func (a *Client) ListInstallations(params *ListInstallationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListInstallationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListInstallationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listInstallations",
		Method:             "GET",
		PathPattern:        "/installedapps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListInstallationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListInstallationsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}