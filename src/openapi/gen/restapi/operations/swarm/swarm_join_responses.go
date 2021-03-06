// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// SwarmJoinOKCode is the HTTP code returned for type SwarmJoinOK
const SwarmJoinOKCode int = 200

/*SwarmJoinOK no error

swagger:response swarmJoinOK
*/
type SwarmJoinOK struct {
}

// NewSwarmJoinOK creates SwarmJoinOK with default headers values
func NewSwarmJoinOK() *SwarmJoinOK {

	return &SwarmJoinOK{}
}

// WriteResponse to the client
func (o *SwarmJoinOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SwarmJoinBadRequestCode is the HTTP code returned for type SwarmJoinBadRequest
const SwarmJoinBadRequestCode int = 400

/*SwarmJoinBadRequest bad parameter

swagger:response swarmJoinBadRequest
*/
type SwarmJoinBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSwarmJoinBadRequest creates SwarmJoinBadRequest with default headers values
func NewSwarmJoinBadRequest() *SwarmJoinBadRequest {

	return &SwarmJoinBadRequest{}
}

// WithPayload adds the payload to the swarm join bad request response
func (o *SwarmJoinBadRequest) WithPayload(payload *models.ErrorResponse) *SwarmJoinBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the swarm join bad request response
func (o *SwarmJoinBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SwarmJoinBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SwarmJoinInternalServerErrorCode is the HTTP code returned for type SwarmJoinInternalServerError
const SwarmJoinInternalServerErrorCode int = 500

/*SwarmJoinInternalServerError server error

swagger:response swarmJoinInternalServerError
*/
type SwarmJoinInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSwarmJoinInternalServerError creates SwarmJoinInternalServerError with default headers values
func NewSwarmJoinInternalServerError() *SwarmJoinInternalServerError {

	return &SwarmJoinInternalServerError{}
}

// WithPayload adds the payload to the swarm join internal server error response
func (o *SwarmJoinInternalServerError) WithPayload(payload *models.ErrorResponse) *SwarmJoinInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the swarm join internal server error response
func (o *SwarmJoinInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SwarmJoinInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SwarmJoinServiceUnavailableCode is the HTTP code returned for type SwarmJoinServiceUnavailable
const SwarmJoinServiceUnavailableCode int = 503

/*SwarmJoinServiceUnavailable node is already part of a swarm

swagger:response swarmJoinServiceUnavailable
*/
type SwarmJoinServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSwarmJoinServiceUnavailable creates SwarmJoinServiceUnavailable with default headers values
func NewSwarmJoinServiceUnavailable() *SwarmJoinServiceUnavailable {

	return &SwarmJoinServiceUnavailable{}
}

// WithPayload adds the payload to the swarm join service unavailable response
func (o *SwarmJoinServiceUnavailable) WithPayload(payload *models.ErrorResponse) *SwarmJoinServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the swarm join service unavailable response
func (o *SwarmJoinServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SwarmJoinServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
