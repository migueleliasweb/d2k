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
	"github.com/go-openapi/validate"
)

// ContainerWaitHandlerFunc turns a function with the right signature into a container wait handler
type ContainerWaitHandlerFunc func(ContainerWaitParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ContainerWaitHandlerFunc) Handle(params ContainerWaitParams) middleware.Responder {
	return fn(params)
}

// ContainerWaitHandler interface for that can handle valid container wait params
type ContainerWaitHandler interface {
	Handle(ContainerWaitParams) middleware.Responder
}

// NewContainerWait creates a new http.Handler for the container wait operation
func NewContainerWait(ctx *middleware.Context, handler ContainerWaitHandler) *ContainerWait {
	return &ContainerWait{Context: ctx, Handler: handler}
}

/* ContainerWait swagger:route POST /containers/{id}/wait Container containerWait

Wait for a container

Block until a container stops, then returns the exit code.

*/
type ContainerWait struct {
	Context *middleware.Context
	Handler ContainerWaitHandler
}

func (o *ContainerWait) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewContainerWaitParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ContainerWaitOKBody ContainerWaitResponse
//
// OK response to ContainerWait operation
//
// swagger:model ContainerWaitOKBody
type ContainerWaitOKBody struct {

	// error
	Error *ContainerWaitOKBodyError `json:"Error,omitempty"`

	// Exit code of the container
	// Required: true
	StatusCode int64 `json:"StatusCode"`
}

// Validate validates this container wait o k body
func (o *ContainerWaitOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerWaitOKBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerWaitOK" + "." + "Error")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerWaitOKBody) validateStatusCode(formats strfmt.Registry) error {

	if err := validate.Required("containerWaitOK"+"."+"StatusCode", "body", int64(o.StatusCode)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this container wait o k body based on the context it is used
func (o *ContainerWaitOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerWaitOKBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerWaitOK" + "." + "Error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ContainerWaitOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerWaitOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerWaitOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ContainerWaitOKBodyError container waiting error, if any
//
// swagger:model ContainerWaitOKBodyError
type ContainerWaitOKBodyError struct {

	// Details of an error
	Message string `json:"Message,omitempty"`
}

// Validate validates this container wait o k body error
func (o *ContainerWaitOKBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container wait o k body error based on context it is used
func (o *ContainerWaitOKBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerWaitOKBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerWaitOKBodyError) UnmarshalBinary(b []byte) error {
	var res ContainerWaitOKBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
