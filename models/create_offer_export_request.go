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

// CreateOfferExportRequest create offer export request
//
// swagger:model CreateOfferExportRequest
type CreateOfferExportRequest struct {

	// The file format in which to return the export.
	// Required: true
	// Enum: [CSV]
	Format *string `json:"format"`
}

// Validate validates this create offer export request
func (m *CreateOfferExportRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createOfferExportRequestTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CSV"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOfferExportRequestTypeFormatPropEnum = append(createOfferExportRequestTypeFormatPropEnum, v)
	}
}

const (

	// CreateOfferExportRequestFormatCSV captures enum value "CSV"
	CreateOfferExportRequestFormatCSV string = "CSV"
)

// prop value enum
func (m *CreateOfferExportRequest) validateFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createOfferExportRequestTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateOfferExportRequest) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("format", "body", m.Format); err != nil {
		return err
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", *m.Format); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateOfferExportRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOfferExportRequest) UnmarshalBinary(b []byte) error {
	var res CreateOfferExportRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
