# \CRMApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmNotesGet**](CRMApi.md#CrmNotesGet) | **Get** /crm/notes | Get all notes
[**CrmNotesIdDelete**](CRMApi.md#CrmNotesIdDelete) | **Delete** /crm/notes/{id} | Delete a note
[**CrmNotesIdGet**](CRMApi.md#CrmNotesIdGet) | **Get** /crm/notes/{id} | Get a note
[**CrmNotesIdPatch**](CRMApi.md#CrmNotesIdPatch) | **Patch** /crm/notes/{id} | Update a note
[**CrmNotesPost**](CRMApi.md#CrmNotesPost) | **Post** /crm/notes | Create a note
[**CrmTasksGet**](CRMApi.md#CrmTasksGet) | **Get** /crm/tasks | Get all tasks
[**CrmTasksIdDelete**](CRMApi.md#CrmTasksIdDelete) | **Delete** /crm/tasks/{id} | Delete a task
[**CrmTasksIdGet**](CRMApi.md#CrmTasksIdGet) | **Get** /crm/tasks/{id} | Get a task
[**CrmTasksIdPatch**](CRMApi.md#CrmTasksIdPatch) | **Patch** /crm/tasks/{id} | Update a task
[**CrmTasksPost**](CRMApi.md#CrmTasksPost) | **Post** /crm/tasks | Create a task
[**CrmTasktypesGet**](CRMApi.md#CrmTasktypesGet) | **Get** /crm/tasktypes | Get all task types


# **CrmNotesGet**
> NoteList CrmNotesGet(ctx, optional)
Get all notes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrmNotesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrmNotesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **optional.String**| Filter by note entity type | 
 **entityIds** | **optional.String**| Filter by note entity IDs | 
 **dateFrom** | **optional.Int32**| dateFrom to date range filter type (timestamp in milliseconds) | 
 **dateTo** | **optional.Int32**| dateTo to date range filter type (timestamp in milliseconds) | 
 **offset** | **optional.Int64**| Index of the first document of the page | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **sort** | **optional.String**| Sort the results in the ascending/descending order. Default order is **descending** by creation if &#x60;sort&#x60; is not passed | 

### Return type

[**NoteList**](NoteList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmNotesIdDelete**
> CrmNotesIdDelete(ctx, id)
Delete a note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Note ID to delete | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmNotesIdGet**
> Note CrmNotesIdGet(ctx, id)
Get a note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Note ID to get | 

### Return type

[**Note**](Note.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmNotesIdPatch**
> CrmNotesIdPatch(ctx, id, body)
Update a note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Note ID to update | 
  **body** | [**NoteData**](NoteData.md)| Note data to update a note | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmNotesPost**
> NoteId CrmNotesPost(ctx, body)
Create a note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NoteData**](NoteData.md)| Note data to create a note. | 

### Return type

[**NoteId**](NoteId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksGet**
> TaskList CrmTasksGet(ctx, optional)
Get all tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrmTasksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrmTasksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterType** | **optional.String**| Filter by task type (ID) | 
 **filterStatus** | **optional.String**| Filter by task status | 
 **filterDate** | **optional.String**| Filter by date | 
 **filterAssignTo** | **optional.String**| Filter by assignTo id | 
 **filterContacts** | **optional.String**| Filter by contact ids | 
 **filterDeals** | **optional.String**| Filter by deals ids | 
 **filterCompanies** | **optional.String**| Filter by companies ids | 
 **dateFrom** | **optional.Int32**| dateFrom to date range filter type (timestamp in milliseconds) | 
 **dateTo** | **optional.Int32**| dateTo to date range filter type (timestamp in milliseconds) | 
 **offset** | **optional.Int64**| Index of the first document of the page | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **sort** | **optional.String**| Sort the results in the ascending/descending order. Default order is **descending** by creation if &#x60;sort&#x60; is not passed | 
 **sortBy** | **optional.String**| The field used to sort field names. | 

### Return type

[**TaskList**](TaskList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksIdDelete**
> CrmTasksIdDelete(ctx, id)
Delete a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksIdGet**
> Task CrmTasksIdGet(ctx, id)
Get a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Task**](Task.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksIdPatch**
> CrmTasksIdPatch(ctx, id, body)
Update a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**Body6**](Body6.md)| Updated task details. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksPost**
> InlineResponse2011 CrmTasksPost(ctx, body)
Create a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body5**](Body5.md)| Task name. | 

### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasktypesGet**
> TaskTypes CrmTasktypesGet(ctx, )
Get all task types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TaskTypes**](TaskTypes.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

