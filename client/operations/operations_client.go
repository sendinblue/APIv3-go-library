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
AddContactToList adds existing contacts to a list
*/
func (a *Client) AddContactToList(params *AddContactToListParams, authInfo runtime.ClientAuthInfoWriter) (*AddContactToListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddContactToListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addContactToList",
		Method:             "POST",
		PathPattern:        "/contacts/lists/{listId}/contacts/add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddContactToListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddContactToListCreated), nil

}

/*
CreateAttribute creates contact attribute
*/
func (a *Client) CreateAttribute(params *CreateAttributeParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAttributeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAttributeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAttribute",
		Method:             "POST",
		PathPattern:        "/contacts/attributes/{attributeCategory}/{attributeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAttributeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAttributeCreated), nil

}

/*
CreateFolder creates a folder
*/
func (a *Client) CreateFolder(params *CreateFolderParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFolderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFolderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createFolder",
		Method:             "POST",
		PathPattern:        "/contacts/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateFolderCreated), nil

}

/*
CreateList creates a list
*/
func (a *Client) CreateList(params *CreateListParams, authInfo runtime.ClientAuthInfoWriter) (*CreateListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createList",
		Method:             "POST",
		PathPattern:        "/contacts/lists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateListCreated), nil

}

/*
DeleteAttribute deletes an attribute
*/
func (a *Client) DeleteAttribute(params *DeleteAttributeParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAttributeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAttributeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAttribute",
		Method:             "DELETE",
		PathPattern:        "/contacts/attributes/{attributeCategory}/{attributeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAttributeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAttributeNoContent), nil

}

/*
DeleteFolder deletes a folder and all its lists
*/
func (a *Client) DeleteFolder(params *DeleteFolderParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFolderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFolderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFolder",
		Method:             "DELETE",
		PathPattern:        "/contacts/folders/{folderId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFolderNoContent), nil

}

/*
DeleteList deletes a list
*/
func (a *Client) DeleteList(params *DeleteListParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteListNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteList",
		Method:             "DELETE",
		PathPattern:        "/contacts/lists/{listId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteListNoContent), nil

}

/*
GetAttributes lists all attributes
*/
func (a *Client) GetAttributes(params *GetAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAttributes",
		Method:             "GET",
		PathPattern:        "/contacts/attributes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAttributesOK), nil

}

/*
GetContactsFromList gets the contacts in a list
*/
func (a *Client) GetContactsFromList(params *GetContactsFromListParams, authInfo runtime.ClientAuthInfoWriter) (*GetContactsFromListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContactsFromListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getContactsFromList",
		Method:             "GET",
		PathPattern:        "/contacts/lists/{listId}/contacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetContactsFromListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetContactsFromListOK), nil

}

/*
GetFolder returns folder details
*/
func (a *Client) GetFolder(params *GetFolderParams, authInfo runtime.ClientAuthInfoWriter) (*GetFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFolderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFolder",
		Method:             "GET",
		PathPattern:        "/contacts/folders/{folderId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFolderOK), nil

}

/*
GetFolderLists gets the lists in a folder
*/
func (a *Client) GetFolderLists(params *GetFolderListsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFolderListsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFolderListsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFolderLists",
		Method:             "GET",
		PathPattern:        "/contacts/folders/{folderId}/lists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFolderListsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFolderListsOK), nil

}

/*
GetFolders gets all the folders
*/
func (a *Client) GetFolders(params *GetFoldersParams, authInfo runtime.ClientAuthInfoWriter) (*GetFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoldersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFolders",
		Method:             "GET",
		PathPattern:        "/contacts/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoldersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFoldersOK), nil

}

/*
GetList gets the details of a list
*/
func (a *Client) GetList(params *GetListParams, authInfo runtime.ClientAuthInfoWriter) (*GetListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getList",
		Method:             "GET",
		PathPattern:        "/contacts/lists/{listId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetListOK), nil

}

/*
GetLists gets all the lists
*/
func (a *Client) GetLists(params *GetListsParams, authInfo runtime.ClientAuthInfoWriter) (*GetListsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetListsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLists",
		Method:             "GET",
		PathPattern:        "/contacts/lists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetListsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetListsOK), nil

}

/*
RemoveContactToList removes existing contacts from a list
*/
func (a *Client) RemoveContactToList(params *RemoveContactToListParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveContactToListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveContactToListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeContactToList",
		Method:             "POST",
		PathPattern:        "/contacts/lists/{listId}/contacts/remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveContactToListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveContactToListCreated), nil

}

/*
UpdateAttribute updates contact attribute
*/
func (a *Client) UpdateAttribute(params *UpdateAttributeParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAttributeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAttributeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAttribute",
		Method:             "PUT",
		PathPattern:        "/contacts/attributes/{attributeCategory}/{attributeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAttributeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAttributeNoContent), nil

}

/*
UpdateFolder updates a contact folder
*/
func (a *Client) UpdateFolder(params *UpdateFolderParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFolderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFolderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFolder",
		Method:             "PUT",
		PathPattern:        "/contacts/folders/{folderId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateFolderNoContent), nil

}

/*
UpdateList updates a list
*/
func (a *Client) UpdateList(params *UpdateListParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateListNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateList",
		Method:             "PUT",
		PathPattern:        "/contacts/lists/{listId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateListNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
