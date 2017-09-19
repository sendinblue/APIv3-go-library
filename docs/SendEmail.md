# SendEmail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTo** | **[]string** | Email addresses of the recipients | [default to null]
**EmailBcc** | **[]string** | Email addresses of the recipients in bcc | [optional] [default to null]
**EmailCc** | **[]string** | Email addresses of the recipients in cc | [optional] [default to null]
**ReplyTo** | **string** | Email on which campaign recipients will be able to reply to | [optional] [default to null]
**AttachmentUrl** | **string** | Absolute url of the attachment (no local file). Extension allowed: gif, png, bmp, cgm, jpg, jpeg, tif, tiff, rtf, txt, css, shtml, html, htm, csv, zip, pdf, xml, ods, doc, docx, docm, ics, xls, xlsx, ppt, tar, and ez | [optional] [default to null]
**Attachment** | [**[]SendEmailAttachment**](SendEmailAttachment.md) | Pass the base64 content of the attachment. Extension allowed: gif, png, bmp, cgm, jpg, jpeg, tif, tiff, rtf, txt, css, shtml, html, htm, csv, zip, pdf, xml, ods, doc, docx, docm, ics, xls, xlsx, ppt, tar, and ez | [optional] [default to null]
**Headers** | **map[string]string** |  | [optional] [default to null]
**Attributes** | **map[string]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


