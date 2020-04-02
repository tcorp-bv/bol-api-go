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

// ProductLabelsRequest The product labels to retrieve.
//
// swagger:model ProductLabelsRequest
type ProductLabelsRequest struct {

	// The printer format to create labels for, defaults to AVERY_J8159 if not provided.
	// Enum: [AVERY_J8159 AVERY_J8160 AVERY_3474 DYMO_99012 BROTHER_DK11208D ZEBRA_Z_PERFORM_1000T]
	Format string `json:"format,omitempty"`

	// product labels
	// Required: true
	// Max Items: 2147483647
	// Min Items: 1
	ProductLabels []*ProductLabel `json:"productLabels"`
}

// Validate validates this product labels request
func (m *ProductLabelsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var productLabelsRequestTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AVERY_J8159","AVERY_J8160","AVERY_3474","DYMO_99012","BROTHER_DK11208D","ZEBRA_Z_PERFORM_1000T"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productLabelsRequestTypeFormatPropEnum = append(productLabelsRequestTypeFormatPropEnum, v)
	}
}

const (

	// ProductLabelsRequestFormatAVERYJ8159 captures enum value "AVERY_J8159"
	ProductLabelsRequestFormatAVERYJ8159 string = "AVERY_J8159"

	// ProductLabelsRequestFormatAVERYJ8160 captures enum value "AVERY_J8160"
	ProductLabelsRequestFormatAVERYJ8160 string = "AVERY_J8160"

	// ProductLabelsRequestFormatAVERY3474 captures enum value "AVERY_3474"
	ProductLabelsRequestFormatAVERY3474 string = "AVERY_3474"

	// ProductLabelsRequestFormatDYMO99012 captures enum value "DYMO_99012"
	ProductLabelsRequestFormatDYMO99012 string = "DYMO_99012"

	// ProductLabelsRequestFormatBROTHERDK11208D captures enum value "BROTHER_DK11208D"
	ProductLabelsRequestFormatBROTHERDK11208D string = "BROTHER_DK11208D"

	// ProductLabelsRequestFormatZEBRAZPERFORM1000T captures enum value "ZEBRA_Z_PERFORM_1000T"
	ProductLabelsRequestFormatZEBRAZPERFORM1000T string = "ZEBRA_Z_PERFORM_1000T"
)

// prop value enum
func (m *ProductLabelsRequest) validateFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productLabelsRequestTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProductLabelsRequest) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *ProductLabelsRequest) validateProductLabels(formats strfmt.Registry) error {

	if err := validate.Required("productLabels", "body", m.ProductLabels); err != nil {
		return err
	}

	iProductLabelsSize := int64(len(m.ProductLabels))

	if err := validate.MinItems("productLabels", "body", iProductLabelsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("productLabels", "body", iProductLabelsSize, 2147483647); err != nil {
		return err
	}

	for i := 0; i < len(m.ProductLabels); i++ {
		if swag.IsZero(m.ProductLabels[i]) { // not required
			continue
		}

		if m.ProductLabels[i] != nil {
			if err := m.ProductLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductLabelsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductLabelsRequest) UnmarshalBinary(b []byte) error {
	var res ProductLabelsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
