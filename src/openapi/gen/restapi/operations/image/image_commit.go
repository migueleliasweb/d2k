// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ImageCommitHandlerFunc turns a function with the right signature into a image commit handler
type ImageCommitHandlerFunc func(ImageCommitParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ImageCommitHandlerFunc) Handle(params ImageCommitParams) middleware.Responder {
	return fn(params)
}

// ImageCommitHandler interface for that can handle valid image commit params
type ImageCommitHandler interface {
	Handle(ImageCommitParams) middleware.Responder
}

// NewImageCommit creates a new http.Handler for the image commit operation
func NewImageCommit(ctx *middleware.Context, handler ImageCommitHandler) *ImageCommit {
	return &ImageCommit{Context: ctx, Handler: handler}
}

/* ImageCommit swagger:route POST /commit Image imageCommit

Create a new image from a container

*/
type ImageCommit struct {
	Context *middleware.Context
	Handler ImageCommitHandler
}

func (o *ImageCommit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewImageCommitParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
