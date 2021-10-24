// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PluginEnableHandlerFunc turns a function with the right signature into a plugin enable handler
type PluginEnableHandlerFunc func(PluginEnableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PluginEnableHandlerFunc) Handle(params PluginEnableParams) middleware.Responder {
	return fn(params)
}

// PluginEnableHandler interface for that can handle valid plugin enable params
type PluginEnableHandler interface {
	Handle(PluginEnableParams) middleware.Responder
}

// NewPluginEnable creates a new http.Handler for the plugin enable operation
func NewPluginEnable(ctx *middleware.Context, handler PluginEnableHandler) *PluginEnable {
	return &PluginEnable{Context: ctx, Handler: handler}
}

/* PluginEnable swagger:route POST /plugins/{name}/enable Plugin pluginEnable

Enable a plugin

*/
type PluginEnable struct {
	Context *middleware.Context
	Handler PluginEnableHandler
}

func (o *PluginEnable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPluginEnableParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
