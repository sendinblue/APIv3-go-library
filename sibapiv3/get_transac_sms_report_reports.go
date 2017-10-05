/*
 * SendinBlue API
 *
 * SendinBlue provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/sendinblue  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |
 *
 * OpenAPI spec version: 3.0.0
 * Contact: contact@sendinblue.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package sibapiv3

import (
	"time"
)

type GetTransacSmsReportReports struct {

	// Date for which statistics are retrieved
	Date time.Time `json:"date"`

	// Tag specified in request
	Tag string `json:"tag"`

	// Number of requests for the date
	Requests int64 `json:"requests"`

	// Number of delivered SMS for the date
	Delivered int64 `json:"delivered"`

	// Number of hardbounces for the date
	HardBounces int64 `json:"hardBounces"`

	// Number of softbounces for the date
	SoftBounces int64 `json:"softBounces"`

	// Number of blocked contact for the date
	Blocked int64 `json:"blocked"`

	// Number of unsubscription for the date
	Unsubscribed int64 `json:"unsubscribed"`

	// Number of answered SMS for the date
	Replied int64 `json:"replied"`

	// Number of accepted for the date
	Accepted int64 `json:"accepted"`

	// Number of rejected for the date
	Rejected int64 `json:"rejected"`
}
