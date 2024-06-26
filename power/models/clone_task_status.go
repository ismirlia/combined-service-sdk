// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloneTaskStatus clone task status
//
// swagger:model CloneTaskStatus
type CloneTaskStatus struct {

	// List of cloned volumes created from the clone volumes task
	ClonedVolumes []*ClonedVolume `json:"clonedVolumes"`

	// The reason the clone volumes task has failed
	FailedReason string `json:"failedReason,omitempty"`

	// Snapshot completion percentage
	// Required: true
	PercentComplete *int64 `json:"percentComplete"`

	// Status of the clone volumes task
	// Required: true
	// Enum: ["running","completed","failed","unknown"]
	Status *string `json:"status"`
}

// Validate validates this clone task status
func (m *CloneTaskStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClonedVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentComplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloneTaskStatus) validateClonedVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.ClonedVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.ClonedVolumes); i++ {
		if swag.IsZero(m.ClonedVolumes[i]) { // not required
			continue
		}

		if m.ClonedVolumes[i] != nil {
			if err := m.ClonedVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clonedVolumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clonedVolumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloneTaskStatus) validatePercentComplete(formats strfmt.Registry) error {

	if err := validate.Required("percentComplete", "body", m.PercentComplete); err != nil {
		return err
	}

	return nil
}

var cloneTaskStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["running","completed","failed","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloneTaskStatusTypeStatusPropEnum = append(cloneTaskStatusTypeStatusPropEnum, v)
	}
}

const (

	// CloneTaskStatusStatusRunning captures enum value "running"
	CloneTaskStatusStatusRunning string = "running"

	// CloneTaskStatusStatusCompleted captures enum value "completed"
	CloneTaskStatusStatusCompleted string = "completed"

	// CloneTaskStatusStatusFailed captures enum value "failed"
	CloneTaskStatusStatusFailed string = "failed"

	// CloneTaskStatusStatusUnknown captures enum value "unknown"
	CloneTaskStatusStatusUnknown string = "unknown"
)

// prop value enum
func (m *CloneTaskStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloneTaskStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CloneTaskStatus) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this clone task status based on the context it is used
func (m *CloneTaskStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClonedVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloneTaskStatus) contextValidateClonedVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClonedVolumes); i++ {

		if m.ClonedVolumes[i] != nil {

			if swag.IsZero(m.ClonedVolumes[i]) { // not required
				return nil
			}

			if err := m.ClonedVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clonedVolumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clonedVolumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloneTaskStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloneTaskStatus) UnmarshalBinary(b []byte) error {
	var res CloneTaskStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
