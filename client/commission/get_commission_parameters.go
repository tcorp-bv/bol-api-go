// Code generated by go-swagger; DO NOT EDIT.

package commission

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

// NewGetCommissionParams creates a new GetCommissionParams object
// with the default values initialized.
func NewGetCommissionParams() *GetCommissionParams {
	var (
		conditionDefault = string("NEW")
	)
	return &GetCommissionParams{
		Condition: &conditionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCommissionParamsWithTimeout creates a new GetCommissionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCommissionParamsWithTimeout(timeout time.Duration) *GetCommissionParams {
	var (
		conditionDefault = string("NEW")
	)
	return &GetCommissionParams{
		Condition: &conditionDefault,

		timeout: timeout,
	}
}

// NewGetCommissionParamsWithContext creates a new GetCommissionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCommissionParamsWithContext(ctx context.Context) *GetCommissionParams {
	var (
		conditionDefault = string("NEW")
	)
	return &GetCommissionParams{
		Condition: &conditionDefault,

		Context: ctx,
	}
}

// NewGetCommissionParamsWithHTTPClient creates a new GetCommissionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCommissionParamsWithHTTPClient(client *http.Client) *GetCommissionParams {
	var (
		conditionDefault = string("NEW")
	)
	return &GetCommissionParams{
		Condition:  &conditionDefault,
		HTTPClient: client,
	}
}

/*GetCommissionParams contains all the parameters to send to the API endpoint
for the get commission operation typically these are written to a http.Request
*/
type GetCommissionParams struct {

	/*Condition
	  The condition of the offer.

	*/
	Condition *string
	/*Ean
	  The EAN number associated with this product.

	*/
	Ean string
	/*Price
	  The price of the product with a period as a decimal separator. The price should always have two decimals precision.

	*/
	Price *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get commission params
func (o *GetCommissionParams) WithTimeout(timeout time.Duration) *GetCommissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get commission params
func (o *GetCommissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get commission params
func (o *GetCommissionParams) WithContext(ctx context.Context) *GetCommissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get commission params
func (o *GetCommissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get commission params
func (o *GetCommissionParams) WithHTTPClient(client *http.Client) *GetCommissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get commission params
func (o *GetCommissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCondition adds the condition to the get commission params
func (o *GetCommissionParams) WithCondition(condition *string) *GetCommissionParams {
	o.SetCondition(condition)
	return o
}

// SetCondition adds the condition to the get commission params
func (o *GetCommissionParams) SetCondition(condition *string) {
	o.Condition = condition
}

// WithEan adds the ean to the get commission params
func (o *GetCommissionParams) WithEan(ean string) *GetCommissionParams {
	o.SetEan(ean)
	return o
}

// SetEan adds the ean to the get commission params
func (o *GetCommissionParams) SetEan(ean string) {
	o.Ean = ean
}

// WithPrice adds the price to the get commission params
func (o *GetCommissionParams) WithPrice(price *float64) *GetCommissionParams {
	o.SetPrice(price)
	return o
}

// SetPrice adds the price to the get commission params
func (o *GetCommissionParams) SetPrice(price *float64) {
	o.Price = price
}

// WriteToRequest writes these params to a swagger request
func (o *GetCommissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Condition != nil {

		// query param condition
		var qrCondition string
		if o.Condition != nil {
			qrCondition = *o.Condition
		}
		qCondition := qrCondition
		if qCondition != "" {
			if err := r.SetQueryParam("condition", qCondition); err != nil {
				return err
			}
		}

	}

	// path param ean
	if err := r.SetPathParam("ean", o.Ean); err != nil {
		return err
	}

	if o.Price != nil {

		// query param price
		var qrPrice float64
		if o.Price != nil {
			qrPrice = *o.Price
		}
		qPrice := swag.FormatFloat64(qrPrice)
		if qPrice != "" {
			if err := r.SetQueryParam("price", qPrice); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}