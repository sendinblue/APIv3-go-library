# UpdateSmtpTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | Tag of the template | [optional] [default to null]
**Sender** | [**UpdateSmtpTemplateSender**](UpdateSmtpTemplateSender.md) |  | [optional] [default to null]
**TemplateName** | **string** | Name of the template | [optional] [default to null]
**HtmlContent** | **string** | Required if htmlUrl is empty. Body of the message (HTML must have more than 10 characters) | [optional] [default to null]
**HtmlUrl** | **string** | Required if htmlContent is empty. URL to the body of the email (HTML) | [optional] [default to null]
**Subject** | **string** | Subject of the email | [optional] [default to null]
**ReplyTo** | **string** | Email on which campaign recipients will be able to reply to | [optional] [default to null]
**ToField** | **string** | To personalize the «To» Field, e.g. if you want to include the first name and last name of your recipient, add [FNAME] [LNAME]. These attributes must already exist in contacts database | [optional] [default to null]
**AttachmentUrl** | **string** | Absolute url of the attachment (no local file). Extension allowed: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff and rtf | [optional] [default to null]
**IsActive** | **bool** | Status of the template. isActive &#x3D; false means template is inactive, isActive &#x3D; true means template is active | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


