# CreateEmailCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | Tag of the campaign | [optional] [default to null]
**Sender** | [**CreateEmailCampaignSender**](CreateEmailCampaignSender.md) |  | [optional] [default to null]
**Name** | **string** | Name of the campaign | [default to null]
**HtmlContent** | **string** | Mandatory if htmlUrl is empty. Body of the message (HTML) | [optional] [default to null]
**HtmlUrl** | **string** | Mandatory if htmlContent is empty. Url to the message (HTML) | [optional] [default to null]
**ScheduledAt** | **string** | Sending date and time (YYYY-MM-DD HH:mm:ss) | [optional] [default to null]
**Subject** | **string** | Subject of the campaign | [default to null]
**ReplyTo** | **string** | Email on which the campaign recipients will be able to reply to | [optional] [default to null]
**ToField** | **string** | To personalize the «To» Field, e.g. if you want to include the first name and last name of your recipient, use [FNAME] [LNAME]. These attributes must already exist in your contact database | [optional] [default to null]
**Recipients** | [**CreateEmailCampaignRecipients**](CreateEmailCampaignRecipients.md) |  | [optional] [default to null]
**AttachmentUrl** | **string** | Absolute url of the attachment (no local file). Extensions allowed xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff and rtf | [optional] [default to null]
**InlineImageActivation** | **bool** | Use true to embedded the images in your email. Final size of the email should be less than 4MB. Campaigns with embedded images can not be sent to more than 5000 contacts | [optional] [default to null]
**MirrorActive** | **bool** | Use true to enable the mirror link | [optional] [default to null]
**Recurring** | **bool** | For trigger campagins use false to make sure a contact receives the same campaign only once | [optional] [default to null]
**Type_** | **string** | Type of the campaign | [default to null]
**Footer** | **string** | Footer of the email campaign | [optional] [default to null]
**Header** | **string** | Header of the email campaign | [optional] [default to null]
**UtmCampaign** | **string** | Customize the utm_campaign value. If this field is empty, the campaign name will be used. Only alphanumeric characters and spaces are allowed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


