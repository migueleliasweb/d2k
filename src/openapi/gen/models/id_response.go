// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IDResponse Response to an API call that returns just an Id
//
// swagger:model IdResponse
type IDResponse struct {

	// The id of the newly created object.
	// Required: true
	ID string `json:"Id"`
}

// Validate validates this Id response
func (m *IDResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDResponse) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("Id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Id response based on context it is used
func (m *IDResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IDResponse) UnmarshalBinary(b []byte) error {
	var res IDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
