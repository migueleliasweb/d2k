// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SwarmJoinHandlerFunc turns a function with the right signature into a swarm join handler
type SwarmJoinHandlerFunc func(SwarmJoinParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SwarmJoinHandlerFunc) Handle(params SwarmJoinParams) middleware.Responder {
	return fn(params)
}

// SwarmJoinHandler interface for that can handle valid swarm join params
type SwarmJoinHandler interface {
	Handle(SwarmJoinParams) middleware.Responder
}

// NewSwarmJoin creates a new http.Handler for the swarm join operation
func NewSwarmJoin(ctx *middleware.Context, handler SwarmJoinHandler) *SwarmJoin {
	return &SwarmJoin{Context: ctx, Handler: handler}
}

/* SwarmJoin swagger:route POST /swarm/join Swarm swarmJoin

Join an existing swarm

*/
type SwarmJoin struct {
	Context *middleware.Context
	Handler SwarmJoinHandler
}

func (o *SwarmJoin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSwarmJoinParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SwarmJoinBody swarm join body
// Example: {"AdvertiseAddr":"192.168.1.1:2377","JoinToken":"SWMTKN-1-3pu6hszjas19xyp7ghgosyx9k8atbfcr8p2is99znpy26u2lkl-7p73s1dx5in4tatdymyhg9hu2","ListenAddr":"0.0.0.0:2377","RemoteAddrs":["node1:2377"]}
//
// swagger:model SwarmJoinBody
type SwarmJoinBody struct {

	// Externally reachable address advertised to other nodes. This
	// can either be an address/port combination in the form
	// `192.168.1.1:4567`, or an interface followed by a port number,
	// like `eth0:4567`. If the port number is omitted, the port
	// number from the listen address is used. If `AdvertiseAddr` is
	// not specified, it will be automatically detected when possible.
	//
	AdvertiseAddr string `json:"AdvertiseAddr,omitempty"`

	// Address or interface to use for data path traffic (format:
	// `<ip|interface>`), for example,  `192.168.1.1`, or an interface,
	// like `eth0`. If `DataPathAddr` is unspecified, the same addres
	// as `AdvertiseAddr` is used.
	//
	// The `DataPathAddr` specifies the address that global scope
	// network drivers will publish towards other nodes in order to
	// reach the containers running on this node. Using this parameter
	// it is possible to separate the container data traffic from the
	// management traffic of the cluster.
	//
	DataPathAddr string `json:"DataPathAddr,omitempty"`

	// Secret token for joining this swarm.
	JoinToken string `json:"JoinToken,omitempty"`

	// Listen address used for inter-manager communication if the node
	// gets promoted to manager, as well as determining the networking
	// interface used for the VXLAN Tunnel Endpoint (VTEP).
	//
	ListenAddr string `json:"ListenAddr,omitempty"`

	// Addresses of manager nodes already participating in the swarm.
	//
	RemoteAddrs []string `json:"RemoteAddrs"`
}

// Validate validates this swarm join body
func (o *SwarmJoinBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this swarm join body based on context it is used
func (o *SwarmJoinBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SwarmJoinBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SwarmJoinBody) UnmarshalBinary(b []byte) error {
	var res SwarmJoinBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
