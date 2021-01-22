# {{classname}}

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailCampaign**](EmailCampaignsApi.md#CreateEmailCampaign) | **Post** /emailCampaigns | Create an email campaign
[**DeleteEmailCampaign**](EmailCampaignsApi.md#DeleteEmailCampaign) | **Delete** /emailCampaigns/{campaignId} | Delete an email campaign
[**EmailExportRecipients**](EmailCampaignsApi.md#EmailExportRecipients) | **Post** /emailCampaigns/{campaignId}/exportRecipients | Export the recipients of an email campaign
[**GetAbTestCampaignResult**](EmailCampaignsApi.md#GetAbTestCampaignResult) | **Get** /emailCampaigns/{campaignId}/abTestCampaignResult | Get an A/B test email campaign results
[**GetEmailCampaign**](EmailCampaignsApi.md#GetEmailCampaign) | **Get** /emailCampaigns/{campaignId} | Get an email campaign report
[**GetEmailCampaigns**](EmailCampaignsApi.md#GetEmailCampaigns) | **Get** /emailCampaigns | Return all your created email campaigns
[**GetSharedTemplateUrl**](EmailCampaignsApi.md#GetSharedTemplateUrl) | **Get** /emailCampaigns/{campaignId}/sharedUrl | Get a shared template url
[**SendEmailCampaignNow**](EmailCampaignsApi.md#SendEmailCampaignNow) | **Post** /emailCampaigns/{campaignId}/sendNow | Send an email campaign immediately, based on campaignId
[**SendReport**](EmailCampaignsApi.md#SendReport) | **Post** /emailCampaigns/{campaignId}/sendReport | Send the report of a campaign
[**SendTestEmail**](EmailCampaignsApi.md#SendTestEmail) | **Post** /emailCampaigns/{campaignId}/sendTest | Send an email campaign to your test list
[**UpdateCampaignStatus**](EmailCampaignsApi.md#UpdateCampaignStatus) | **Put** /emailCampaigns/{campaignId}/status | Update an email campaign status
[**UpdateEmailCampaign**](EmailCampaignsApi.md#UpdateEmailCampaign) | **Put** /emailCampaigns/{campaignId} | Update an email campaign
[**UploadImageToGallery**](EmailCampaignsApi.md#UploadImageToGallery) | **Post** /emailCampaigns/images | Upload an image to your account&#x27;s image gallery

# **CreateEmailCampaign**
> CreateModel CreateEmailCampaign(ctx, body)
Create an email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEmailCampaign**](CreateEmailCampaign.md)| Values to create a campaign | 

### Return type

[**CreateModel**](createModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmailCampaign**
> DeleteEmailCampaign(ctx, campaignId)
Delete an email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmailExportRecipients**
> CreatedProcessId EmailExportRecipients(ctx, campaignId, optional)
Export the recipients of an email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 
 **optional** | ***EmailCampaignsApiEmailExportRecipientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailCampaignsApiEmailExportRecipientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EmailExportRecipients**](EmailExportRecipients.md)| Values to send for a recipient export request | 

### Return type

[**CreatedProcessId**](createdProcessId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAbTestCampaignResult**
> AbTestCampaignResult GetAbTestCampaignResult(ctx, campaignId)
Get an A/B test email campaign results

Obtain winning version of an A/B test email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the A/B test campaign | 

### Return type

[**AbTestCampaignResult**](abTestCampaignResult.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailCampaign**
> GetEmailCampaign GetEmailCampaign(ctx, campaignId)
Get an email campaign report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 

### Return type

[**GetEmailCampaign**](getEmailCampaign.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailCampaigns**
> GetEmailCampaigns GetEmailCampaigns(ctx, optional)
Return all your created email campaigns

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailCampaignsApiGetEmailCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailCampaignsApiGetEmailCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Filter on the type of the campaigns | 
 **status** | **optional.String**| Filter on the status of the campaign | 
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the sent email campaigns. Prefer to pass your timezone in date-time format for accurate result ( only available if either &#x27;status&#x27; not passed and if passed is set to &#x27;sent&#x27; ) | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the sent email campaigns. Prefer to pass your timezone in date-time format for accurate result ( only available if either &#x27;status&#x27; not passed and if passed is set to &#x27;sent&#x27; ) | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 500]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetEmailCampaigns**](getEmailCampaigns.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, applications/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSharedTemplateUrl**
> GetSharedTemplateUrl GetSharedTemplateUrl(ctx, campaignId)
Get a shared template url

Get a unique URL to share & import an email template from one Sendinblue account to another.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign or template | 

### Return type

[**GetSharedTemplateUrl**](getSharedTemplateUrl.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendEmailCampaignNow**
> SendEmailCampaignNow(ctx, campaignId)
Send an email campaign immediately, based on campaignId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendReport**
> SendReport(ctx, body, campaignId)
Send the report of a campaign

A PDF will be sent to the specified email addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendReport**](SendReport.md)| Values for send a report | 
  **campaignId** | **int64**| Id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestEmail**
> SendTestEmail(ctx, body, campaignId)
Send an email campaign to your test list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendTestEmail**](SendTestEmail.md)|  | 
  **campaignId** | **int64**| Id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaignStatus**
> UpdateCampaignStatus(ctx, body, campaignId)
Update an email campaign status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCampaignStatus**](UpdateCampaignStatus.md)| Status of the campaign | 
  **campaignId** | **int64**| Id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmailCampaign**
> UpdateEmailCampaign(ctx, body, campaignId)
Update an email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEmailCampaign**](UpdateEmailCampaign.md)| Values to update a campaign | 
  **campaignId** | **int64**| Id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadImageToGallery**
> UploadImageToGallery(ctx, body)
Upload an image to your account's image gallery

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UploadImageToGallery**](UploadImageToGallery.md)| Parameters to upload an image | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

