# GetEmailEventReportEvents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address which generates the event | [default to null]
**Date** | **string** | Date on which the event has been generated | [default to null]
**Subject** | **string** | Subject of the event | [optional] [default to null]
**MessageId** | **string** | Message ID which generated the event | [default to null]
**Event** | **string** | Event which occured | [default to null]
**Reason** | **string** | Reason of bounce (only availble if the event is hardbounce or softbounce) | [default to null]
**Tag** | **string** | Tag of the email which generated the event | [default to null]
**Ip** | **string** | IP from which the user has opened the email or clicked on the link (only availble if the event is opened or clicks) | [optional] [default to null]
**Link** | **string** | The link which is sent to the user (only availble if the event is requests or opened or clicks) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


