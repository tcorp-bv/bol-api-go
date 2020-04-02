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

// PurchasableShippingLabelsResponse purchasable shipping labels response
//
// swagger:model PurchasableShippingLabelsResponse
type PurchasableShippingLabelsResponse struct {

	// purchasable shipping labels
	// Required: true
	PurchasableShippingLabels []*PurchasableShippingLabel `json:"purchasableShippingLabels"`
}

// Validate validates this purchasable shipping labels response
func (m *PurchasableShippingLabelsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePurchasableShippingLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurchasableShippingLabelsResponse) validatePurchasableShippingLabels(formats strfmt.Registry) error {

	if err := validate.Required("purchasableShippingLabels", "body", m.PurchasableShippingLabels); err != nil {
		return err
	}

	for i := 0; i < len(m.PurchasableShippingLabels); i++ {
		if swag.IsZero(m.PurchasableShippingLabels[i]) { // not required
			continue
		}

		if m.PurchasableShippingLabels[i] != nil {
			if err := m.PurchasableShippingLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("purchasableShippingLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PurchasableShippingLabelsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PurchasableShippingLabelsResponse) UnmarshalBinary(b []byte) error {
	var res PurchasableShippingLabelsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
