# sib_api_v3_sdk.TransactionalEmailsApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockNewDomain**](TransactionalEmailsApi.md#BlockNewDomain) | **Post** /smtp/blockedDomains | Add a new domain to the list of blocked domains
[**CreateSmtpTemplate**](TransactionalEmailsApi.md#CreateSmtpTemplate) | **Post** /smtp/templates | Create an email template
[**DeleteBlockedDomain**](TransactionalEmailsApi.md#DeleteBlockedDomain) | **Delete** /smtp/blockedDomains/{domain} | Unblock an existing domain from the list of blocked domains
[**DeleteHardbounces**](TransactionalEmailsApi.md#DeleteHardbounces) | **Post** /smtp/deleteHardbounces | Delete hardbounces
[**DeleteSmtpTemplate**](TransactionalEmailsApi.md#DeleteSmtpTemplate) | **Delete** /smtp/templates/{templateId} | Delete an inactive email template
[**GetAggregatedSmtpReport**](TransactionalEmailsApi.md#GetAggregatedSmtpReport) | **Get** /smtp/statistics/aggregatedReport | Get your transactional email activity aggregated over a period of time
[**GetBlockedDomains**](TransactionalEmailsApi.md#GetBlockedDomains) | **Get** /smtp/blockedDomains | Get the list of blocked domains
[**GetEmailEventReport**](TransactionalEmailsApi.md#GetEmailEventReport) | **Get** /smtp/statistics/events | Get all your transactional email activity (unaggregated events)
[**GetSmtpReport**](TransactionalEmailsApi.md#GetSmtpReport) | **Get** /smtp/statistics/reports | Get your transactional email activity aggregated per day
[**GetSmtpTemplate**](TransactionalEmailsApi.md#GetSmtpTemplate) | **Get** /smtp/templates/{templateId} | Returns the template information
[**GetSmtpTemplates**](TransactionalEmailsApi.md#GetSmtpTemplates) | **Get** /smtp/templates | Get the list of email templates
[**GetTransacBlockedContacts**](TransactionalEmailsApi.md#GetTransacBlockedContacts) | **Get** /smtp/blockedContacts | Get the list of blocked or unsubscribed transactional contacts
[**GetTransacEmailContent**](TransactionalEmailsApi.md#GetTransacEmailContent) | **Get** /smtp/emails/{uuid} | Get the personalized content of a sent transactional email
[**GetTransacEmailsList**](TransactionalEmailsApi.md#GetTransacEmailsList) | **Get** /smtp/emails | Get the list of transactional emails on the basis of allowed filters
[**SendTemplate**](TransactionalEmailsApi.md#SendTemplate) | **Post** /smtp/templates/{templateId}/send | Send a template
[**SendTestTemplate**](TransactionalEmailsApi.md#SendTestTemplate) | **Post** /smtp/templates/{templateId}/sendTest | Send a template to your test list
[**SendTransacEmail**](TransactionalEmailsApi.md#SendTransacEmail) | **Post** /smtp/email | Send a transactional email
[**SmtpBlockedContactsEmailDelete**](TransactionalEmailsApi.md#SmtpBlockedContactsEmailDelete) | **Delete** /smtp/blockedContacts/{email} | Unblock or resubscribe a transactional contact
[**SmtpLogMessageIdDelete**](TransactionalEmailsApi.md#SmtpLogMessageIdDelete) | **Delete** /smtp/log/{messageId} | Delete an SMTP transactional log
[**UpdateSmtpTemplate**](TransactionalEmailsApi.md#UpdateSmtpTemplate) | **Put** /smtp/templates/{templateId} | Update an email template

# **BlockNewDomain**
> BlockNewDomain(ctx, body)
Add a new domain to the list of blocked domains

Blocks a new domain in order to avoid messages being sent to the same

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BlockDomain**](BlockDomain.md)|  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSmtpTemplate**
> CreateModel CreateSmtpTemplate(ctx, body)
Create an email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSmtpTemplate**](CreateSmtpTemplate.md)| values to update in transactional email template | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlockedDomain**
> DeleteBlockedDomain(ctx, domain)
Unblock an existing domain from the list of blocked domains

Unblocks an existing domain from the list of blocked domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| The name of the domain to be deleted | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHardbounces**
> DeleteHardbounces(ctx, optional)
Delete hardbounces

Delete hardbounces. To use carefully (e.g. in case of temporary ISP failures)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionalEmailsApiDeleteHardbouncesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionalEmailsApiDeleteHardbouncesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DeleteHardbounces**](DeleteHardbounces.md)| values to delete hardbounces | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmtpTemplate**
> DeleteSmtpTemplate(ctx, templateId)
Delete an inactive email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**| id of the template | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedSmtpReport**
> GetAggregatedReport GetAggregatedSmtpReport(ctx, optional)
Get your transactional email activity aggregated over a period of time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionalEmailsApiGetAggregatedSmtpReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionalEmailsApiGetAggregatedSmtpReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD). Must be lower than equal to endDate | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with &#x27;startDate&#x27; and &#x27;endDate&#x27; | 
 **tag** | **optional.String**| Tag of the emails | 

### Return type

[**GetAggregatedReport**](GetAggregatedReport.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockedDomains**
> GetBlockedDomains GetBlockedDomains(ctx, )
Get the list of blocked domains

Get the list of blocked domains

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetBlockedDomains**](GetBlockedDomains.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailEventReport**
> GetEmailEventReport GetEmailEventReport(ctx, optional)
Get all your transactional email activity (unaggregated events)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionalEmailsApiGetEmailEventReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionalEmailsApiGetEmailEventReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number limitation for the result returned | [default to 50]
 **offset** | **optional.Int64**| Beginning point in the list to retrieve from. | [default to 0]
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD). Must be lower than equal to endDate | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with &#x27;startDate&#x27; and &#x27;endDate&#x27; | 
 **email** | **optional.String**| Filter the report for a specific email addresses | 
 **event** | **optional.String**| Filter the report for a specific event type | 
 **tags** | **optional.String**| Filter the report for tags (serialized and urlencoded array) | 
 **messageId** | **optional.String**| Filter on a specific message id | 
 **templateId** | **optional.Int64**| Filter on a specific template id | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetEmailEventReport**](GetEmailEventReport.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpReport**
> GetReports GetSmtpReport(ctx, optional)
Get your transactional email activity aggregated per day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionalEmailsApiGetSmtpReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionalEmailsApiGetSmtpReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 10]
 **offset** | **optional.Int64**| Index of the first document on the page | [default to 0]
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD) | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD) | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with &#x27;startDate&#x27; and &#x27;endDate&#x27; | 
 **tag** | **optional.String**| Tag of the emails | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetReports**](GetReports.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpTemplate**
> GetSmtpTemplateOverview GetSmtpTemplate(ctx, templateId)
Returns the template information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**| id of the template | 

### Return type

[**GetSmtpTemplateOverview**](GetSmtpTemplateOverview.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpTemplates**
> GetSmtpTemplates GetSmtpTemplates(ctx, optional)
Get the list of email templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionalEmailsApiGetSmtpTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionalEmailsApiGetSmtpTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateStatus** | **optional.Bool**| Filter on the status of the template. Active &#x3D; true, inactive &#x3D; false | 
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetSmtpTemplates**](GetSmtpTemplates.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacBlockedContacts**
> GetTransacBlockedContacts GetTransacBlockedContacts(ctx, optional)
Get the list of blocked or unsubscribed transactional contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionalEmailsApiGetTransacBlockedContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionalEmailsApiGetTransacBlockedContactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) from which you want to fetch the blocked or unsubscribed contacts | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) till which you want to fetch the blocked or unsubscribed contacts | 
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document on the page | [default to 0]
 **senders** | [**optional.Interface of []string**](string.md)| Comma separated list of emails of the senders from which contacts are blocked or unsubscribed | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetTransacBlockedContacts**](GetTransacBlockedContacts.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacEmailContent**
> GetTransacEmailContent GetTransacEmailContent(ctx, uuid)
Get the personalized content of a sent transactional email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Unique id of the transactional email that has been sent to a particular contact | 

### Return type

[**GetTransacEmailContent**](GetTransacEmailContent.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacEmailsList**
> GetTransacEmailsList GetTransacEmailsList(ctx, optional)
Get the list of transactional emails on the basis of allowed filters

This endpoint will show the list of emails for past 30 days by default. To retrieve emails before that time, please pass startDate and endDate in query filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionalEmailsApiGetTransacEmailsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionalEmailsApiGetTransacEmailsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| Mandatory if templateId and messageId are not passed in query filters. Email address to which transactional email has been sent. | 
 **templateId** | **optional.Int64**| Mandatory if email and messageId are not passed in query filters. Id of the template that was used to compose transactional email. | 
 **messageId** | **optional.String**| Mandatory if templateId and email are not passed in query filters. Message ID of the transactional email sent. | 
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) from which you want to fetch the list. Maximum time period that can be selected is one month. | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) till which you want to fetch the list. Maximum time period that can be selected is one month. | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetTransacEmailsList**](GetTransacEmailsList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTemplate**
> SendTemplateEmail SendTemplate(ctx, body, templateId)
Send a template

This endpoint is deprecated. Prefer v3/smtp/email instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendEmail**](SendEmail.md)|  | 
  **templateId** | **int64**| Id of the template | 

### Return type

[**SendTemplateEmail**](SendTemplateEmail.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestTemplate**
> SendTestTemplate(ctx, body, templateId)
Send a template to your test list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendTestEmail**](SendTestEmail.md)|  | 
  **templateId** | **int64**| Id of the template | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTransacEmail**
> CreateSmtpEmail SendTransacEmail(ctx, body)
Send a transactional email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendSmtpEmail**](SendSmtpEmail.md)| Values to send a transactional email | 

### Return type

[**CreateSmtpEmail**](CreateSmtpEmail.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmtpBlockedContactsEmailDelete**
> SmtpBlockedContactsEmailDelete(ctx, email)
Unblock or resubscribe a transactional contact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| contact email (urlencoded) to unblock. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmtpLogMessageIdDelete**
> SmtpLogMessageIdDelete(ctx, messageId)
Delete an SMTP transactional log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **messageId** | **string**| MessageId of the transactional log to delete | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmtpTemplate**
> UpdateSmtpTemplate(ctx, body, templateId)
Update an email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSmtpTemplate**](UpdateSmtpTemplate.md)| values to update in transactional email template | 
  **templateId** | **int64**| id of the template | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

