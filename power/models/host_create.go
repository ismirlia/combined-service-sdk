// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostCreate Parameters to add a host to an existing host group
//
// swagger:model HostCreate
type HostCreate struct {

	// ID of the host group to which the host should be added
	// Required: true
	HostGroupID *string `json:"hostGroupID"`

	// Hosts to be added
	// Required: true
	Hosts []*AddHost `json:"hosts"`
}

// Validate validates this host create
func (m *HostCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostCreate) validateHostGroupID(formats strfmt.Registry) error {

	if err := validate.Required("hostGroupID", "body", m.HostGroupID); err != nil {
		return err
	}

	return nil
}

func (m *HostCreate) validateHosts(formats strfmt.Registry) error {

	if err := validate.Required("hosts", "body", m.Hosts); err != nil {
		return err
	}

	for i := 0; i < len(m.Hosts); i++ {
		if swag.IsZero(m.Hosts[i]) { // not required
			continue
		}

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this host create based on the context it is used
func (m *HostCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostCreate) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hosts); i++ {

		if m.Hosts[i] != nil {

			if swag.IsZero(m.Hosts[i]) { // not required
				return nil
			}

			if err := m.Hosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostCreate) UnmarshalBinary(b []byte) error {
	var res HostCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
