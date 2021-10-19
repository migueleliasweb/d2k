// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ImageGetAllHandlerFunc turns a function with the right signature into a image get all handler
type ImageGetAllHandlerFunc func(ImageGetAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ImageGetAllHandlerFunc) Handle(params ImageGetAllParams) middleware.Responder {
	return fn(params)
}

// ImageGetAllHandler interface for that can handle valid image get all params
type ImageGetAllHandler interface {
	Handle(ImageGetAllParams) middleware.Responder
}

// NewImageGetAll creates a new http.Handler for the image get all operation
func NewImageGetAll(ctx *middleware.Context, handler ImageGetAllHandler) *ImageGetAll {
	return &ImageGetAll{Context: ctx, Handler: handler}
}

/* ImageGetAll swagger:route GET /images/get Image imageGetAll

Export several images

Get a tarball containing all images and metadata for several image
repositories.

For each value of the `names` parameter: if it is a specific name and
tag (e.g. `ubuntu:latest`), then only that image (and its parents) are
returned; if it is an image ID, similarly only that image (and its parents)
are returned and there would be no names referenced in the 'repositories'
file for this image ID.

For details on the format, see the [export image endpoint](#operation/ImageGet).


*/
type ImageGetAll struct {
	Context *middleware.Context
	Handler ImageGetAllHandler
}

func (o *ImageGetAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewImageGetAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}