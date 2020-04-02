// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tcorp-bv/bol-api-go/models"
)

// CancelOrderReader is a Reader for the CancelOrder structure.
type CancelOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCancelOrderAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelOrderAccepted creates a CancelOrderAccepted with default headers values
func NewCancelOrderAccepted() *CancelOrderAccepted {
	return &CancelOrderAccepted{}
}

/*CancelOrderAccepted handles this case with default header values.

Accepted: Successfully scheduled the request for processing.
*/
type CancelOrderAccepted struct {
	Payload *models.ProcessStatus
}

func (o *CancelOrderAccepted) Error() string {
	return fmt.Sprintf("[PUT /retailer/orders/{order-item-id}/cancellation][%d] cancelOrderAccepted  %+v", 202, o.Payload)
}

func (o *CancelOrderAccepted) GetPayload() *models.ProcessStatus {
	return o.Payload
}

func (o *CancelOrderAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelOrderBadRequest creates a CancelOrderBadRequest with default headers values
func NewCancelOrderBadRequest() *CancelOrderBadRequest {
	return &CancelOrderBadRequest{}
}

/*CancelOrderBadRequest handles this case with default header values.

Bad request: The sent request does not meet the API specification. Please check the error message(s) for more information.
*/
type CancelOrderBadRequest struct {
	Payload *models.Problem
}

func (o *CancelOrderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /retailer/orders/{order-item-id}/cancellation][%d] cancelOrderBadRequest  %+v", 400, o.Payload)
}

func (o *CancelOrderBadRequest) GetPayload() *models.Problem {
	return o.Payload
}

func (o *CancelOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
