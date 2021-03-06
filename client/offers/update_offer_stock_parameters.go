// Code generated by go-swagger; DO NOT EDIT.

package offers

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

// NewUpdateOfferStockParams creates a new UpdateOfferStockParams object
// with the default values initialized.
func NewUpdateOfferStockParams() *UpdateOfferStockParams {
	var ()
	return &UpdateOfferStockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOfferStockParamsWithTimeout creates a new UpdateOfferStockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOfferStockParamsWithTimeout(timeout time.Duration) *UpdateOfferStockParams {
	var ()
	return &UpdateOfferStockParams{

		timeout: timeout,
	}
}

// NewUpdateOfferStockParamsWithContext creates a new UpdateOfferStockParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOfferStockParamsWithContext(ctx context.Context) *UpdateOfferStockParams {
	var ()
	return &UpdateOfferStockParams{

		Context: ctx,
	}
}

// NewUpdateOfferStockParamsWithHTTPClient creates a new UpdateOfferStockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOfferStockParamsWithHTTPClient(client *http.Client) *UpdateOfferStockParams {
	var ()
	return &UpdateOfferStockParams{
		HTTPClient: client,
	}
}

/*UpdateOfferStockParams contains all the parameters to send to the API endpoint
for the update offer stock operation typically these are written to a http.Request
*/
type UpdateOfferStockParams struct {

	/*Body*/
	Body *models.UpdateOfferStockRequest
	/*OfferID
	  Unique identifier for an offer.

	*/
	OfferID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update offer stock params
func (o *UpdateOfferStockParams) WithTimeout(timeout time.Duration) *UpdateOfferStockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update offer stock params
func (o *UpdateOfferStockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update offer stock params
func (o *UpdateOfferStockParams) WithContext(ctx context.Context) *UpdateOfferStockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update offer stock params
func (o *UpdateOfferStockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update offer stock params
func (o *UpdateOfferStockParams) WithHTTPClient(client *http.Client) *UpdateOfferStockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update offer stock params
func (o *UpdateOfferStockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update offer stock params
func (o *UpdateOfferStockParams) WithBody(body *models.UpdateOfferStockRequest) *UpdateOfferStockParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update offer stock params
func (o *UpdateOfferStockParams) SetBody(body *models.UpdateOfferStockRequest) {
	o.Body = body
}

// WithOfferID adds the offerID to the update offer stock params
func (o *UpdateOfferStockParams) WithOfferID(offerID string) *UpdateOfferStockParams {
	o.SetOfferID(offerID)
	return o
}

// SetOfferID adds the offerId to the update offer stock params
func (o *UpdateOfferStockParams) SetOfferID(offerID string) {
	o.OfferID = offerID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOfferStockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param offer-id
	if err := r.SetPathParam("offer-id", o.OfferID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
