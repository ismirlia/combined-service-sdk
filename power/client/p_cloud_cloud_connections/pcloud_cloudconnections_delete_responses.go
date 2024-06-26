// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

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

// PcloudCloudconnectionsDeleteReader is a Reader for the PcloudCloudconnectionsDelete structure.
type PcloudCloudconnectionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudconnectionsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsDeleteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudconnectionsDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}] pcloud.cloudconnections.delete", response, response.Code())
	}
}

// NewPcloudCloudconnectionsDeleteOK creates a PcloudCloudconnectionsDeleteOK with default headers values
func NewPcloudCloudconnectionsDeleteOK() *PcloudCloudconnectionsDeleteOK {
	return &PcloudCloudconnectionsDeleteOK{}
}

/*
PcloudCloudconnectionsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudconnections delete o k response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections delete o k response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete o k response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections delete o k response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections delete o k response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudconnections delete o k response
func (o *PcloudCloudconnectionsDeleteOK) Code() int {
	return 200
}

func (o *PcloudCloudconnectionsDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteOK %s", 200, payload)
}

func (o *PcloudCloudconnectionsDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteOK %s", 200, payload)
}

func (o *PcloudCloudconnectionsDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteAccepted creates a PcloudCloudconnectionsDeleteAccepted with default headers values
func NewPcloudCloudconnectionsDeleteAccepted() *PcloudCloudconnectionsDeleteAccepted {
	return &PcloudCloudconnectionsDeleteAccepted{}
}

/*
PcloudCloudconnectionsDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsDeleteAccepted struct {
	Payload *models.JobReference
}

// IsSuccess returns true when this pcloud cloudconnections delete accepted response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections delete accepted response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete accepted response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections delete accepted response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections delete accepted response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud cloudconnections delete accepted response
func (o *PcloudCloudconnectionsDeleteAccepted) Code() int {
	return 202
}

func (o *PcloudCloudconnectionsDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteAccepted %s", 202, payload)
}

func (o *PcloudCloudconnectionsDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteAccepted %s", 202, payload)
}

func (o *PcloudCloudconnectionsDeleteAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteBadRequest creates a PcloudCloudconnectionsDeleteBadRequest with default headers values
func NewPcloudCloudconnectionsDeleteBadRequest() *PcloudCloudconnectionsDeleteBadRequest {
	return &PcloudCloudconnectionsDeleteBadRequest{}
}

/*
PcloudCloudconnectionsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections delete bad request response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections delete bad request response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete bad request response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections delete bad request response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections delete bad request response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudconnections delete bad request response
func (o *PcloudCloudconnectionsDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudconnectionsDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteBadRequest %s", 400, payload)
}

func (o *PcloudCloudconnectionsDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteBadRequest %s", 400, payload)
}

func (o *PcloudCloudconnectionsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteUnauthorized creates a PcloudCloudconnectionsDeleteUnauthorized with default headers values
func NewPcloudCloudconnectionsDeleteUnauthorized() *PcloudCloudconnectionsDeleteUnauthorized {
	return &PcloudCloudconnectionsDeleteUnauthorized{}
}

/*
PcloudCloudconnectionsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections delete unauthorized response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections delete unauthorized response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete unauthorized response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections delete unauthorized response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections delete unauthorized response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudconnections delete unauthorized response
func (o *PcloudCloudconnectionsDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudconnectionsDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudCloudconnectionsDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudCloudconnectionsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteForbidden creates a PcloudCloudconnectionsDeleteForbidden with default headers values
func NewPcloudCloudconnectionsDeleteForbidden() *PcloudCloudconnectionsDeleteForbidden {
	return &PcloudCloudconnectionsDeleteForbidden{}
}

/*
PcloudCloudconnectionsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudconnectionsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections delete forbidden response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections delete forbidden response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete forbidden response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections delete forbidden response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections delete forbidden response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudconnections delete forbidden response
func (o *PcloudCloudconnectionsDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudCloudconnectionsDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteForbidden %s", 403, payload)
}

func (o *PcloudCloudconnectionsDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteForbidden %s", 403, payload)
}

func (o *PcloudCloudconnectionsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteNotFound creates a PcloudCloudconnectionsDeleteNotFound with default headers values
func NewPcloudCloudconnectionsDeleteNotFound() *PcloudCloudconnectionsDeleteNotFound {
	return &PcloudCloudconnectionsDeleteNotFound{}
}

/*
PcloudCloudconnectionsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections delete not found response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections delete not found response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete not found response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections delete not found response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections delete not found response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudconnections delete not found response
func (o *PcloudCloudconnectionsDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudCloudconnectionsDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteNotFound %s", 404, payload)
}

func (o *PcloudCloudconnectionsDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteNotFound %s", 404, payload)
}

func (o *PcloudCloudconnectionsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteRequestTimeout creates a PcloudCloudconnectionsDeleteRequestTimeout with default headers values
func NewPcloudCloudconnectionsDeleteRequestTimeout() *PcloudCloudconnectionsDeleteRequestTimeout {
	return &PcloudCloudconnectionsDeleteRequestTimeout{}
}

/*
PcloudCloudconnectionsDeleteRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsDeleteRequestTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections delete request timeout response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections delete request timeout response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete request timeout response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections delete request timeout response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections delete request timeout response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the pcloud cloudconnections delete request timeout response
func (o *PcloudCloudconnectionsDeleteRequestTimeout) Code() int {
	return 408
}

func (o *PcloudCloudconnectionsDeleteRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteRequestTimeout %s", 408, payload)
}

func (o *PcloudCloudconnectionsDeleteRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteRequestTimeout %s", 408, payload)
}

func (o *PcloudCloudconnectionsDeleteRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteGone creates a PcloudCloudconnectionsDeleteGone with default headers values
func NewPcloudCloudconnectionsDeleteGone() *PcloudCloudconnectionsDeleteGone {
	return &PcloudCloudconnectionsDeleteGone{}
}

/*
PcloudCloudconnectionsDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudconnectionsDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections delete gone response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections delete gone response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete gone response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections delete gone response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections delete gone response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the pcloud cloudconnections delete gone response
func (o *PcloudCloudconnectionsDeleteGone) Code() int {
	return 410
}

func (o *PcloudCloudconnectionsDeleteGone) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteGone %s", 410, payload)
}

func (o *PcloudCloudconnectionsDeleteGone) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteGone %s", 410, payload)
}

func (o *PcloudCloudconnectionsDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteInternalServerError creates a PcloudCloudconnectionsDeleteInternalServerError with default headers values
func NewPcloudCloudconnectionsDeleteInternalServerError() *PcloudCloudconnectionsDeleteInternalServerError {
	return &PcloudCloudconnectionsDeleteInternalServerError{}
}

/*
PcloudCloudconnectionsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections delete internal server error response has a 2xx status code
func (o *PcloudCloudconnectionsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections delete internal server error response has a 3xx status code
func (o *PcloudCloudconnectionsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections delete internal server error response has a 4xx status code
func (o *PcloudCloudconnectionsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections delete internal server error response has a 5xx status code
func (o *PcloudCloudconnectionsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections delete internal server error response a status code equal to that given
func (o *PcloudCloudconnectionsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudconnections delete internal server error response
func (o *PcloudCloudconnectionsDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudconnectionsDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudCloudconnectionsDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudCloudconnectionsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
