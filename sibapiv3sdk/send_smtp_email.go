/* 
 * SendinBlue API
 *
 * SendinBlue provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/sendinblue  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  | 
 *
 * OpenAPI spec version: 3.0.0
 * Contact: contact@sendinblue.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package sibapiv3sdk

type SendSmtpEmail struct {

	Sender SendSmtpEmailSender `json:"sender,omitempty"`

	// Email addresses and names of the recipients
	To []SendSmtpEmailTo `json:"to"`

	// Email addresses and names of the recipients in bcc
	Bcc []SendSmtpEmailBcc `json:"bcc,omitempty"`

	// Email addresses and names of the recipients in cc
	Cc []SendSmtpEmailCc `json:"cc,omitempty"`

	// HTML body of the message
	HtmlContent string `json:"htmlContent"`

	// Plain Text body of the message
	TextContent string `json:"textContent,omitempty"`

	// Subject of the message
	Subject string `json:"subject"`

	ReplyTo SendSmtpEmailReplyTo `json:"replyTo,omitempty"`

	// Pass the absolute URL (no local file) or the base64 content of the attachment. Name can be used in both cases to define the attachment name. It is mandatory in case of content. Extension allowed: gif, png, bmp, cgm, jpg, jpeg, tif, tiff, rtf, txt, css, shtml, html, htm, csv, zip, pdf, xml, ods, doc, docx, docm, ics, xls, xlsx, ppt, tar, and ez
	Attachment []SendSmtpEmailAttachment `json:"attachment,omitempty"`

	Headers map[string]string `json:"headers,omitempty"`
}
