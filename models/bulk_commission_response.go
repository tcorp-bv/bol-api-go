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

// BulkCommissionResponse bulk commission response
//
// swagger:model BulkCommissionResponse
type BulkCommissionResponse struct {

	// commissions
	// Required: true
	Commissions []*Commission `json:"commissions"`
}

// Validate validates this bulk commission response
func (m *BulkCommissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkCommissionResponse) validateCommissions(formats strfmt.Registry) error {

	if err := validate.Required("commissions", "body", m.Commissions); err != nil {
		return err
	}

	for i := 0; i < len(m.Commissions); i++ {
		if swag.IsZero(m.Commissions[i]) { // not required
			continue
		}

		if m.Commissions[i] != nil {
			if err := m.Commissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkCommissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkCommissionResponse) UnmarshalBinary(b []byte) error {
	var res BulkCommissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
