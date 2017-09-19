# GetExtendedContactDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the contact for which you requested the details | [default to null]
**Id** | **int32** | ID of the contact for which you requested the details | [default to null]
**EmailBlacklisted** | **bool** | Blacklist status for email campaigns (true&#x3D;blacklisted, false&#x3D;not blacklisted) | [default to null]
**SmsBlacklisted** | **bool** | Blacklist status for SMS campaigns (true&#x3D;blacklisted, false&#x3D;not blacklisted) | [default to null]
**ModifiedAt** | **string** | Last modification date of the contact (YYYY-MM-DD HH:mm:ss) | [default to null]
**ListIds** | **[]int32** |  | [default to null]
**ListUnsubscribed** | **[]int32** |  | [optional] [default to null]
**Attributes** | **map[string]string** |  | [default to null]
**Statistics** | [**GetExtendedContactDetailsStatistics**](GetExtendedContactDetailsStatistics.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


