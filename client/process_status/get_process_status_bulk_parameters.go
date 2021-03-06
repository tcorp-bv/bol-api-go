// Code generated by go-swagger; DO NOT EDIT.

package process_status

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

// NewGetProcessStatusBulkParams creates a new GetProcessStatusBulkParams object
// with the default values initialized.
func NewGetProcessStatusBulkParams() *GetProcessStatusBulkParams {
	var ()
	return &GetProcessStatusBulkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessStatusBulkParamsWithTimeout creates a new GetProcessStatusBulkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProcessStatusBulkParamsWithTimeout(timeout time.Duration) *GetProcessStatusBulkParams {
	var ()
	return &GetProcessStatusBulkParams{

		timeout: timeout,
	}
}

// NewGetProcessStatusBulkParamsWithContext creates a new GetProcessStatusBulkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProcessStatusBulkParamsWithContext(ctx context.Context) *GetProcessStatusBulkParams {
	var ()
	return &GetProcessStatusBulkParams{

		Context: ctx,
	}
}

// NewGetProcessStatusBulkParamsWithHTTPClient creates a new GetProcessStatusBulkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProcessStatusBulkParamsWithHTTPClient(client *http.Client) *GetProcessStatusBulkParams {
	var ()
	return &GetProcessStatusBulkParams{
		HTTPClient: client,
	}
}

/*GetProcessStatusBulkParams contains all the parameters to send to the API endpoint
for the get process status bulk operation typically these are written to a http.Request
*/
type GetProcessStatusBulkParams struct {

	/*Body*/
	Body *models.BulkProcessStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get process status bulk params
func (o *GetProcessStatusBulkParams) WithTimeout(timeout time.Duration) *GetProcessStatusBulkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get process status bulk params
func (o *GetProcessStatusBulkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get process status bulk params
func (o *GetProcessStatusBulkParams) WithContext(ctx context.Context) *GetProcessStatusBulkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get process status bulk params
func (o *GetProcessStatusBulkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get process status bulk params
func (o *GetProcessStatusBulkParams) WithHTTPClient(client *http.Client) *GetProcessStatusBulkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get process status bulk params
func (o *GetProcessStatusBulkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get process status bulk params
func (o *GetProcessStatusBulkParams) WithBody(body *models.BulkProcessStatusRequest) *GetProcessStatusBulkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get process status bulk params
func (o *GetProcessStatusBulkParams) SetBody(body *models.BulkProcessStatusRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessStatusBulkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
