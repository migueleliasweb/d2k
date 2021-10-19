// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ContainerUnpauseNoContentCode is the HTTP code returned for type ContainerUnpauseNoContent
const ContainerUnpauseNoContentCode int = 204

/*ContainerUnpauseNoContent no error

swagger:response containerUnpauseNoContent
*/
type ContainerUnpauseNoContent struct {
}

// NewContainerUnpauseNoContent creates ContainerUnpauseNoContent with default headers values
func NewContainerUnpauseNoContent() *ContainerUnpauseNoContent {

	return &ContainerUnpauseNoContent{}
}

// WriteResponse to the client
func (o *ContainerUnpauseNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ContainerUnpauseNotFoundCode is the HTTP code returned for type ContainerUnpauseNotFound
const ContainerUnpauseNotFoundCode int = 404

/*ContainerUnpauseNotFound no such container

swagger:response containerUnpauseNotFound
*/
type ContainerUnpauseNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerUnpauseNotFound creates ContainerUnpauseNotFound with default headers values
func NewContainerUnpauseNotFound() *ContainerUnpauseNotFound {

	return &ContainerUnpauseNotFound{}
}

// WithPayload adds the payload to the container unpause not found response
func (o *ContainerUnpauseNotFound) WithPayload(payload *models.ErrorResponse) *ContainerUnpauseNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container unpause not found response
func (o *ContainerUnpauseNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerUnpauseNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerUnpauseInternalServerErrorCode is the HTTP code returned for type ContainerUnpauseInternalServerError
const ContainerUnpauseInternalServerErrorCode int = 500

/*ContainerUnpauseInternalServerError server error

swagger:response containerUnpauseInternalServerError
*/
type ContainerUnpauseInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerUnpauseInternalServerError creates ContainerUnpauseInternalServerError with default headers values
func NewContainerUnpauseInternalServerError() *ContainerUnpauseInternalServerError {

	return &ContainerUnpauseInternalServerError{}
}

// WithPayload adds the payload to the container unpause internal server error response
func (o *ContainerUnpauseInternalServerError) WithPayload(payload *models.ErrorResponse) *ContainerUnpauseInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container unpause internal server error response
func (o *ContainerUnpauseInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerUnpauseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}