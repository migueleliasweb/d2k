// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PluginDeleteHandlerFunc turns a function with the right signature into a plugin delete handler
type PluginDeleteHandlerFunc func(PluginDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PluginDeleteHandlerFunc) Handle(params PluginDeleteParams) middleware.Responder {
	return fn(params)
}

// PluginDeleteHandler interface for that can handle valid plugin delete params
type PluginDeleteHandler interface {
	Handle(PluginDeleteParams) middleware.Responder
}

// NewPluginDelete creates a new http.Handler for the plugin delete operation
func NewPluginDelete(ctx *middleware.Context, handler PluginDeleteHandler) *PluginDelete {
	return &PluginDelete{Context: ctx, Handler: handler}
}

/* PluginDelete swagger:route DELETE /plugins/{name} Plugin pluginDelete

Remove a plugin

*/
type PluginDelete struct {
	Context *middleware.Context
	Handler PluginDeleteHandler
}

func (o *PluginDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPluginDeleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}