# \ContactsApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContactToList**](ContactsApi.md#AddContactToList) | **Post** /contacts/lists/{listId}/contacts/add | Add existing contacts to a list
[**CreateAttribute**](ContactsApi.md#CreateAttribute) | **Post** /contacts/attributes | Creates contact attributes
[**CreateContact**](ContactsApi.md#CreateContact) | **Post** /contacts | Create a contact
[**CreateFolder**](ContactsApi.md#CreateFolder) | **Post** /contacts/folders | Create a folder
[**CreateList**](ContactsApi.md#CreateList) | **Post** /contacts/lists | Create a list
[**DeleteAttribute**](ContactsApi.md#DeleteAttribute) | **Delete** /contacts/attributes/{attributeId} | Deletes an attribute
[**DeleteFolder**](ContactsApi.md#DeleteFolder) | **Delete** /contacts/folders/{folderId} | Delete a folder (and all its lists)
[**DeleteList**](ContactsApi.md#DeleteList) | **Delete** /contacts/lists/{listId} | Delete a list
[**GetAttributes**](ContactsApi.md#GetAttributes) | **Get** /contacts/attributes | Lists all attributes
[**GetContactInfo**](ContactsApi.md#GetContactInfo) | **Get** /contacts/{email} | Retrieves contact informations
[**GetContactStats**](ContactsApi.md#GetContactStats) | **Get** /contacts/{email}/campaignStats | Get the campaigns statistics for a contact
[**GetContacts**](ContactsApi.md#GetContacts) | **Get** /contacts | Get all the contacts
[**GetContactsFromList**](ContactsApi.md#GetContactsFromList) | **Get** /contacts/lists/{listId}/contacts | Get the contacts in a list
[**GetFolder**](ContactsApi.md#GetFolder) | **Get** /contacts/folders/{folderId} | Returns folder details
[**GetFolderLists**](ContactsApi.md#GetFolderLists) | **Get** /contacts/folders/{folderId}/lists | Get the lists in a folder
[**GetFolders**](ContactsApi.md#GetFolders) | **Get** /contacts/folders | Get all the folders
[**GetList**](ContactsApi.md#GetList) | **Get** /contacts/lists/{listId} | Get the details of a list
[**GetLists**](ContactsApi.md#GetLists) | **Get** /contacts/lists | Get all the lists
[**ImportContacts**](ContactsApi.md#ImportContacts) | **Post** /contacts/import | Import contacts
[**RemoveContactToList**](ContactsApi.md#RemoveContactToList) | **Post** /contacts/lists/{listId}/contacts/remove | Remove existing contacts from a list
[**RequestContactExport**](ContactsApi.md#RequestContactExport) | **Post** /contacts/export | Export contacts
[**UpdateContact**](ContactsApi.md#UpdateContact) | **Put** /contacts/{email} | Updates a contact
[**UpdateFolder**](ContactsApi.md#UpdateFolder) | **Put** /contacts/folders/{folderId} | Update a contact folder
[**UpdateList**](ContactsApi.md#UpdateList) | **Put** /contacts/lists/{listId} | Update a list


# **AddContactToList**
> PostContactInfo AddContactToList($listId, $contactEmails)

Add existing contacts to a list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| Id of the list | 
 **contactEmails** | [**AddRemoveContactToList**](AddRemoveContactToList.md)| Emails addresses of the contacts | 

### Return type

[**PostContactInfo**](PostContactInfo.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAttribute**
> CreateModel CreateAttribute($createAttribute)

Creates contact attributes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAttribute** | [**CreateAttribute**](CreateAttribute.md)| Values to create an attribute | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContact**
> CreateModel CreateContact($createContact)

Create a contact


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContact** | [**CreateContact**](CreateContact.md)| Values to create a contact | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> CreateModel CreateFolder($name)

Create a folder


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | [**CreaUpdateFolder**](CreaUpdateFolder.md)| Name of the folder | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateList**
> CreateModel CreateList($createList)

Create a list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createList** | [**CreateList**](CreateList.md)| Values to create a list | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAttribute**
> DeleteAttribute($attributeId)

Deletes an attribute


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeId** | **string**| id of the attribute | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> DeleteFolder($folderId)

Delete a folder (and all its lists)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string**| Id of the folder | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteList**
> DeleteList($listId)

Delete a list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| Id of the list | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttributes**
> GetAttributes GetAttributes()

Lists all attributes


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetAttributes**](GetAttributes.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactInfo**
> GetExtendedContactDetails GetContactInfo($email)

Retrieves contact informations


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**| Email (urlencoded) of the contact | 

### Return type

[**GetExtendedContactDetails**](GetExtendedContactDetails.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactStats**
> GetContactCampaignStats GetContactStats($email)

Get the campaigns statistics for a contact


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**| Email address (urlencoded) of the contact | 

### Return type

[**GetContactCampaignStats**](GetContactCampaignStats.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContacts**
> GetContacts GetContacts($limit, $offset)

Get all the contacts


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Number of documents per page | [optional] [default to 50]
 **offset** | **int32**| Index of the first document of the page | [optional] [default to 0]

### Return type

[**GetContacts**](GetContacts.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactsFromList**
> GetContacts GetContactsFromList($listId, $modifiedSince, $limit, $offset)

Get the contacts in a list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| Id of the list | 
 **modifiedSince** | **string**| Filter the contacts modified after a given date (YYYY-MM-DD HH:mm:ss) | [optional] 
 **limit** | **int32**| Number of documents per page | [optional] [default to 50]
 **offset** | **int32**| Index of the first document of the page | [optional] [default to 0]

### Return type

[**GetContacts**](GetContacts.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolder**
> GetFolder GetFolder($folderId)

Returns folder details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string**| id of the folder | 

### Return type

[**GetFolder**](GetFolder.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolderLists**
> GetFolderLists GetFolderLists($folderId, $limit, $offset)

Get the lists in a folder


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string**| Id of the folder | 
 **limit** | **int32**| Number of documents per page | [optional] [default to 10]
 **offset** | **int32**| Index of the first document of the page | [optional] [default to 0]

### Return type

[**GetFolderLists**](GetFolderLists.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolders**
> GetFolders GetFolders($limit, $offset)

Get all the folders


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Number of documents per page | [default to 10]
 **offset** | **int32**| Index of the first document of the page | [default to 0]

### Return type

[**GetFolders**](GetFolders.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetList**
> GetExtendedList GetList($listId)

Get the details of a list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| Id of the list | 

### Return type

[**GetExtendedList**](GetExtendedList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLists**
> GetLists GetLists($limit, $offset)

Get all the lists


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Number of documents per page | [optional] [default to 10]
 **offset** | **int32**| Index of the first document of the page | [optional] [default to 0]

### Return type

[**GetLists**](GetLists.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportContacts**
> CreatedProcessId ImportContacts($requestContactImport)

Import contacts

It returns the background process ID which on completion calls the notify URL that you have set in the input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestContactImport** | [**RequestContactImport**](RequestContactImport.md)| Values to import contacts in Sendinblue. To know more about the expected format, please have a look at &#x60;&#x60;https://help.sendinblue.com/hc/en-us/articles/209499265-Build-contacts-lists-for-your-email-marketing-campaigns&#x60;&#x60; | 

### Return type

[**CreatedProcessId**](CreatedProcessId.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveContactToList**
> PostContactInfo RemoveContactToList($listId, $contactEmails)

Remove existing contacts from a list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| Id of the list | 
 **contactEmails** | [**AddRemoveContactToList**](AddRemoveContactToList.md)| Emails adresses of the contact | 

### Return type

[**PostContactInfo**](PostContactInfo.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestContactExport**
> CreatedProcessId RequestContactExport($requestContactExport)

Export contacts

It returns the background process ID which on completion calls the notify URL that you have set in the input. File will be available in csv.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestContactExport** | [**RequestContactExport**](RequestContactExport.md)| Values to request a contact export | 

### Return type

[**CreatedProcessId**](CreatedProcessId.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContact**
> UpdateContact($email, $updateContact)

Updates a contact


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**| Email (urlencoded) of the contact | 
 **updateContact** | [**UpdateContact**](UpdateContact.md)| Values to update a contact | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFolder**
> UpdateFolder($folderId, $name)

Update a contact folder


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string**| Id of the folder | 
 **name** | [**CreaUpdateFolder**](CreaUpdateFolder.md)| Name of the folder | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateList**
> UpdateList($listId, $updateList)

Update a list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| Id of the list | 
 **updateList** | [**UpdateList**](UpdateList.md)| Values to update a list | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

