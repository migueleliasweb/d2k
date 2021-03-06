// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ContainerRenameHandlerFunc turns a function with the right signature into a container rename handler
type ContainerRenameHandlerFunc func(ContainerRenameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ContainerRenameHandlerFunc) Handle(params ContainerRenameParams) middleware.Responder {
	return fn(params)
}

// ContainerRenameHandler interface for that can handle valid container rename params
type ContainerRenameHandler interface {
	Handle(ContainerRenameParams) middleware.Responder
}

// NewContainerRename creates a new http.Handler for the container rename operation
func NewContainerRename(ctx *middleware.Context, handler ContainerRenameHandler) *ContainerRename {
	return &ContainerRename{Context: ctx, Handler: handler}
}

/* ContainerRename swagger:route POST /containers/{id}/rename Container containerRename

Rename a container

*/
type ContainerRename struct {
	Context *middleware.Context
	Handler ContainerRenameHandler
}

func (o *ContainerRename) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewContainerRenameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
