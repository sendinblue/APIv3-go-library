# \AttributesApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttribute**](AttributesApi.md#CreateAttribute) | **Post** /contacts/attributes | Creates contact attributes
[**DeleteAttribute**](AttributesApi.md#DeleteAttribute) | **Delete** /contacts/attributes/{attributeId} | Deletes an attribute
[**GetAttributes**](AttributesApi.md#GetAttributes) | **Get** /contacts/attributes | Lists all attributes


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

