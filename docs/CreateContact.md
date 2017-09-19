# CreateContact

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the user. Mandatory if &#x60;attributes.sms&#x60; is not passed | [optional] [default to null]
**Attributes** | [**interface{}**](interface{}.md) | Values of the attributes to fill. The attributes must exist in you contact database | [optional] [default to null]
**EmailBlacklisted** | **bool** | Blacklist the contact for emails (emailBlacklisted &#x3D; true) | [optional] [default to null]
**SmsBlacklisted** | **bool** | Blacklist the contact for SMS (smsBlacklisted &#x3D; true) | [optional] [default to null]
**ListIds** | **[]int32** | Ids of the lists to add the contact to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


