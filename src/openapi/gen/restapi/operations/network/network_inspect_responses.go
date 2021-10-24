// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// NetworkInspectOKCode is the HTTP code returned for type NetworkInspectOK
const NetworkInspectOKCode int = 200

/*NetworkInspectOK No error

swagger:response networkInspectOK
*/
type NetworkInspectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Network `json:"body,omitempty"`
}

// NewNetworkInspectOK creates NetworkInspectOK with default headers values
func NewNetworkInspectOK() *NetworkInspectOK {

	return &NetworkInspectOK{}
}

// WithPayload adds the payload to the network inspect o k response
func (o *NetworkInspectOK) WithPayload(payload *models.Network) *NetworkInspectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network inspect o k response
func (o *NetworkInspectOK) SetPayload(payload *models.Network) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkInspectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NetworkInspectNotFoundCode is the HTTP code returned for type NetworkInspectNotFound
const NetworkInspectNotFoundCode int = 404

/*NetworkInspectNotFound Network not found

swagger:response networkInspectNotFound
*/
type NetworkInspectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkInspectNotFound creates NetworkInspectNotFound with default headers values
func NewNetworkInspectNotFound() *NetworkInspectNotFound {

	return &NetworkInspectNotFound{}
}

// WithPayload adds the payload to the network inspect not found response
func (o *NetworkInspectNotFound) WithPayload(payload *models.ErrorResponse) *NetworkInspectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network inspect not found response
func (o *NetworkInspectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkInspectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NetworkInspectInternalServerErrorCode is the HTTP code returned for type NetworkInspectInternalServerError
const NetworkInspectInternalServerErrorCode int = 500

/*NetworkInspectInternalServerError Server error

swagger:response networkInspectInternalServerError
*/
type NetworkInspectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkInspectInternalServerError creates NetworkInspectInternalServerError with default headers values
func NewNetworkInspectInternalServerError() *NetworkInspectInternalServerError {

	return &NetworkInspectInternalServerError{}
}

// WithPayload adds the payload to the network inspect internal server error response
func (o *NetworkInspectInternalServerError) WithPayload(payload *models.ErrorResponse) *NetworkInspectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network inspect internal server error response
func (o *NetworkInspectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkInspectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
