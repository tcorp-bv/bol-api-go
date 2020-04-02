// Code generated by go-swagger; DO NOT EDIT.

package commission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tcorp-bv/bol-api-go/models"
)

// GetCommissionReader is a Reader for the GetCommission structure.
type GetCommissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCommissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCommissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCommissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCommissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCommissionOK creates a GetCommissionOK with default headers values
func NewGetCommissionOK() *GetCommissionOK {
	return &GetCommissionOK{}
}

/*GetCommissionOK handles this case with default header values.

Ok: Successfully processed the request.
*/
type GetCommissionOK struct {
	Payload *models.Commission
}

func (o *GetCommissionOK) Error() string {
	return fmt.Sprintf("[GET /retailer/commission/{ean}][%d] getCommissionOK  %+v", 200, o.Payload)
}

func (o *GetCommissionOK) GetPayload() *models.Commission {
	return o.Payload
}

func (o *GetCommissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Commission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommissionBadRequest creates a GetCommissionBadRequest with default headers values
func NewGetCommissionBadRequest() *GetCommissionBadRequest {
	return &GetCommissionBadRequest{}
}

/*GetCommissionBadRequest handles this case with default header values.

Bad request: The sent request does not meet the API specification. Please check the error message(s) for more information.
*/
type GetCommissionBadRequest struct {
	Payload *models.Problem
}

func (o *GetCommissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /retailer/commission/{ean}][%d] getCommissionBadRequest  %+v", 400, o.Payload)
}

func (o *GetCommissionBadRequest) GetPayload() *models.Problem {
	return o.Payload
}

func (o *GetCommissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommissionNotFound creates a GetCommissionNotFound with default headers values
func NewGetCommissionNotFound() *GetCommissionNotFound {
	return &GetCommissionNotFound{}
}

/*GetCommissionNotFound handles this case with default header values.

Not found: The requested item could not be found.
*/
type GetCommissionNotFound struct {
	Payload *models.Problem
}

func (o *GetCommissionNotFound) Error() string {
	return fmt.Sprintf("[GET /retailer/commission/{ean}][%d] getCommissionNotFound  %+v", 404, o.Payload)
}

func (o *GetCommissionNotFound) GetPayload() *models.Problem {
	return o.Payload
}

func (o *GetCommissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}