# \SendersApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSender**](SendersApi.md#CreateSender) | **Post** /senders | Create a new sender
[**DeleteSender**](SendersApi.md#DeleteSender) | **Delete** /senders/{senderId} | Delete a sender
[**GetIps**](SendersApi.md#GetIps) | **Get** /senders/ips | Return all the dedicated IPs for your account
[**GetIpsFromSender**](SendersApi.md#GetIpsFromSender) | **Get** /senders/{senderId}/ips | Return all the dedicated IPs for a sender
[**GetSenders**](SendersApi.md#GetSenders) | **Get** /senders | Get the list of all your senders
[**UpdateSender**](SendersApi.md#UpdateSender) | **Put** /senders/{senderId} | Update a sender


# **CreateSender**
> CreateSenderModel CreateSender($sender)

Create a new sender


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sender** | [**CreateSender**](CreateSender.md)| sender&#39;s name | [optional] 

### Return type

[**CreateSenderModel**](CreateSenderModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSender**
> DeleteSender($senderId)

Delete a sender


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderId** | **string**| Id of the sender | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIps**
> GetIps GetIps()

Return all the dedicated IPs for your account


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetIps**](GetIps.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpsFromSender**
> GetIpsFromSender GetIpsFromSender($senderId)

Return all the dedicated IPs for a sender


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderId** | **string**| Id of the sender | 

### Return type

[**GetIpsFromSender**](getIpsFromSender.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenders**
> GetSendersList GetSenders($ip, $domain)

Get the list of all your senders


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string**| Filter your senders for a specific ip (available for dedicated IP usage only) | [optional] 
 **domain** | **string**| Filter your senders for a specific domain | [optional] 

### Return type

[**GetSendersList**](GetSendersList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSender**
> UpdateSender($senderId, $sender)

Update a sender


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **senderId** | **string**| Id of the sender | 
 **sender** | [**UpdateSender**](UpdateSender.md)| sender&#39;s name | [optional] 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

