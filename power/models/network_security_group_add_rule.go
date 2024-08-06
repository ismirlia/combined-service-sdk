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

// NetworkSecurityGroupAddRule network security group add rule
//
// swagger:model NetworkSecurityGroupAddRule
type NetworkSecurityGroupAddRule struct {

	// The action to take if the rule matches network traffic
	// Required: true
	// Enum: ["allow","deny"]
	Action *string `json:"action"`

	// destination ports
	DestinationPorts *NetworkSecurityGroupRulePort `json:"destinationPorts,omitempty"`

	// The direction of the network traffic
	// Required: true
	// Enum: ["inbound","outbound"]
	Direction *string `json:"direction"`

	// The unique name of the Network Security Group rule
	// Required: true
	Name *string `json:"name"`

	// protocol
	// Required: true
	Protocol *NetworkSecurityGroupRuleProtocol `json:"protocol"`

	// remote
	// Required: true
	Remote *NetworkSecurityGroupRuleRemote `json:"remote"`

	// source ports
	SourcePorts *NetworkSecurityGroupRulePort `json:"sourcePorts,omitempty"`
}

// Validate validates this network security group add rule
func (m *NetworkSecurityGroupAddRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkSecurityGroupAddRuleTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkSecurityGroupAddRuleTypeActionPropEnum = append(networkSecurityGroupAddRuleTypeActionPropEnum, v)
	}
}

const (

	// NetworkSecurityGroupAddRuleActionAllow captures enum value "allow"
	NetworkSecurityGroupAddRuleActionAllow string = "allow"

	// NetworkSecurityGroupAddRuleActionDeny captures enum value "deny"
	NetworkSecurityGroupAddRuleActionDeny string = "deny"
)

// prop value enum
func (m *NetworkSecurityGroupAddRule) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkSecurityGroupAddRuleTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkSecurityGroupAddRule) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSecurityGroupAddRule) validateDestinationPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationPorts) { // not required
		return nil
	}

	if m.DestinationPorts != nil {
		if err := m.DestinationPorts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationPorts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationPorts")
			}
			return err
		}
	}

	return nil
}

var networkSecurityGroupAddRuleTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkSecurityGroupAddRuleTypeDirectionPropEnum = append(networkSecurityGroupAddRuleTypeDirectionPropEnum, v)
	}
}

const (

	// NetworkSecurityGroupAddRuleDirectionInbound captures enum value "inbound"
	NetworkSecurityGroupAddRuleDirectionInbound string = "inbound"

	// NetworkSecurityGroupAddRuleDirectionOutbound captures enum value "outbound"
	NetworkSecurityGroupAddRuleDirectionOutbound string = "outbound"
)

// prop value enum
func (m *NetworkSecurityGroupAddRule) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkSecurityGroupAddRuleTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkSecurityGroupAddRule) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSecurityGroupAddRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSecurityGroupAddRule) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if m.Protocol != nil {
		if err := m.Protocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupAddRule) validateRemote(formats strfmt.Registry) error {

	if err := validate.Required("remote", "body", m.Remote); err != nil {
		return err
	}

	if m.Remote != nil {
		if err := m.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remote")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupAddRule) validateSourcePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.SourcePorts) { // not required
		return nil
	}

	if m.SourcePorts != nil {
		if err := m.SourcePorts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcePorts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourcePorts")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network security group add rule based on the context it is used
func (m *NetworkSecurityGroupAddRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourcePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSecurityGroupAddRule) contextValidateDestinationPorts(ctx context.Context, formats strfmt.Registry) error {

	if m.DestinationPorts != nil {

		if swag.IsZero(m.DestinationPorts) { // not required
			return nil
		}

		if err := m.DestinationPorts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationPorts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationPorts")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupAddRule) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.Protocol != nil {

		if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupAddRule) contextValidateRemote(ctx context.Context, formats strfmt.Registry) error {

	if m.Remote != nil {

		if err := m.Remote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remote")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkSecurityGroupAddRule) contextValidateSourcePorts(ctx context.Context, formats strfmt.Registry) error {

	if m.SourcePorts != nil {

		if swag.IsZero(m.SourcePorts) { // not required
			return nil
		}

		if err := m.SourcePorts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcePorts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourcePorts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSecurityGroupAddRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSecurityGroupAddRule) UnmarshalBinary(b []byte) error {
	var res NetworkSecurityGroupAddRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}