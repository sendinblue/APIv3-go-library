# \TransactionalSMSApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmsEvents**](TransactionalSMSApi.md#GetSmsEvents) | **Get** /transactionalSMS/statistics/events | Get all the SMS activity (unaggregated events)
[**GetTransacAggregatedSmsReport**](TransactionalSMSApi.md#GetTransacAggregatedSmsReport) | **Get** /transactionalSMS/statistics/aggregatedReport | Get your SMS activity aggregated over a period of time
[**GetTransacSmsReport**](TransactionalSMSApi.md#GetTransacSmsReport) | **Get** /transactionalSMS/statistics/reports | Get your SMS activity aggregated per day
[**SendTransacSms**](TransactionalSMSApi.md#SendTransacSms) | **Post** /transactionalSMS/sms | Send the SMS campaign to the specified mobile number


# **GetSmsEvents**
> GetSmsEventReport GetSmsEvents($limit, $startDate, $endDate, $offset, $days, $phoneNumber, $event, $tags)

Get all the SMS activity (unaggregated events)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Number of documents per page | [optional] [default to 50]
 **startDate** | **string**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report | [optional] 
 **endDate** | **string**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report | [optional] 
 **offset** | **int32**| Index of the first document of the page | [optional] [default to 0]
 **days** | **int32**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | [optional] 
 **phoneNumber** | **string**| Filter the report for a specific phone number | [optional] 
 **event** | **string**| Filter the report for specific events | [optional] 
 **tags** | **string**| Filter the report for specific tags passed as a serialized urlencoded array | [optional] 

### Return type

[**GetSmsEventReport**](GetSmsEventReport.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacAggregatedSmsReport**
> GetTransacAggregatedSmsReport GetTransacAggregatedSmsReport($startDate, $endDate, $days, $tag)

Get your SMS activity aggregated over a period of time


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report | [optional] 
 **endDate** | **string**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report | [optional] 
 **days** | **int32**| Number of days in the past including today (positive integer). Not compatible with startDate and endDate | [optional] 
 **tag** | **string**| Filter on a tag | [optional] 

### Return type

[**GetTransacAggregatedSmsReport**](GetTransacAggregatedSmsReport.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacSmsReport**
> GetTransacSmsReport GetTransacSmsReport($startDate, $endDate, $days, $tag)

Get your SMS activity aggregated per day


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report | [optional] 
 **endDate** | **string**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report | [optional] 
 **days** | **int32**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | [optional] 
 **tag** | **string**| Filter on a tag | [optional] 

### Return type

[**GetTransacSmsReport**](GetTransacSmsReport.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTransacSms**
> SendSms SendTransacSms($sendTransacSms)

Send the SMS campaign to the specified mobile number


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendTransacSms** | [**SendTransacSms**](SendTransacSms.md)| Values to send a transactional SMS | 

### Return type

[**SendSms**](SendSms.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

