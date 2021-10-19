// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeCreateHandlerFunc turns a function with the right signature into a volume create handler
type VolumeCreateHandlerFunc func(VolumeCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VolumeCreateHandlerFunc) Handle(params VolumeCreateParams) middleware.Responder {
	return fn(params)
}

// VolumeCreateHandler interface for that can handle valid volume create params
type VolumeCreateHandler interface {
	Handle(VolumeCreateParams) middleware.Responder
}

// NewVolumeCreate creates a new http.Handler for the volume create operation
func NewVolumeCreate(ctx *middleware.Context, handler VolumeCreateHandler) *VolumeCreate {
	return &VolumeCreate{Context: ctx, Handler: handler}
}

/* VolumeCreate swagger:route POST /volumes/create Volume volumeCreate

Create a volume

*/
type VolumeCreate struct {
	Context *middleware.Context
	Handler VolumeCreateHandler
}

func (o *VolumeCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVolumeCreateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// VolumeCreateBody VolumeConfig
//
// Volume configuration
// Example: {"Driver":"custom","Labels":{"com.example.some-label":"some-value","com.example.some-other-label":"some-other-value"},"Name":"tardis"}
//
// swagger:model VolumeCreateBody
type VolumeCreateBody struct {

	// Name of the volume driver to use.
	Driver string `json:"Driver,omitempty"`

	// A mapping of driver options and values. These options are
	// passed directly to the driver and are driver specific.
	//
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// The new volume's name. If not specified, Docker generates a name.
	//
	Name string `json:"Name,omitempty"`
}

// Validate validates this volume create body
func (o *VolumeCreateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume create body based on context it is used
func (o *VolumeCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeCreateBody) UnmarshalBinary(b []byte) error {
	var res VolumeCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}