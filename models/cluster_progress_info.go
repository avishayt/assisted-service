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

// ClusterProgressInfo cluster progress info
//
// swagger:model cluster-progress-info
type ClusterProgressInfo struct {

	// progress info
	ProgressInfo string `json:"progress_info,omitempty" gorm:"type:varchar(2048)"`

	// Time at which the cluster install progress was last updated.
	// Format: date-time
	ProgressUpdatedAt strfmt.DateTime `json:"progress_updated_at,omitempty" gorm:"type:timestamp with time zone"`
}

// Validate validates this cluster progress info
func (m *ClusterProgressInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProgressUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterProgressInfo) validateProgressUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ProgressUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("progress_updated_at", "body", "date-time", m.ProgressUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterProgressInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterProgressInfo) UnmarshalBinary(b []byte) error {
	var res ClusterProgressInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
