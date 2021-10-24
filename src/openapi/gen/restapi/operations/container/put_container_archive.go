// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutContainerArchiveHandlerFunc turns a function with the right signature into a put container archive handler
type PutContainerArchiveHandlerFunc func(PutContainerArchiveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutContainerArchiveHandlerFunc) Handle(params PutContainerArchiveParams) middleware.Responder {
	return fn(params)
}

// PutContainerArchiveHandler interface for that can handle valid put container archive params
type PutContainerArchiveHandler interface {
	Handle(PutContainerArchiveParams) middleware.Responder
}

// NewPutContainerArchive creates a new http.Handler for the put container archive operation
func NewPutContainerArchive(ctx *middleware.Context, handler PutContainerArchiveHandler) *PutContainerArchive {
	return &PutContainerArchive{Context: ctx, Handler: handler}
}

/* PutContainerArchive swagger:route PUT /containers/{id}/archive Container putContainerArchive

Extract an archive of files or folders to a directory in a container

Upload a tar archive to be extracted to a path in the filesystem of container id.

*/
type PutContainerArchive struct {
	Context *middleware.Context
	Handler PutContainerArchiveHandler
}

func (o *PutContainerArchive) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutContainerArchiveParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
