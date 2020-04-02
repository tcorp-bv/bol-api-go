// Code generated by go-swagger; DO NOT EDIT.

package inbounds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tcorp-bv/bol-api-go/models"
)

// GetDeliveryWindowsReader is a Reader for the GetDeliveryWindows structure.
type GetDeliveryWindowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeliveryWindowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeliveryWindowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeliveryWindowsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeliveryWindowsOK creates a GetDeliveryWindowsOK with default headers values
func NewGetDeliveryWindowsOK() *GetDeliveryWindowsOK {
	return &GetDeliveryWindowsOK{}
}

/*GetDeliveryWindowsOK handles this case with default header values.

Ok: Successfully processed the request.
*/
type GetDeliveryWindowsOK struct {
	Payload *models.DeliveryWindowsForInboundShipments
}

func (o *GetDeliveryWindowsOK) Error() string {
	return fmt.Sprintf("[GET /retailer/inbounds/delivery-windows][%d] getDeliveryWindowsOK  %+v", 200, o.Payload)
}

func (o *GetDeliveryWindowsOK) GetPayload() *models.DeliveryWindowsForInboundShipments {
	return o.Payload
}

func (o *GetDeliveryWindowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeliveryWindowsForInboundShipments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryWindowsBadRequest creates a GetDeliveryWindowsBadRequest with default headers values
func NewGetDeliveryWindowsBadRequest() *GetDeliveryWindowsBadRequest {
	return &GetDeliveryWindowsBadRequest{}
}

/*GetDeliveryWindowsBadRequest handles this case with default header values.

Bad request: The sent request does not meet the API specification. Please check the error message(s) for more information.
*/
type GetDeliveryWindowsBadRequest struct {
	Payload *models.Problem
}

func (o *GetDeliveryWindowsBadRequest) Error() string {
	return fmt.Sprintf("[GET /retailer/inbounds/delivery-windows][%d] getDeliveryWindowsBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeliveryWindowsBadRequest) GetPayload() *models.Problem {
	return o.Payload
}

func (o *GetDeliveryWindowsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}