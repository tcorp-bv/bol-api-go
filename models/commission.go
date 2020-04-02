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

// Commission commission
//
// swagger:model Commission
type Commission struct {

	// The condition of the offer.
	Condition string `json:"condition,omitempty"`

	// The EAN number associated with this product.
	Ean string `json:"ean,omitempty"`

	// Fixed fee.
	FixedAmount float64 `json:"fixedAmount,omitempty"`

	// A percentage of the offer price. The percentage can vary per product category.
	Percentage float64 `json:"percentage,omitempty"`

	// The price for this product with two decimals precision. The price includes VAT.
	Price float64 `json:"price,omitempty"`

	// reductions
	// Required: true
	Reductions []*Reduction `json:"reductions"`

	// Total applicable fee calculated based on the offer price provided.
	TotalCost float64 `json:"totalCost,omitempty"`

	// Total applicable fee calculated based on the offer price if you do not meet the maximum price criteria.
	TotalCostWithoutReduction float64 `json:"totalCostWithoutReduction,omitempty"`
}

// Validate validates this commission
func (m *Commission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReductions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Commission) validateReductions(formats strfmt.Registry) error {

	if err := validate.Required("reductions", "body", m.Reductions); err != nil {
		return err
	}

	for i := 0; i < len(m.Reductions); i++ {
		if swag.IsZero(m.Reductions[i]) { // not required
			continue
		}

		if m.Reductions[i] != nil {
			if err := m.Reductions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reductions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Commission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Commission) UnmarshalBinary(b []byte) error {
	var res Commission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
