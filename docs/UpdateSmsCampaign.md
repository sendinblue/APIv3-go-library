# UpdateSmsCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the campaign | [optional] [default to null]
**Sender** | **string** | Name of the sender. The number of characters is limited to 11 | [optional] [default to null]
**Content** | **string** | Content of the message. The maximum characters used per SMS is 160, if used more than that, it will be counted as more than one SMS | [optional] [default to null]
**Recipients** | [**CreateSmsCampaignRecipients**](CreateSmsCampaignRecipients.md) |  | [optional] [default to null]
**ScheduledAt** | **string** | Date and time on which the campaign has to run (YYYY-MM-DD HH:mm:ss) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


