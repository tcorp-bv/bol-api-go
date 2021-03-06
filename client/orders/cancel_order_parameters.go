// Code generated by go-swagger; DO NOT EDIT.

package orders

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

	"github.com/tcorp-bv/bol-api-go/models"
)

// NewCancelOrderParams creates a new CancelOrderParams object
// with the default values initialized.
func NewCancelOrderParams() *CancelOrderParams {
	var ()
	return &CancelOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelOrderParamsWithTimeout creates a new CancelOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelOrderParamsWithTimeout(timeout time.Duration) *CancelOrderParams {
	var ()
	return &CancelOrderParams{

		timeout: timeout,
	}
}

// NewCancelOrderParamsWithContext creates a new CancelOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelOrderParamsWithContext(ctx context.Context) *CancelOrderParams {
	var ()
	return &CancelOrderParams{

		Context: ctx,
	}
}

// NewCancelOrderParamsWithHTTPClient creates a new CancelOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelOrderParamsWithHTTPClient(client *http.Client) *CancelOrderParams {
	var ()
	return &CancelOrderParams{
		HTTPClient: client,
	}
}

/*CancelOrderParams contains all the parameters to send to the API endpoint
for the cancel order operation typically these are written to a http.Request
*/
type CancelOrderParams struct {

	/*Body*/
	Body *models.Cancellation
	/*OrderItemID
	  The id of the order item to cancel.

	*/
	OrderItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel order params
func (o *CancelOrderParams) WithTimeout(timeout time.Duration) *CancelOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel order params
func (o *CancelOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel order params
func (o *CancelOrderParams) WithContext(ctx context.Context) *CancelOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel order params
func (o *CancelOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel order params
func (o *CancelOrderParams) WithHTTPClient(client *http.Client) *CancelOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel order params
func (o *CancelOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cancel order params
func (o *CancelOrderParams) WithBody(body *models.Cancellation) *CancelOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cancel order params
func (o *CancelOrderParams) SetBody(body *models.Cancellation) {
	o.Body = body
}

// WithOrderItemID adds the orderItemID to the cancel order params
func (o *CancelOrderParams) WithOrderItemID(orderItemID string) *CancelOrderParams {
	o.SetOrderItemID(orderItemID)
	return o
}

// SetOrderItemID adds the orderItemId to the cancel order params
func (o *CancelOrderParams) SetOrderItemID(orderItemID string) {
	o.OrderItemID = orderItemID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param order-item-id
	if err := r.SetPathParam("order-item-id", o.OrderItemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
