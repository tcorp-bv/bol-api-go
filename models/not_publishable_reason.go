// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotPublishableReason not publishable reason
//
// swagger:model NotPublishableReason
type NotPublishableReason struct {

	// Error code signalling that the offer is invalid.
	Code string `json:"code,omitempty"`

	// Error message describing the reason the offer is invalid.
	Description string `json:"description,omitempty"`
}

// Validate validates this not publishable reason
func (m *NotPublishableReason) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotPublishableReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotPublishableReason) UnmarshalBinary(b []byte) error {
	var res NotPublishableReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}