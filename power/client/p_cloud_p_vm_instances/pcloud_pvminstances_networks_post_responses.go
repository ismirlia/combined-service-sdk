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

// PcloudPvminstancesNetworksPostReader is a Reader for the PcloudPvminstancesNetworksPost structure.
type PcloudPvminstancesNetworksPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesNetworksPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPcloudPvminstancesNetworksPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesNetworksPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesNetworksPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesNetworksPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesNetworksPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesNetworksPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPvminstancesNetworksPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesNetworksPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks] pcloud.pvminstances.networks.post", response, response.Code())
	}
}

// NewPcloudPvminstancesNetworksPostCreated creates a PcloudPvminstancesNetworksPostCreated with default headers values
func NewPcloudPvminstancesNetworksPostCreated() *PcloudPvminstancesNetworksPostCreated {
	return &PcloudPvminstancesNetworksPostCreated{}
}

/*
PcloudPvminstancesNetworksPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudPvminstancesNetworksPostCreated struct {
	Payload *models.PVMInstanceNetwork
}

// IsSuccess returns true when this pcloud pvminstances networks post created response has a 2xx status code
func (o *PcloudPvminstancesNetworksPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances networks post created response has a 3xx status code
func (o *PcloudPvminstancesNetworksPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks post created response has a 4xx status code
func (o *PcloudPvminstancesNetworksPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances networks post created response has a 5xx status code
func (o *PcloudPvminstancesNetworksPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks post created response a status code equal to that given
func (o *PcloudPvminstancesNetworksPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pcloud pvminstances networks post created response
func (o *PcloudPvminstancesNetworksPostCreated) Code() int {
	return 201
}

func (o *PcloudPvminstancesNetworksPostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostCreated %s", 201, payload)
}

func (o *PcloudPvminstancesNetworksPostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostCreated %s", 201, payload)
}

func (o *PcloudPvminstancesNetworksPostCreated) GetPayload() *models.PVMInstanceNetwork {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstanceNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksPostBadRequest creates a PcloudPvminstancesNetworksPostBadRequest with default headers values
func NewPcloudPvminstancesNetworksPostBadRequest() *PcloudPvminstancesNetworksPostBadRequest {
	return &PcloudPvminstancesNetworksPostBadRequest{}
}

/*
PcloudPvminstancesNetworksPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesNetworksPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks post bad request response has a 2xx status code
func (o *PcloudPvminstancesNetworksPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks post bad request response has a 3xx status code
func (o *PcloudPvminstancesNetworksPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks post bad request response has a 4xx status code
func (o *PcloudPvminstancesNetworksPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks post bad request response has a 5xx status code
func (o *PcloudPvminstancesNetworksPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks post bad request response a status code equal to that given
func (o *PcloudPvminstancesNetworksPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances networks post bad request response
func (o *PcloudPvminstancesNetworksPostBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesNetworksPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesNetworksPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesNetworksPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksPostUnauthorized creates a PcloudPvminstancesNetworksPostUnauthorized with default headers values
func NewPcloudPvminstancesNetworksPostUnauthorized() *PcloudPvminstancesNetworksPostUnauthorized {
	return &PcloudPvminstancesNetworksPostUnauthorized{}
}

/*
PcloudPvminstancesNetworksPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesNetworksPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks post unauthorized response has a 2xx status code
func (o *PcloudPvminstancesNetworksPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks post unauthorized response has a 3xx status code
func (o *PcloudPvminstancesNetworksPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks post unauthorized response has a 4xx status code
func (o *PcloudPvminstancesNetworksPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks post unauthorized response has a 5xx status code
func (o *PcloudPvminstancesNetworksPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks post unauthorized response a status code equal to that given
func (o *PcloudPvminstancesNetworksPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances networks post unauthorized response
func (o *PcloudPvminstancesNetworksPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesNetworksPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesNetworksPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesNetworksPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksPostForbidden creates a PcloudPvminstancesNetworksPostForbidden with default headers values
func NewPcloudPvminstancesNetworksPostForbidden() *PcloudPvminstancesNetworksPostForbidden {
	return &PcloudPvminstancesNetworksPostForbidden{}
}

/*
PcloudPvminstancesNetworksPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesNetworksPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks post forbidden response has a 2xx status code
func (o *PcloudPvminstancesNetworksPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks post forbidden response has a 3xx status code
func (o *PcloudPvminstancesNetworksPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks post forbidden response has a 4xx status code
func (o *PcloudPvminstancesNetworksPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks post forbidden response has a 5xx status code
func (o *PcloudPvminstancesNetworksPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks post forbidden response a status code equal to that given
func (o *PcloudPvminstancesNetworksPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances networks post forbidden response
func (o *PcloudPvminstancesNetworksPostForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesNetworksPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesNetworksPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesNetworksPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksPostNotFound creates a PcloudPvminstancesNetworksPostNotFound with default headers values
func NewPcloudPvminstancesNetworksPostNotFound() *PcloudPvminstancesNetworksPostNotFound {
	return &PcloudPvminstancesNetworksPostNotFound{}
}

/*
PcloudPvminstancesNetworksPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesNetworksPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks post not found response has a 2xx status code
func (o *PcloudPvminstancesNetworksPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks post not found response has a 3xx status code
func (o *PcloudPvminstancesNetworksPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks post not found response has a 4xx status code
func (o *PcloudPvminstancesNetworksPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks post not found response has a 5xx status code
func (o *PcloudPvminstancesNetworksPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks post not found response a status code equal to that given
func (o *PcloudPvminstancesNetworksPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances networks post not found response
func (o *PcloudPvminstancesNetworksPostNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesNetworksPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesNetworksPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesNetworksPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksPostConflict creates a PcloudPvminstancesNetworksPostConflict with default headers values
func NewPcloudPvminstancesNetworksPostConflict() *PcloudPvminstancesNetworksPostConflict {
	return &PcloudPvminstancesNetworksPostConflict{}
}

/*
PcloudPvminstancesNetworksPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesNetworksPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks post conflict response has a 2xx status code
func (o *PcloudPvminstancesNetworksPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks post conflict response has a 3xx status code
func (o *PcloudPvminstancesNetworksPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks post conflict response has a 4xx status code
func (o *PcloudPvminstancesNetworksPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks post conflict response has a 5xx status code
func (o *PcloudPvminstancesNetworksPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks post conflict response a status code equal to that given
func (o *PcloudPvminstancesNetworksPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud pvminstances networks post conflict response
func (o *PcloudPvminstancesNetworksPostConflict) Code() int {
	return 409
}

func (o *PcloudPvminstancesNetworksPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostConflict %s", 409, payload)
}

func (o *PcloudPvminstancesNetworksPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostConflict %s", 409, payload)
}

func (o *PcloudPvminstancesNetworksPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksPostUnprocessableEntity creates a PcloudPvminstancesNetworksPostUnprocessableEntity with default headers values
func NewPcloudPvminstancesNetworksPostUnprocessableEntity() *PcloudPvminstancesNetworksPostUnprocessableEntity {
	return &PcloudPvminstancesNetworksPostUnprocessableEntity{}
}

/*
PcloudPvminstancesNetworksPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesNetworksPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks post unprocessable entity response has a 2xx status code
func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks post unprocessable entity response has a 3xx status code
func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks post unprocessable entity response has a 4xx status code
func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks post unprocessable entity response has a 5xx status code
func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks post unprocessable entity response a status code equal to that given
func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud pvminstances networks post unprocessable entity response
func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksPostInternalServerError creates a PcloudPvminstancesNetworksPostInternalServerError with default headers values
func NewPcloudPvminstancesNetworksPostInternalServerError() *PcloudPvminstancesNetworksPostInternalServerError {
	return &PcloudPvminstancesNetworksPostInternalServerError{}
}

/*
PcloudPvminstancesNetworksPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesNetworksPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks post internal server error response has a 2xx status code
func (o *PcloudPvminstancesNetworksPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks post internal server error response has a 3xx status code
func (o *PcloudPvminstancesNetworksPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks post internal server error response has a 4xx status code
func (o *PcloudPvminstancesNetworksPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances networks post internal server error response has a 5xx status code
func (o *PcloudPvminstancesNetworksPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances networks post internal server error response a status code equal to that given
func (o *PcloudPvminstancesNetworksPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances networks post internal server error response
func (o *PcloudPvminstancesNetworksPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesNetworksPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesNetworksPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks][%d] pcloudPvminstancesNetworksPostInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesNetworksPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
