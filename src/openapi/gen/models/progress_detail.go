// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProgressDetail progress detail
//
// swagger:model ProgressDetail
type ProgressDetail struct {

	// current
	Current int64 `json:"current,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this progress detail
func (m *ProgressDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this progress detail based on context it is used
func (m *ProgressDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProgressDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressDetail) UnmarshalBinary(b []byte) error {
	var res ProgressDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
