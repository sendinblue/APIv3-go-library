// Code generated by go-swagger; DO NOT EDIT.

package transactional_sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new transactional sms API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transactional sms API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSMSEvents gets all the SMS activity unaggregated events
*/
func (a *Client) GetSMSEvents(params *GetSMSEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSMSEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSMSEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSmsEvents",
		Method:             "GET",
		PathPattern:        "/transactionalSMS/statistics/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSMSEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSMSEventsOK), nil

}

/*
GetTransacAggregatedSMSReport gets your SMS activity aggregated over a period of time
*/
func (a *Client) GetTransacAggregatedSMSReport(params *GetTransacAggregatedSMSReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetTransacAggregatedSMSReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransacAggregatedSMSReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTransacAggregatedSmsReport",
		Method:             "GET",
		PathPattern:        "/transactionalSMS/statistics/aggregatedReport",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransacAggregatedSMSReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransacAggregatedSMSReportOK), nil

}

/*
GetTransacSMSReport gets your SMS activity aggregated per day
*/
func (a *Client) GetTransacSMSReport(params *GetTransacSMSReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetTransacSMSReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransacSMSReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTransacSmsReport",
		Method:             "GET",
		PathPattern:        "/transactionalSMS/statistics/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransacSMSReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransacSMSReportOK), nil

}

/*
SendTransacSMS sends the SMS campaign to the specified mobile number
*/
func (a *Client) SendTransacSMS(params *SendTransacSMSParams, authInfo runtime.ClientAuthInfoWriter) (*SendTransacSMSCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendTransacSMSParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendTransacSms",
		Method:             "POST",
		PathPattern:        "/transactionalSMS/sms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendTransacSMSReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendTransacSMSCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
