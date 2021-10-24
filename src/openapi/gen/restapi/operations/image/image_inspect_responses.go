// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ImageInspectOKCode is the HTTP code returned for type ImageInspectOK
const ImageInspectOKCode int = 200

/*ImageInspectOK No error

swagger:response imageInspectOK
*/
type ImageInspectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Image `json:"body,omitempty"`
}

// NewImageInspectOK creates ImageInspectOK with default headers values
func NewImageInspectOK() *ImageInspectOK {

	return &ImageInspectOK{}
}

// WithPayload adds the payload to the image inspect o k response
func (o *ImageInspectOK) WithPayload(payload *models.Image) *ImageInspectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image inspect o k response
func (o *ImageInspectOK) SetPayload(payload *models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageInspectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImageInspectNotFoundCode is the HTTP code returned for type ImageInspectNotFound
const ImageInspectNotFoundCode int = 404

/*ImageInspectNotFound No such image

swagger:response imageInspectNotFound
*/
type ImageInspectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageInspectNotFound creates ImageInspectNotFound with default headers values
func NewImageInspectNotFound() *ImageInspectNotFound {

	return &ImageInspectNotFound{}
}

// WithPayload adds the payload to the image inspect not found response
func (o *ImageInspectNotFound) WithPayload(payload *models.ErrorResponse) *ImageInspectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image inspect not found response
func (o *ImageInspectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageInspectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImageInspectInternalServerErrorCode is the HTTP code returned for type ImageInspectInternalServerError
const ImageInspectInternalServerErrorCode int = 500

/*ImageInspectInternalServerError Server error

swagger:response imageInspectInternalServerError
*/
type ImageInspectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageInspectInternalServerError creates ImageInspectInternalServerError with default headers values
func NewImageInspectInternalServerError() *ImageInspectInternalServerError {

	return &ImageInspectInternalServerError{}
}

// WithPayload adds the payload to the image inspect internal server error response
func (o *ImageInspectInternalServerError) WithPayload(payload *models.ErrorResponse) *ImageInspectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image inspect internal server error response
func (o *ImageInspectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageInspectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
