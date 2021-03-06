// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Inbound inbound
//
// swagger:model Inbound
type Inbound struct {

	// The number of announced BSKU‘s.
	// Required: true
	AnnouncedBSKUs *int64 `json:"announcedBSKUs"`

	// The number of announced items.
	// Required: true
	AnnouncedQuantity *int64 `json:"announcedQuantity"`

	// The date the inbound shipment was created in ISO 8601 format.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// Transporter for the inbound shipment.
	// Required: true
	FbbTransporter *Transporter `json:"fbbTransporter"`

	// A unique identifier for an inbound shipment.
	// Required: true
	ID *int64 `json:"id"`

	// Indicates whether the inbound will be labeled by bol.com or not.
	// Required: true
	LabellingService *bool `json:"labellingService"`

	// List of products.
	// Required: true
	Products []*Product `json:"products"`

	// The number of received BSKU‘s.
	// Required: true
	ReceivedBSKUs *int64 `json:"receivedBSKUs"`

	// The number of received items.
	// Required: true
	ReceivedQuantity *int64 `json:"receivedQuantity"`

	// A user defined reference to identify the inbound shipment.
	// Required: true
	Reference *string `json:"reference"`

	// The current state of the inbound shipment.
	// Required: true
	// Enum: [Draft PreAnnounced ArrivedAtWH Cancelled]
	State *string `json:"state"`

	// List of state transitions.
	// Required: true
	StateTransitions []*StateTransition `json:"stateTransitions"`

	// The chosen timeslot for the inbound shipment.
	TimeSlot *TimeSlot `json:"timeSlot,omitempty"`
}

// Validate validates this inbound
func (m *Inbound) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnouncedBSKUs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnnouncedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFbbTransporter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabellingService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivedBSKUs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateTransitions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeSlot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Inbound) validateAnnouncedBSKUs(formats strfmt.Registry) error {

	if err := validate.Required("announcedBSKUs", "body", m.AnnouncedBSKUs); err != nil {
		return err
	}

	return nil
}

func (m *Inbound) validateAnnouncedQuantity(formats strfmt.Registry) error {

	if err := validate.Required("announcedQuantity", "body", m.AnnouncedQuantity); err != nil {
		return err
	}

	return nil
}

func (m *Inbound) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Inbound) validateFbbTransporter(formats strfmt.Registry) error {

	if err := validate.Required("fbbTransporter", "body", m.FbbTransporter); err != nil {
		return err
	}

	if m.FbbTransporter != nil {
		if err := m.FbbTransporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fbbTransporter")
			}
			return err
		}
	}

	return nil
}

func (m *Inbound) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Inbound) validateLabellingService(formats strfmt.Registry) error {

	if err := validate.Required("labellingService", "body", m.LabellingService); err != nil {
		return err
	}

	return nil
}

func (m *Inbound) validateProducts(formats strfmt.Registry) error {

	if err := validate.Required("products", "body", m.Products); err != nil {
		return err
	}

	for i := 0; i < len(m.Products); i++ {
		if swag.IsZero(m.Products[i]) { // not required
			continue
		}

		if m.Products[i] != nil {
			if err := m.Products[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Inbound) validateReceivedBSKUs(formats strfmt.Registry) error {

	if err := validate.Required("receivedBSKUs", "body", m.ReceivedBSKUs); err != nil {
		return err
	}

	return nil
}

func (m *Inbound) validateReceivedQuantity(formats strfmt.Registry) error {

	if err := validate.Required("receivedQuantity", "body", m.ReceivedQuantity); err != nil {
		return err
	}

	return nil
}

func (m *Inbound) validateReference(formats strfmt.Registry) error {

	if err := validate.Required("reference", "body", m.Reference); err != nil {
		return err
	}

	return nil
}

var inboundTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Draft","PreAnnounced","ArrivedAtWH","Cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inboundTypeStatePropEnum = append(inboundTypeStatePropEnum, v)
	}
}

const (

	// InboundStateDraft captures enum value "Draft"
	InboundStateDraft string = "Draft"

	// InboundStatePreAnnounced captures enum value "PreAnnounced"
	InboundStatePreAnnounced string = "PreAnnounced"

	// InboundStateArrivedAtWH captures enum value "ArrivedAtWH"
	InboundStateArrivedAtWH string = "ArrivedAtWH"

	// InboundStateCancelled captures enum value "Cancelled"
	InboundStateCancelled string = "Cancelled"
)

// prop value enum
func (m *Inbound) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, inboundTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Inbound) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *Inbound) validateStateTransitions(formats strfmt.Registry) error {

	if err := validate.Required("stateTransitions", "body", m.StateTransitions); err != nil {
		return err
	}

	for i := 0; i < len(m.StateTransitions); i++ {
		if swag.IsZero(m.StateTransitions[i]) { // not required
			continue
		}

		if m.StateTransitions[i] != nil {
			if err := m.StateTransitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stateTransitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Inbound) validateTimeSlot(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeSlot) { // not required
		return nil
	}

	if m.TimeSlot != nil {
		if err := m.TimeSlot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeSlot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Inbound) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Inbound) UnmarshalBinary(b []byte) error {
	var res Inbound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
