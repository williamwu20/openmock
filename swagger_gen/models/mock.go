// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Mock mock
//
// swagger:model Mock
type Mock struct {

	// for behaviors, the actions this mock would do when the expect is met
	Actions []*ActionDispatcher `json:"actions"`

	// expect
	Expect *Expect `json:"expect,omitempty"`

	// for behaviors, makes this behavior extend a specified AbstractBehavior
	Extend string `json:"extend,omitempty"`

	// Unique key for the item in OM's model
	// Pattern: [\w_\-\.]+
	Key string `json:"key,omitempty"`

	// The type of item this is. possible types are: Behavior - creates a new mock behavior  AbstractBehavior - allows behaviors to use common features from this item Template - used in template language rendering to do fancy stuff
	//
	// Enum: [Behavior AbstractBehavior Template]
	Kind string `json:"kind,omitempty"`

	// a go template to be embedded in other templates
	Template string `json:"template,omitempty"`

	// Arbitrary values that can be used in go templates rendered by this item
	Values interface{} `json:"values,omitempty"`
}

// Validate validates this mock
func (m *Mock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mock) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Mock) validateExpect(formats strfmt.Registry) error {

	if swag.IsZero(m.Expect) { // not required
		return nil
	}

	if m.Expect != nil {
		if err := m.Expect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expect")
			}
			return err
		}
	}

	return nil
}

func (m *Mock) validateKey(formats strfmt.Registry) error {

	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := validate.Pattern("key", "body", string(m.Key), `[\w_\-\.]+`); err != nil {
		return err
	}

	return nil
}

var mockTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Behavior","AbstractBehavior","Template"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mockTypeKindPropEnum = append(mockTypeKindPropEnum, v)
	}
}

const (

	// MockKindBehavior captures enum value "Behavior"
	MockKindBehavior string = "Behavior"

	// MockKindAbstractBehavior captures enum value "AbstractBehavior"
	MockKindAbstractBehavior string = "AbstractBehavior"

	// MockKindTemplate captures enum value "Template"
	MockKindTemplate string = "Template"
)

// prop value enum
func (m *Mock) validateKindEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mockTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Mock) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Mock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mock) UnmarshalBinary(b []byte) error {
	var res Mock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
