// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_p_p_placement_groups

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

// PcloudSppplacementgroupsGetReader is a Reader for the PcloudSppplacementgroupsGet structure.
type PcloudSppplacementgroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSppplacementgroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSppplacementgroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSppplacementgroupsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSppplacementgroupsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudSppplacementgroupsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudSppplacementgroupsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSppplacementgroupsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}] pcloud.sppplacementgroups.get", response, response.Code())
	}
}

// NewPcloudSppplacementgroupsGetOK creates a PcloudSppplacementgroupsGetOK with default headers values
func NewPcloudSppplacementgroupsGetOK() *PcloudSppplacementgroupsGetOK {
	return &PcloudSppplacementgroupsGetOK{}
}

/*
PcloudSppplacementgroupsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSppplacementgroupsGetOK struct {
	Payload *models.SPPPlacementGroup
}

// IsSuccess returns true when this pcloud sppplacementgroups get o k response has a 2xx status code
func (o *PcloudSppplacementgroupsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sppplacementgroups get o k response has a 3xx status code
func (o *PcloudSppplacementgroupsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups get o k response has a 4xx status code
func (o *PcloudSppplacementgroupsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sppplacementgroups get o k response has a 5xx status code
func (o *PcloudSppplacementgroupsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups get o k response a status code equal to that given
func (o *PcloudSppplacementgroupsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud sppplacementgroups get o k response
func (o *PcloudSppplacementgroupsGetOK) Code() int {
	return 200
}

func (o *PcloudSppplacementgroupsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetOK %s", 200, payload)
}

func (o *PcloudSppplacementgroupsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetOK %s", 200, payload)
}

func (o *PcloudSppplacementgroupsGetOK) GetPayload() *models.SPPPlacementGroup {
	return o.Payload
}

func (o *PcloudSppplacementgroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SPPPlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsGetBadRequest creates a PcloudSppplacementgroupsGetBadRequest with default headers values
func NewPcloudSppplacementgroupsGetBadRequest() *PcloudSppplacementgroupsGetBadRequest {
	return &PcloudSppplacementgroupsGetBadRequest{}
}

/*
PcloudSppplacementgroupsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSppplacementgroupsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups get bad request response has a 2xx status code
func (o *PcloudSppplacementgroupsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups get bad request response has a 3xx status code
func (o *PcloudSppplacementgroupsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups get bad request response has a 4xx status code
func (o *PcloudSppplacementgroupsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups get bad request response has a 5xx status code
func (o *PcloudSppplacementgroupsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups get bad request response a status code equal to that given
func (o *PcloudSppplacementgroupsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud sppplacementgroups get bad request response
func (o *PcloudSppplacementgroupsGetBadRequest) Code() int {
	return 400
}

func (o *PcloudSppplacementgroupsGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetBadRequest %s", 400, payload)
}

func (o *PcloudSppplacementgroupsGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetBadRequest %s", 400, payload)
}

func (o *PcloudSppplacementgroupsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsGetUnauthorized creates a PcloudSppplacementgroupsGetUnauthorized with default headers values
func NewPcloudSppplacementgroupsGetUnauthorized() *PcloudSppplacementgroupsGetUnauthorized {
	return &PcloudSppplacementgroupsGetUnauthorized{}
}

/*
PcloudSppplacementgroupsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSppplacementgroupsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups get unauthorized response has a 2xx status code
func (o *PcloudSppplacementgroupsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups get unauthorized response has a 3xx status code
func (o *PcloudSppplacementgroupsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups get unauthorized response has a 4xx status code
func (o *PcloudSppplacementgroupsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups get unauthorized response has a 5xx status code
func (o *PcloudSppplacementgroupsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups get unauthorized response a status code equal to that given
func (o *PcloudSppplacementgroupsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud sppplacementgroups get unauthorized response
func (o *PcloudSppplacementgroupsGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudSppplacementgroupsGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetUnauthorized %s", 401, payload)
}

func (o *PcloudSppplacementgroupsGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetUnauthorized %s", 401, payload)
}

func (o *PcloudSppplacementgroupsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsGetForbidden creates a PcloudSppplacementgroupsGetForbidden with default headers values
func NewPcloudSppplacementgroupsGetForbidden() *PcloudSppplacementgroupsGetForbidden {
	return &PcloudSppplacementgroupsGetForbidden{}
}

/*
PcloudSppplacementgroupsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudSppplacementgroupsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups get forbidden response has a 2xx status code
func (o *PcloudSppplacementgroupsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups get forbidden response has a 3xx status code
func (o *PcloudSppplacementgroupsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups get forbidden response has a 4xx status code
func (o *PcloudSppplacementgroupsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups get forbidden response has a 5xx status code
func (o *PcloudSppplacementgroupsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups get forbidden response a status code equal to that given
func (o *PcloudSppplacementgroupsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud sppplacementgroups get forbidden response
func (o *PcloudSppplacementgroupsGetForbidden) Code() int {
	return 403
}

func (o *PcloudSppplacementgroupsGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetForbidden %s", 403, payload)
}

func (o *PcloudSppplacementgroupsGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetForbidden %s", 403, payload)
}

func (o *PcloudSppplacementgroupsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsGetNotFound creates a PcloudSppplacementgroupsGetNotFound with default headers values
func NewPcloudSppplacementgroupsGetNotFound() *PcloudSppplacementgroupsGetNotFound {
	return &PcloudSppplacementgroupsGetNotFound{}
}

/*
PcloudSppplacementgroupsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudSppplacementgroupsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups get not found response has a 2xx status code
func (o *PcloudSppplacementgroupsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups get not found response has a 3xx status code
func (o *PcloudSppplacementgroupsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups get not found response has a 4xx status code
func (o *PcloudSppplacementgroupsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sppplacementgroups get not found response has a 5xx status code
func (o *PcloudSppplacementgroupsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sppplacementgroups get not found response a status code equal to that given
func (o *PcloudSppplacementgroupsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud sppplacementgroups get not found response
func (o *PcloudSppplacementgroupsGetNotFound) Code() int {
	return 404
}

func (o *PcloudSppplacementgroupsGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetNotFound %s", 404, payload)
}

func (o *PcloudSppplacementgroupsGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetNotFound %s", 404, payload)
}

func (o *PcloudSppplacementgroupsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSppplacementgroupsGetInternalServerError creates a PcloudSppplacementgroupsGetInternalServerError with default headers values
func NewPcloudSppplacementgroupsGetInternalServerError() *PcloudSppplacementgroupsGetInternalServerError {
	return &PcloudSppplacementgroupsGetInternalServerError{}
}

/*
PcloudSppplacementgroupsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSppplacementgroupsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sppplacementgroups get internal server error response has a 2xx status code
func (o *PcloudSppplacementgroupsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sppplacementgroups get internal server error response has a 3xx status code
func (o *PcloudSppplacementgroupsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sppplacementgroups get internal server error response has a 4xx status code
func (o *PcloudSppplacementgroupsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sppplacementgroups get internal server error response has a 5xx status code
func (o *PcloudSppplacementgroupsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud sppplacementgroups get internal server error response a status code equal to that given
func (o *PcloudSppplacementgroupsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud sppplacementgroups get internal server error response
func (o *PcloudSppplacementgroupsGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudSppplacementgroupsGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetInternalServerError %s", 500, payload)
}

func (o *PcloudSppplacementgroupsGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/spp-placement-groups/{spp_placement_group_id}][%d] pcloudSppplacementgroupsGetInternalServerError %s", 500, payload)
}

func (o *PcloudSppplacementgroupsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSppplacementgroupsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
