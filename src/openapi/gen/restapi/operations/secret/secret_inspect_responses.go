// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// SecretInspectOKCode is the HTTP code returned for type SecretInspectOK
const SecretInspectOKCode int = 200

/*SecretInspectOK no error

swagger:response secretInspectOK
*/
type SecretInspectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Secret `json:"body,omitempty"`
}

// NewSecretInspectOK creates SecretInspectOK with default headers values
func NewSecretInspectOK() *SecretInspectOK {

	return &SecretInspectOK{}
}

// WithPayload adds the payload to the secret inspect o k response
func (o *SecretInspectOK) WithPayload(payload *models.Secret) *SecretInspectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret inspect o k response
func (o *SecretInspectOK) SetPayload(payload *models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretInspectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretInspectNotFoundCode is the HTTP code returned for type SecretInspectNotFound
const SecretInspectNotFoundCode int = 404

/*SecretInspectNotFound secret not found

swagger:response secretInspectNotFound
*/
type SecretInspectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretInspectNotFound creates SecretInspectNotFound with default headers values
func NewSecretInspectNotFound() *SecretInspectNotFound {

	return &SecretInspectNotFound{}
}

// WithPayload adds the payload to the secret inspect not found response
func (o *SecretInspectNotFound) WithPayload(payload *models.ErrorResponse) *SecretInspectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret inspect not found response
func (o *SecretInspectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretInspectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretInspectInternalServerErrorCode is the HTTP code returned for type SecretInspectInternalServerError
const SecretInspectInternalServerErrorCode int = 500

/*SecretInspectInternalServerError server error

swagger:response secretInspectInternalServerError
*/
type SecretInspectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretInspectInternalServerError creates SecretInspectInternalServerError with default headers values
func NewSecretInspectInternalServerError() *SecretInspectInternalServerError {

	return &SecretInspectInternalServerError{}
}

// WithPayload adds the payload to the secret inspect internal server error response
func (o *SecretInspectInternalServerError) WithPayload(payload *models.ErrorResponse) *SecretInspectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret inspect internal server error response
func (o *SecretInspectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretInspectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretInspectServiceUnavailableCode is the HTTP code returned for type SecretInspectServiceUnavailable
const SecretInspectServiceUnavailableCode int = 503

/*SecretInspectServiceUnavailable node is not part of a swarm

swagger:response secretInspectServiceUnavailable
*/
type SecretInspectServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretInspectServiceUnavailable creates SecretInspectServiceUnavailable with default headers values
func NewSecretInspectServiceUnavailable() *SecretInspectServiceUnavailable {

	return &SecretInspectServiceUnavailable{}
}

// WithPayload adds the payload to the secret inspect service unavailable response
func (o *SecretInspectServiceUnavailable) WithPayload(payload *models.ErrorResponse) *SecretInspectServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret inspect service unavailable response
func (o *SecretInspectServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretInspectServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
