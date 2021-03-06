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

// Condition condition
//
// swagger:model Condition
type Condition struct {

	// The category of the condition. If not given NEW or SECONDHAND is derived from NAME.
	// Enum: [NEW SECONDHAND]
	Category string `json:"category,omitempty"`

	// The description of the condition of the product. Only allowed if name is not NEW and may not contain e-mail addresses.
	// Max Length: 2000
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// The condition of the offered product.
	// Required: true
	// Enum: [NEW AS_NEW GOOD REASONABLE MODERATE]
	Name *string `json:"name"`
}

// Validate validates this condition
func (m *Condition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conditionTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEW","SECONDHAND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeCategoryPropEnum = append(conditionTypeCategoryPropEnum, v)
	}
}

const (

	// ConditionCategoryNEW captures enum value "NEW"
	ConditionCategoryNEW string = "NEW"

	// ConditionCategorySECONDHAND captures enum value "SECONDHAND"
	ConditionCategorySECONDHAND string = "SECONDHAND"
)

// prop value enum
func (m *Condition) validateCategoryEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, conditionTypeCategoryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Condition) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *Condition) validateComment(formats strfmt.Registry) error {

	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", string(*m.Comment), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", string(*m.Comment), 2000); err != nil {
		return err
	}

	return nil
}

var conditionTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEW","AS_NEW","GOOD","REASONABLE","MODERATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeNamePropEnum = append(conditionTypeNamePropEnum, v)
	}
}

const (

	// ConditionNameNEW captures enum value "NEW"
	ConditionNameNEW string = "NEW"

	// ConditionNameASNEW captures enum value "AS_NEW"
	ConditionNameASNEW string = "AS_NEW"

	// ConditionNameGOOD captures enum value "GOOD"
	ConditionNameGOOD string = "GOOD"

	// ConditionNameREASONABLE captures enum value "REASONABLE"
	ConditionNameREASONABLE string = "REASONABLE"

	// ConditionNameMODERATE captures enum value "MODERATE"
	ConditionNameMODERATE string = "MODERATE"
)

// prop value enum
func (m *Condition) validateNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, conditionTypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Condition) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Condition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Condition) UnmarshalBinary(b []byte) error {
	var res Condition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
