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
	"encoding/json"
	"net/url"
	"strings"
	"time"
)

type TransactionalSMSApi struct {
	Configuration *Configuration
}

func NewTransactionalSMSApi() *TransactionalSMSApi {
	configuration := NewConfiguration()
	return &TransactionalSMSApi{
		Configuration: configuration,
	}
}

func NewTransactionalSMSApiWithBasePath(basePath string) *TransactionalSMSApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &TransactionalSMSApi{
		Configuration: configuration,
	}
}

/**
 * Get all the SMS activity (unaggregated events)
 *
 * @param limit Number of documents per page
 * @param startDate Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report
 * @param endDate Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report
 * @param offset Index of the first document of the page
 * @param days Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39;
 * @param phoneNumber Filter the report for a specific phone number
 * @param event Filter the report for specific events
 * @param tags Filter the report for specific tags passed as a serialized urlencoded array
 * @return *GetSmsEventReport
 */
func (a TransactionalSMSApi) GetSmsEvents(limit int64, startDate time.Time, endDate time.Time, offset int64, days int32, phoneNumber string, event string, tags string) (*GetSmsEventReport, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/transactionalSMS/statistics/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(api-key)' required
	// set key with prefix in header
	localVarHeaderParams["api-key"] = a.Configuration.GetAPIKeyWithPrefix("api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("startDate", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("endDate", a.Configuration.APIClient.ParameterToString(endDate, ""))
	localVarQueryParams.Add("offset", a.Configuration.APIClient.ParameterToString(offset, ""))
	localVarQueryParams.Add("days", a.Configuration.APIClient.ParameterToString(days, ""))
	localVarQueryParams.Add("phoneNumber", a.Configuration.APIClient.ParameterToString(phoneNumber, ""))
	localVarQueryParams.Add("event", a.Configuration.APIClient.ParameterToString(event, ""))
	localVarQueryParams.Add("tags", a.Configuration.APIClient.ParameterToString(tags, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetSmsEventReport)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetSmsEvents", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Get your SMS activity aggregated over a period of time
 *
 * @param startDate Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report
 * @param endDate Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report
 * @param days Number of days in the past including today (positive integer). Not compatible with startDate and endDate
 * @param tag Filter on a tag
 * @return *GetTransacAggregatedSmsReport
 */
func (a TransactionalSMSApi) GetTransacAggregatedSmsReport(startDate time.Time, endDate time.Time, days int32, tag string) (*GetTransacAggregatedSmsReport, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/transactionalSMS/statistics/aggregatedReport"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(api-key)' required
	// set key with prefix in header
	localVarHeaderParams["api-key"] = a.Configuration.GetAPIKeyWithPrefix("api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("startDate", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("endDate", a.Configuration.APIClient.ParameterToString(endDate, ""))
	localVarQueryParams.Add("days", a.Configuration.APIClient.ParameterToString(days, ""))
	localVarQueryParams.Add("tag", a.Configuration.APIClient.ParameterToString(tag, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetTransacAggregatedSmsReport)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetTransacAggregatedSmsReport", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Get your SMS activity aggregated per day
 *
 * @param startDate Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report
 * @param endDate Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report
 * @param days Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39;
 * @param tag Filter on a tag
 * @return *GetTransacSmsReport
 */
func (a TransactionalSMSApi) GetTransacSmsReport(startDate time.Time, endDate time.Time, days int32, tag string) (*GetTransacSmsReport, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/transactionalSMS/statistics/reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(api-key)' required
	// set key with prefix in header
	localVarHeaderParams["api-key"] = a.Configuration.GetAPIKeyWithPrefix("api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("startDate", a.Configuration.APIClient.ParameterToString(startDate, ""))
	localVarQueryParams.Add("endDate", a.Configuration.APIClient.ParameterToString(endDate, ""))
	localVarQueryParams.Add("days", a.Configuration.APIClient.ParameterToString(days, ""))
	localVarQueryParams.Add("tag", a.Configuration.APIClient.ParameterToString(tag, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetTransacSmsReport)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetTransacSmsReport", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Send the SMS campaign to the specified mobile number
 *
 * @param sendTransacSms Values to send a transactional SMS
 * @return *SendSms
 */
func (a TransactionalSMSApi) SendTransacSms(sendTransacSms SendTransacSms) (*SendSms, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/transactionalSMS/sms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(api-key)' required
	// set key with prefix in header
	localVarHeaderParams["api-key"] = a.Configuration.GetAPIKeyWithPrefix("api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &sendTransacSms
	var successPayload = new(SendSms)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "SendTransacSms", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}
