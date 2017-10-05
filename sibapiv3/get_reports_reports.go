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

type GetReportsReports struct {

	// Date of the statistics
	Date time.Time `json:"date"`

	// Reminder of the specified tag (only available if a specific tag has been specified in the request)
	Tag string `json:"tag"`

	// Number of requests for the date
	Requests int64 `json:"requests"`

	// Number of delivered emails for the date
	Delivered int64 `json:"delivered"`

	// Number of hardbounces for the date
	HardBounces int64 `json:"hardBounces"`

	// Number of softbounces for the date
	SoftBounces int64 `json:"softBounces"`

	// Number of clicks for the date
	Clicks int64 `json:"clicks"`

	// Number of unique clicks for the date
	UniqueClicks int64 `json:"uniqueClicks"`

	// Number of openings for the date
	Opens int64 `json:"opens"`

	// Number of unique openings for the date
	UniqueOpens int64 `json:"uniqueOpens"`

	// Number of complaints (spam reports) for the date
	SpamReports int64 `json:"spamReports"`

	// Number of blocked emails for the date
	Blocked int64 `json:"blocked"`

	// Number of invalid emails for the date
	Invalid int64 `json:"invalid"`
}
