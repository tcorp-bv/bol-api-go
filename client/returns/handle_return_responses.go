// Code generated by go-swagger; DO NOT EDIT.

package returns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tcorp-bv/bol-api-go/models"
)

// HandleReturnReader is a Reader for the HandleReturn structure.
type HandleReturnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleReturnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewHandleReturnAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHandleReturnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHandleReturnAccepted creates a HandleReturnAccepted with default headers values
func NewHandleReturnAccepted() *HandleReturnAccepted {
	return &HandleReturnAccepted{}
}

/*HandleReturnAccepted handles this case with default header values.

Accepted: Successfully scheduled the request for processing.
*/
type HandleReturnAccepted struct {
	Payload *models.ProcessStatus
}

func (o *HandleReturnAccepted) Error() string {
	return fmt.Sprintf("[PUT /retailer/returns/{rma-id}][%d] handleReturnAccepted  %+v", 202, o.Payload)
}

func (o *HandleReturnAccepted) GetPayload() *models.ProcessStatus {
	return o.Payload
}

func (o *HandleReturnAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandleReturnBadRequest creates a HandleReturnBadRequest with default headers values
func NewHandleReturnBadRequest() *HandleReturnBadRequest {
	return &HandleReturnBadRequest{}
}

/*HandleReturnBadRequest handles this case with default header values.

Bad request: The sent request does not meet the API specification. Please check the error message(s) for more information.
*/
type HandleReturnBadRequest struct {
	Payload *models.Problem
}

func (o *HandleReturnBadRequest) Error() string {
	return fmt.Sprintf("[PUT /retailer/returns/{rma-id}][%d] handleReturnBadRequest  %+v", 400, o.Payload)
}

func (o *HandleReturnBadRequest) GetPayload() *models.Problem {
	return o.Payload
}

func (o *HandleReturnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
