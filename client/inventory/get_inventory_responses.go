// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tcorp-bv/bol-api-go/models"
)

// GetInventoryReader is a Reader for the GetInventory structure.
type GetInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInventoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInventoryOK creates a GetInventoryOK with default headers values
func NewGetInventoryOK() *GetInventoryOK {
	return &GetInventoryOK{}
}

/*GetInventoryOK handles this case with default header values.

Ok: Successfully processed the request.
*/
type GetInventoryOK struct {
	Payload *models.InventoryResponse
}

func (o *GetInventoryOK) Error() string {
	return fmt.Sprintf("[GET /retailer/inventory][%d] getInventoryOK  %+v", 200, o.Payload)
}

func (o *GetInventoryOK) GetPayload() *models.InventoryResponse {
	return o.Payload
}

func (o *GetInventoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryBadRequest creates a GetInventoryBadRequest with default headers values
func NewGetInventoryBadRequest() *GetInventoryBadRequest {
	return &GetInventoryBadRequest{}
}

/*GetInventoryBadRequest handles this case with default header values.

Bad request: The sent request does not meet the API specification. Please check the error message(s) for more information.
*/
type GetInventoryBadRequest struct {
	Payload *models.Problem
}

func (o *GetInventoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /retailer/inventory][%d] getInventoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetInventoryBadRequest) GetPayload() *models.Problem {
	return o.Payload
}

func (o *GetInventoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}