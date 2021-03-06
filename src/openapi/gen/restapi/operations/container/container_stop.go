// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ContainerStopHandlerFunc turns a function with the right signature into a container stop handler
type ContainerStopHandlerFunc func(ContainerStopParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ContainerStopHandlerFunc) Handle(params ContainerStopParams) middleware.Responder {
	return fn(params)
}

// ContainerStopHandler interface for that can handle valid container stop params
type ContainerStopHandler interface {
	Handle(ContainerStopParams) middleware.Responder
}

// NewContainerStop creates a new http.Handler for the container stop operation
func NewContainerStop(ctx *middleware.Context, handler ContainerStopHandler) *ContainerStop {
	return &ContainerStop{Context: ctx, Handler: handler}
}

/* ContainerStop swagger:route POST /containers/{id}/stop Container containerStop

Stop a container

*/
type ContainerStop struct {
	Context *middleware.Context
	Handler ContainerStopHandler
}

func (o *ContainerStop) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewContainerStopParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
