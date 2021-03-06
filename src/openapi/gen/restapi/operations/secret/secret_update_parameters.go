// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
)

// NewSecretUpdateParams creates a new SecretUpdateParams object
//
// There are no default values defined in the spec.
func NewSecretUpdateParams() SecretUpdateParams {

	return SecretUpdateParams{}
}

// SecretUpdateParams contains all the bound params for the secret update operation
// typically these are obtained from a http.Request
//
// swagger:parameters SecretUpdate
type SecretUpdateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The spec of the secret to update. Currently, only the Labels field
	can be updated. All other fields must remain unchanged from the
	[SecretInspect endpoint](#operation/SecretInspect) response values.

	  In: body
	*/
	Body *models.SecretSpec
	/*The ID or name of the secret
	  Required: true
	  In: path
	*/
	ID string
	/*The version number of the secret object being updated. This is
	required to avoid conflicting writes.

	  Required: true
	  In: query
	*/
	Version int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSecretUpdateParams() beforehand.
func (o *SecretUpdateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.SecretSpec
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *SecretUpdateParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}

// bindVersion binds and validates parameter Version from query.
func (o *SecretUpdateParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("version", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("version", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("version", "query", "int64", raw)
	}
	o.Version = value

	return nil
}
