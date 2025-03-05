// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// PcloudPvminstancesConsolePostReader is a Reader for the PcloudPvminstancesConsolePost structure.
type PcloudPvminstancesConsolePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesConsolePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPcloudPvminstancesConsolePostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesConsolePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesConsolePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesConsolePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesConsolePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesConsolePostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPvminstancesConsolePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesConsolePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console] pcloud.pvminstances.console.post", response, response.Code())
	}
}

// NewPcloudPvminstancesConsolePostCreated creates a PcloudPvminstancesConsolePostCreated with default headers values
func NewPcloudPvminstancesConsolePostCreated() *PcloudPvminstancesConsolePostCreated {
	return &PcloudPvminstancesConsolePostCreated{}
}

/*
PcloudPvminstancesConsolePostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudPvminstancesConsolePostCreated struct {
	Payload *models.PVMInstanceConsole
}

// IsSuccess returns true when this pcloud pvminstances console post created response has a 2xx status code
func (o *PcloudPvminstancesConsolePostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances console post created response has a 3xx status code
func (o *PcloudPvminstancesConsolePostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console post created response has a 4xx status code
func (o *PcloudPvminstancesConsolePostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances console post created response has a 5xx status code
func (o *PcloudPvminstancesConsolePostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console post created response a status code equal to that given
func (o *PcloudPvminstancesConsolePostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pcloud pvminstances console post created response
func (o *PcloudPvminstancesConsolePostCreated) Code() int {
	return 201
}

func (o *PcloudPvminstancesConsolePostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostCreated %s", 201, payload)
}

func (o *PcloudPvminstancesConsolePostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostCreated %s", 201, payload)
}

func (o *PcloudPvminstancesConsolePostCreated) GetPayload() *models.PVMInstanceConsole {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstanceConsole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostBadRequest creates a PcloudPvminstancesConsolePostBadRequest with default headers values
func NewPcloudPvminstancesConsolePostBadRequest() *PcloudPvminstancesConsolePostBadRequest {
	return &PcloudPvminstancesConsolePostBadRequest{}
}

/*
PcloudPvminstancesConsolePostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesConsolePostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console post bad request response has a 2xx status code
func (o *PcloudPvminstancesConsolePostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console post bad request response has a 3xx status code
func (o *PcloudPvminstancesConsolePostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console post bad request response has a 4xx status code
func (o *PcloudPvminstancesConsolePostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console post bad request response has a 5xx status code
func (o *PcloudPvminstancesConsolePostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console post bad request response a status code equal to that given
func (o *PcloudPvminstancesConsolePostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances console post bad request response
func (o *PcloudPvminstancesConsolePostBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesConsolePostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesConsolePostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesConsolePostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostUnauthorized creates a PcloudPvminstancesConsolePostUnauthorized with default headers values
func NewPcloudPvminstancesConsolePostUnauthorized() *PcloudPvminstancesConsolePostUnauthorized {
	return &PcloudPvminstancesConsolePostUnauthorized{}
}

/*
PcloudPvminstancesConsolePostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesConsolePostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console post unauthorized response has a 2xx status code
func (o *PcloudPvminstancesConsolePostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console post unauthorized response has a 3xx status code
func (o *PcloudPvminstancesConsolePostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console post unauthorized response has a 4xx status code
func (o *PcloudPvminstancesConsolePostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console post unauthorized response has a 5xx status code
func (o *PcloudPvminstancesConsolePostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console post unauthorized response a status code equal to that given
func (o *PcloudPvminstancesConsolePostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances console post unauthorized response
func (o *PcloudPvminstancesConsolePostUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesConsolePostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesConsolePostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesConsolePostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostForbidden creates a PcloudPvminstancesConsolePostForbidden with default headers values
func NewPcloudPvminstancesConsolePostForbidden() *PcloudPvminstancesConsolePostForbidden {
	return &PcloudPvminstancesConsolePostForbidden{}
}

/*
PcloudPvminstancesConsolePostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesConsolePostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console post forbidden response has a 2xx status code
func (o *PcloudPvminstancesConsolePostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console post forbidden response has a 3xx status code
func (o *PcloudPvminstancesConsolePostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console post forbidden response has a 4xx status code
func (o *PcloudPvminstancesConsolePostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console post forbidden response has a 5xx status code
func (o *PcloudPvminstancesConsolePostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console post forbidden response a status code equal to that given
func (o *PcloudPvminstancesConsolePostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances console post forbidden response
func (o *PcloudPvminstancesConsolePostForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesConsolePostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesConsolePostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesConsolePostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostNotFound creates a PcloudPvminstancesConsolePostNotFound with default headers values
func NewPcloudPvminstancesConsolePostNotFound() *PcloudPvminstancesConsolePostNotFound {
	return &PcloudPvminstancesConsolePostNotFound{}
}

/*
PcloudPvminstancesConsolePostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesConsolePostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console post not found response has a 2xx status code
func (o *PcloudPvminstancesConsolePostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console post not found response has a 3xx status code
func (o *PcloudPvminstancesConsolePostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console post not found response has a 4xx status code
func (o *PcloudPvminstancesConsolePostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console post not found response has a 5xx status code
func (o *PcloudPvminstancesConsolePostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console post not found response a status code equal to that given
func (o *PcloudPvminstancesConsolePostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances console post not found response
func (o *PcloudPvminstancesConsolePostNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesConsolePostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesConsolePostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesConsolePostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostConflict creates a PcloudPvminstancesConsolePostConflict with default headers values
func NewPcloudPvminstancesConsolePostConflict() *PcloudPvminstancesConsolePostConflict {
	return &PcloudPvminstancesConsolePostConflict{}
}

/*
PcloudPvminstancesConsolePostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesConsolePostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console post conflict response has a 2xx status code
func (o *PcloudPvminstancesConsolePostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console post conflict response has a 3xx status code
func (o *PcloudPvminstancesConsolePostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console post conflict response has a 4xx status code
func (o *PcloudPvminstancesConsolePostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console post conflict response has a 5xx status code
func (o *PcloudPvminstancesConsolePostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console post conflict response a status code equal to that given
func (o *PcloudPvminstancesConsolePostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud pvminstances console post conflict response
func (o *PcloudPvminstancesConsolePostConflict) Code() int {
	return 409
}

func (o *PcloudPvminstancesConsolePostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostConflict %s", 409, payload)
}

func (o *PcloudPvminstancesConsolePostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostConflict %s", 409, payload)
}

func (o *PcloudPvminstancesConsolePostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostUnprocessableEntity creates a PcloudPvminstancesConsolePostUnprocessableEntity with default headers values
func NewPcloudPvminstancesConsolePostUnprocessableEntity() *PcloudPvminstancesConsolePostUnprocessableEntity {
	return &PcloudPvminstancesConsolePostUnprocessableEntity{}
}

/*
PcloudPvminstancesConsolePostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesConsolePostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console post unprocessable entity response has a 2xx status code
func (o *PcloudPvminstancesConsolePostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console post unprocessable entity response has a 3xx status code
func (o *PcloudPvminstancesConsolePostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console post unprocessable entity response has a 4xx status code
func (o *PcloudPvminstancesConsolePostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console post unprocessable entity response has a 5xx status code
func (o *PcloudPvminstancesConsolePostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console post unprocessable entity response a status code equal to that given
func (o *PcloudPvminstancesConsolePostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud pvminstances console post unprocessable entity response
func (o *PcloudPvminstancesConsolePostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudPvminstancesConsolePostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudPvminstancesConsolePostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudPvminstancesConsolePostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostInternalServerError creates a PcloudPvminstancesConsolePostInternalServerError with default headers values
func NewPcloudPvminstancesConsolePostInternalServerError() *PcloudPvminstancesConsolePostInternalServerError {
	return &PcloudPvminstancesConsolePostInternalServerError{}
}

/*
PcloudPvminstancesConsolePostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesConsolePostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console post internal server error response has a 2xx status code
func (o *PcloudPvminstancesConsolePostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console post internal server error response has a 3xx status code
func (o *PcloudPvminstancesConsolePostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console post internal server error response has a 4xx status code
func (o *PcloudPvminstancesConsolePostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances console post internal server error response has a 5xx status code
func (o *PcloudPvminstancesConsolePostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances console post internal server error response a status code equal to that given
func (o *PcloudPvminstancesConsolePostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances console post internal server error response
func (o *PcloudPvminstancesConsolePostInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesConsolePostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesConsolePostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesConsolePostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
