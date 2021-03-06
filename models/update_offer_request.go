// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateOfferRequest update offer request
//
// swagger:model UpdateOfferRequest
type UpdateOfferRequest struct {

	// fulfilment
	// Required: true
	Fulfilment *Fulfilment `json:"fulfilment"`

	// Indicates whether or not you want to put this offer for sale on the bol.com website. Defaults to false.
	OnHoldByRetailer bool `json:"onHoldByRetailer,omitempty"`

	// A user-defined reference that helps you identify this particular offer when receiving data from us. This element can optionally be left empty and has a maximum amount of 20 characters.
	// Max Length: 20
	// Min Length: 0
	ReferenceCode *string `json:"referenceCode,omitempty"`

	// In case the item is not known to bol.com you can use this field to identify this particular product. Note: in case the product is known to bol.com, the unknown product title will not be stored.
	// Max Length: 500
	// Min Length: 0
	UnknownProductTitle *string `json:"unknownProductTitle,omitempty"`
}

// Validate validates this update offer request
func (m *UpdateOfferRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFulfilment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnknownProductTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateOfferRequest) validateFulfilment(formats strfmt.Registry) error {

	if err := validate.Required("fulfilment", "body", m.Fulfilment); err != nil {
		return err
	}

	if m.Fulfilment != nil {
		if err := m.Fulfilment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fulfilment")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateOfferRequest) validateReferenceCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceCode) { // not required
		return nil
	}

	if err := validate.MinLength("referenceCode", "body", string(*m.ReferenceCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("referenceCode", "body", string(*m.ReferenceCode), 20); err != nil {
		return err
	}

	return nil
}

func (m *UpdateOfferRequest) validateUnknownProductTitle(formats strfmt.Registry) error {

	if swag.IsZero(m.UnknownProductTitle) { // not required
		return nil
	}

	if err := validate.MinLength("unknownProductTitle", "body", string(*m.UnknownProductTitle), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("unknownProductTitle", "body", string(*m.UnknownProductTitle), 500); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateOfferRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateOfferRequest) UnmarshalBinary(b []byte) error {
	var res UpdateOfferRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
