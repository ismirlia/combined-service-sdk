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

// RemoveVirtualSerialNumber remove virtual serial number
//
// swagger:model RemoveVirtualSerialNumber
type RemoveVirtualSerialNumber struct {

	// Provide a reserved Virtual Serial Number
	// Required: true
	Serial *string `json:"serial"`
}

// Validate validates this remove virtual serial number
func (m *RemoveVirtualSerialNumber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveVirtualSerialNumber) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", m.Serial); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this remove virtual serial number based on context it is used
func (m *RemoveVirtualSerialNumber) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoveVirtualSerialNumber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveVirtualSerialNumber) UnmarshalBinary(b []byte) error {
	var res RemoveVirtualSerialNumber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
