// Code generated by go-swagger; DO NOT EDIT.

package shipments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetShipmentParams creates a new GetShipmentParams object
// with the default values initialized.
func NewGetShipmentParams() *GetShipmentParams {
	var ()
	return &GetShipmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetShipmentParamsWithTimeout creates a new GetShipmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetShipmentParamsWithTimeout(timeout time.Duration) *GetShipmentParams {
	var ()
	return &GetShipmentParams{

		timeout: timeout,
	}
}

// NewGetShipmentParamsWithContext creates a new GetShipmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetShipmentParamsWithContext(ctx context.Context) *GetShipmentParams {
	var ()
	return &GetShipmentParams{

		Context: ctx,
	}
}

// NewGetShipmentParamsWithHTTPClient creates a new GetShipmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetShipmentParamsWithHTTPClient(client *http.Client) *GetShipmentParams {
	var ()
	return &GetShipmentParams{
		HTTPClient: client,
	}
}

/*GetShipmentParams contains all the parameters to send to the API endpoint
for the get shipment operation typically these are written to a http.Request
*/
type GetShipmentParams struct {

	/*ShipmentID
	  The id of the shipment.

	*/
	ShipmentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get shipment params
func (o *GetShipmentParams) WithTimeout(timeout time.Duration) *GetShipmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get shipment params
func (o *GetShipmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get shipment params
func (o *GetShipmentParams) WithContext(ctx context.Context) *GetShipmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get shipment params
func (o *GetShipmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get shipment params
func (o *GetShipmentParams) WithHTTPClient(client *http.Client) *GetShipmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get shipment params
func (o *GetShipmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithShipmentID adds the shipmentID to the get shipment params
func (o *GetShipmentParams) WithShipmentID(shipmentID int64) *GetShipmentParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the get shipment params
func (o *GetShipmentParams) SetShipmentID(shipmentID int64) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetShipmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param shipment-id
	if err := r.SetPathParam("shipment-id", swag.FormatInt64(o.ShipmentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
