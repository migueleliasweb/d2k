// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SwarmUnlockkeyHandlerFunc turns a function with the right signature into a swarm unlockkey handler
type SwarmUnlockkeyHandlerFunc func(SwarmUnlockkeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SwarmUnlockkeyHandlerFunc) Handle(params SwarmUnlockkeyParams) middleware.Responder {
	return fn(params)
}

// SwarmUnlockkeyHandler interface for that can handle valid swarm unlockkey params
type SwarmUnlockkeyHandler interface {
	Handle(SwarmUnlockkeyParams) middleware.Responder
}

// NewSwarmUnlockkey creates a new http.Handler for the swarm unlockkey operation
func NewSwarmUnlockkey(ctx *middleware.Context, handler SwarmUnlockkeyHandler) *SwarmUnlockkey {
	return &SwarmUnlockkey{Context: ctx, Handler: handler}
}

/* SwarmUnlockkey swagger:route GET /swarm/unlockkey Swarm swarmUnlockkey

Get the unlock key

*/
type SwarmUnlockkey struct {
	Context *middleware.Context
	Handler SwarmUnlockkeyHandler
}

func (o *SwarmUnlockkey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSwarmUnlockkeyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SwarmUnlockkeyOKBody UnlockKeyResponse
// Example: {"UnlockKey":"SWMKEY-1-7c37Cc8654o6p38HnroywCi19pllOnGtbdZEgtKxZu8"}
//
// swagger:model SwarmUnlockkeyOKBody
type SwarmUnlockkeyOKBody struct {

	// The swarm's unlock key.
	UnlockKey string `json:"UnlockKey,omitempty"`
}

// Validate validates this swarm unlockkey o k body
func (o *SwarmUnlockkeyOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this swarm unlockkey o k body based on context it is used
func (o *SwarmUnlockkeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SwarmUnlockkeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SwarmUnlockkeyOKBody) UnmarshalBinary(b []byte) error {
	var res SwarmUnlockkeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
