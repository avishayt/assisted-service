// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewInstallHostsParams creates a new InstallHostsParams object
// with the default values initialized.
func NewInstallHostsParams() *InstallHostsParams {
	var ()
	return &InstallHostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstallHostsParamsWithTimeout creates a new InstallHostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstallHostsParamsWithTimeout(timeout time.Duration) *InstallHostsParams {
	var ()
	return &InstallHostsParams{

		timeout: timeout,
	}
}

// NewInstallHostsParamsWithContext creates a new InstallHostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstallHostsParamsWithContext(ctx context.Context) *InstallHostsParams {
	var ()
	return &InstallHostsParams{

		Context: ctx,
	}
}

// NewInstallHostsParamsWithHTTPClient creates a new InstallHostsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstallHostsParamsWithHTTPClient(client *http.Client) *InstallHostsParams {
	var ()
	return &InstallHostsParams{
		HTTPClient: client,
	}
}

/*InstallHostsParams contains all the parameters to send to the API endpoint
for the install hosts operation typically these are written to a http.Request
*/
type InstallHostsParams struct {

	/*ClusterID
	  The existing cluster whose hosts should be added.

	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the install hosts params
func (o *InstallHostsParams) WithTimeout(timeout time.Duration) *InstallHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install hosts params
func (o *InstallHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install hosts params
func (o *InstallHostsParams) WithContext(ctx context.Context) *InstallHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install hosts params
func (o *InstallHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install hosts params
func (o *InstallHostsParams) WithHTTPClient(client *http.Client) *InstallHostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install hosts params
func (o *InstallHostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the install hosts params
func (o *InstallHostsParams) WithClusterID(clusterID strfmt.UUID) *InstallHostsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the install hosts params
func (o *InstallHostsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *InstallHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
