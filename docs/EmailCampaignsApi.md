# \EmailCampaignsApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailCampaign**](EmailCampaignsApi.md#CreateEmailCampaign) | **Post** /emailCampaigns | Create an email campaign
[**DeleteEmailCampaigns**](EmailCampaignsApi.md#DeleteEmailCampaigns) | **Delete** /emailCampaigns/{campaignId} | Delete an email campaign
[**EmailExportRecipients**](EmailCampaignsApi.md#EmailExportRecipients) | **Post** /emailCampaigns/{campaignId}/exportRecipients | Export the recipients of a campaign
[**GetEmailCampaign**](EmailCampaignsApi.md#GetEmailCampaign) | **Get** /emailCampaigns/{campaignId} | Get campaign informations
[**GetEmailCampaigns**](EmailCampaignsApi.md#GetEmailCampaigns) | **Get** /emailCampaigns | Return all your created campaigns
[**SendEmailCampaignNow**](EmailCampaignsApi.md#SendEmailCampaignNow) | **Post** /emailCampaigns/{campaignId}/sendNow | Send an email campaign id of the campaign immediately
[**SendReport**](EmailCampaignsApi.md#SendReport) | **Post** /emailCampaigns/{campaignId}/sendReport | Send the report of a campaigns
[**SendTestEmail**](EmailCampaignsApi.md#SendTestEmail) | **Post** /emailCampaigns/{campaignId}/sendTest | Send an email campaign to your test list
[**UpdateCampaignStatus**](EmailCampaignsApi.md#UpdateCampaignStatus) | **Put** /emailCampaigns/{campaignId}/status | Update a campaign status
[**UpdateEmailCampaigns**](EmailCampaignsApi.md#UpdateEmailCampaigns) | **Put** /emailCampaigns/{campaignId} | Update a campaign


# **CreateEmailCampaign**
> CreateModel CreateEmailCampaign($emailCampaigns)

Create an email campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailCampaigns** | [**CreateEmailCampaign**](CreateEmailCampaign.md)| Values to create a campaign | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmailCampaigns**
> DeleteEmailCampaigns($campaignId)

Delete an email campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| id of the campaign | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmailExportRecipients**
> CreatedProcessId EmailExportRecipients($campaignId, $recipientExport)

Export the recipients of a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| Id of the campaign | 
 **recipientExport** | [**EmailExportRecipients**](EmailExportRecipients.md)| Values to send for a recipient export request | [optional] 

### Return type

[**CreatedProcessId**](CreatedProcessId.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailCampaign**
> GetEmailCampaign GetEmailCampaign($campaignId)

Get campaign informations


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| Id of the campaign | 

### Return type

[**GetEmailCampaign**](GetEmailCampaign.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailCampaigns**
> GetEmailCampaigns GetEmailCampaigns($type_, $status, $limit, $offset)

Return all your created campaigns


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| Filter on the type of the campaigns | [optional] 
 **status** | **string**| Filter on the status of the campaign | [optional] 
 **limit** | **int32**| Number of documents per page | [optional] [default to 500]
 **offset** | **int32**| Index of the first document in the page | [optional] [default to 0]

### Return type

[**GetEmailCampaigns**](GetEmailCampaigns.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendEmailCampaignNow**
> SendEmailCampaignNow($campaignId)

Send an email campaign id of the campaign immediately


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| Id of the campaign | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendReport**
> SendReport($campaignId, $sendReport)

Send the report of a campaigns

A PDF will be sent to the specified email addresses


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| Id of the campaign | 
 **sendReport** | [**SendReport**](SendReport.md)| Values for send a report | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestEmail**
> SendTestEmail($campaignId, $emailTo)

Send an email campaign to your test list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| Id of the campaign | 
 **emailTo** | [**SendTestEmail**](SendTestEmail.md)|  | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaignStatus**
> UpdateCampaignStatus($campaignId, $status)

Update a campaign status


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| Id of the campaign | 
 **status** | [**UpdateCampaignStatus**](UpdateCampaignStatus.md)| Status of the campaign | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmailCampaigns**
> UpdateEmailCampaigns($campaignId, $emailCampaign)

Update a campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| Id of the campaign | 
 **emailCampaign** | [**UpdateEmailCampaign**](UpdateEmailCampaign.md)| Values to update a campaign | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

