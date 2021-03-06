// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// DeregisterHostReader is a Reader for the DeregisterHost structure.
type DeregisterHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeregisterHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeregisterHostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeregisterHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeregisterHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeregisterHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeregisterHostNoContent creates a DeregisterHostNoContent with default headers values
func NewDeregisterHostNoContent() *DeregisterHostNoContent {
	return &DeregisterHostNoContent{}
}

/*DeregisterHostNoContent handles this case with default header values.

Success.
*/
type DeregisterHostNoContent struct {
}

func (o *DeregisterHostNoContent) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{cluster_id}/hosts/{host_id}][%d] deregisterHostNoContent ", 204)
}

func (o *DeregisterHostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeregisterHostBadRequest creates a DeregisterHostBadRequest with default headers values
func NewDeregisterHostBadRequest() *DeregisterHostBadRequest {
	return &DeregisterHostBadRequest{}
}

/*DeregisterHostBadRequest handles this case with default header values.

Error.
*/
type DeregisterHostBadRequest struct {
	Payload *models.Error
}

func (o *DeregisterHostBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{cluster_id}/hosts/{host_id}][%d] deregisterHostBadRequest  %+v", 400, o.Payload)
}

func (o *DeregisterHostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeregisterHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeregisterHostNotFound creates a DeregisterHostNotFound with default headers values
func NewDeregisterHostNotFound() *DeregisterHostNotFound {
	return &DeregisterHostNotFound{}
}

/*DeregisterHostNotFound handles this case with default header values.

Error.
*/
type DeregisterHostNotFound struct {
	Payload *models.Error
}

func (o *DeregisterHostNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{cluster_id}/hosts/{host_id}][%d] deregisterHostNotFound  %+v", 404, o.Payload)
}

func (o *DeregisterHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeregisterHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeregisterHostInternalServerError creates a DeregisterHostInternalServerError with default headers values
func NewDeregisterHostInternalServerError() *DeregisterHostInternalServerError {
	return &DeregisterHostInternalServerError{}
}

/*DeregisterHostInternalServerError handles this case with default header values.

Error.
*/
type DeregisterHostInternalServerError struct {
	Payload *models.Error
}

func (o *DeregisterHostInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{cluster_id}/hosts/{host_id}][%d] deregisterHostInternalServerError  %+v", 500, o.Payload)
}

func (o *DeregisterHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeregisterHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
