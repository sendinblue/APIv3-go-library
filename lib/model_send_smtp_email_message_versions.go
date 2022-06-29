/*
 * SendinBlue API
 *
 * SendinBlue provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/sendinblue  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  |
 *
 * API version: 3.0.0
 * Contact: contact@sendinblue.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package lib

type SendSmtpEmailMessageVersions struct {
	// List of email addresses and names (_optional_) of the recipients. For example, [{\"name\":\"Jimmy\", \"email\":\"jimmy98@example.com\"}, {\"name\":\"Joe\", \"email\":\"joe@example.com\"}]
	To []SendSmtpEmailTo1 `json:"to"`
	// Pass the set of attributes to customize the template. For example, {\"FNAME\":\"Joe\", \"LNAME\":\"Doe\"}. It's considered only if template is in New Template Language format.
	Params map[string]interface{} `json:"params,omitempty"`
	// List of email addresses and names (optional) of the recipients in bcc
	Bcc []SendSmtpEmailBcc `json:"bcc,omitempty"`
	// List of email addresses and names (optional) of the recipients in cc
	Cc      []SendSmtpEmailCc      `json:"cc,omitempty"`
	ReplyTo *SendSmtpEmailReplyTo1 `json:"replyTo,omitempty"`
	// Custom subject specific to message version
	Subject string `json:"subject,omitempty"`
}
