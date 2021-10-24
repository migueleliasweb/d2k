// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewImagePushParams creates a new ImagePushParams object
//
// There are no default values defined in the spec.
func NewImagePushParams() ImagePushParams {

	return ImagePushParams{}
}

// ImagePushParams contains all the bound params for the image push operation
// typically these are obtained from a http.Request
//
// swagger:parameters ImagePush
type ImagePushParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A base64url-encoded auth configuration.

	Refer to the [authentication section](#section/Authentication) for
	details.

	  Required: true
	  In: header
	*/
	XRegistryAuth string
	/*Image name or ID.
	  Required: true
	  In: path
	*/
	Name string
	/*The tag to associate with the image on the registry.
	  In: query
	*/
	Tag *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewImagePushParams() beforehand.
func (o *ImagePushParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXRegistryAuth(r.Header[http.CanonicalHeaderKey("X-Registry-Auth")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qTag, qhkTag, _ := qs.GetOK("tag")
	if err := o.bindTag(qTag, qhkTag, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRegistryAuth binds and validates parameter XRegistryAuth from header.
func (o *ImagePushParams) bindXRegistryAuth(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Registry-Auth", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Registry-Auth", "header", raw); err != nil {
		return err
	}
	o.XRegistryAuth = raw

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ImagePushParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Name = raw

	return nil
}

// bindTag binds and validates parameter Tag from query.
func (o *ImagePushParams) bindTag(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Tag = &raw

	return nil
}
