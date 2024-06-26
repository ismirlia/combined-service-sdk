// Code generated by go-swagger; DO NOT EDIT.

package host_groups

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

// V1HostsGetReader is a Reader for the V1HostsGet structure.
type V1HostsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1HostsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1HostsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1HostsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1HostsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1HostsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1HostsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewV1HostsGetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/hosts] v1.hosts.get", response, response.Code())
	}
}

// NewV1HostsGetOK creates a V1HostsGetOK with default headers values
func NewV1HostsGetOK() *V1HostsGetOK {
	return &V1HostsGetOK{}
}

/*
V1HostsGetOK describes a response with status code 200, with default header values.

OK
*/
type V1HostsGetOK struct {
	Payload models.HostList
}

// IsSuccess returns true when this v1 hosts get o k response has a 2xx status code
func (o *V1HostsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 hosts get o k response has a 3xx status code
func (o *V1HostsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts get o k response has a 4xx status code
func (o *V1HostsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hosts get o k response has a 5xx status code
func (o *V1HostsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts get o k response a status code equal to that given
func (o *V1HostsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 hosts get o k response
func (o *V1HostsGetOK) Code() int {
	return 200
}

func (o *V1HostsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetOK %s", 200, payload)
}

func (o *V1HostsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetOK %s", 200, payload)
}

func (o *V1HostsGetOK) GetPayload() models.HostList {
	return o.Payload
}

func (o *V1HostsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsGetBadRequest creates a V1HostsGetBadRequest with default headers values
func NewV1HostsGetBadRequest() *V1HostsGetBadRequest {
	return &V1HostsGetBadRequest{}
}

/*
V1HostsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1HostsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts get bad request response has a 2xx status code
func (o *V1HostsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts get bad request response has a 3xx status code
func (o *V1HostsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts get bad request response has a 4xx status code
func (o *V1HostsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hosts get bad request response has a 5xx status code
func (o *V1HostsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts get bad request response a status code equal to that given
func (o *V1HostsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 hosts get bad request response
func (o *V1HostsGetBadRequest) Code() int {
	return 400
}

func (o *V1HostsGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetBadRequest %s", 400, payload)
}

func (o *V1HostsGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetBadRequest %s", 400, payload)
}

func (o *V1HostsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsGetUnauthorized creates a V1HostsGetUnauthorized with default headers values
func NewV1HostsGetUnauthorized() *V1HostsGetUnauthorized {
	return &V1HostsGetUnauthorized{}
}

/*
V1HostsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1HostsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts get unauthorized response has a 2xx status code
func (o *V1HostsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts get unauthorized response has a 3xx status code
func (o *V1HostsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts get unauthorized response has a 4xx status code
func (o *V1HostsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hosts get unauthorized response has a 5xx status code
func (o *V1HostsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts get unauthorized response a status code equal to that given
func (o *V1HostsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 hosts get unauthorized response
func (o *V1HostsGetUnauthorized) Code() int {
	return 401
}

func (o *V1HostsGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetUnauthorized %s", 401, payload)
}

func (o *V1HostsGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetUnauthorized %s", 401, payload)
}

func (o *V1HostsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsGetForbidden creates a V1HostsGetForbidden with default headers values
func NewV1HostsGetForbidden() *V1HostsGetForbidden {
	return &V1HostsGetForbidden{}
}

/*
V1HostsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1HostsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts get forbidden response has a 2xx status code
func (o *V1HostsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts get forbidden response has a 3xx status code
func (o *V1HostsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts get forbidden response has a 4xx status code
func (o *V1HostsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 hosts get forbidden response has a 5xx status code
func (o *V1HostsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 hosts get forbidden response a status code equal to that given
func (o *V1HostsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 hosts get forbidden response
func (o *V1HostsGetForbidden) Code() int {
	return 403
}

func (o *V1HostsGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetForbidden %s", 403, payload)
}

func (o *V1HostsGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetForbidden %s", 403, payload)
}

func (o *V1HostsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsGetInternalServerError creates a V1HostsGetInternalServerError with default headers values
func NewV1HostsGetInternalServerError() *V1HostsGetInternalServerError {
	return &V1HostsGetInternalServerError{}
}

/*
V1HostsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1HostsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts get internal server error response has a 2xx status code
func (o *V1HostsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts get internal server error response has a 3xx status code
func (o *V1HostsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts get internal server error response has a 4xx status code
func (o *V1HostsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hosts get internal server error response has a 5xx status code
func (o *V1HostsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 hosts get internal server error response a status code equal to that given
func (o *V1HostsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 hosts get internal server error response
func (o *V1HostsGetInternalServerError) Code() int {
	return 500
}

func (o *V1HostsGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetInternalServerError %s", 500, payload)
}

func (o *V1HostsGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetInternalServerError %s", 500, payload)
}

func (o *V1HostsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1HostsGetGatewayTimeout creates a V1HostsGetGatewayTimeout with default headers values
func NewV1HostsGetGatewayTimeout() *V1HostsGetGatewayTimeout {
	return &V1HostsGetGatewayTimeout{}
}

/*
V1HostsGetGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. Request is still processing and taking longer than expected.
*/
type V1HostsGetGatewayTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 hosts get gateway timeout response has a 2xx status code
func (o *V1HostsGetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 hosts get gateway timeout response has a 3xx status code
func (o *V1HostsGetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 hosts get gateway timeout response has a 4xx status code
func (o *V1HostsGetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 hosts get gateway timeout response has a 5xx status code
func (o *V1HostsGetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 hosts get gateway timeout response a status code equal to that given
func (o *V1HostsGetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the v1 hosts get gateway timeout response
func (o *V1HostsGetGatewayTimeout) Code() int {
	return 504
}

func (o *V1HostsGetGatewayTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetGatewayTimeout %s", 504, payload)
}

func (o *V1HostsGetGatewayTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts][%d] v1HostsGetGatewayTimeout %s", 504, payload)
}

func (o *V1HostsGetGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1HostsGetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
