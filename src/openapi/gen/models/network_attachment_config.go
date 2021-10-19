// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkAttachmentConfig Specifies how a service should be attached to a particular network.
//
//
// swagger:model NetworkAttachmentConfig
type NetworkAttachmentConfig struct {

	// Discoverable alternate names for the service on this network.
	//
	Aliases []string `json:"Aliases"`

	// Driver attachment options for the network target.
	//
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// The target network for attachment. Must be a network name or ID.
	//
	Target string `json:"Target,omitempty"`
}

// Validate validates this network attachment config
func (m *NetworkAttachmentConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network attachment config based on context it is used
func (m *NetworkAttachmentConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkAttachmentConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkAttachmentConfig) UnmarshalBinary(b []byte) error {
	var res NetworkAttachmentConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}