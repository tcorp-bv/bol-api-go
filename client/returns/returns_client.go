// Code generated by go-swagger; DO NOT EDIT.

package returns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new returns API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for returns API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetReturn(params *GetReturnParams) (*GetReturnOK, error)

	GetReturns(params *GetReturnsParams) (*GetReturnsOK, error)

	HandleReturn(params *HandleReturnParams) (*HandleReturnAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetReturn gets a return by rma id

  Disclaimer: We recommend using the most recent version for returns which is version 4. We discourage using version 3 as of January 16th. Retrieve a return based on the rma id.
*/
func (a *Client) GetReturn(params *GetReturnParams) (*GetReturnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReturnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-return",
		Method:             "GET",
		PathPattern:        "/retailer/returns/{rma-id}",
		ProducesMediaTypes: []string{"application/vnd.retailer.v3+json", "application/vnd.retailer.v3+xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReturnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReturnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-return: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReturns gets returns

  Disclaimer: We recommend using the most recent version for returns which is version 4. We discourage using version 3 as of January 16th. Get a paginated list of all returns, which are sorted by date in descending order.
*/
func (a *Client) GetReturns(params *GetReturnsParams) (*GetReturnsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReturnsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-returns",
		Method:             "GET",
		PathPattern:        "/retailer/returns",
		ProducesMediaTypes: []string{"application/vnd.retailer.v3+json", "application/vnd.retailer.v3+xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReturnsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReturnsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-returns: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HandleReturn handles a return

  Disclaimer: We recommend using the most recent version for returns which is version 4. We discourage using version 3 as of January 16th. Allows the user to handle a return. This can be to either handle an open return, or change the handlingResult of an already handled return. The latter is only possible in limited scenarios, please check the returns documentation for a complete list.
*/
func (a *Client) HandleReturn(params *HandleReturnParams) (*HandleReturnAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHandleReturnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "handle-return",
		Method:             "PUT",
		PathPattern:        "/retailer/returns/{rma-id}",
		ProducesMediaTypes: []string{"application/vnd.retailer.v3+json", "application/vnd.retailer.v3+xml"},
		ConsumesMediaTypes: []string{"application/vnd.retailer.v3+json", "application/vnd.retailer.v3+xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HandleReturnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HandleReturnAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for handle-return: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
