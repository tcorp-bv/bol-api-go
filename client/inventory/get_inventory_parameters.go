// Code generated by go-swagger; DO NOT EDIT.

package inventory

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

// NewGetInventoryParams creates a new GetInventoryParams object
// with the default values initialized.
func NewGetInventoryParams() *GetInventoryParams {
	var (
		pageDefault = int32(1)
	)
	return &GetInventoryParams{
		Page: &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryParamsWithTimeout creates a new GetInventoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInventoryParamsWithTimeout(timeout time.Duration) *GetInventoryParams {
	var (
		pageDefault = int32(1)
	)
	return &GetInventoryParams{
		Page: &pageDefault,

		timeout: timeout,
	}
}

// NewGetInventoryParamsWithContext creates a new GetInventoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInventoryParamsWithContext(ctx context.Context) *GetInventoryParams {
	var (
		pageDefault = int32(1)
	)
	return &GetInventoryParams{
		Page: &pageDefault,

		Context: ctx,
	}
}

// NewGetInventoryParamsWithHTTPClient creates a new GetInventoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInventoryParamsWithHTTPClient(client *http.Client) *GetInventoryParams {
	var (
		pageDefault = int32(1)
	)
	return &GetInventoryParams{
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*GetInventoryParams contains all the parameters to send to the API endpoint
for the get inventory operation typically these are written to a http.Request
*/
type GetInventoryParams struct {

	/*Page
	  The requested page number with a pagesize of 50

	*/
	Page *int32
	/*Quantity
	  Filter inventory by providing a range of quantity (min-range)-(max-range).

	*/
	Quantity []string
	/*Query
	  Filter inventory by EAN or product title.

	*/
	Query *string
	/*State
	  Filter inventory by stock type (saleable or unsaleable).

	*/
	State *string
	/*Stock
	  Filter inventory by stock level.

	*/
	Stock *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get inventory params
func (o *GetInventoryParams) WithTimeout(timeout time.Duration) *GetInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory params
func (o *GetInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory params
func (o *GetInventoryParams) WithContext(ctx context.Context) *GetInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory params
func (o *GetInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory params
func (o *GetInventoryParams) WithHTTPClient(client *http.Client) *GetInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory params
func (o *GetInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get inventory params
func (o *GetInventoryParams) WithPage(page *int32) *GetInventoryParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get inventory params
func (o *GetInventoryParams) SetPage(page *int32) {
	o.Page = page
}

// WithQuantity adds the quantity to the get inventory params
func (o *GetInventoryParams) WithQuantity(quantity []string) *GetInventoryParams {
	o.SetQuantity(quantity)
	return o
}

// SetQuantity adds the quantity to the get inventory params
func (o *GetInventoryParams) SetQuantity(quantity []string) {
	o.Quantity = quantity
}

// WithQuery adds the query to the get inventory params
func (o *GetInventoryParams) WithQuery(query *string) *GetInventoryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get inventory params
func (o *GetInventoryParams) SetQuery(query *string) {
	o.Query = query
}

// WithState adds the state to the get inventory params
func (o *GetInventoryParams) WithState(state *string) *GetInventoryParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get inventory params
func (o *GetInventoryParams) SetState(state *string) {
	o.State = state
}

// WithStock adds the stock to the get inventory params
func (o *GetInventoryParams) WithStock(stock *string) *GetInventoryParams {
	o.SetStock(stock)
	return o
}

// SetStock adds the stock to the get inventory params
func (o *GetInventoryParams) SetStock(stock *string) {
	o.Stock = stock
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	valuesQuantity := o.Quantity

	joinedQuantity := swag.JoinByFormat(valuesQuantity, "multi")
	// query array param quantity
	if err := r.SetQueryParam("quantity", joinedQuantity...); err != nil {
		return err
	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if o.Stock != nil {

		// query param stock
		var qrStock string
		if o.Stock != nil {
			qrStock = *o.Stock
		}
		qStock := qrStock
		if qStock != "" {
			if err := r.SetQueryParam("stock", qStock); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
