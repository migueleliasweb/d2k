// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NodeListHandlerFunc turns a function with the right signature into a node list handler
type NodeListHandlerFunc func(NodeListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NodeListHandlerFunc) Handle(params NodeListParams) middleware.Responder {
	return fn(params)
}

// NodeListHandler interface for that can handle valid node list params
type NodeListHandler interface {
	Handle(NodeListParams) middleware.Responder
}

// NewNodeList creates a new http.Handler for the node list operation
func NewNodeList(ctx *middleware.Context, handler NodeListHandler) *NodeList {
	return &NodeList{Context: ctx, Handler: handler}
}

/* NodeList swagger:route GET /nodes Node nodeList

List nodes

*/
type NodeList struct {
	Context *middleware.Context
	Handler NodeListHandler
}

func (o *NodeList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewNodeListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
