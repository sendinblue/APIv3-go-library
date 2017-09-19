# SendTransacSms

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | Name of the sender. Only alphanumeric characters. No more than 11 characters | [default to null]
**Recipient** | **string** | Mobile number to send SMS with the country code | [default to null]
**Content** | **string** | Content of the message. If more than 160 characters long, multiple text messages will be sent | [default to null]
**Tag** | **string** | Tag of the message | [optional] [default to null]
**WebUrl** | **string** | Webhook to call for each event triggered by the message (delivered etc.) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


