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

// NewPutOfferParams creates a new PutOfferParams object
// with the default values initialized.
func NewPutOfferParams() *PutOfferParams {
	var ()
	return &PutOfferParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOfferParamsWithTimeout creates a new PutOfferParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOfferParamsWithTimeout(timeout time.Duration) *PutOfferParams {
	var ()
	return &PutOfferParams{

		timeout: timeout,
	}
}

// NewPutOfferParamsWithContext creates a new PutOfferParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOfferParamsWithContext(ctx context.Context) *PutOfferParams {
	var ()
	return &PutOfferParams{

		Context: ctx,
	}
}

// NewPutOfferParamsWithHTTPClient creates a new PutOfferParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOfferParamsWithHTTPClient(client *http.Client) *PutOfferParams {
	var ()
	return &PutOfferParams{
		HTTPClient: client,
	}
}

/*PutOfferParams contains all the parameters to send to the API endpoint
for the put offer operation typically these are written to a http.Request
*/
type PutOfferParams struct {

	/*Body*/
	Body *models.UpdateOfferRequest
	/*OfferID
	  Unique identifier for an offer.

	*/
	OfferID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put offer params
func (o *PutOfferParams) WithTimeout(timeout time.Duration) *PutOfferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put offer params
func (o *PutOfferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put offer params
func (o *PutOfferParams) WithContext(ctx context.Context) *PutOfferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put offer params
func (o *PutOfferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put offer params
func (o *PutOfferParams) WithHTTPClient(client *http.Client) *PutOfferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put offer params
func (o *PutOfferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put offer params
func (o *PutOfferParams) WithBody(body *models.UpdateOfferRequest) *PutOfferParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put offer params
func (o *PutOfferParams) SetBody(body *models.UpdateOfferRequest) {
	o.Body = body
}

// WithOfferID adds the offerID to the put offer params
func (o *PutOfferParams) WithOfferID(offerID string) *PutOfferParams {
	o.SetOfferID(offerID)
	return o
}

// SetOfferID adds the offerId to the put offer params
func (o *PutOfferParams) SetOfferID(offerID string) {
	o.OfferID = offerID
}

// WriteToRequest writes these params to a swagger request
func (o *PutOfferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
