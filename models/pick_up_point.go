// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PickUpPoint pick up point
//
// swagger:model PickUpPoint
type PickUpPoint struct {

	// The code of the pickup point location in case you want to ship items to pickup points.
	// Enum: [PUP_AH_NL]
	Code string `json:"code,omitempty"`
}

// Validate validates this pick up point
func (m *PickUpPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pickUpPointTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUP_AH_NL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pickUpPointTypeCodePropEnum = append(pickUpPointTypeCodePropEnum, v)
	}
}

const (

	// PickUpPointCodePUPAHNL captures enum value "PUP_AH_NL"
	PickUpPointCodePUPAHNL string = "PUP_AH_NL"
)

// prop value enum
func (m *PickUpPoint) validateCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pickUpPointTypeCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PickUpPoint) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PickUpPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PickUpPoint) UnmarshalBinary(b []byte) error {
	var res PickUpPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
