// Code generated by go-swagger; DO NOT EDIT.

package deviceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// PutDeviceProfileTranslationsReader is a Reader for the PutDeviceProfileTranslations structure.
type PutDeviceProfileTranslationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDeviceProfileTranslationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutDeviceProfileTranslationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutDeviceProfileTranslationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutDeviceProfileTranslationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutDeviceProfileTranslationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPutDeviceProfileTranslationsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPutDeviceProfileTranslationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutDeviceProfileTranslationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutDeviceProfileTranslationsOK creates a PutDeviceProfileTranslationsOK with default headers values
func NewPutDeviceProfileTranslationsOK() *PutDeviceProfileTranslationsOK {
	return &PutDeviceProfileTranslationsOK{}
}

/*PutDeviceProfileTranslationsOK handles this case with default header values.

An device profile.
*/
type PutDeviceProfileTranslationsOK struct {
	Payload *models.DeviceProfileTranslations
}

func (o *PutDeviceProfileTranslationsOK) Error() string {
	return fmt.Sprintf("[PUT /deviceprofiles/{deviceProfileId}/i18n/{locale}][%d] putDeviceProfileTranslationsOK  %+v", 200, o.Payload)
}

func (o *PutDeviceProfileTranslationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceProfileTranslations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDeviceProfileTranslationsBadRequest creates a PutDeviceProfileTranslationsBadRequest with default headers values
func NewPutDeviceProfileTranslationsBadRequest() *PutDeviceProfileTranslationsBadRequest {
	return &PutDeviceProfileTranslationsBadRequest{}
}

/*PutDeviceProfileTranslationsBadRequest handles this case with default header values.

Bad request
*/
type PutDeviceProfileTranslationsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PutDeviceProfileTranslationsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /deviceprofiles/{deviceProfileId}/i18n/{locale}][%d] putDeviceProfileTranslationsBadRequest  %+v", 400, o.Payload)
}

func (o *PutDeviceProfileTranslationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDeviceProfileTranslationsUnauthorized creates a PutDeviceProfileTranslationsUnauthorized with default headers values
func NewPutDeviceProfileTranslationsUnauthorized() *PutDeviceProfileTranslationsUnauthorized {
	return &PutDeviceProfileTranslationsUnauthorized{}
}

/*PutDeviceProfileTranslationsUnauthorized handles this case with default header values.

Unauthorized
*/
type PutDeviceProfileTranslationsUnauthorized struct {
}

func (o *PutDeviceProfileTranslationsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /deviceprofiles/{deviceProfileId}/i18n/{locale}][%d] putDeviceProfileTranslationsUnauthorized ", 401)
}

func (o *PutDeviceProfileTranslationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDeviceProfileTranslationsForbidden creates a PutDeviceProfileTranslationsForbidden with default headers values
func NewPutDeviceProfileTranslationsForbidden() *PutDeviceProfileTranslationsForbidden {
	return &PutDeviceProfileTranslationsForbidden{}
}

/*PutDeviceProfileTranslationsForbidden handles this case with default header values.

Forbidden
*/
type PutDeviceProfileTranslationsForbidden struct {
}

func (o *PutDeviceProfileTranslationsForbidden) Error() string {
	return fmt.Sprintf("[PUT /deviceprofiles/{deviceProfileId}/i18n/{locale}][%d] putDeviceProfileTranslationsForbidden ", 403)
}

func (o *PutDeviceProfileTranslationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDeviceProfileTranslationsUnprocessableEntity creates a PutDeviceProfileTranslationsUnprocessableEntity with default headers values
func NewPutDeviceProfileTranslationsUnprocessableEntity() *PutDeviceProfileTranslationsUnprocessableEntity {
	return &PutDeviceProfileTranslationsUnprocessableEntity{}
}

/*PutDeviceProfileTranslationsUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PutDeviceProfileTranslationsUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *PutDeviceProfileTranslationsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /deviceprofiles/{deviceProfileId}/i18n/{locale}][%d] putDeviceProfileTranslationsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PutDeviceProfileTranslationsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDeviceProfileTranslationsTooManyRequests creates a PutDeviceProfileTranslationsTooManyRequests with default headers values
func NewPutDeviceProfileTranslationsTooManyRequests() *PutDeviceProfileTranslationsTooManyRequests {
	return &PutDeviceProfileTranslationsTooManyRequests{}
}

/*PutDeviceProfileTranslationsTooManyRequests handles this case with default header values.

Too many requests
*/
type PutDeviceProfileTranslationsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *PutDeviceProfileTranslationsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /deviceprofiles/{deviceProfileId}/i18n/{locale}][%d] putDeviceProfileTranslationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutDeviceProfileTranslationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDeviceProfileTranslationsDefault creates a PutDeviceProfileTranslationsDefault with default headers values
func NewPutDeviceProfileTranslationsDefault(code int) *PutDeviceProfileTranslationsDefault {
	return &PutDeviceProfileTranslationsDefault{
		_statusCode: code,
	}
}

/*PutDeviceProfileTranslationsDefault handles this case with default header values.

Unexpected error
*/
type PutDeviceProfileTranslationsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the put device profile translations default response
func (o *PutDeviceProfileTranslationsDefault) Code() int {
	return o._statusCode
}

func (o *PutDeviceProfileTranslationsDefault) Error() string {
	return fmt.Sprintf("[PUT /deviceprofiles/{deviceProfileId}/i18n/{locale}][%d] putDeviceProfileTranslations default  %+v", o._statusCode, o.Payload)
}

func (o *PutDeviceProfileTranslationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}