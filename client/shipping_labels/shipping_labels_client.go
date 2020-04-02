// Code generated by go-swagger; DO NOT EDIT.

package shipping_labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new shipping labels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for shipping labels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetShippingLabels(params *GetShippingLabelsParams) (*GetShippingLabelsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetShippingLabels gets purchasable shipping labels by order item id

  Retrieves all available shipping labels based on the supplied order item.
*/
func (a *Client) GetShippingLabels(params *GetShippingLabelsParams) (*GetShippingLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShippingLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-shipping-labels",
		Method:             "GET",
		PathPattern:        "/retailer/purchasable-shippinglabels/{order-item-id}",
		ProducesMediaTypes: []string{"application/vnd.retailer.v3+json", "application/vnd.retailer.v3+xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetShippingLabelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetShippingLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-shipping-labels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
