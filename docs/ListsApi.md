# \ListsApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContactToList**](ListsApi.md#AddContactToList) | **Post** /contacts/lists/{listId}/contacts/add | Add existing contacts to a list
[**CreateList**](ListsApi.md#CreateList) | **Post** /contacts/lists | Create a list
[**DeleteList**](ListsApi.md#DeleteList) | **Delete** /contacts/lists/{listId} | Delete a list
[**GetContactsFromList**](ListsApi.md#GetContactsFromList) | **Get** /contacts/lists/{listId}/contacts | Get the contacts in a list
[**GetFolderLists**](ListsApi.md#GetFolderLists) | **Get** /contacts/folders/{folderId}/lists | Get the lists in a folder
[**GetList**](ListsApi.md#GetList) | **Get** /contacts/lists/{listId} | Get the details of a list
[**GetLists**](ListsApi.md#GetLists) | **Get** /contacts/lists | Get all the lists
[**RemoveContactToList**](ListsApi.md#RemoveContactToList) | **Post** /contacts/lists/{listId}/contacts/remove | Remove existing contacts from a list
[**UpdateList**](ListsApi.md#UpdateList) | **Put** /contacts/lists/{listId} | Update a list


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

