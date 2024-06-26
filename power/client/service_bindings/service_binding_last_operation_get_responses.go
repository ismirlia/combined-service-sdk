// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBindingLastOperationGetReader is a Reader for the ServiceBindingLastOperationGet structure.
type ServiceBindingLastOperationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingLastOperationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBindingLastOperationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBindingLastOperationGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBindingLastOperationGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBindingLastOperationGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBindingLastOperationGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewServiceBindingLastOperationGetGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation] serviceBinding.lastOperation.get", response, response.Code())
	}
}

// NewServiceBindingLastOperationGetOK creates a ServiceBindingLastOperationGetOK with default headers values
func NewServiceBindingLastOperationGetOK() *ServiceBindingLastOperationGetOK {
	return &ServiceBindingLastOperationGetOK{}
}

/*
ServiceBindingLastOperationGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBindingLastOperationGetOK struct {
	Payload *models.LastOperationResource
}

// IsSuccess returns true when this service binding last operation get o k response has a 2xx status code
func (o *ServiceBindingLastOperationGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service binding last operation get o k response has a 3xx status code
func (o *ServiceBindingLastOperationGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get o k response has a 4xx status code
func (o *ServiceBindingLastOperationGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service binding last operation get o k response has a 5xx status code
func (o *ServiceBindingLastOperationGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get o k response a status code equal to that given
func (o *ServiceBindingLastOperationGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service binding last operation get o k response
func (o *ServiceBindingLastOperationGetOK) Code() int {
	return 200
}

func (o *ServiceBindingLastOperationGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetOK %s", 200, payload)
}

func (o *ServiceBindingLastOperationGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetOK %s", 200, payload)
}

func (o *ServiceBindingLastOperationGetOK) GetPayload() *models.LastOperationResource {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LastOperationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetBadRequest creates a ServiceBindingLastOperationGetBadRequest with default headers values
func NewServiceBindingLastOperationGetBadRequest() *ServiceBindingLastOperationGetBadRequest {
	return &ServiceBindingLastOperationGetBadRequest{}
}

/*
ServiceBindingLastOperationGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBindingLastOperationGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding last operation get bad request response has a 2xx status code
func (o *ServiceBindingLastOperationGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding last operation get bad request response has a 3xx status code
func (o *ServiceBindingLastOperationGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get bad request response has a 4xx status code
func (o *ServiceBindingLastOperationGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding last operation get bad request response has a 5xx status code
func (o *ServiceBindingLastOperationGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get bad request response a status code equal to that given
func (o *ServiceBindingLastOperationGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service binding last operation get bad request response
func (o *ServiceBindingLastOperationGetBadRequest) Code() int {
	return 400
}

func (o *ServiceBindingLastOperationGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetBadRequest %s", 400, payload)
}

func (o *ServiceBindingLastOperationGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetBadRequest %s", 400, payload)
}

func (o *ServiceBindingLastOperationGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetUnauthorized creates a ServiceBindingLastOperationGetUnauthorized with default headers values
func NewServiceBindingLastOperationGetUnauthorized() *ServiceBindingLastOperationGetUnauthorized {
	return &ServiceBindingLastOperationGetUnauthorized{}
}

/*
ServiceBindingLastOperationGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBindingLastOperationGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding last operation get unauthorized response has a 2xx status code
func (o *ServiceBindingLastOperationGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding last operation get unauthorized response has a 3xx status code
func (o *ServiceBindingLastOperationGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get unauthorized response has a 4xx status code
func (o *ServiceBindingLastOperationGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding last operation get unauthorized response has a 5xx status code
func (o *ServiceBindingLastOperationGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get unauthorized response a status code equal to that given
func (o *ServiceBindingLastOperationGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service binding last operation get unauthorized response
func (o *ServiceBindingLastOperationGetUnauthorized) Code() int {
	return 401
}

func (o *ServiceBindingLastOperationGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetUnauthorized %s", 401, payload)
}

func (o *ServiceBindingLastOperationGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetUnauthorized %s", 401, payload)
}

func (o *ServiceBindingLastOperationGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetForbidden creates a ServiceBindingLastOperationGetForbidden with default headers values
func NewServiceBindingLastOperationGetForbidden() *ServiceBindingLastOperationGetForbidden {
	return &ServiceBindingLastOperationGetForbidden{}
}

/*
ServiceBindingLastOperationGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBindingLastOperationGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding last operation get forbidden response has a 2xx status code
func (o *ServiceBindingLastOperationGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding last operation get forbidden response has a 3xx status code
func (o *ServiceBindingLastOperationGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get forbidden response has a 4xx status code
func (o *ServiceBindingLastOperationGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding last operation get forbidden response has a 5xx status code
func (o *ServiceBindingLastOperationGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get forbidden response a status code equal to that given
func (o *ServiceBindingLastOperationGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service binding last operation get forbidden response
func (o *ServiceBindingLastOperationGetForbidden) Code() int {
	return 403
}

func (o *ServiceBindingLastOperationGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetForbidden %s", 403, payload)
}

func (o *ServiceBindingLastOperationGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetForbidden %s", 403, payload)
}

func (o *ServiceBindingLastOperationGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetNotFound creates a ServiceBindingLastOperationGetNotFound with default headers values
func NewServiceBindingLastOperationGetNotFound() *ServiceBindingLastOperationGetNotFound {
	return &ServiceBindingLastOperationGetNotFound{}
}

/*
ServiceBindingLastOperationGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBindingLastOperationGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding last operation get not found response has a 2xx status code
func (o *ServiceBindingLastOperationGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding last operation get not found response has a 3xx status code
func (o *ServiceBindingLastOperationGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get not found response has a 4xx status code
func (o *ServiceBindingLastOperationGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding last operation get not found response has a 5xx status code
func (o *ServiceBindingLastOperationGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get not found response a status code equal to that given
func (o *ServiceBindingLastOperationGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service binding last operation get not found response
func (o *ServiceBindingLastOperationGetNotFound) Code() int {
	return 404
}

func (o *ServiceBindingLastOperationGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetNotFound %s", 404, payload)
}

func (o *ServiceBindingLastOperationGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetNotFound %s", 404, payload)
}

func (o *ServiceBindingLastOperationGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetGone creates a ServiceBindingLastOperationGetGone with default headers values
func NewServiceBindingLastOperationGetGone() *ServiceBindingLastOperationGetGone {
	return &ServiceBindingLastOperationGetGone{}
}

/*
ServiceBindingLastOperationGetGone describes a response with status code 410, with default header values.

Gone
*/
type ServiceBindingLastOperationGetGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding last operation get gone response has a 2xx status code
func (o *ServiceBindingLastOperationGetGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding last operation get gone response has a 3xx status code
func (o *ServiceBindingLastOperationGetGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get gone response has a 4xx status code
func (o *ServiceBindingLastOperationGetGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding last operation get gone response has a 5xx status code
func (o *ServiceBindingLastOperationGetGone) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get gone response a status code equal to that given
func (o *ServiceBindingLastOperationGetGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the service binding last operation get gone response
func (o *ServiceBindingLastOperationGetGone) Code() int {
	return 410
}

func (o *ServiceBindingLastOperationGetGone) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetGone %s", 410, payload)
}

func (o *ServiceBindingLastOperationGetGone) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetGone %s", 410, payload)
}

func (o *ServiceBindingLastOperationGetGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
