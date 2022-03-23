// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_shared_proc_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudSharedprocpoolsPostReader is a Reader for the PcloudSharedprocpoolsPost structure.
type PcloudSharedprocpoolsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSharedprocpoolsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSharedprocpoolsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSharedprocpoolsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudSharedprocpoolsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudSharedprocpoolsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSharedprocpoolsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudSharedprocpoolsPostOK creates a PcloudSharedprocpoolsPostOK with default headers values
func NewPcloudSharedprocpoolsPostOK() *PcloudSharedprocpoolsPostOK {
	return &PcloudSharedprocpoolsPostOK{}
}

/* PcloudSharedprocpoolsPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSharedprocpoolsPostOK struct {
	Payload *models.SharedProcPool
}

func (o *PcloudSharedprocpoolsPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-proc-pools][%d] pcloudSharedprocpoolsPostOK  %+v", 200, o.Payload)
}
func (o *PcloudSharedprocpoolsPostOK) GetPayload() *models.SharedProcPool {
	return o.Payload
}

func (o *PcloudSharedprocpoolsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SharedProcPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocpoolsPostBadRequest creates a PcloudSharedprocpoolsPostBadRequest with default headers values
func NewPcloudSharedprocpoolsPostBadRequest() *PcloudSharedprocpoolsPostBadRequest {
	return &PcloudSharedprocpoolsPostBadRequest{}
}

/* PcloudSharedprocpoolsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSharedprocpoolsPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudSharedprocpoolsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-proc-pools][%d] pcloudSharedprocpoolsPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudSharedprocpoolsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocpoolsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocpoolsPostConflict creates a PcloudSharedprocpoolsPostConflict with default headers values
func NewPcloudSharedprocpoolsPostConflict() *PcloudSharedprocpoolsPostConflict {
	return &PcloudSharedprocpoolsPostConflict{}
}

/* PcloudSharedprocpoolsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudSharedprocpoolsPostConflict struct {
	Payload *models.Error
}

func (o *PcloudSharedprocpoolsPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-proc-pools][%d] pcloudSharedprocpoolsPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudSharedprocpoolsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocpoolsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocpoolsPostUnprocessableEntity creates a PcloudSharedprocpoolsPostUnprocessableEntity with default headers values
func NewPcloudSharedprocpoolsPostUnprocessableEntity() *PcloudSharedprocpoolsPostUnprocessableEntity {
	return &PcloudSharedprocpoolsPostUnprocessableEntity{}
}

/* PcloudSharedprocpoolsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudSharedprocpoolsPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudSharedprocpoolsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-proc-pools][%d] pcloudSharedprocpoolsPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudSharedprocpoolsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocpoolsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocpoolsPostInternalServerError creates a PcloudSharedprocpoolsPostInternalServerError with default headers values
func NewPcloudSharedprocpoolsPostInternalServerError() *PcloudSharedprocpoolsPostInternalServerError {
	return &PcloudSharedprocpoolsPostInternalServerError{}
}

/* PcloudSharedprocpoolsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSharedprocpoolsPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudSharedprocpoolsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-proc-pools][%d] pcloudSharedprocpoolsPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudSharedprocpoolsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocpoolsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}