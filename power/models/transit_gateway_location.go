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

// TransitGatewayLocation The PER enabled PowerVS Service Location
//
// swagger:model TransitGatewayLocation
type TransitGatewayLocation struct {

	// The Location of the PowerVS Service
	// Example: dal12
	// Required: true
	Location *string `json:"location"`

	// Location Type of the PowerVS Service
	// Example: data-center
	// Required: true
	// Enum: ["region","data-center","zone"]
	LocationType *string `json:"locationType"`

	// The PowerVS Location URL path to access specific service instance information
	// Example: https://us-south.power-iaas.cloud.ibm.com
	// Required: true
	LocationURL *string `json:"locationUrl"`
}

// Validate validates this transit gateway location
func (m *TransitGatewayLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransitGatewayLocation) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	return nil
}

var transitGatewayLocationTypeLocationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["region","data-center","zone"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transitGatewayLocationTypeLocationTypePropEnum = append(transitGatewayLocationTypeLocationTypePropEnum, v)
	}
}

const (

	// TransitGatewayLocationLocationTypeRegion captures enum value "region"
	TransitGatewayLocationLocationTypeRegion string = "region"

	// TransitGatewayLocationLocationTypeDataDashCenter captures enum value "data-center"
	TransitGatewayLocationLocationTypeDataDashCenter string = "data-center"

	// TransitGatewayLocationLocationTypeZone captures enum value "zone"
	TransitGatewayLocationLocationTypeZone string = "zone"
)

// prop value enum
func (m *TransitGatewayLocation) validateLocationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transitGatewayLocationTypeLocationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransitGatewayLocation) validateLocationType(formats strfmt.Registry) error {

	if err := validate.Required("locationType", "body", m.LocationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateLocationTypeEnum("locationType", "body", *m.LocationType); err != nil {
		return err
	}

	return nil
}

func (m *TransitGatewayLocation) validateLocationURL(formats strfmt.Registry) error {

	if err := validate.Required("locationUrl", "body", m.LocationURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transit gateway location based on context it is used
func (m *TransitGatewayLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransitGatewayLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransitGatewayLocation) UnmarshalBinary(b []byte) error {
	var res TransitGatewayLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
