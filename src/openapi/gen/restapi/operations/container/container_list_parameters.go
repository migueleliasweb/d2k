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
	"github.com/go-openapi/swag"
)

// NewContainerListParams creates a new ContainerListParams object
// with the default values initialized.
func NewContainerListParams() ContainerListParams {

	var (
		// initialize parameters with default values

		allDefault = bool(false)

		sizeDefault = bool(false)
	)

	return ContainerListParams{
		All: &allDefault,

		Size: &sizeDefault,
	}
}

// ContainerListParams contains all the bound params for the container list operation
// typically these are obtained from a http.Request
//
// swagger:parameters ContainerList
type ContainerListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Return all containers. By default, only running containers are shown.

	  In: query
	  Default: false
	*/
	All *bool
	/*Filters to process on the container list, encoded as JSON (a
	`map[string][]string`). For example, `{"status": ["paused"]}` will
	only return paused containers.

	Available filters:

	- `ancestor`=(`<image-name>[:<tag>]`, `<image id>`, or `<image@digest>`)
	- `before`=(`<container id>` or `<container name>`)
	- `expose`=(`<port>[/<proto>]`|`<startport-endport>/[<proto>]`)
	- `exited=<int>` containers with exit code of `<int>`
	- `health`=(`starting`|`healthy`|`unhealthy`|`none`)
	- `id=<ID>` a container's ID
	- `isolation=`(`default`|`process`|`hyperv`) (Windows daemon only)
	- `is-task=`(`true`|`false`)
	- `label=key` or `label="key=value"` of a container label
	- `name=<name>` a container's name
	- `network`=(`<network id>` or `<network name>`)
	- `publish`=(`<port>[/<proto>]`|`<startport-endport>/[<proto>]`)
	- `since`=(`<container id>` or `<container name>`)
	- `status=`(`created`|`restarting`|`running`|`removing`|`paused`|`exited`|`dead`)
	- `volume`=(`<volume name>` or `<mount point destination>`)

	  In: query
	*/
	Filters *string
	/*Return this number of most recently created containers, including
	non-running ones.

	  In: query
	*/
	Limit *int64
	/*Return the size of container as fields `SizeRw` and `SizeRootFs`.

	  In: query
	  Default: false
	*/
	Size *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewContainerListParams() beforehand.
func (o *ContainerListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAll, qhkAll, _ := qs.GetOK("all")
	if err := o.bindAll(qAll, qhkAll, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilters, qhkFilters, _ := qs.GetOK("filters")
	if err := o.bindFilters(qFilters, qhkFilters, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qSize, qhkSize, _ := qs.GetOK("size")
	if err := o.bindSize(qSize, qhkSize, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAll binds and validates parameter All from query.
func (o *ContainerListParams) bindAll(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewContainerListParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("all", "query", "bool", raw)
	}
	o.All = &value

	return nil
}

// bindFilters binds and validates parameter Filters from query.
func (o *ContainerListParams) bindFilters(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindLimit binds and validates parameter Limit from query.
func (o *ContainerListParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindSize binds and validates parameter Size from query.
func (o *ContainerListParams) bindSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewContainerListParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("size", "query", "bool", raw)
	}
	o.Size = &value

	return nil
}
