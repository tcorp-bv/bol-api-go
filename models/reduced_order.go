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

// ReducedOrder An order.
//
// swagger:model ReducedOrder
type ReducedOrder struct {

	// The date and time in ISO 8601 format when the order was placed.
	// Format: date-time
	DateTimeOrderPlaced strfmt.DateTime `json:"dateTimeOrderPlaced,omitempty"`

	// The order id.
	OrderID string `json:"orderId,omitempty"`

	// order items
	// Required: true
	OrderItems []*ReducedOrderItem `json:"orderItems"`
}

// Validate validates this reduced order
func (m *ReducedOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTimeOrderPlaced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReducedOrder) validateDateTimeOrderPlaced(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeOrderPlaced) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTimeOrderPlaced", "body", "date-time", m.DateTimeOrderPlaced.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReducedOrder) validateOrderItems(formats strfmt.Registry) error {

	if err := validate.Required("orderItems", "body", m.OrderItems); err != nil {
		return err
	}

	for i := 0; i < len(m.OrderItems); i++ {
		if swag.IsZero(m.OrderItems[i]) { // not required
			continue
		}

		if m.OrderItems[i] != nil {
			if err := m.OrderItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReducedOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReducedOrder) UnmarshalBinary(b []byte) error {
	var res ReducedOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
