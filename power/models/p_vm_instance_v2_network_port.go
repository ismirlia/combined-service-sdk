// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstanceV2NetworkPort PVM's Port information
//
// swagger:model PVMInstanceV2NetworkPort
type PVMInstanceV2NetworkPort struct {

	// Unique Port ID
	ID string `json:"id,omitempty"`

	// Dynamic Host Configuration Protocol {IPv4, IPv6}
	// Enum: ["IPv4","IPv6"]
	IPProtocol string `json:"ipProtocol,omitempty"`

	// The mac address of the network interface
	MacAddress string `json:"macAddress,omitempty"`

	// The private ip address
	PrivateIP string `json:"privateIP,omitempty"`

	// The type of ip allocation {dhcp, static}
	// Enum: ["dhcp","static"]
	Type string `json:"type,omitempty"`
}

// Validate validates this p VM instance v2 network port
func (m *PVMInstanceV2NetworkPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pVmInstanceV2NetworkPortTypeIPProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceV2NetworkPortTypeIPProtocolPropEnum = append(pVmInstanceV2NetworkPortTypeIPProtocolPropEnum, v)
	}
}

const (

	// PVMInstanceV2NetworkPortIPProtocolIPV4 captures enum value "IPv4"
	PVMInstanceV2NetworkPortIPProtocolIPV4 string = "IPv4"

	// PVMInstanceV2NetworkPortIPProtocolIPV6 captures enum value "IPv6"
	PVMInstanceV2NetworkPortIPProtocolIPV6 string = "IPv6"
)

// prop value enum
func (m *PVMInstanceV2NetworkPort) validateIPProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceV2NetworkPortTypeIPProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceV2NetworkPort) validateIPProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.IPProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPProtocolEnum("ipProtocol", "body", m.IPProtocol); err != nil {
		return err
	}

	return nil
}

var pVmInstanceV2NetworkPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dhcp","static"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceV2NetworkPortTypeTypePropEnum = append(pVmInstanceV2NetworkPortTypeTypePropEnum, v)
	}
}

const (

	// PVMInstanceV2NetworkPortTypeDhcp captures enum value "dhcp"
	PVMInstanceV2NetworkPortTypeDhcp string = "dhcp"

	// PVMInstanceV2NetworkPortTypeStatic captures enum value "static"
	PVMInstanceV2NetworkPortTypeStatic string = "static"
)

// prop value enum
func (m *PVMInstanceV2NetworkPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceV2NetworkPortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceV2NetworkPort) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p VM instance v2 network port based on context it is used
func (m *PVMInstanceV2NetworkPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceV2NetworkPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceV2NetworkPort) UnmarshalBinary(b []byte) error {
	var res PVMInstanceV2NetworkPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
