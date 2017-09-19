# \SMSCampaignsApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSMSCampaign**](SMSCampaignsApi.md#CreateSMSCampaign) | **Post** /smsCampaigns | Creates a SMS campaign
[**DeleteSMSCampaigns**](SMSCampaignsApi.md#DeleteSMSCampaigns) | **Delete** /smsCampaigns/{campaignId} | Delete the SMS campaign
[**GetSMSCampaigns**](SMSCampaignsApi.md#GetSMSCampaigns) | **Get** /smsCampaigns | Returns the informations for all your created SMS campaigns
[**GetSmsCampaign**](SMSCampaignsApi.md#GetSmsCampaign) | **Get** /smsCampaigns/{campaignId} | Get a SMS campaign
[**RequestSMSRecipientExport**](SMSCampaignsApi.md#RequestSMSRecipientExport) | **Post** /smsCampaigns/{campaignId}/exportRecipients | Exports the recipients of the specified campaign.
[**SendSMSCampaignNow**](SMSCampaignsApi.md#SendSMSCampaignNow) | **Post** /smsCampaigns/{campaignId}/sendNow | Send your SMS campaign immediately
[**SendSMSReport**](SMSCampaignsApi.md#SendSMSReport) | **Post** /smsCampaigns/{campaignId}/sendReport | Send report of SMS campaigns
[**SendTestSms**](SMSCampaignsApi.md#SendTestSms) | **Post** /smsCampaigns/{campaignId}/sendTest | Send an SMS
[**UpdateSMSCampaignStatus**](SMSCampaignsApi.md#UpdateSMSCampaignStatus) | **Put** /smsCampaigns/{campaignId}/status | Update the campaign status
[**UpdateSmsCampaign**](SMSCampaignsApi.md#UpdateSmsCampaign) | **Put** /smsCampaigns/{campaignId} | Updates a SMS campaign


# **CreateSMSCampaign**
> CreateModel CreateSMSCampaign($createSmsCampaign)

Creates a SMS campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSmsCampaign** | [**CreateSmsCampaign**](CreateSmsCampaign.md)| Values to create an SMS Campaign | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSMSCampaigns**
> DeleteSMSCampaigns($campaignId)

Delete the SMS campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| id of the SMS campaign | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSMSCampaigns**
> GetSmsCampaigns GetSMSCampaigns($status, $limit, $offset)

Returns the informations for all your created SMS campaigns


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string**| Status of campaign. | [optional] 
 **limit** | **int32**| Number limitation for the result returned | [optional] [default to 500]
 **offset** | **int32**| Beginning point in the list to retrieve from. | [optional] [default to 0]

### Return type

[**GetSmsCampaigns**](GetSmsCampaigns.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmsCampaign**
> GetSmsCampaign GetSmsCampaign($campaignId, $getSmsCampaign)

Get a SMS campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| id of the SMS campaign | 
 **getSmsCampaign** | [**GetSmsCampaign**](GetSmsCampaign.md)| Values to update an SMS Campaign | 

### Return type

[**GetSmsCampaign**](GetSmsCampaign.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSMSRecipientExport**
> CreatedProcessId RequestSMSRecipientExport($campaignId, $recipientExport)

Exports the recipients of the specified campaign.

It returns the background process ID which on completion calls the notify URL that you have set in the input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| id of the campaign | 
 **recipientExport** | [**RequestSmsRecipientExport**](RequestSmsRecipientExport.md)| Values to send for a recipient export request | [optional] 

### Return type

[**CreatedProcessId**](CreatedProcessId.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendSMSCampaignNow**
> SendSMSCampaignNow($campaignId)

Send your SMS campaign immediately


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

# **SendSMSReport**
> SendSMSReport($campaignId, $sendReport)

Send report of SMS campaigns

Send report of Sent and Archived campaign, to the specified email addresses, with respective data and a pdf attachment in detail.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| id of the campaign | 
 **sendReport** | [**SendReport**](SendReport.md)| Values for send a report | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestSms**
> SendTestSms($campaignId, $sendTestSms)

Send an SMS


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| Id of the SMS campaign | 
 **sendTestSms** | [**SendTestSms**](SendTestSms.md)| Mobile number to which send the test | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSMSCampaignStatus**
> UpdateSMSCampaignStatus($campaignId, $status)

Update the campaign status


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| id of the campaign | 
 **status** | [**UpdateCampaignStatus**](UpdateCampaignStatus.md)| Status of the campaign. | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmsCampaign**
> UpdateSmsCampaign($campaignId, $updateSmsCampaign)

Updates a SMS campaign


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **string**| id of the SMS campaign | 
 **updateSmsCampaign** | [**UpdateSmsCampaign**](UpdateSmsCampaign.md)| Values to update an SMS Campaign | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

