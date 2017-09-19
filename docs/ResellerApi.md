# \ResellerApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredits**](ResellerApi.md#AddCredits) | **Post** /reseller/children/{childId}/credits/add | Add Email and/or SMS credits to a specific child account
[**AssociateIpToChild**](ResellerApi.md#AssociateIpToChild) | **Post** /reseller/children/{childId}/ips/associate | Associate a dedicated IP to the child
[**CreateResellerChild**](ResellerApi.md#CreateResellerChild) | **Post** /reseller/children | Creates a reseller child
[**DeleteResellerChild**](ResellerApi.md#DeleteResellerChild) | **Delete** /reseller/children/{childId} | Deletes a single reseller child based on the childId supplied
[**DissociateIpFromChild**](ResellerApi.md#DissociateIpFromChild) | **Post** /reseller/children/{childId}/ips/dissociate | Dissociate a dedicated IP to the child
[**GetChildInfo**](ResellerApi.md#GetChildInfo) | **Get** /reseller/children/{childId} | Gets the info about a specific child account
[**GetResellerChilds**](ResellerApi.md#GetResellerChilds) | **Get** /reseller/children | Gets the list of all reseller&#39;s children accounts
[**RemoveCredits**](ResellerApi.md#RemoveCredits) | **Post** /reseller/children/{childId}/credits/remove | Remove Email and/or SMS credits from a specific child account
[**UpdateResellerChild**](ResellerApi.md#UpdateResellerChild) | **Put** /reseller/children/{childId} | Updates infos of reseller&#39;s child based on the childId supplied


# **AddCredits**
> RemainingCreditModel AddCredits($childId, $addCredits)

Add Email and/or SMS credits to a specific child account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childId** | **int32**| id of reseller&#39;s child | 
 **addCredits** | [**AddCredits**](AddCredits.md)| Values to post to add credit to a specific child account | 

### Return type

[**RemainingCreditModel**](RemainingCreditModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssociateIpToChild**
> AssociateIpToChild($childId, $ipId)

Associate a dedicated IP to the child


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childId** | **int32**| id of reseller&#39;s child | 
 **ipId** | [**ManageIp**](ManageIp.md)| IP&#39;s id | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateResellerChild**
> CreateModel CreateResellerChild($resellerChild)

Creates a reseller child


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerChild** | [**CreateChild**](CreateChild.md)| reseller child to add | [optional] 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResellerChild**
> DeleteResellerChild($childId)

Deletes a single reseller child based on the childId supplied


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childId** | **int32**| id of reseller&#39;s child | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DissociateIpFromChild**
> DissociateIpFromChild($childId, $ipId)

Dissociate a dedicated IP to the child


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childId** | **int32**| id of reseller&#39;s child | 
 **ipId** | [**ManageIp**](ManageIp.md)| IP&#39;s id | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildInfo**
> GetChildInfo GetChildInfo($childId)

Gets the info about a specific child account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childId** | **int32**| id of reseller&#39;s child | 

### Return type

[**GetChildInfo**](GetChildInfo.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResellerChilds**
> GetChildrenList GetResellerChilds()

Gets the list of all reseller's children accounts


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetChildrenList**](GetChildrenList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCredits**
> RemainingCreditModel RemoveCredits($childId, $removeCredits)

Remove Email and/or SMS credits from a specific child account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childId** | **int32**| id of reseller&#39;s child | 
 **removeCredits** | [**RemoveCredits**](RemoveCredits.md)| Values to post to remove email or SMS credits from a specific child account | 

### Return type

[**RemainingCreditModel**](RemainingCreditModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResellerChild**
> UpdateResellerChild($childId, $resellerChild)

Updates infos of reseller's child based on the childId supplied


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childId** | **int32**| id of reseller&#39;s child | 
 **resellerChild** | [**UpdateChild**](UpdateChild.md)| values to update in child profile | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

