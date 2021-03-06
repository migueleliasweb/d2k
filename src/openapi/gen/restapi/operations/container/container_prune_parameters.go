// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewContainerPruneParams creates a new ContainerPruneParams object
//
// There are no default values defined in the spec.
func NewContainerPruneParams() ContainerPruneParams {

	return ContainerPruneParams{}
}

// ContainerPruneParams contains all the bound params for the container prune operation
// typically these are obtained from a http.Request
//
// swagger:parameters ContainerPrune
type ContainerPruneParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Filters to process on the prune list, encoded as JSON (a `map[string][]string`).

	Available filters:
	- `until=<timestamp>` Prune containers created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	- `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune containers with (or without, in case `label!=...` is used) the specified labels.

	  In: query
	*/
	Filters *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewContainerPruneParams() beforehand.
func (o *ContainerPruneParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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
func (o *ContainerPruneParams) bindFilters(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
