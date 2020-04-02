// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PurchasableShippingLabel purchasable shipping label
//
// swagger:model PurchasableShippingLabel
type PurchasableShippingLabel struct {

	// The discount of the item that has been sold.
	Discount float64 `json:"discount,omitempty"`

	// The type of the label, representing the way an item is being transported.
	LabelType string `json:"labelType,omitempty"`

	// The dimensions of a package.
	MaxDimensions string `json:"maxDimensions,omitempty"`

	// The weight of a package.
	MaxWeight string `json:"maxWeight,omitempty"`

	// The price that is charged to the partner for the shipping label, including VAT.
	PurchasePrice float64 `json:"purchasePrice,omitempty"`

	// The price the item has been sold for.
	RetailerPrice float64 `json:"retailerPrice,omitempty"`

	// A unique code referring to the used shipping label for this shipment.
	ShippingLabelCode string `json:"shippingLabelCode,omitempty"`

	// A code representing the transporter which is being used for transportation.
	TransporterCode string `json:"transporterCode,omitempty"`
}

// Validate validates this purchasable shipping label
func (m *PurchasableShippingLabel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PurchasableShippingLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PurchasableShippingLabel) UnmarshalBinary(b []byte) error {
	var res PurchasableShippingLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
