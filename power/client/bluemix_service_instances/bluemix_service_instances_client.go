// Code generated by go-swagger; DO NOT EDIT.

package bluemix_service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new bluemix service instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new bluemix service instances API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new bluemix service instances API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for bluemix service instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BluemixServiceInstanceGet(params *BluemixServiceInstanceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BluemixServiceInstanceGetOK, error)

	BluemixServiceInstancePut(params *BluemixServiceInstancePutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BluemixServiceInstancePutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
BluemixServiceInstanceGet gets the current state information associated with the service instance
*/
func (a *Client) BluemixServiceInstanceGet(params *BluemixServiceInstanceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BluemixServiceInstanceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBluemixServiceInstanceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "bluemix.serviceInstance.get",
		Method:             "GET",
		PathPattern:        "/bluemix_v1/service_instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BluemixServiceInstanceGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BluemixServiceInstanceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bluemix.serviceInstance.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BluemixServiceInstancePut updates disable or enable the state of a provisioned service instance
*/
func (a *Client) BluemixServiceInstancePut(params *BluemixServiceInstancePutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BluemixServiceInstancePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBluemixServiceInstancePutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "bluemix.serviceInstance.put",
		Method:             "PUT",
		PathPattern:        "/bluemix_v1/service_instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BluemixServiceInstancePutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BluemixServiceInstancePutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bluemix.serviceInstance.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
