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

// Billing details of an order.
type OrderBilling struct {
	// Full billing address.
	Address string `json:"address,omitempty"`
	// Exact city of the address.
	City string `json:"city,omitempty"`
	// Billing country 2-letter ISO code.
	CountryCode string `json:"countryCode,omitempty"`
	// Phone number to contact for further details about the order, Mandatory if \"email\" field is not passed.
	Phone string `json:"phone,omitempty"`
	// Postcode for delivery and billing.
	PostCode string `json:"postCode,omitempty"`
	// How the visitor will pay for the item(s), e.g. paypal, check, etc.
	PaymentMethod string `json:"paymentMethod,omitempty"`
	// Exact region (state/province) for delivery and billing.
	Region string `json:"region,omitempty"`
}
