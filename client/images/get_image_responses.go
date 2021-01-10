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

// GetImageReader is a Reader for the GetImage structure.
type GetImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetImageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetImageOK creates a GetImageOK with default headers values
func NewGetImageOK() *GetImageOK {
	return &GetImageOK{}
}

/*GetImageOK handles this case with default header values.

Success.
*/
type GetImageOK struct {
	Payload *models.Image
}

func (o *GetImageOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images/{image_id}][%d] getImageOK  %+v", 200, o.Payload)
}

func (o *GetImageOK) GetPayload() *models.Image {
	return o.Payload
}

func (o *GetImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageUnauthorized creates a GetImageUnauthorized with default headers values
func NewGetImageUnauthorized() *GetImageUnauthorized {
	return &GetImageUnauthorized{}
}

/*GetImageUnauthorized handles this case with default header values.

Unauthorized.
*/
type GetImageUnauthorized struct {
	Payload *models.InfraError
}

func (o *GetImageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images/{image_id}][%d] getImageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetImageUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageForbidden creates a GetImageForbidden with default headers values
func NewGetImageForbidden() *GetImageForbidden {
	return &GetImageForbidden{}
}

/*GetImageForbidden handles this case with default header values.

Forbidden.
*/
type GetImageForbidden struct {
	Payload *models.InfraError
}

func (o *GetImageForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images/{image_id}][%d] getImageForbidden  %+v", 403, o.Payload)
}

func (o *GetImageForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageNotFound creates a GetImageNotFound with default headers values
func NewGetImageNotFound() *GetImageNotFound {
	return &GetImageNotFound{}
}

/*GetImageNotFound handles this case with default header values.

Error.
*/
type GetImageNotFound struct {
	Payload *models.Error
}

func (o *GetImageNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images/{image_id}][%d] getImageNotFound  %+v", 404, o.Payload)
}

func (o *GetImageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageInternalServerError creates a GetImageInternalServerError with default headers values
func NewGetImageInternalServerError() *GetImageInternalServerError {
	return &GetImageInternalServerError{}
}

/*GetImageInternalServerError handles this case with default header values.

Error.
*/
type GetImageInternalServerError struct {
	Payload *models.Error
}

func (o *GetImageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images/{image_id}][%d] getImageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetImageInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageServiceUnavailable creates a GetImageServiceUnavailable with default headers values
func NewGetImageServiceUnavailable() *GetImageServiceUnavailable {
	return &GetImageServiceUnavailable{}
}

/*GetImageServiceUnavailable handles this case with default header values.

Unavailable.
*/
type GetImageServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetImageServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/images/{image_id}][%d] getImageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetImageServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetImageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
