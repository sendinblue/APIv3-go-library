# \AttributesApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttribute**](AttributesApi.md#CreateAttribute) | **Post** /contacts/attributes/{attributeCategory}/{attributeName} | Create contact attribute
[**DeleteAttribute**](AttributesApi.md#DeleteAttribute) | **Delete** /contacts/attributes/{attributeCategory}/{attributeName} | Delete an attribute
[**GetAttributes**](AttributesApi.md#GetAttributes) | **Get** /contacts/attributes | List all attributes
[**UpdateAttribute**](AttributesApi.md#UpdateAttribute) | **Put** /contacts/attributes/{attributeCategory}/{attributeName} | Update contact attribute


# **CreateAttribute**
> CreateAttribute(ctx, attributeCategory, attributeName, createAttribute)
Create contact attribute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attributeCategory** | **string**| Category of the attribute | 
  **attributeName** | **string**| Name of the attribute | 
  **createAttribute** | [**CreateAttribute**](CreateAttribute.md)| Values to create an attribute | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAttribute**
> DeleteAttribute(ctx, attributeCategory, attributeName)
Delete an attribute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attributeCategory** | **string**| Category of the attribute | 
  **attributeName** | **string**| Name of the existing attribute | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttributes**
> GetAttributes GetAttributes(ctx, )
List all attributes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAttributes**](GetAttributes.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAttribute**
> UpdateAttribute(ctx, attributeCategory, attributeName, updateAttribute)
Update contact attribute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attributeCategory** | **string**| Category of the attribute | 
  **attributeName** | **string**| Name of the existing attribute | 
  **updateAttribute** | [**UpdateAttribute**](UpdateAttribute.md)| Values to update an attribute | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
