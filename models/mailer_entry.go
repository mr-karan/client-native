// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// MailerEntry Mailer Entry
//
// Mailer entry of a Mailers section
//
// swagger:model mailer_entry
type MailerEntry struct {

	// address
	// Required: true
	// Pattern: ^\S+$
	Address string `json:"address"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_]+$
	Name string `json:"name"`

	// port
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port int64 `json:"port"`
}

// Validate validates this mailer entry
func (m *MailerEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MailerEntry) validateAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("address", "body", m.Address); err != nil {
		return err
	}

	if err := validate.Pattern("address", "body", m.Address, `^\S+$`); err != nil {
		return err
	}

	return nil
}

func (m *MailerEntry) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9-_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *MailerEntry) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", int64(m.Port)); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", m.Port, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", m.Port, 65535, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mailer entry based on context it is used
func (m *MailerEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MailerEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MailerEntry) UnmarshalBinary(b []byte) error {
	var res MailerEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
