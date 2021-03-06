// Code generated by go-swagger; DO NOT EDIT.

package returns

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

// NewGetReturnsParams creates a new GetReturnsParams object
// with the default values initialized.
func NewGetReturnsParams() *GetReturnsParams {
	var (
		fulfilmentMethodDefault = string("FBR")
		handledDefault          = bool(false)
		pageDefault             = int32(1)
	)
	return &GetReturnsParams{
		FulfilmentMethod: &fulfilmentMethodDefault,
		Handled:          &handledDefault,
		Page:             &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReturnsParamsWithTimeout creates a new GetReturnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReturnsParamsWithTimeout(timeout time.Duration) *GetReturnsParams {
	var (
		fulfilmentMethodDefault = string("FBR")
		handledDefault          = bool(false)
		pageDefault             = int32(1)
	)
	return &GetReturnsParams{
		FulfilmentMethod: &fulfilmentMethodDefault,
		Handled:          &handledDefault,
		Page:             &pageDefault,

		timeout: timeout,
	}
}

// NewGetReturnsParamsWithContext creates a new GetReturnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReturnsParamsWithContext(ctx context.Context) *GetReturnsParams {
	var (
		fulfilmentMethodDefault = string("FBR")
		handledDefault          = bool(false)
		pageDefault             = int32(1)
	)
	return &GetReturnsParams{
		FulfilmentMethod: &fulfilmentMethodDefault,
		Handled:          &handledDefault,
		Page:             &pageDefault,

		Context: ctx,
	}
}

// NewGetReturnsParamsWithHTTPClient creates a new GetReturnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReturnsParamsWithHTTPClient(client *http.Client) *GetReturnsParams {
	var (
		fulfilmentMethodDefault = string("FBR")
		handledDefault          = bool(false)
		pageDefault             = int32(1)
	)
	return &GetReturnsParams{
		FulfilmentMethod: &fulfilmentMethodDefault,
		Handled:          &handledDefault,
		Page:             &pageDefault,
		HTTPClient:       client,
	}
}

/*GetReturnsParams contains all the parameters to send to the API endpoint
for the get returns operation typically these are written to a http.Request
*/
type GetReturnsParams struct {

	/*FulfilmentMethod
	  The fulfilment method. Fulfilled by the retailer (FBR) or fulfilled by bol.com (FBB).

	*/
	FulfilmentMethod *string
	/*Handled
	  The status of the returns you wish to see, shows either handled or unhandled returns.

	*/
	Handled *bool
	/*Page
	  The requested page number with a pagesize of 50. The requested page number with a pagesize of 50 returns (within one return there can be multiple rma id's).

	*/
	Page *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get returns params
func (o *GetReturnsParams) WithTimeout(timeout time.Duration) *GetReturnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get returns params
func (o *GetReturnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get returns params
func (o *GetReturnsParams) WithContext(ctx context.Context) *GetReturnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get returns params
func (o *GetReturnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get returns params
func (o *GetReturnsParams) WithHTTPClient(client *http.Client) *GetReturnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get returns params
func (o *GetReturnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFulfilmentMethod adds the fulfilmentMethod to the get returns params
func (o *GetReturnsParams) WithFulfilmentMethod(fulfilmentMethod *string) *GetReturnsParams {
	o.SetFulfilmentMethod(fulfilmentMethod)
	return o
}

// SetFulfilmentMethod adds the fulfilmentMethod to the get returns params
func (o *GetReturnsParams) SetFulfilmentMethod(fulfilmentMethod *string) {
	o.FulfilmentMethod = fulfilmentMethod
}

// WithHandled adds the handled to the get returns params
func (o *GetReturnsParams) WithHandled(handled *bool) *GetReturnsParams {
	o.SetHandled(handled)
	return o
}

// SetHandled adds the handled to the get returns params
func (o *GetReturnsParams) SetHandled(handled *bool) {
	o.Handled = handled
}

// WithPage adds the page to the get returns params
func (o *GetReturnsParams) WithPage(page *int32) *GetReturnsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get returns params
func (o *GetReturnsParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetReturnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FulfilmentMethod != nil {

		// query param fulfilment-method
		var qrFulfilmentMethod string
		if o.FulfilmentMethod != nil {
			qrFulfilmentMethod = *o.FulfilmentMethod
		}
		qFulfilmentMethod := qrFulfilmentMethod
		if qFulfilmentMethod != "" {
			if err := r.SetQueryParam("fulfilment-method", qFulfilmentMethod); err != nil {
				return err
			}
		}

	}

	if o.Handled != nil {

		// query param handled
		var qrHandled bool
		if o.Handled != nil {
			qrHandled = *o.Handled
		}
		qHandled := swag.FormatBool(qrHandled)
		if qHandled != "" {
			if err := r.SetQueryParam("handled", qHandled); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
