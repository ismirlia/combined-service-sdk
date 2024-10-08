// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CapabilitiesDetails capabilities details
//
// swagger:model CapabilitiesDetails
type CapabilitiesDetails struct {

	// Disaster Recovery Information
	// Required: true
	DisasterRecovery *DisasterRecovery `json:"disasterRecovery"`

	// Datacenter System Types Information
	// Required: true
	SupportedSystems *SupportedSystems `json:"supportedSystems"`
}

// Validate validates this capabilities details
func (m *CapabilitiesDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisasterRecovery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedSystems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapabilitiesDetails) validateDisasterRecovery(formats strfmt.Registry) error {

	if err := validate.Required("disasterRecovery", "body", m.DisasterRecovery); err != nil {
		return err
	}

	if m.DisasterRecovery != nil {
		if err := m.DisasterRecovery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disasterRecovery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disasterRecovery")
			}
			return err
		}
	}

	return nil
}

func (m *CapabilitiesDetails) validateSupportedSystems(formats strfmt.Registry) error {

	if err := validate.Required("supportedSystems", "body", m.SupportedSystems); err != nil {
		return err
	}

	if m.SupportedSystems != nil {
		if err := m.SupportedSystems.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supportedSystems")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("supportedSystems")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this capabilities details based on the context it is used
func (m *CapabilitiesDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisasterRecovery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportedSystems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapabilitiesDetails) contextValidateDisasterRecovery(ctx context.Context, formats strfmt.Registry) error {

	if m.DisasterRecovery != nil {

		if err := m.DisasterRecovery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disasterRecovery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disasterRecovery")
			}
			return err
		}
	}

	return nil
}

func (m *CapabilitiesDetails) contextValidateSupportedSystems(ctx context.Context, formats strfmt.Registry) error {

	if m.SupportedSystems != nil {

		if err := m.SupportedSystems.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supportedSystems")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("supportedSystems")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapabilitiesDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapabilitiesDetails) UnmarshalBinary(b []byte) error {
	var res CapabilitiesDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
