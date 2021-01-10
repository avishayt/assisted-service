// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// ListImagesReader is a Reader for the ListImages structure.
type ListImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListImagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListImagesOK creates a ListImagesOK with default headers values
func NewListImagesOK() *ListImagesOK {
	return &ListImagesOK{}
}

/*ListImagesOK handles this case with default header values.

Success.
*/
type ListImagesOK struct {
	Payload models.ImageList
}

func (o *ListImagesOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images][%d] listImagesOK  %+v", 200, o.Payload)
}

func (o *ListImagesOK) GetPayload() models.ImageList {
	return o.Payload
}

func (o *ListImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesUnauthorized creates a ListImagesUnauthorized with default headers values
func NewListImagesUnauthorized() *ListImagesUnauthorized {
	return &ListImagesUnauthorized{}
}

/*ListImagesUnauthorized handles this case with default header values.

Unauthorized.
*/
type ListImagesUnauthorized struct {
	Payload *models.InfraError
}

func (o *ListImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images][%d] listImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListImagesUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *ListImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesForbidden creates a ListImagesForbidden with default headers values
func NewListImagesForbidden() *ListImagesForbidden {
	return &ListImagesForbidden{}
}

/*ListImagesForbidden handles this case with default header values.

Forbidden.
*/
type ListImagesForbidden struct {
	Payload *models.InfraError
}

func (o *ListImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images][%d] listImagesForbidden  %+v", 403, o.Payload)
}

func (o *ListImagesForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *ListImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesNotFound creates a ListImagesNotFound with default headers values
func NewListImagesNotFound() *ListImagesNotFound {
	return &ListImagesNotFound{}
}

/*ListImagesNotFound handles this case with default header values.

Error.
*/
type ListImagesNotFound struct {
	Payload *models.Error
}

func (o *ListImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images][%d] listImagesNotFound  %+v", 404, o.Payload)
}

func (o *ListImagesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesInternalServerError creates a ListImagesInternalServerError with default headers values
func NewListImagesInternalServerError() *ListImagesInternalServerError {
	return &ListImagesInternalServerError{}
}

/*ListImagesInternalServerError handles this case with default header values.

Error.
*/
type ListImagesInternalServerError struct {
	Payload *models.Error
}

func (o *ListImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images][%d] listImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListImagesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesServiceUnavailable creates a ListImagesServiceUnavailable with default headers values
func NewListImagesServiceUnavailable() *ListImagesServiceUnavailable {
	return &ListImagesServiceUnavailable{}
}

/*ListImagesServiceUnavailable handles this case with default header values.

Unavailable.
*/
type ListImagesServiceUnavailable struct {
	Payload *models.Error
}

func (o *ListImagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images][%d] listImagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListImagesServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListImagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
