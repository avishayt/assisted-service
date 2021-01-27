// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StaticIPV6Config static ipv6 config
//
// swagger:model static-ipv6-config
type StaticIPV6Config struct {

	// dns
	DNS string `json:"dns,omitempty"`

	// gateway
	// Pattern: ^(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}$
	Gateway string `json:"gateway,omitempty"`

	// ip
	// Pattern: ^(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}$
	IP string `json:"ip,omitempty"`

	// mask
	// Pattern: ^([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$
	Mask string `json:"mask,omitempty"`
}

// Validate validates this static ipv6 config
func (m *StaticIPV6Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StaticIPV6Config) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if err := validate.Pattern("gateway", "body", string(m.Gateway), `^(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}$`); err != nil {
		return err
	}

	return nil
}

func (m *StaticIPV6Config) validateIP(formats strfmt.Registry) error {

	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if err := validate.Pattern("ip", "body", string(m.IP), `^(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}$`); err != nil {
		return err
	}

	return nil
}

func (m *StaticIPV6Config) validateMask(formats strfmt.Registry) error {

	if swag.IsZero(m.Mask) { // not required
		return nil
	}

	if err := validate.Pattern("mask", "body", string(m.Mask), `^([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StaticIPV6Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StaticIPV6Config) UnmarshalBinary(b []byte) error {
	var res StaticIPV6Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}