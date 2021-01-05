// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewListEventsParams creates a new ListEventsParams object
// with the default values initialized.
func NewListEventsParams() *ListEventsParams {
	var ()
	return &ListEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListEventsParamsWithTimeout creates a new ListEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEventsParamsWithTimeout(timeout time.Duration) *ListEventsParams {
	var ()
	return &ListEventsParams{

		timeout: timeout,
	}
}

// NewListEventsParamsWithContext creates a new ListEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListEventsParamsWithContext(ctx context.Context) *ListEventsParams {
	var ()
	return &ListEventsParams{

		Context: ctx,
	}
}

// NewListEventsParamsWithHTTPClient creates a new ListEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEventsParamsWithHTTPClient(client *http.Client) *ListEventsParams {
	var ()
	return &ListEventsParams{
		HTTPClient: client,
	}
}

/*ListEventsParams contains all the parameters to send to the API endpoint
for the list events operation typically these are written to a http.Request
*/
type ListEventsParams struct {

	/*ClusterID
	  The cluster to return events for.

	*/
	ClusterID strfmt.UUID
	/*HostID
	  A host in the specified cluster to return events for.

	*/
	HostID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list events params
func (o *ListEventsParams) WithTimeout(timeout time.Duration) *ListEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list events params
func (o *ListEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list events params
func (o *ListEventsParams) WithContext(ctx context.Context) *ListEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list events params
func (o *ListEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list events params
func (o *ListEventsParams) WithHTTPClient(client *http.Client) *ListEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list events params
func (o *ListEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list events params
func (o *ListEventsParams) WithClusterID(clusterID strfmt.UUID) *ListEventsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list events params
func (o *ListEventsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostID adds the hostID to the list events params
func (o *ListEventsParams) WithHostID(hostID *strfmt.UUID) *ListEventsParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the list events params
func (o *ListEventsParams) SetHostID(hostID *strfmt.UUID) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *ListEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.HostID != nil {

		// query param host_id
		var qrHostID strfmt.UUID
		if o.HostID != nil {
			qrHostID = *o.HostID
		}
		qHostID := qrHostID.String()
		if qHostID != "" {
			if err := r.SetQueryParam("host_id", qHostID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
