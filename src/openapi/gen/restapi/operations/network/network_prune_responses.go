// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// NetworkPruneOKCode is the HTTP code returned for type NetworkPruneOK
const NetworkPruneOKCode int = 200

/*NetworkPruneOK No error

swagger:response networkPruneOK
*/
type NetworkPruneOK struct {

	/*
	  In: Body
	*/
	Payload *NetworkPruneOKBody `json:"body,omitempty"`
}

// NewNetworkPruneOK creates NetworkPruneOK with default headers values
func NewNetworkPruneOK() *NetworkPruneOK {

	return &NetworkPruneOK{}
}

// WithPayload adds the payload to the network prune o k response
func (o *NetworkPruneOK) WithPayload(payload *NetworkPruneOKBody) *NetworkPruneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network prune o k response
func (o *NetworkPruneOK) SetPayload(payload *NetworkPruneOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkPruneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NetworkPruneInternalServerErrorCode is the HTTP code returned for type NetworkPruneInternalServerError
const NetworkPruneInternalServerErrorCode int = 500

/*NetworkPruneInternalServerError Server error

swagger:response networkPruneInternalServerError
*/
type NetworkPruneInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkPruneInternalServerError creates NetworkPruneInternalServerError with default headers values
func NewNetworkPruneInternalServerError() *NetworkPruneInternalServerError {

	return &NetworkPruneInternalServerError{}
}

// WithPayload adds the payload to the network prune internal server error response
func (o *NetworkPruneInternalServerError) WithPayload(payload *models.ErrorResponse) *NetworkPruneInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network prune internal server error response
func (o *NetworkPruneInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkPruneInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
