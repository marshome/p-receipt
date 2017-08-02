// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Word A word representation.
// swagger:model Word
type Word struct {

	// The bounding box for the word.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `json:"boundingBox,omitempty"`

	// Additional information detected for the word.
	Property *TextProperty `json:"property,omitempty"`

	// List of symbols in the word.
	// The order of the symbols follows the natural reading order.
	Symbols []*Symbol `json:"symbols,omitempty"`
}

// Validate validates this word
func (m *Word) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoundingBox(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProperty(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSymbols(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Word) validateBoundingBox(formats strfmt.Registry) error {

	if swag.IsZero(m.BoundingBox) { // not required
		return nil
	}

	if m.BoundingBox != nil {

		if err := m.BoundingBox.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("boundingBox")
			}
			return err
		}
	}

	return nil
}

func (m *Word) validateProperty(formats strfmt.Registry) error {

	if swag.IsZero(m.Property) { // not required
		return nil
	}

	if m.Property != nil {

		if err := m.Property.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("property")
			}
			return err
		}
	}

	return nil
}

func (m *Word) validateSymbols(formats strfmt.Registry) error {

	if swag.IsZero(m.Symbols) { // not required
		return nil
	}

	for i := 0; i < len(m.Symbols); i++ {

		if swag.IsZero(m.Symbols[i]) { // not required
			continue
		}

		if m.Symbols[i] != nil {

			if err := m.Symbols[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("symbols" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Word) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Word) UnmarshalBinary(b []byte) error {
	var res Word
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
