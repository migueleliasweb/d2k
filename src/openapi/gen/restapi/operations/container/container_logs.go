// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ContainerLogsHandlerFunc turns a function with the right signature into a container logs handler
type ContainerLogsHandlerFunc func(ContainerLogsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ContainerLogsHandlerFunc) Handle(params ContainerLogsParams) middleware.Responder {
	return fn(params)
}

// ContainerLogsHandler interface for that can handle valid container logs params
type ContainerLogsHandler interface {
	Handle(ContainerLogsParams) middleware.Responder
}

// NewContainerLogs creates a new http.Handler for the container logs operation
func NewContainerLogs(ctx *middleware.Context, handler ContainerLogsHandler) *ContainerLogs {
	return &ContainerLogs{Context: ctx, Handler: handler}
}

/* ContainerLogs swagger:route GET /containers/{id}/logs Container containerLogs

Get container logs

Get `stdout` and `stderr` logs from a container.

Note: This endpoint works only for containers with the `json-file` or
`journald` logging driver.


*/
type ContainerLogs struct {
	Context *middleware.Context
	Handler ContainerLogsHandler
}

func (o *ContainerLogs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewContainerLogsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
