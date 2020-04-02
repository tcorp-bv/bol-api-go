// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReducedShipmentItem reduced shipment item
//
// swagger:model ReducedShipmentItem
type ReducedShipmentItem struct {

	// A unique identifier for the order this shipment is related to.
	OrderID string `json:"orderId,omitempty"`

	// A unique identifier for the item of the order that was shipped in this shipment.
	OrderItemID string `json:"orderItemId,omitempty"`
}

// Validate validates this reduced shipment item
func (m *ReducedShipmentItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReducedShipmentItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReducedShipmentItem) UnmarshalBinary(b []byte) error {
	var res ReducedShipmentItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}