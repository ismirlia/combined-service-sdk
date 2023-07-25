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

// NetworkReference network reference
//
// swagger:model NetworkReference
type NetworkReference struct {

	// Network communication configuration (for satellite locations only)
	//   * `internal-only` - network is only used for internal host communication
	//   * `outbound-only` - network will be capable of egress traffic
	//   * `bidirectional-static-route` - network will be capable of ingress and egress traffic via static routes
	//   * `bidirectional-bgp` - network will be capable of ingress and egress traffic via bgp configuration
	//   * `bidirectional-l2out` - network will be capable of ingress and egress traffic via l2out ACI configuration
	//
	// Enum: [internal-only outbound-only bidirectional-static-route bidirectional-bgp bidirectional-l2out]
	AccessConfig string `json:"accessConfig,omitempty"`

	// DHCP Managed Network
	DhcpManaged bool `json:"dhcpManaged,omitempty"`

	// Link to Network resource
	// Required: true
	Href *string `json:"href"`

	// Enable MTU Jumbo Network (for multi-zone locations only)
	Jumbo bool `json:"jumbo,omitempty"`

	// Maximum transmission unit (for satellite locations only)
	Mtu int64 `json:"mtu,omitempty"`

	// Network Name
	// Required: true
	Name *string `json:"name"`

	// Unique Network ID
	// Required: true
	NetworkID *string `json:"networkID"`

	// Type of Network {vlan, pub-vlan}
	// Required: true
	// Enum: [vlan pub-vlan dhcp-vlan]
	Type *string `json:"type"`

	// VLAN ID
	// Required: true
	VlanID *float64 `json:"vlanID"`
}

// Validate validates this network reference
func (m *NetworkReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkReferenceTypeAccessConfigPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internal-only","outbound-only","bidirectional-static-route","bidirectional-bgp","bidirectional-l2out"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkReferenceTypeAccessConfigPropEnum = append(networkReferenceTypeAccessConfigPropEnum, v)
	}
}

const (

	// NetworkReferenceAccessConfigInternalDashOnly captures enum value "internal-only"
	NetworkReferenceAccessConfigInternalDashOnly string = "internal-only"

	// NetworkReferenceAccessConfigOutboundDashOnly captures enum value "outbound-only"
	NetworkReferenceAccessConfigOutboundDashOnly string = "outbound-only"

	// NetworkReferenceAccessConfigBidirectionalDashStaticDashRoute captures enum value "bidirectional-static-route"
	NetworkReferenceAccessConfigBidirectionalDashStaticDashRoute string = "bidirectional-static-route"

	// NetworkReferenceAccessConfigBidirectionalDashBgp captures enum value "bidirectional-bgp"
	NetworkReferenceAccessConfigBidirectionalDashBgp string = "bidirectional-bgp"

	// NetworkReferenceAccessConfigBidirectionalDashL2out captures enum value "bidirectional-l2out"
	NetworkReferenceAccessConfigBidirectionalDashL2out string = "bidirectional-l2out"
)

// prop value enum
func (m *NetworkReference) validateAccessConfigEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkReferenceTypeAccessConfigPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkReference) validateAccessConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessConfig) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessConfigEnum("accessConfig", "body", m.AccessConfig); err != nil {
		return err
	}

	return nil
}

func (m *NetworkReference) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *NetworkReference) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NetworkReference) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("networkID", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

var networkReferenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vlan","pub-vlan","dhcp-vlan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkReferenceTypeTypePropEnum = append(networkReferenceTypeTypePropEnum, v)
	}
}

const (

	// NetworkReferenceTypeVlan captures enum value "vlan"
	NetworkReferenceTypeVlan string = "vlan"

	// NetworkReferenceTypePubDashVlan captures enum value "pub-vlan"
	NetworkReferenceTypePubDashVlan string = "pub-vlan"

	// NetworkReferenceTypeDhcpDashVlan captures enum value "dhcp-vlan"
	NetworkReferenceTypeDhcpDashVlan string = "dhcp-vlan"
)

// prop value enum
func (m *NetworkReference) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkReferenceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkReference) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *NetworkReference) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlanID", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network reference based on context it is used
func (m *NetworkReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkReference) UnmarshalBinary(b []byte) error {
	var res NetworkReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
