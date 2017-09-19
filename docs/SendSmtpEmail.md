# SendSmtpEmail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | [**SendSmtpEmailSender**](SendSmtpEmailSender.md) |  | [optional] [default to null]
**To** | [**[]SendSmtpEmailTo**](SendSmtpEmailTo.md) | Email addresses and names of the recipients | [default to null]
**Bcc** | [**[]SendSmtpEmailBcc**](SendSmtpEmailBcc.md) | Email addresses and names of the recipients in bcc | [optional] [default to null]
**Cc** | [**[]SendSmtpEmailCc**](SendSmtpEmailCc.md) | Email addresses and names of the recipients in cc | [optional] [default to null]
**HtmlContent** | **string** | HTML body of the message | [default to null]
**TextContent** | **string** | Plain Text body of the message | [optional] [default to null]
**Subject** | **string** | Subject of the message | [default to null]
**ReplyTo** | [**SendSmtpEmailReplyTo**](SendSmtpEmailReplyTo.md) |  | [optional] [default to null]
**Attachment** | [**[]SendSmtpEmailAttachment**](SendSmtpEmailAttachment.md) | Pass the absolute URL (no local file) or the base64 content of the attachment. Name can be used in both cases to define the attachment name. It is mandatory in case of content. Extension allowed: gif, png, bmp, cgm, jpg, jpeg, tif, tiff, rtf, txt, css, shtml, html, htm, csv, zip, pdf, xml, ods, doc, docx, docm, ics, xls, xlsx, ppt, tar, and ez | [optional] [default to null]
**Headers** | **map[string]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


