// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewNetworkListParams creates a new NetworkListParams object
//
// There are no default values defined in the spec.
func NewNetworkListParams() NetworkListParams {

	return NetworkListParams{}
}

// NetworkListParams contains all the bound params for the network list operation
// typically these are obtained from a http.Request
//
// swagger:parameters NetworkList
type NetworkListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*JSON encoded value of the filters (a `map[string][]string`) to process
	on the networks list.

	Available filters:

	- `dangling=<boolean>` When set to `true` (or `1`), returns all
	   networks that are not in use by a container. When set to `false`
	   (or `0`), only networks that are in use by one or more
	   containers are returned.
	- `driver=<driver-name>` Matches a network's driver.
	- `id=<network-id>` Matches all or part of a network ID.
	- `label=<key>` or `label=<key>=<value>` of a network label.
	- `name=<network-name>` Matches all or part of a network name.
	- `scope=["swarm"|"global"|"local"]` Filters networks by scope (`swarm`, `global`, or `local`).
	- `type=["custom"|"builtin"]` Filters networks by type. The `custom` keyword returns all user-defined networks.

	  In: query
	*/
	Filters *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNetworkListParams() beforehand.
func (o *NetworkListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFilters, qhkFilters, _ := qs.GetOK("filters")
	if err := o.bindFilters(qFilters, qhkFilters, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFilters binds and validates parameter Filters from query.
func (o *NetworkListParams) bindFilters(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Filters = &raw

	return nil
}
