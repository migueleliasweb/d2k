// Code generated by go-swagger; DO NOT EDIT.

package exec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExecStartHandlerFunc turns a function with the right signature into a exec start handler
type ExecStartHandlerFunc func(ExecStartParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ExecStartHandlerFunc) Handle(params ExecStartParams) middleware.Responder {
	return fn(params)
}

// ExecStartHandler interface for that can handle valid exec start params
type ExecStartHandler interface {
	Handle(ExecStartParams) middleware.Responder
}

// NewExecStart creates a new http.Handler for the exec start operation
func NewExecStart(ctx *middleware.Context, handler ExecStartHandler) *ExecStart {
	return &ExecStart{Context: ctx, Handler: handler}
}

/* ExecStart swagger:route POST /exec/{id}/start Exec execStart

Start an exec instance

Starts a previously set up exec instance. If detach is true, this endpoint
returns immediately after starting the command. Otherwise, it sets up an
interactive session with the command.


*/
type ExecStart struct {
	Context *middleware.Context
	Handler ExecStartHandler
}

func (o *ExecStart) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewExecStartParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ExecStartBody exec start body
// Example: {"Detach":false,"Tty":false}
//
// swagger:model ExecStartBody
type ExecStartBody struct {

	// Detach from the command.
	Detach bool `json:"Detach,omitempty"`

	// Allocate a pseudo-TTY.
	Tty bool `json:"Tty,omitempty"`
}

// Validate validates this exec start body
func (o *ExecStartBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exec start body based on context it is used
func (o *ExecStartBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExecStartBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExecStartBody) UnmarshalBinary(b []byte) error {
	var res ExecStartBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}