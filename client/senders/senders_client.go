// Code generated by go-swagger; DO NOT EDIT.

package senders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new senders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for senders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSender creates a new sender
*/
func (a *Client) CreateSender(params *CreateSenderParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSenderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSenderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSender",
		Method:             "POST",
		PathPattern:        "/senders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSenderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSenderCreated), nil

}

/*
DeleteSender deletes a sender
*/
func (a *Client) DeleteSender(params *DeleteSenderParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSenderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSenderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSender",
		Method:             "DELETE",
		PathPattern:        "/senders/{senderId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSenderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSenderNoContent), nil

}

/*
GetIps returns all the dedicated ips for your account
*/
func (a *Client) GetIps(params *GetIpsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIpsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getIps",
		Method:             "GET",
		PathPattern:        "/senders/ips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIpsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIpsOK), nil

}

/*
GetIpsFromSender returns all the dedicated ips for a sender
*/
func (a *Client) GetIpsFromSender(params *GetIpsFromSenderParams, authInfo runtime.ClientAuthInfoWriter) (*GetIpsFromSenderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIpsFromSenderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getIpsFromSender",
		Method:             "GET",
		PathPattern:        "/senders/{senderId}/ips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIpsFromSenderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIpsFromSenderOK), nil

}

/*
GetSenders gets the list of all your senders
*/
func (a *Client) GetSenders(params *GetSendersParams, authInfo runtime.ClientAuthInfoWriter) (*GetSendersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSendersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSenders",
		Method:             "GET",
		PathPattern:        "/senders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSendersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSendersOK), nil

}

/*
UpdateSender updates a sender
*/
func (a *Client) UpdateSender(params *UpdateSenderParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSenderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSenderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSender",
		Method:             "PUT",
		PathPattern:        "/senders/{senderId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSenderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSenderNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
