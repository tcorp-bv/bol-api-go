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

// ShipmentItem shipment item
//
// swagger:model ShipmentItem
type ShipmentItem struct {

	// The EAN number associated with this product.
	Ean string `json:"ean,omitempty"`

	// Specifies whether this shipment has been fulfilled by the retailer (FBR) or fulfilled by bol.com (FBB). Defaults to FBR.
	// Enum: [FBR FBB]
	FulfilmentMethod string `json:"fulfilmentMethod,omitempty"`

	// The date and time in ISO 8601 format when the order was promised to be delivered.
	// Format: date-time
	LatestDeliveryDate strfmt.DateTime `json:"latestDeliveryDate,omitempty"`

	// Condition of the offer.
	// Enum: [NEW AS_NEW GOOD REASONABLE MODERATE]
	OfferCondition string `json:"offerCondition,omitempty"`

	// The total price for this order item id (item price multiplied by the quantity).
	OfferPrice float64 `json:"offerPrice,omitempty"`

	// The reference provided by the retailer through the offer API (referred to as ‘referenceCode’).
	OfferReference string `json:"offerReference,omitempty"`

	// The date and time in ISO 8601 format when the order was placed.
	// Format: date-time
	OrderDate strfmt.DateTime `json:"orderDate,omitempty"`

	// A unique identifier for the order this shipment is related to.
	OrderID string `json:"orderId,omitempty"`

	// A unique identifier for the item of the order that was shipped in this shipment.
	OrderItemID string `json:"orderItemId,omitempty"`

	// Amount of ordered products for this order item id.
	Quantity int32 `json:"quantity,omitempty"`

	// The product title.
	Title string `json:"title,omitempty"`
}

// Validate validates this shipment item
func (m *ShipmentItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFulfilmentMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var shipmentItemTypeFulfilmentMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FBR","FBB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shipmentItemTypeFulfilmentMethodPropEnum = append(shipmentItemTypeFulfilmentMethodPropEnum, v)
	}
}

const (

	// ShipmentItemFulfilmentMethodFBR captures enum value "FBR"
	ShipmentItemFulfilmentMethodFBR string = "FBR"

	// ShipmentItemFulfilmentMethodFBB captures enum value "FBB"
	ShipmentItemFulfilmentMethodFBB string = "FBB"
)

// prop value enum
func (m *ShipmentItem) validateFulfilmentMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, shipmentItemTypeFulfilmentMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ShipmentItem) validateFulfilmentMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.FulfilmentMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateFulfilmentMethodEnum("fulfilmentMethod", "body", m.FulfilmentMethod); err != nil {
		return err
	}

	return nil
}

func (m *ShipmentItem) validateLatestDeliveryDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LatestDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("latestDeliveryDate", "body", "date-time", m.LatestDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var shipmentItemTypeOfferConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEW","AS_NEW","GOOD","REASONABLE","MODERATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shipmentItemTypeOfferConditionPropEnum = append(shipmentItemTypeOfferConditionPropEnum, v)
	}
}

const (

	// ShipmentItemOfferConditionNEW captures enum value "NEW"
	ShipmentItemOfferConditionNEW string = "NEW"

	// ShipmentItemOfferConditionASNEW captures enum value "AS_NEW"
	ShipmentItemOfferConditionASNEW string = "AS_NEW"

	// ShipmentItemOfferConditionGOOD captures enum value "GOOD"
	ShipmentItemOfferConditionGOOD string = "GOOD"

	// ShipmentItemOfferConditionREASONABLE captures enum value "REASONABLE"
	ShipmentItemOfferConditionREASONABLE string = "REASONABLE"

	// ShipmentItemOfferConditionMODERATE captures enum value "MODERATE"
	ShipmentItemOfferConditionMODERATE string = "MODERATE"
)

// prop value enum
func (m *ShipmentItem) validateOfferConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, shipmentItemTypeOfferConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ShipmentItem) validateOfferCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.OfferCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateOfferConditionEnum("offerCondition", "body", m.OfferCondition); err != nil {
		return err
	}

	return nil
}

func (m *ShipmentItem) validateOrderDate(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderDate) { // not required
		return nil
	}

	if err := validate.FormatOf("orderDate", "body", "date-time", m.OrderDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShipmentItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipmentItem) UnmarshalBinary(b []byte) error {
	var res ShipmentItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
