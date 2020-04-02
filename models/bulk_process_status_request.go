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

// BulkProcessStatusRequest bulk process status request
//
// swagger:model BulkProcessStatusRequest
type BulkProcessStatusRequest struct {

	// process status queries
	// Required: true
	// Max Items: 1000
	// Min Items: 1
	ProcessStatusQueries []*ProcessStatusID `json:"processStatusQueries"`
}

// Validate validates this bulk process status request
func (m *BulkProcessStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessStatusQueries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkProcessStatusRequest) validateProcessStatusQueries(formats strfmt.Registry) error {

	if err := validate.Required("processStatusQueries", "body", m.ProcessStatusQueries); err != nil {
		return err
	}

	iProcessStatusQueriesSize := int64(len(m.ProcessStatusQueries))

	if err := validate.MinItems("processStatusQueries", "body", iProcessStatusQueriesSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("processStatusQueries", "body", iProcessStatusQueriesSize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(m.ProcessStatusQueries); i++ {
		if swag.IsZero(m.ProcessStatusQueries[i]) { // not required
			continue
		}

		if m.ProcessStatusQueries[i] != nil {
			if err := m.ProcessStatusQueries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processStatusQueries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkProcessStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkProcessStatusRequest) UnmarshalBinary(b []byte) error {
	var res BulkProcessStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}