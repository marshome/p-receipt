// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/marshome/p-receipt/api/receipt/server/models"
)

// ReceiptsExtractOKCode is the HTTP code returned for type ReceiptsExtractOK
const ReceiptsExtractOKCode int = 200

/*ReceiptsExtractOK Response

swagger:response receiptsExtractOK
*/
type ReceiptsExtractOK struct {
	/*cors
	  Required: true
	*/
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`

	/*
	  In: Body
	*/
	Payload *models.ReceiptExtractResponse `json:"body,omitempty"`
}

// NewReceiptsExtractOK creates ReceiptsExtractOK with default headers values
func NewReceiptsExtractOK() *ReceiptsExtractOK {
	return &ReceiptsExtractOK{}
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the receipts extract o k response
func (o *ReceiptsExtractOK) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *ReceiptsExtractOK {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the receipts extract o k response
func (o *ReceiptsExtractOK) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WithPayload adds the payload to the receipts extract o k response
func (o *ReceiptsExtractOK) WithPayload(payload *models.ReceiptExtractResponse) *ReceiptsExtractOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receipts extract o k response
func (o *ReceiptsExtractOK) SetPayload(payload *models.ReceiptExtractResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiptsExtractOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReceiptsExtractBadRequestCode is the HTTP code returned for type ReceiptsExtractBadRequest
const ReceiptsExtractBadRequestCode int = 400

/*ReceiptsExtractBadRequest invalid argument

swagger:response receiptsExtractBadRequest
*/
type ReceiptsExtractBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewReceiptsExtractBadRequest creates ReceiptsExtractBadRequest with default headers values
func NewReceiptsExtractBadRequest() *ReceiptsExtractBadRequest {
	return &ReceiptsExtractBadRequest{}
}

// WithPayload adds the payload to the receipts extract bad request response
func (o *ReceiptsExtractBadRequest) WithPayload(payload string) *ReceiptsExtractBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receipts extract bad request response
func (o *ReceiptsExtractBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiptsExtractBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ReceiptsExtractInternalServerErrorCode is the HTTP code returned for type ReceiptsExtractInternalServerError
const ReceiptsExtractInternalServerErrorCode int = 500

/*ReceiptsExtractInternalServerError internal

swagger:response receiptsExtractInternalServerError
*/
type ReceiptsExtractInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewReceiptsExtractInternalServerError creates ReceiptsExtractInternalServerError with default headers values
func NewReceiptsExtractInternalServerError() *ReceiptsExtractInternalServerError {
	return &ReceiptsExtractInternalServerError{}
}

// WithPayload adds the payload to the receipts extract internal server error response
func (o *ReceiptsExtractInternalServerError) WithPayload(payload string) *ReceiptsExtractInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receipts extract internal server error response
func (o *ReceiptsExtractInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiptsExtractInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
