// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPutContainerArchiveParams creates a new PutContainerArchiveParams object
//
// There are no default values defined in the spec.
func NewPutContainerArchiveParams() PutContainerArchiveParams {

	return PutContainerArchiveParams{}
}

// PutContainerArchiveParams contains all the bound params for the put container archive operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutContainerArchive
type PutContainerArchiveParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*If `1`, `true`, then it will copy UID/GID maps to the dest file or
	dir

	  In: query
	*/
	CopyUIDGID *string
	/*ID or name of the container
	  Required: true
	  In: path
	*/
	ID string
	/*The input stream must be a tar archive compressed with one of the
	following algorithms: `identity` (no compression), `gzip`, `bzip2`,
	or `xz`.

	  Required: true
	  In: body
	*/
	InputStream io.ReadCloser
	/*If `1`, `true`, or `True` then it will be an error if unpacking the
	given content would cause an existing directory to be replaced with
	a non-directory and vice versa.

	  In: query
	*/
	NoOverwriteDirNonDir *string
	/*Path to a directory in the container to extract the archive’s contents into.
	  Required: true
	  In: query
	*/
	Path string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutContainerArchiveParams() beforehand.
func (o *PutContainerArchiveParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCopyUIDGID, qhkCopyUIDGID, _ := qs.GetOK("copyUIDGID")
	if err := o.bindCopyUIDGID(qCopyUIDGID, qhkCopyUIDGID, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		o.InputStream = r.Body
	} else {
		res = append(res, errors.Required("inputStream", "body", ""))
	}

	qNoOverwriteDirNonDir, qhkNoOverwriteDirNonDir, _ := qs.GetOK("noOverwriteDirNonDir")
	if err := o.bindNoOverwriteDirNonDir(qNoOverwriteDirNonDir, qhkNoOverwriteDirNonDir, route.Formats); err != nil {
		res = append(res, err)
	}

	qPath, qhkPath, _ := qs.GetOK("path")
	if err := o.bindPath(qPath, qhkPath, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCopyUIDGID binds and validates parameter CopyUIDGID from query.
func (o *PutContainerArchiveParams) bindCopyUIDGID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.CopyUIDGID = &raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PutContainerArchiveParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}

// bindNoOverwriteDirNonDir binds and validates parameter NoOverwriteDirNonDir from query.
func (o *PutContainerArchiveParams) bindNoOverwriteDirNonDir(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.NoOverwriteDirNonDir = &raw

	return nil
}

// bindPath binds and validates parameter Path from query.
func (o *PutContainerArchiveParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("path", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("path", "query", raw); err != nil {
		return err
	}
	o.Path = raw

	return nil
}
