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

// ReturnItem return item
//
// swagger:model ReturnItem
type ReturnItem struct {

	// customer details
	CustomerDetails *CustomerDetails `json:"customerDetails,omitempty"`

	// The EAN number associated with this product.
	Ean string `json:"ean,omitempty"`

	// Specifies whether this shipment has been fulfilled by the retailer (FBR) or fulfilled by bol.com (FBB). Defaults to FBR.
	FulfilmentMethod string `json:"fulfilmentMethod,omitempty"`

	// Indicates if this return item has been handled (by the retailer).
	Handled bool `json:"handled,omitempty"`

	// The handling result requested by the retailer.
	HandlingResult string `json:"handlingResult,omitempty"`

	// The id of the customer order this return item is in.
	OrderID string `json:"orderId,omitempty"`

	// The date and time in ISO 8601 format when the return was processed.
	// Format: date-time
	ProcessingDateTime strfmt.DateTime `json:"processingDateTime,omitempty"`

	// The processing result of the return.
	ProcessingResult string `json:"processingResult,omitempty"`

	// The quantity that is returned by the customer. Note: this can be greater than 1 in case the customer ordered a quantity greater than 1 of the same product in the same customer order
	Quantity int32 `json:"quantity,omitempty"`

	// The date and time in ISO 8601 format when this return was registered.
	// Format: date-time
	RegistrationDateTime strfmt.DateTime `json:"registrationDateTime,omitempty"`

	// The reason why the customer returned this product.
	ReturnReason string `json:"returnReason,omitempty"`

	// Additional details from the customer as to why this item was returned.
	ReturnReasonComments string `json:"returnReasonComments,omitempty"`

	// The RMA (Return Merchandise Authorization) id that identifies this particular return.
	RmaID int64 `json:"rmaId,omitempty"`

	// The product title.
	Title string `json:"title,omitempty"`

	// The track and trace code that is associated with this transport.
	TrackAndTrace string `json:"trackAndTrace,omitempty"`
}

// Validate validates this return item
func (m *ReturnItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrationDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnItem) validateCustomerDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerDetails) { // not required
		return nil
	}

	if m.CustomerDetails != nil {
		if err := m.CustomerDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerDetails")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnItem) validateProcessingDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("processingDateTime", "body", "date-time", m.ProcessingDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnItem) validateRegistrationDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistrationDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("registrationDateTime", "body", "date-time", m.RegistrationDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnItem) UnmarshalBinary(b []byte) error {
	var res ReturnItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
