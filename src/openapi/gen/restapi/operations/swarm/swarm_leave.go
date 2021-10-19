// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SwarmLeaveHandlerFunc turns a function with the right signature into a swarm leave handler
type SwarmLeaveHandlerFunc func(SwarmLeaveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SwarmLeaveHandlerFunc) Handle(params SwarmLeaveParams) middleware.Responder {
	return fn(params)
}

// SwarmLeaveHandler interface for that can handle valid swarm leave params
type SwarmLeaveHandler interface {
	Handle(SwarmLeaveParams) middleware.Responder
}

// NewSwarmLeave creates a new http.Handler for the swarm leave operation
func NewSwarmLeave(ctx *middleware.Context, handler SwarmLeaveHandler) *SwarmLeave {
	return &SwarmLeave{Context: ctx, Handler: handler}
}

/* SwarmLeave swagger:route POST /swarm/leave Swarm swarmLeave

Leave a swarm

*/
type SwarmLeave struct {
	Context *middleware.Context
	Handler SwarmLeaveHandler
}

func (o *SwarmLeave) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSwarmLeaveParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}