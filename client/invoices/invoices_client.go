// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new invoices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for invoices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetInvoice(params *GetInvoiceParams) (*GetInvoiceOK, error)

	GetInvoiceSpecification(params *GetInvoiceSpecificationParams) (*GetInvoiceSpecificationOK, error)

	GetInvoices(params *GetInvoicesParams) (*GetInvoicesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetInvoice gets an invoice by invoice id

  Gets an invoice by invoice id. The available media types differ per invoice and are listed within the response from the ‘GET all invoices’ call. Note: the media types listed in the response must be given in our standard API format.
*/
func (a *Client) GetInvoice(params *GetInvoiceParams) (*GetInvoiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-invoice",
		Method:             "GET",
		PathPattern:        "/retailer/invoices/{invoice-id}",
		ProducesMediaTypes: []string{"application/vnd.retailer.v3+json", "application/vnd.retailer.v3+pdf", "application/vnd.retailer.v3+xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvoiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-invoice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInvoiceSpecification gets an invoice specification by invoice id

  Gets an invoice specification for an invoice with a paginated list of its transactions. The available media types differ per invoice specification and are listed within the response from the ‘GET all invoices’ call. Note, the media types listed in the response must be given in our standard API format.
*/
func (a *Client) GetInvoiceSpecification(params *GetInvoiceSpecificationParams) (*GetInvoiceSpecificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceSpecificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-invoice-specification",
		Method:             "GET",
		PathPattern:        "/retailer/invoices/{invoice-id}/specification",
		ProducesMediaTypes: []string{"application/vnd.retailer.v3+json", "application/vnd.retailer.v3+openxmlformats-officedocument.spreadsheetml.sheet", "application/vnd.retailer.v3+pdf", "application/vnd.retailer.v3+xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvoiceSpecificationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceSpecificationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-invoice-specification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInvoices gets all invoices

  Gets a list of invoices, by default from the past 4 weeks. The optional period-start and period-end parameters can be used together to retrieve invoices from a specific date range in the past, the period can be no longer than 31 days. Invoices and their specifications can be downloaded separately in different media formats with the ‘GET an invoice by id’ and the ‘GET an invoice specification by id’ calls. The available media types differ per invoice and are listed per invoice within the response. Note: the media types listed in the response must be given in our standard API format.
*/
func (a *Client) GetInvoices(params *GetInvoicesParams) (*GetInvoicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-invoices",
		Method:             "GET",
		PathPattern:        "/retailer/invoices",
		ProducesMediaTypes: []string{"application/vnd.retailer.v3+json", "application/vnd.retailer.v3+xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvoicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-invoices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
