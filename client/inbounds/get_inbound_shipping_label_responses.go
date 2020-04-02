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

// GetInboundShippingLabelReader is a Reader for the GetInboundShippingLabel structure.
type GetInboundShippingLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInboundShippingLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInboundShippingLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInboundShippingLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInboundShippingLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInboundShippingLabelOK creates a GetInboundShippingLabelOK with default headers values
func NewGetInboundShippingLabelOK() *GetInboundShippingLabelOK {
	return &GetInboundShippingLabelOK{}
}

/*GetInboundShippingLabelOK handles this case with default header values.

Ok: Successfully processed the request.
*/
type GetInboundShippingLabelOK struct {
	Payload []strfmt.Base64
}

func (o *GetInboundShippingLabelOK) Error() string {
	return fmt.Sprintf("[GET /retailer/inbounds/{inbound-id}/shippinglabel][%d] getInboundShippingLabelOK  %+v", 200, o.Payload)
}

func (o *GetInboundShippingLabelOK) GetPayload() []strfmt.Base64 {
	return o.Payload
}

func (o *GetInboundShippingLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInboundShippingLabelBadRequest creates a GetInboundShippingLabelBadRequest with default headers values
func NewGetInboundShippingLabelBadRequest() *GetInboundShippingLabelBadRequest {
	return &GetInboundShippingLabelBadRequest{}
}

/*GetInboundShippingLabelBadRequest handles this case with default header values.

Bad request: The sent request does not meet the API specification. Please check the error message(s) for more information.
*/
type GetInboundShippingLabelBadRequest struct {
	Payload *models.Problem
}

func (o *GetInboundShippingLabelBadRequest) Error() string {
	return fmt.Sprintf("[GET /retailer/inbounds/{inbound-id}/shippinglabel][%d] getInboundShippingLabelBadRequest  %+v", 400, o.Payload)
}

func (o *GetInboundShippingLabelBadRequest) GetPayload() *models.Problem {
	return o.Payload
}

func (o *GetInboundShippingLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInboundShippingLabelNotFound creates a GetInboundShippingLabelNotFound with default headers values
func NewGetInboundShippingLabelNotFound() *GetInboundShippingLabelNotFound {
	return &GetInboundShippingLabelNotFound{}
}

/*GetInboundShippingLabelNotFound handles this case with default header values.

Not found: The requested item could not be found.
*/
type GetInboundShippingLabelNotFound struct {
	Payload *models.Problem
}

func (o *GetInboundShippingLabelNotFound) Error() string {
	return fmt.Sprintf("[GET /retailer/inbounds/{inbound-id}/shippinglabel][%d] getInboundShippingLabelNotFound  %+v", 404, o.Payload)
}

func (o *GetInboundShippingLabelNotFound) GetPayload() *models.Problem {
	return o.Payload
}

func (o *GetInboundShippingLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}