// Code generated by go-swagger; DO NOT EDIT.

package reseller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new reseller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reseller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddCredits adds email and or s m s credits to a specific child account
*/
func (a *Client) AddCredits(params *AddCreditsParams, authInfo runtime.ClientAuthInfoWriter) (*AddCreditsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddCreditsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addCredits",
		Method:             "POST",
		PathPattern:        "/reseller/children/{childId}/credits/add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddCreditsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddCreditsOK), nil

}

/*
AssociateIPToChild associates a dedicated IP to the child
*/
func (a *Client) AssociateIPToChild(params *AssociateIPToChildParams, authInfo runtime.ClientAuthInfoWriter) (*AssociateIPToChildNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssociateIPToChildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "associateIpToChild",
		Method:             "POST",
		PathPattern:        "/reseller/children/{childId}/ips/associate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssociateIPToChildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AssociateIPToChildNoContent), nil

}

/*
CreateResellerChild creates a reseller child
*/
func (a *Client) CreateResellerChild(params *CreateResellerChildParams, authInfo runtime.ClientAuthInfoWriter) (*CreateResellerChildCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateResellerChildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createResellerChild",
		Method:             "POST",
		PathPattern:        "/reseller/children",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateResellerChildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateResellerChildCreated), nil

}

/*
DeleteResellerChild deletes a single reseller child based on the child Id supplied
*/
func (a *Client) DeleteResellerChild(params *DeleteResellerChildParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteResellerChildNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteResellerChildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteResellerChild",
		Method:             "DELETE",
		PathPattern:        "/reseller/children/{childId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteResellerChildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteResellerChildNoContent), nil

}

/*
DissociateIPFromChild dissociates a dedicated IP to the child
*/
func (a *Client) DissociateIPFromChild(params *DissociateIPFromChildParams, authInfo runtime.ClientAuthInfoWriter) (*DissociateIPFromChildNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDissociateIPFromChildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dissociateIpFromChild",
		Method:             "POST",
		PathPattern:        "/reseller/children/{childId}/ips/dissociate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DissociateIPFromChildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DissociateIPFromChildNoContent), nil

}

/*
GetChildInfo gets the info about a specific child account
*/
func (a *Client) GetChildInfo(params *GetChildInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetChildInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChildInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getChildInfo",
		Method:             "GET",
		PathPattern:        "/reseller/children/{childId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetChildInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetChildInfoOK), nil

}

/*
GetResellerChilds gets the list of all reseller s children accounts
*/
func (a *Client) GetResellerChilds(params *GetResellerChildsParams, authInfo runtime.ClientAuthInfoWriter) (*GetResellerChildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResellerChildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResellerChilds",
		Method:             "GET",
		PathPattern:        "/reseller/children",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResellerChildsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetResellerChildsOK), nil

}

/*
RemoveCredits removes email and or s m s credits from a specific child account
*/
func (a *Client) RemoveCredits(params *RemoveCreditsParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveCreditsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveCreditsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeCredits",
		Method:             "POST",
		PathPattern:        "/reseller/children/{childId}/credits/remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveCreditsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveCreditsOK), nil

}

/*
UpdateResellerChild updates infos of reseller s child based on the child Id supplied
*/
func (a *Client) UpdateResellerChild(params *UpdateResellerChildParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateResellerChildNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateResellerChildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateResellerChild",
		Method:             "PUT",
		PathPattern:        "/reseller/children/{childId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateResellerChildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateResellerChildNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
