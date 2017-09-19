# \FoldersApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FoldersApi.md#CreateFolder) | **Post** /contacts/folders | Create a folder
[**DeleteFolder**](FoldersApi.md#DeleteFolder) | **Delete** /contacts/folders/{folderId} | Delete a folder (and all its lists)
[**GetFolder**](FoldersApi.md#GetFolder) | **Get** /contacts/folders/{folderId} | Returns folder details
[**GetFolderLists**](FoldersApi.md#GetFolderLists) | **Get** /contacts/folders/{folderId}/lists | Get the lists in a folder
[**GetFolders**](FoldersApi.md#GetFolders) | **Get** /contacts/folders | Get all the folders
[**UpdateFolder**](FoldersApi.md#UpdateFolder) | **Put** /contacts/folders/{folderId} | Update a contact folder


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

# **GetFolder**
> GetFolder GetFolder($folderId)

Returns folder details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string**| id of the folder | 

### Return type

[**GetFolder**](getFolder.md)

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

