// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DetectedBreak Detected start or end of a structural component.
// swagger:model DetectedBreak
type DetectedBreak struct {

	// True if break prepends the element.
	IsPrefix bool `json:"isPrefix,omitempty"`

	// Detected break type.
	Type string `json:"type,omitempty"`
}

// Validate validates this detected break
func (m *DetectedBreak) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DetectedBreak) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetectedBreak) UnmarshalBinary(b []byte) error {
	var res DetectedBreak
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}