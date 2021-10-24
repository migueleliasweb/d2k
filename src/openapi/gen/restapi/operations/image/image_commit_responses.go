// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ImageCommitCreatedCode is the HTTP code returned for type ImageCommitCreated
const ImageCommitCreatedCode int = 201

/*ImageCommitCreated no error

swagger:response imageCommitCreated
*/
type ImageCommitCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IDResponse `json:"body,omitempty"`
}

// NewImageCommitCreated creates ImageCommitCreated with default headers values
func NewImageCommitCreated() *ImageCommitCreated {

	return &ImageCommitCreated{}
}

// WithPayload adds the payload to the image commit created response
func (o *ImageCommitCreated) WithPayload(payload *models.IDResponse) *ImageCommitCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image commit created response
func (o *ImageCommitCreated) SetPayload(payload *models.IDResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageCommitCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImageCommitNotFoundCode is the HTTP code returned for type ImageCommitNotFound
const ImageCommitNotFoundCode int = 404

/*ImageCommitNotFound no such container

swagger:response imageCommitNotFound
*/
type ImageCommitNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageCommitNotFound creates ImageCommitNotFound with default headers values
func NewImageCommitNotFound() *ImageCommitNotFound {

	return &ImageCommitNotFound{}
}

// WithPayload adds the payload to the image commit not found response
func (o *ImageCommitNotFound) WithPayload(payload *models.ErrorResponse) *ImageCommitNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image commit not found response
func (o *ImageCommitNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageCommitNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImageCommitInternalServerErrorCode is the HTTP code returned for type ImageCommitInternalServerError
const ImageCommitInternalServerErrorCode int = 500

/*ImageCommitInternalServerError server error

swagger:response imageCommitInternalServerError
*/
type ImageCommitInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageCommitInternalServerError creates ImageCommitInternalServerError with default headers values
func NewImageCommitInternalServerError() *ImageCommitInternalServerError {

	return &ImageCommitInternalServerError{}
}

// WithPayload adds the payload to the image commit internal server error response
func (o *ImageCommitInternalServerError) WithPayload(payload *models.ErrorResponse) *ImageCommitInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image commit internal server error response
func (o *ImageCommitInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageCommitInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
