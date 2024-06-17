// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// V1SnapshotsGetReader is a Reader for the V1SnapshotsGet structure.
type V1SnapshotsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SnapshotsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SnapshotsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1SnapshotsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1SnapshotsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1SnapshotsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1SnapshotsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1SnapshotsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/snapshots/{snapshot_id}] v1.snapshots.get", response, response.Code())
	}
}

// NewV1SnapshotsGetOK creates a V1SnapshotsGetOK with default headers values
func NewV1SnapshotsGetOK() *V1SnapshotsGetOK {
	return &V1SnapshotsGetOK{}
}

/*
V1SnapshotsGetOK describes a response with status code 200, with default header values.

OK
*/
type V1SnapshotsGetOK struct {
	Payload *models.SnapshotV1
}

// IsSuccess returns true when this v1 snapshots get o k response has a 2xx status code
func (o *V1SnapshotsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 snapshots get o k response has a 3xx status code
func (o *V1SnapshotsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 snapshots get o k response has a 4xx status code
func (o *V1SnapshotsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 snapshots get o k response has a 5xx status code
func (o *V1SnapshotsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 snapshots get o k response a status code equal to that given
func (o *V1SnapshotsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 snapshots get o k response
func (o *V1SnapshotsGetOK) Code() int {
	return 200
}

func (o *V1SnapshotsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetOK %s", 200, payload)
}

func (o *V1SnapshotsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetOK %s", 200, payload)
}

func (o *V1SnapshotsGetOK) GetPayload() *models.SnapshotV1 {
	return o.Payload
}

func (o *V1SnapshotsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SnapshotsGetBadRequest creates a V1SnapshotsGetBadRequest with default headers values
func NewV1SnapshotsGetBadRequest() *V1SnapshotsGetBadRequest {
	return &V1SnapshotsGetBadRequest{}
}

/*
V1SnapshotsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1SnapshotsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 snapshots get bad request response has a 2xx status code
func (o *V1SnapshotsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 snapshots get bad request response has a 3xx status code
func (o *V1SnapshotsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 snapshots get bad request response has a 4xx status code
func (o *V1SnapshotsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 snapshots get bad request response has a 5xx status code
func (o *V1SnapshotsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 snapshots get bad request response a status code equal to that given
func (o *V1SnapshotsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 snapshots get bad request response
func (o *V1SnapshotsGetBadRequest) Code() int {
	return 400
}

func (o *V1SnapshotsGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetBadRequest %s", 400, payload)
}

func (o *V1SnapshotsGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetBadRequest %s", 400, payload)
}

func (o *V1SnapshotsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SnapshotsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SnapshotsGetUnauthorized creates a V1SnapshotsGetUnauthorized with default headers values
func NewV1SnapshotsGetUnauthorized() *V1SnapshotsGetUnauthorized {
	return &V1SnapshotsGetUnauthorized{}
}

/*
V1SnapshotsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1SnapshotsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 snapshots get unauthorized response has a 2xx status code
func (o *V1SnapshotsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 snapshots get unauthorized response has a 3xx status code
func (o *V1SnapshotsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 snapshots get unauthorized response has a 4xx status code
func (o *V1SnapshotsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 snapshots get unauthorized response has a 5xx status code
func (o *V1SnapshotsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 snapshots get unauthorized response a status code equal to that given
func (o *V1SnapshotsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 snapshots get unauthorized response
func (o *V1SnapshotsGetUnauthorized) Code() int {
	return 401
}

func (o *V1SnapshotsGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetUnauthorized %s", 401, payload)
}

func (o *V1SnapshotsGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetUnauthorized %s", 401, payload)
}

func (o *V1SnapshotsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SnapshotsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SnapshotsGetForbidden creates a V1SnapshotsGetForbidden with default headers values
func NewV1SnapshotsGetForbidden() *V1SnapshotsGetForbidden {
	return &V1SnapshotsGetForbidden{}
}

/*
V1SnapshotsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1SnapshotsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 snapshots get forbidden response has a 2xx status code
func (o *V1SnapshotsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 snapshots get forbidden response has a 3xx status code
func (o *V1SnapshotsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 snapshots get forbidden response has a 4xx status code
func (o *V1SnapshotsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 snapshots get forbidden response has a 5xx status code
func (o *V1SnapshotsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 snapshots get forbidden response a status code equal to that given
func (o *V1SnapshotsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 snapshots get forbidden response
func (o *V1SnapshotsGetForbidden) Code() int {
	return 403
}

func (o *V1SnapshotsGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetForbidden %s", 403, payload)
}

func (o *V1SnapshotsGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetForbidden %s", 403, payload)
}

func (o *V1SnapshotsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SnapshotsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SnapshotsGetNotFound creates a V1SnapshotsGetNotFound with default headers values
func NewV1SnapshotsGetNotFound() *V1SnapshotsGetNotFound {
	return &V1SnapshotsGetNotFound{}
}

/*
V1SnapshotsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1SnapshotsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 snapshots get not found response has a 2xx status code
func (o *V1SnapshotsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 snapshots get not found response has a 3xx status code
func (o *V1SnapshotsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 snapshots get not found response has a 4xx status code
func (o *V1SnapshotsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 snapshots get not found response has a 5xx status code
func (o *V1SnapshotsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 snapshots get not found response a status code equal to that given
func (o *V1SnapshotsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 snapshots get not found response
func (o *V1SnapshotsGetNotFound) Code() int {
	return 404
}

func (o *V1SnapshotsGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetNotFound %s", 404, payload)
}

func (o *V1SnapshotsGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetNotFound %s", 404, payload)
}

func (o *V1SnapshotsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SnapshotsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1SnapshotsGetInternalServerError creates a V1SnapshotsGetInternalServerError with default headers values
func NewV1SnapshotsGetInternalServerError() *V1SnapshotsGetInternalServerError {
	return &V1SnapshotsGetInternalServerError{}
}

/*
V1SnapshotsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1SnapshotsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 snapshots get internal server error response has a 2xx status code
func (o *V1SnapshotsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 snapshots get internal server error response has a 3xx status code
func (o *V1SnapshotsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 snapshots get internal server error response has a 4xx status code
func (o *V1SnapshotsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 snapshots get internal server error response has a 5xx status code
func (o *V1SnapshotsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 snapshots get internal server error response a status code equal to that given
func (o *V1SnapshotsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 snapshots get internal server error response
func (o *V1SnapshotsGetInternalServerError) Code() int {
	return 500
}

func (o *V1SnapshotsGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetInternalServerError %s", 500, payload)
}

func (o *V1SnapshotsGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/snapshots/{snapshot_id}][%d] v1SnapshotsGetInternalServerError %s", 500, payload)
}

func (o *V1SnapshotsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1SnapshotsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}