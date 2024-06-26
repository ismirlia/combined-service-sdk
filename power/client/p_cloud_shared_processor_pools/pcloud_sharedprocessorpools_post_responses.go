// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_shared_processor_pools

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

// PcloudSharedprocessorpoolsPostReader is a Reader for the PcloudSharedprocessorpoolsPost structure.
type PcloudSharedprocessorpoolsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSharedprocessorpoolsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudSharedprocessorpoolsPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSharedprocessorpoolsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSharedprocessorpoolsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudSharedprocessorpoolsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudSharedprocessorpoolsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudSharedprocessorpoolsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudSharedprocessorpoolsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSharedprocessorpoolsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools] pcloud.sharedprocessorpools.post", response, response.Code())
	}
}

// NewPcloudSharedprocessorpoolsPostAccepted creates a PcloudSharedprocessorpoolsPostAccepted with default headers values
func NewPcloudSharedprocessorpoolsPostAccepted() *PcloudSharedprocessorpoolsPostAccepted {
	return &PcloudSharedprocessorpoolsPostAccepted{}
}

/*
PcloudSharedprocessorpoolsPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudSharedprocessorpoolsPostAccepted struct {
	Payload *models.SharedProcessorPool
}

// IsSuccess returns true when this pcloud sharedprocessorpools post accepted response has a 2xx status code
func (o *PcloudSharedprocessorpoolsPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sharedprocessorpools post accepted response has a 3xx status code
func (o *PcloudSharedprocessorpoolsPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools post accepted response has a 4xx status code
func (o *PcloudSharedprocessorpoolsPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sharedprocessorpools post accepted response has a 5xx status code
func (o *PcloudSharedprocessorpoolsPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools post accepted response a status code equal to that given
func (o *PcloudSharedprocessorpoolsPostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud sharedprocessorpools post accepted response
func (o *PcloudSharedprocessorpoolsPostAccepted) Code() int {
	return 202
}

func (o *PcloudSharedprocessorpoolsPostAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostAccepted %s", 202, payload)
}

func (o *PcloudSharedprocessorpoolsPostAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostAccepted %s", 202, payload)
}

func (o *PcloudSharedprocessorpoolsPostAccepted) GetPayload() *models.SharedProcessorPool {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SharedProcessorPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsPostBadRequest creates a PcloudSharedprocessorpoolsPostBadRequest with default headers values
func NewPcloudSharedprocessorpoolsPostBadRequest() *PcloudSharedprocessorpoolsPostBadRequest {
	return &PcloudSharedprocessorpoolsPostBadRequest{}
}

/*
PcloudSharedprocessorpoolsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSharedprocessorpoolsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools post bad request response has a 2xx status code
func (o *PcloudSharedprocessorpoolsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools post bad request response has a 3xx status code
func (o *PcloudSharedprocessorpoolsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools post bad request response has a 4xx status code
func (o *PcloudSharedprocessorpoolsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools post bad request response has a 5xx status code
func (o *PcloudSharedprocessorpoolsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools post bad request response a status code equal to that given
func (o *PcloudSharedprocessorpoolsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud sharedprocessorpools post bad request response
func (o *PcloudSharedprocessorpoolsPostBadRequest) Code() int {
	return 400
}

func (o *PcloudSharedprocessorpoolsPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostBadRequest %s", 400, payload)
}

func (o *PcloudSharedprocessorpoolsPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostBadRequest %s", 400, payload)
}

func (o *PcloudSharedprocessorpoolsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsPostUnauthorized creates a PcloudSharedprocessorpoolsPostUnauthorized with default headers values
func NewPcloudSharedprocessorpoolsPostUnauthorized() *PcloudSharedprocessorpoolsPostUnauthorized {
	return &PcloudSharedprocessorpoolsPostUnauthorized{}
}

/*
PcloudSharedprocessorpoolsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSharedprocessorpoolsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools post unauthorized response has a 2xx status code
func (o *PcloudSharedprocessorpoolsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools post unauthorized response has a 3xx status code
func (o *PcloudSharedprocessorpoolsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools post unauthorized response has a 4xx status code
func (o *PcloudSharedprocessorpoolsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools post unauthorized response has a 5xx status code
func (o *PcloudSharedprocessorpoolsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools post unauthorized response a status code equal to that given
func (o *PcloudSharedprocessorpoolsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud sharedprocessorpools post unauthorized response
func (o *PcloudSharedprocessorpoolsPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudSharedprocessorpoolsPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostUnauthorized %s", 401, payload)
}

func (o *PcloudSharedprocessorpoolsPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostUnauthorized %s", 401, payload)
}

func (o *PcloudSharedprocessorpoolsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsPostForbidden creates a PcloudSharedprocessorpoolsPostForbidden with default headers values
func NewPcloudSharedprocessorpoolsPostForbidden() *PcloudSharedprocessorpoolsPostForbidden {
	return &PcloudSharedprocessorpoolsPostForbidden{}
}

/*
PcloudSharedprocessorpoolsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudSharedprocessorpoolsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools post forbidden response has a 2xx status code
func (o *PcloudSharedprocessorpoolsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools post forbidden response has a 3xx status code
func (o *PcloudSharedprocessorpoolsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools post forbidden response has a 4xx status code
func (o *PcloudSharedprocessorpoolsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools post forbidden response has a 5xx status code
func (o *PcloudSharedprocessorpoolsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools post forbidden response a status code equal to that given
func (o *PcloudSharedprocessorpoolsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud sharedprocessorpools post forbidden response
func (o *PcloudSharedprocessorpoolsPostForbidden) Code() int {
	return 403
}

func (o *PcloudSharedprocessorpoolsPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostForbidden %s", 403, payload)
}

func (o *PcloudSharedprocessorpoolsPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostForbidden %s", 403, payload)
}

func (o *PcloudSharedprocessorpoolsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsPostNotFound creates a PcloudSharedprocessorpoolsPostNotFound with default headers values
func NewPcloudSharedprocessorpoolsPostNotFound() *PcloudSharedprocessorpoolsPostNotFound {
	return &PcloudSharedprocessorpoolsPostNotFound{}
}

/*
PcloudSharedprocessorpoolsPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudSharedprocessorpoolsPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools post not found response has a 2xx status code
func (o *PcloudSharedprocessorpoolsPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools post not found response has a 3xx status code
func (o *PcloudSharedprocessorpoolsPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools post not found response has a 4xx status code
func (o *PcloudSharedprocessorpoolsPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools post not found response has a 5xx status code
func (o *PcloudSharedprocessorpoolsPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools post not found response a status code equal to that given
func (o *PcloudSharedprocessorpoolsPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud sharedprocessorpools post not found response
func (o *PcloudSharedprocessorpoolsPostNotFound) Code() int {
	return 404
}

func (o *PcloudSharedprocessorpoolsPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostNotFound %s", 404, payload)
}

func (o *PcloudSharedprocessorpoolsPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostNotFound %s", 404, payload)
}

func (o *PcloudSharedprocessorpoolsPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsPostConflict creates a PcloudSharedprocessorpoolsPostConflict with default headers values
func NewPcloudSharedprocessorpoolsPostConflict() *PcloudSharedprocessorpoolsPostConflict {
	return &PcloudSharedprocessorpoolsPostConflict{}
}

/*
PcloudSharedprocessorpoolsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudSharedprocessorpoolsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools post conflict response has a 2xx status code
func (o *PcloudSharedprocessorpoolsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools post conflict response has a 3xx status code
func (o *PcloudSharedprocessorpoolsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools post conflict response has a 4xx status code
func (o *PcloudSharedprocessorpoolsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools post conflict response has a 5xx status code
func (o *PcloudSharedprocessorpoolsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools post conflict response a status code equal to that given
func (o *PcloudSharedprocessorpoolsPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud sharedprocessorpools post conflict response
func (o *PcloudSharedprocessorpoolsPostConflict) Code() int {
	return 409
}

func (o *PcloudSharedprocessorpoolsPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostConflict %s", 409, payload)
}

func (o *PcloudSharedprocessorpoolsPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostConflict %s", 409, payload)
}

func (o *PcloudSharedprocessorpoolsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsPostUnprocessableEntity creates a PcloudSharedprocessorpoolsPostUnprocessableEntity with default headers values
func NewPcloudSharedprocessorpoolsPostUnprocessableEntity() *PcloudSharedprocessorpoolsPostUnprocessableEntity {
	return &PcloudSharedprocessorpoolsPostUnprocessableEntity{}
}

/*
PcloudSharedprocessorpoolsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudSharedprocessorpoolsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools post unprocessable entity response has a 2xx status code
func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools post unprocessable entity response has a 3xx status code
func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools post unprocessable entity response has a 4xx status code
func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sharedprocessorpools post unprocessable entity response has a 5xx status code
func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sharedprocessorpools post unprocessable entity response a status code equal to that given
func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud sharedprocessorpools post unprocessable entity response
func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsPostInternalServerError creates a PcloudSharedprocessorpoolsPostInternalServerError with default headers values
func NewPcloudSharedprocessorpoolsPostInternalServerError() *PcloudSharedprocessorpoolsPostInternalServerError {
	return &PcloudSharedprocessorpoolsPostInternalServerError{}
}

/*
PcloudSharedprocessorpoolsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSharedprocessorpoolsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sharedprocessorpools post internal server error response has a 2xx status code
func (o *PcloudSharedprocessorpoolsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sharedprocessorpools post internal server error response has a 3xx status code
func (o *PcloudSharedprocessorpoolsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sharedprocessorpools post internal server error response has a 4xx status code
func (o *PcloudSharedprocessorpoolsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sharedprocessorpools post internal server error response has a 5xx status code
func (o *PcloudSharedprocessorpoolsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud sharedprocessorpools post internal server error response a status code equal to that given
func (o *PcloudSharedprocessorpoolsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud sharedprocessorpools post internal server error response
func (o *PcloudSharedprocessorpoolsPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudSharedprocessorpoolsPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostInternalServerError %s", 500, payload)
}

func (o *PcloudSharedprocessorpoolsPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools][%d] pcloudSharedprocessorpoolsPostInternalServerError %s", 500, payload)
}

func (o *PcloudSharedprocessorpoolsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
