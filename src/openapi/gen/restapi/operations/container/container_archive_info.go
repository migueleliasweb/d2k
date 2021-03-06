// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// ContainerArchiveInfoHandlerFunc turns a function with the right signature into a container archive info handler
type ContainerArchiveInfoHandlerFunc func(ContainerArchiveInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ContainerArchiveInfoHandlerFunc) Handle(params ContainerArchiveInfoParams) middleware.Responder {
	return fn(params)
}

// ContainerArchiveInfoHandler interface for that can handle valid container archive info params
type ContainerArchiveInfoHandler interface {
	Handle(ContainerArchiveInfoParams) middleware.Responder
}

// NewContainerArchiveInfo creates a new http.Handler for the container archive info operation
func NewContainerArchiveInfo(ctx *middleware.Context, handler ContainerArchiveInfoHandler) *ContainerArchiveInfo {
	return &ContainerArchiveInfo{Context: ctx, Handler: handler}
}

/* ContainerArchiveInfo swagger:route HEAD /containers/{id}/archive Container containerArchiveInfo

Get information about files in a container

A response header `X-Docker-Container-Path-Stat` is returned, containing
a base64 - encoded JSON object with some filesystem header information
about the path.


*/
type ContainerArchiveInfo struct {
	Context *middleware.Context
	Handler ContainerArchiveInfoHandler
}

func (o *ContainerArchiveInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewContainerArchiveInfoParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ContainerArchiveInfoBadRequestBody container archive info bad request body
//
// swagger:model ContainerArchiveInfoBadRequestBody
type ContainerArchiveInfoBadRequestBody struct {
	models.ErrorResponse

	// The error message. Either "must specify path parameter"
	// (path cannot be empty) or "not a directory" (path was
	// asserted to be a directory but exists as a file).
	//
	Message string `json:"message,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ContainerArchiveInfoBadRequestBody) UnmarshalJSON(raw []byte) error {
	// ContainerArchiveInfoBadRequestBodyAO0
	var containerArchiveInfoBadRequestBodyAO0 models.ErrorResponse
	if err := swag.ReadJSON(raw, &containerArchiveInfoBadRequestBodyAO0); err != nil {
		return err
	}
	o.ErrorResponse = containerArchiveInfoBadRequestBodyAO0

	// ContainerArchiveInfoBadRequestBodyAO1
	var dataContainerArchiveInfoBadRequestBodyAO1 struct {
		Message string `json:"message,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataContainerArchiveInfoBadRequestBodyAO1); err != nil {
		return err
	}

	o.Message = dataContainerArchiveInfoBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ContainerArchiveInfoBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	containerArchiveInfoBadRequestBodyAO0, err := swag.WriteJSON(o.ErrorResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, containerArchiveInfoBadRequestBodyAO0)
	var dataContainerArchiveInfoBadRequestBodyAO1 struct {
		Message string `json:"message,omitempty"`
	}

	dataContainerArchiveInfoBadRequestBodyAO1.Message = o.Message

	jsonDataContainerArchiveInfoBadRequestBodyAO1, errContainerArchiveInfoBadRequestBodyAO1 := swag.WriteJSON(dataContainerArchiveInfoBadRequestBodyAO1)
	if errContainerArchiveInfoBadRequestBodyAO1 != nil {
		return nil, errContainerArchiveInfoBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataContainerArchiveInfoBadRequestBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this container archive info bad request body
func (o *ContainerArchiveInfoBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this container archive info bad request body based on the context it is used
func (o *ContainerArchiveInfoBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerArchiveInfoBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerArchiveInfoBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ContainerArchiveInfoBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
