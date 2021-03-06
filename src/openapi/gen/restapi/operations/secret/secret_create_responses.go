// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// SecretCreateCreatedCode is the HTTP code returned for type SecretCreateCreated
const SecretCreateCreatedCode int = 201

/*SecretCreateCreated no error

swagger:response secretCreateCreated
*/
type SecretCreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IDResponse `json:"body,omitempty"`
}

// NewSecretCreateCreated creates SecretCreateCreated with default headers values
func NewSecretCreateCreated() *SecretCreateCreated {

	return &SecretCreateCreated{}
}

// WithPayload adds the payload to the secret create created response
func (o *SecretCreateCreated) WithPayload(payload *models.IDResponse) *SecretCreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret create created response
func (o *SecretCreateCreated) SetPayload(payload *models.IDResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretCreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretCreateConflictCode is the HTTP code returned for type SecretCreateConflict
const SecretCreateConflictCode int = 409

/*SecretCreateConflict name conflicts with an existing object

swagger:response secretCreateConflict
*/
type SecretCreateConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretCreateConflict creates SecretCreateConflict with default headers values
func NewSecretCreateConflict() *SecretCreateConflict {

	return &SecretCreateConflict{}
}

// WithPayload adds the payload to the secret create conflict response
func (o *SecretCreateConflict) WithPayload(payload *models.ErrorResponse) *SecretCreateConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret create conflict response
func (o *SecretCreateConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretCreateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretCreateInternalServerErrorCode is the HTTP code returned for type SecretCreateInternalServerError
const SecretCreateInternalServerErrorCode int = 500

/*SecretCreateInternalServerError server error

swagger:response secretCreateInternalServerError
*/
type SecretCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretCreateInternalServerError creates SecretCreateInternalServerError with default headers values
func NewSecretCreateInternalServerError() *SecretCreateInternalServerError {

	return &SecretCreateInternalServerError{}
}

// WithPayload adds the payload to the secret create internal server error response
func (o *SecretCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *SecretCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret create internal server error response
func (o *SecretCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretCreateServiceUnavailableCode is the HTTP code returned for type SecretCreateServiceUnavailable
const SecretCreateServiceUnavailableCode int = 503

/*SecretCreateServiceUnavailable node is not part of a swarm

swagger:response secretCreateServiceUnavailable
*/
type SecretCreateServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretCreateServiceUnavailable creates SecretCreateServiceUnavailable with default headers values
func NewSecretCreateServiceUnavailable() *SecretCreateServiceUnavailable {

	return &SecretCreateServiceUnavailable{}
}

// WithPayload adds the payload to the secret create service unavailable response
func (o *SecretCreateServiceUnavailable) WithPayload(payload *models.ErrorResponse) *SecretCreateServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret create service unavailable response
func (o *SecretCreateServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretCreateServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
