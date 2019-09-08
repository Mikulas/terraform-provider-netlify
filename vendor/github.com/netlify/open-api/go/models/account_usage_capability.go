// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AccountUsageCapability account usage capability
// swagger:model accountUsageCapability
type AccountUsageCapability struct {

	// included
	Included int64 `json:"included,omitempty"`

	// used
	Used int64 `json:"used,omitempty"`
}

// Validate validates this account usage capability
func (m *AccountUsageCapability) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountUsageCapability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountUsageCapability) UnmarshalBinary(b []byte) error {
	var res AccountUsageCapability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
