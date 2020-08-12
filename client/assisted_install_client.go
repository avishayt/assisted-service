// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/client/installer"
	"github.com/openshift/assisted-service/client/managed_domains"
	"github.com/openshift/assisted-service/client/versions"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.openshift.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/assisted-install/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	AuthInfo runtime.ClientAuthInfoWriter
}

// New creates a new assisted install HTTP client.
func New(c Config) *AssistedInstall {
	var (
		host     = DefaultHost
		basePath = DefaultBasePath
		schemes  = DefaultSchemes
	)

	if c.URL != nil {
		host = c.URL.Host
		basePath = c.URL.Path
		schemes = []string{c.URL.Scheme}
	}

	transport := rtclient.New(host, basePath, schemes)
	if c.Transport != nil {
		transport.Transport = c.Transport
	}

	cli := new(AssistedInstall)
	cli.Transport = transport
	cli.Installer = installer.New(transport, strfmt.Default, c.AuthInfo)
	cli.ManagedDomains = managed_domains.New(transport, strfmt.Default, c.AuthInfo)
	cli.Versions = versions.New(transport, strfmt.Default, c.AuthInfo)
	return cli
}

// AssistedInstall is a client for assisted install
type AssistedInstall struct {
	Installer      *installer.Client
	ManagedDomains *managed_domains.Client
	Versions       *versions.Client
	Transport      runtime.ClientTransport
}
