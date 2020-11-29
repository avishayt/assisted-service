// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// DownloadPXEArtifactOKCode is the HTTP code returned for type DownloadPXEArtifactOK
const DownloadPXEArtifactOKCode int = 200

/*DownloadPXEArtifactOK Success.

swagger:response downloadPXEArtifactOK
*/
type DownloadPXEArtifactOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadPXEArtifactOK creates DownloadPXEArtifactOK with default headers values
func NewDownloadPXEArtifactOK() *DownloadPXEArtifactOK {

	return &DownloadPXEArtifactOK{}
}

// WithPayload adds the payload to the download p x e artifact o k response
func (o *DownloadPXEArtifactOK) WithPayload(payload io.ReadCloser) *DownloadPXEArtifactOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download p x e artifact o k response
func (o *DownloadPXEArtifactOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadPXEArtifactOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadPXEArtifactTemporaryRedirectCode is the HTTP code returned for type DownloadPXEArtifactTemporaryRedirect
const DownloadPXEArtifactTemporaryRedirectCode int = 307

/*DownloadPXEArtifactTemporaryRedirect Redirect.

swagger:response downloadPXEArtifactTemporaryRedirect
*/
type DownloadPXEArtifactTemporaryRedirect struct {
	/*

	 */
	Location string `json:"Location"`
}

// NewDownloadPXEArtifactTemporaryRedirect creates DownloadPXEArtifactTemporaryRedirect with default headers values
func NewDownloadPXEArtifactTemporaryRedirect() *DownloadPXEArtifactTemporaryRedirect {

	return &DownloadPXEArtifactTemporaryRedirect{}
}

// WithLocation adds the location to the download p x e artifact temporary redirect response
func (o *DownloadPXEArtifactTemporaryRedirect) WithLocation(location string) *DownloadPXEArtifactTemporaryRedirect {
	o.Location = location
	return o
}

// SetLocation sets the location to the download p x e artifact temporary redirect response
func (o *DownloadPXEArtifactTemporaryRedirect) SetLocation(location string) {
	o.Location = location
}

// WriteResponse to the client
func (o *DownloadPXEArtifactTemporaryRedirect) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(307)
}

// DownloadPXEArtifactUnauthorizedCode is the HTTP code returned for type DownloadPXEArtifactUnauthorized
const DownloadPXEArtifactUnauthorizedCode int = 401

/*DownloadPXEArtifactUnauthorized Unauthorized.

swagger:response downloadPXEArtifactUnauthorized
*/
type DownloadPXEArtifactUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadPXEArtifactUnauthorized creates DownloadPXEArtifactUnauthorized with default headers values
func NewDownloadPXEArtifactUnauthorized() *DownloadPXEArtifactUnauthorized {

	return &DownloadPXEArtifactUnauthorized{}
}

// WithPayload adds the payload to the download p x e artifact unauthorized response
func (o *DownloadPXEArtifactUnauthorized) WithPayload(payload *models.InfraError) *DownloadPXEArtifactUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download p x e artifact unauthorized response
func (o *DownloadPXEArtifactUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadPXEArtifactUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadPXEArtifactForbiddenCode is the HTTP code returned for type DownloadPXEArtifactForbidden
const DownloadPXEArtifactForbiddenCode int = 403

/*DownloadPXEArtifactForbidden Forbidden.

swagger:response downloadPXEArtifactForbidden
*/
type DownloadPXEArtifactForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadPXEArtifactForbidden creates DownloadPXEArtifactForbidden with default headers values
func NewDownloadPXEArtifactForbidden() *DownloadPXEArtifactForbidden {

	return &DownloadPXEArtifactForbidden{}
}

// WithPayload adds the payload to the download p x e artifact forbidden response
func (o *DownloadPXEArtifactForbidden) WithPayload(payload *models.InfraError) *DownloadPXEArtifactForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download p x e artifact forbidden response
func (o *DownloadPXEArtifactForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadPXEArtifactForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadPXEArtifactInternalServerErrorCode is the HTTP code returned for type DownloadPXEArtifactInternalServerError
const DownloadPXEArtifactInternalServerErrorCode int = 500

/*DownloadPXEArtifactInternalServerError Error.

swagger:response downloadPXEArtifactInternalServerError
*/
type DownloadPXEArtifactInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadPXEArtifactInternalServerError creates DownloadPXEArtifactInternalServerError with default headers values
func NewDownloadPXEArtifactInternalServerError() *DownloadPXEArtifactInternalServerError {

	return &DownloadPXEArtifactInternalServerError{}
}

// WithPayload adds the payload to the download p x e artifact internal server error response
func (o *DownloadPXEArtifactInternalServerError) WithPayload(payload *models.Error) *DownloadPXEArtifactInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download p x e artifact internal server error response
func (o *DownloadPXEArtifactInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadPXEArtifactInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
