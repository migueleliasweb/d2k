// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerPruneHandlerFunc turns a function with the right signature into a container prune handler
type ContainerPruneHandlerFunc func(ContainerPruneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ContainerPruneHandlerFunc) Handle(params ContainerPruneParams) middleware.Responder {
	return fn(params)
}

// ContainerPruneHandler interface for that can handle valid container prune params
type ContainerPruneHandler interface {
	Handle(ContainerPruneParams) middleware.Responder
}

// NewContainerPrune creates a new http.Handler for the container prune operation
func NewContainerPrune(ctx *middleware.Context, handler ContainerPruneHandler) *ContainerPrune {
	return &ContainerPrune{Context: ctx, Handler: handler}
}

/* ContainerPrune swagger:route POST /containers/prune Container containerPrune

Delete stopped containers

*/
type ContainerPrune struct {
	Context *middleware.Context
	Handler ContainerPruneHandler
}

func (o *ContainerPrune) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewContainerPruneParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ContainerPruneOKBody ContainerPruneResponse
//
// swagger:model ContainerPruneOKBody
type ContainerPruneOKBody struct {

	// Container IDs that were deleted
	ContainersDeleted []string `json:"ContainersDeleted"`

	// Disk space reclaimed in bytes
	SpaceReclaimed int64 `json:"SpaceReclaimed,omitempty"`
}

// Validate validates this container prune o k body
func (o *ContainerPruneOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container prune o k body based on context it is used
func (o *ContainerPruneOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerPruneOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerPruneOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerPruneOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
