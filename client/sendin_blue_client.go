// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/client/account"
	"github.com/sendinblue/APIv3-go-library/client/contacts"
	"github.com/sendinblue/APIv3-go-library/client/email_campaigns"
	"github.com/sendinblue/APIv3-go-library/client/operations"
	"github.com/sendinblue/APIv3-go-library/client/process"
	"github.com/sendinblue/APIv3-go-library/client/reseller"
	"github.com/sendinblue/APIv3-go-library/client/senders"
	"github.com/sendinblue/APIv3-go-library/client/sms_campaigns"
	"github.com/sendinblue/APIv3-go-library/client/smtp"
	"github.com/sendinblue/APIv3-go-library/client/transactional_sms"
	"github.com/sendinblue/APIv3-go-library/client/webhooks"
)

// Default sendin blue HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.sendinblue.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v3"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new sendin blue HTTP client.
func NewHTTPClient(formats strfmt.Registry) *SendinBlue {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new sendin blue HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *SendinBlue {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new sendin blue client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *SendinBlue {
	cli := new(SendinBlue)
	cli.Transport = transport

	cli.Account = account.New(transport, formats)

	cli.Contacts = contacts.New(transport, formats)

	cli.EmailCampaigns = email_campaigns.New(transport, formats)

	cli.Operations = operations.New(transport, formats)

	cli.Process = process.New(transport, formats)

	cli.Reseller = reseller.New(transport, formats)

	cli.Senders = senders.New(transport, formats)

	cli.SmsCampaigns = sms_campaigns.New(transport, formats)

	cli.SMTP = smtp.New(transport, formats)

	cli.TransactionalSms = transactional_sms.New(transport, formats)

	cli.Webhooks = webhooks.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// SendinBlue is a client for sendin blue
type SendinBlue struct {
	Account *account.Client

	Contacts *contacts.Client

	EmailCampaigns *email_campaigns.Client

	Operations *operations.Client

	Process *process.Client

	Reseller *reseller.Client

	Senders *senders.Client

	SmsCampaigns *sms_campaigns.Client

	SMTP *smtp.Client

	TransactionalSms *transactional_sms.Client

	Webhooks *webhooks.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *SendinBlue) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Account.SetTransport(transport)

	c.Contacts.SetTransport(transport)

	c.EmailCampaigns.SetTransport(transport)

	c.Operations.SetTransport(transport)

	c.Process.SetTransport(transport)

	c.Reseller.SetTransport(transport)

	c.Senders.SetTransport(transport)

	c.SmsCampaigns.SetTransport(transport)

	c.SMTP.SetTransport(transport)

	c.TransactionalSms.SetTransport(transport)

	c.Webhooks.SetTransport(transport)

}
