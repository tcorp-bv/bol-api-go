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

// PostInboundReader is a Reader for the PostInbound structure.
type PostInboundReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInboundReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostInboundAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostInboundBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInboundAccepted creates a PostInboundAccepted with default headers values
func NewPostInboundAccepted() *PostInboundAccepted {
	return &PostInboundAccepted{}
}

/*PostInboundAccepted handles this case with default header values.

Accepted: Successfully scheduled the request for processing.
*/
type PostInboundAccepted struct {
	Payload *models.ProcessStatus
}

func (o *PostInboundAccepted) Error() string {
	return fmt.Sprintf("[POST /retailer/inbounds][%d] postInboundAccepted  %+v", 202, o.Payload)
}

func (o *PostInboundAccepted) GetPayload() *models.ProcessStatus {
	return o.Payload
}

func (o *PostInboundAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInboundBadRequest creates a PostInboundBadRequest with default headers values
func NewPostInboundBadRequest() *PostInboundBadRequest {
	return &PostInboundBadRequest{}
}

/*PostInboundBadRequest handles this case with default header values.

Bad request: The sent request does not meet the API specification. Please check the error message(s) for more information.
*/
type PostInboundBadRequest struct {
	Payload *models.Problem
}

func (o *PostInboundBadRequest) Error() string {
	return fmt.Sprintf("[POST /retailer/inbounds][%d] postInboundBadRequest  %+v", 400, o.Payload)
}

func (o *PostInboundBadRequest) GetPayload() *models.Problem {
	return o.Payload
}

func (o *PostInboundBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
