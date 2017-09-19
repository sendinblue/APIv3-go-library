# \ProcessApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProcess**](ProcessApi.md#GetProcess) | **Get** /processes/{processId} | Return the informations for a process
[**GetProcesses**](ProcessApi.md#GetProcesses) | **Get** /processes | Return all the processes for your account


# **GetProcess**
> GetProcess GetProcess($processId)

Return the informations for a process


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Id of the process | 

### Return type

[**GetProcess**](GetProcess.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcesses**
> GetProcesses GetProcesses($limit, $offset)

Return all the processes for your account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Number limitation for the result returned | [optional] [default to 10]
 **offset** | **int32**| Beginning point in the list to retrieve from. | [optional] [default to 0]

### Return type

[**GetProcesses**](GetProcesses.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

