// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Problem Describes a problem that occurred interacting with the API.
//
// swagger:model Problem
type Problem struct {

	// Detailed error message describing in additional detail what caused the service to return this problem.
	Detail string `json:"detail,omitempty"`

	// Host identifier describing the server instance that reported the problem.
	Host string `json:"host,omitempty"`

	// Full URI path of the resource that reported the problem.
	// Format: uri
	Instance strfmt.URI `json:"instance,omitempty"`

	// HTTP status returned from the endpoint causing the problem.
	Status int32 `json:"status,omitempty"`

	// Title describing the nature of the problem.
	Title string `json:"title,omitempty"`

	// Type URI for this problem. Fixed value: https://api.bol.com/problems.
	// Format: uri
	Type strfmt.URI `json:"type,omitempty"`

	// violations
	Violations []*Violation `json:"violations"`
}

// Validate validates this problem
func (m *Problem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViolations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Problem) validateInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.Instance) { // not required
		return nil
	}

	if err := validate.FormatOf("instance", "body", "uri", m.Instance.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Problem) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.FormatOf("type", "body", "uri", m.Type.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Problem) validateViolations(formats strfmt.Registry) error {

	if swag.IsZero(m.Violations) { // not required
		return nil
	}

	for i := 0; i < len(m.Violations); i++ {
		if swag.IsZero(m.Violations[i]) { // not required
			continue
		}

		if m.Violations[i] != nil {
			if err := m.Violations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("violations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Problem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Problem) UnmarshalBinary(b []byte) error {
	var res Problem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
