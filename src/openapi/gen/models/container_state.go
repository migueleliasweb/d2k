// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContainerState ContainerState stores container's running state. It's part of ContainerJSONBase
// and will be returned by the "inspect" command.
//
//
// swagger:model ContainerState
type ContainerState struct {

	// dead
	// Example: false
	Dead bool `json:"Dead,omitempty"`

	// error
	Error string `json:"Error,omitempty"`

	// The last exit code of this container
	// Example: 0
	ExitCode int64 `json:"ExitCode,omitempty"`

	// The time when this container last exited.
	// Example: 2020-01-06T09:07:59.461876391Z
	FinishedAt string `json:"FinishedAt,omitempty"`

	// health
	Health *Health `json:"Health,omitempty"`

	// Whether this container has been killed because it ran out of memory.
	//
	// Example: false
	OOMKilled bool `json:"OOMKilled,omitempty"`

	// Whether this container is paused.
	// Example: false
	Paused bool `json:"Paused,omitempty"`

	// The process ID of this container
	// Example: 1234
	Pid int64 `json:"Pid,omitempty"`

	// Whether this container is restarting.
	// Example: false
	Restarting bool `json:"Restarting,omitempty"`

	// Whether this container is running.
	//
	// Note that a running container can be _paused_. The `Running` and `Paused`
	// booleans are not mutually exclusive:
	//
	// When pausing a container (on Linux), the freezer cgroup is used to suspend
	// all processes in the container. Freezing the process requires the process to
	// be running. As a result, paused containers are both `Running` _and_ `Paused`.
	//
	// Use the `Status` field instead to determine if a container's state is "running".
	//
	// Example: true
	Running bool `json:"Running,omitempty"`

	// The time when this container was last started.
	// Example: 2020-01-06T09:06:59.461876391Z
	StartedAt string `json:"StartedAt,omitempty"`

	// String representation of the container state. Can be one of "created",
	// "running", "paused", "restarting", "removing", "exited", or "dead".
	//
	// Example: running
	// Enum: [created running paused restarting removing exited dead]
	Status string `json:"Status,omitempty"`
}

// Validate validates this container state
func (m *ContainerState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerState) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Health")
			}
			return err
		}
	}

	return nil
}

var containerStateTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["created","running","paused","restarting","removing","exited","dead"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerStateTypeStatusPropEnum = append(containerStateTypeStatusPropEnum, v)
	}
}

const (

	// ContainerStateStatusCreated captures enum value "created"
	ContainerStateStatusCreated string = "created"

	// ContainerStateStatusRunning captures enum value "running"
	ContainerStateStatusRunning string = "running"

	// ContainerStateStatusPaused captures enum value "paused"
	ContainerStateStatusPaused string = "paused"

	// ContainerStateStatusRestarting captures enum value "restarting"
	ContainerStateStatusRestarting string = "restarting"

	// ContainerStateStatusRemoving captures enum value "removing"
	ContainerStateStatusRemoving string = "removing"

	// ContainerStateStatusExited captures enum value "exited"
	ContainerStateStatusExited string = "exited"

	// ContainerStateStatusDead captures enum value "dead"
	ContainerStateStatusDead string = "dead"
)

// prop value enum
func (m *ContainerState) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, containerStateTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContainerState) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this container state based on the context it is used
func (m *ContainerState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerState) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {
		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Health")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerState) UnmarshalBinary(b []byte) error {
	var res ContainerState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
