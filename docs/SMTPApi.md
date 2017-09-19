# \SMTPApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmtpTemplate**](SMTPApi.md#CreateSmtpTemplate) | **Post** /smtp/templates | Create an smtp template
[**DeleteHardbounces**](SMTPApi.md#DeleteHardbounces) | **Post** /smtp/deleteHardbounces | Delete hardbounces
[**GetAggregatedSmtpReport**](SMTPApi.md#GetAggregatedSmtpReport) | **Get** /smtp/statistics/aggregatedReport | Get your SMTP activity aggregated over a period of time
[**GetEmailEventReport**](SMTPApi.md#GetEmailEventReport) | **Get** /smtp/statistics/events | Get all your SMTP activity (unaggregated events)
[**GetSmtpReport**](SMTPApi.md#GetSmtpReport) | **Get** /smtp/statistics/reports | Get your SMTP activity aggregated per day
[**GetSmtpTemplate**](SMTPApi.md#GetSmtpTemplate) | **Get** /smtp/templates/{templateId} | Returns the template informations
[**GetSmtpTemplates**](SMTPApi.md#GetSmtpTemplates) | **Get** /smtp/templates | Get the list of SMTP templates
[**SendTemplate**](SMTPApi.md#SendTemplate) | **Post** /smtp/templates/{templateId}/send | Send a template
[**SendTestTemplate**](SMTPApi.md#SendTestTemplate) | **Post** /smtp/templates/{templateId}/sendTest | Send a template to your test list
[**SendTransacEmail**](SMTPApi.md#SendTransacEmail) | **Post** /smtp/email | Send a transactional email
[**UpdateSmtpTemplate**](SMTPApi.md#UpdateSmtpTemplate) | **Put** /smtp/templates/{templateId} | Updates an smtp templates


# **CreateSmtpTemplate**
> CreateModel CreateSmtpTemplate($smtpTemplate)

Create an smtp template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpTemplate** | [**CreateSmtpTemplate**](CreateSmtpTemplate.md)| values to update in smtp template | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHardbounces**
> DeleteHardbounces($deleteHardbounces)

Delete hardbounces

Delete hardbounces. To use carefully (e.g. in case of temporary ISP failures)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteHardbounces** | [**DeleteHardbounces**](DeleteHardbounces.md)| values to delete hardbounces | [optional] 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedSmtpReport**
> GetAggregatedReport GetAggregatedSmtpReport($startDate, $endDate, $days, $tag)

Get your SMTP activity aggregated over a period of time


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD). Must be lower than equal to endDate | [optional] 
 **endDate** | **string**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate | [optional] 
 **days** | **int32**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | [optional] 
 **tag** | **string**| Tag of the emails | [optional] 

### Return type

[**GetAggregatedReport**](GetAggregatedReport.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailEventReport**
> GetEmailEventReport GetEmailEventReport($limit, $offset, $startDate, $endDate, $days, $email, $event, $tags, $messageId, $templateId)

Get all your SMTP activity (unaggregated events)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Number limitation for the result returned | [optional] [default to 50]
 **offset** | **int32**| Beginning point in the list to retrieve from. | [optional] [default to 0]
 **startDate** | **string**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD). Must be lower than equal to endDate | [optional] 
 **endDate** | **string**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate | [optional] 
 **days** | **int32**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | [optional] 
 **email** | **string**| Filter the report for a specific email addresses | [optional] 
 **event** | **string**| Filter the report for a specific event type | [optional] 
 **tags** | **string**| Filter the report for tags (serialized and urlencoded array) | [optional] 
 **messageId** | **string**| Filter on a specific message id | [optional] 
 **templateId** | **string**| Filter on a specific template id | [optional] 

### Return type

[**GetEmailEventReport**](GetEmailEventReport.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpReport**
> GetReports GetSmtpReport($limit, $offset, $startDate, $endDate, $days, $tag)

Get your SMTP activity aggregated per day


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Number of documents returned per page | [optional] [default to 50]
 **offset** | **int32**| Index of the first document on the page | [optional] [default to 0]
 **startDate** | **string**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD) | [optional] 
 **endDate** | **string**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD) | [optional] 
 **days** | **int32**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | [optional] 
 **tag** | **string**| Tag of the emails | [optional] 

### Return type

[**GetReports**](GetReports.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpTemplate**
> GetSmtpTemplateOverview GetSmtpTemplate($templateId)

Returns the template informations


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string**| id of the template | 

### Return type

[**GetSmtpTemplateOverview**](GetSmtpTemplateOverview.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpTemplates**
> GetSmtpTemplates GetSmtpTemplates($templateStatus, $limit, $offset)

Get the list of SMTP templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateStatus** | **bool**| Filter on the status of the template. Active &#x3D; true, inactive &#x3D; false | [optional] 
 **limit** | **int32**| Number of documents returned per page | [optional] [default to 50]
 **offset** | **int32**| Index of the first document in the page | [optional] [default to 0]

### Return type

[**GetSmtpTemplates**](GetSmtpTemplates.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTemplate**
> SendTemplateEmail SendTemplate($templateId, $sendEmail)

Send a template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string**| Id of the template | 
 **sendEmail** | [**SendEmail**](SendEmail.md)|  | 

### Return type

[**SendTemplateEmail**](SendTemplateEmail.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestTemplate**
> SendTestTemplate($templateId, $sendTestEmail)

Send a template to your test list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string**| Id of the template | 
 **sendTestEmail** | [**SendTestEmail**](SendTestEmail.md)|  | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTransacEmail**
> CreateSmtpEmail SendTransacEmail($sendSmtpEmail)

Send a transactional email


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendSmtpEmail** | [**SendSmtpEmail**](SendSmtpEmail.md)| Values to send a transactional email | 

### Return type

[**CreateSmtpEmail**](CreateSmtpEmail.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmtpTemplate**
> UpdateSmtpTemplate($templateId, $smtpTemplate)

Updates an smtp templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string**| id of the template | 
 **smtpTemplate** | [**UpdateSmtpTemplate**](UpdateSmtpTemplate.md)| values to update in smtp template | 

### Return type

void (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

