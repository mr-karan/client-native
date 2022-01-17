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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TCPCheck TCP Check
//
// swagger:model tcp_check
type TCPCheck struct {

	// action
	// Required: true
	// Enum: [comment connect expect send send-lf send-binary send-binary-lf set-var set-var-fmt unset-var]
	Action string `json:"action"`

	// addr
	// Pattern: ^[^\s]+$
	Addr string `json:"addr,omitempty"`

	// alpn
	// Pattern: ^[^\s]+$
	Alpn string `json:"alpn,omitempty"`

	// check comment
	CheckComment string `json:"check_comment,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// error status
	// Enum: [L7OKC L7RSP L7STS L6RSP L4CON]
	ErrorStatus string `json:"error_status,omitempty"`

	// exclamation mark
	ExclamationMark bool `json:"exclamation_mark,omitempty"`

	// fmt
	Fmt string `json:"fmt,omitempty"`

	// hex fmt
	HexFmt string `json:"hex_fmt,omitempty"`

	// hex string
	HexString string `json:"hex_string,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// linger
	Linger bool `json:"linger,omitempty"`

	// log message
	LogMessage string `json:"log_message,omitempty"`

	// match
	// Pattern: ^[^\s]+$
	// Enum: [string rstring string-lf binary rbinary binary-lf]
	Match string `json:"match,omitempty"`

	// min recv
	MinRecv int64 `json:"min_recv,omitempty"`

	// ok status
	// Enum: [L7OK L7OKC L6OK L4OK]
	OkStatus string `json:"ok_status,omitempty"`

	// on error
	OnError string `json:"on_error,omitempty"`

	// on success
	OnSuccess string `json:"on_success,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// port
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`

	// port string
	PortString string `json:"port_string,omitempty"`

	// proto
	Proto string `json:"proto,omitempty"`

	// send proxy
	SendProxy bool `json:"send_proxy,omitempty"`

	// sni
	Sni string `json:"sni,omitempty"`

	// ssl
	Ssl bool `json:"ssl,omitempty"`

	// status code
	StatusCode string `json:"status-code,omitempty"`

	// tout status
	// Enum: [L7TOUT L6TOUT L4TOUT]
	ToutStatus string `json:"tout_status,omitempty"`

	// var expr
	VarExpr string `json:"var_expr,omitempty"`

	// var fmt
	VarFmt string `json:"var_fmt,omitempty"`

	// var name
	// Pattern: ^[^\s]+$
	VarName string `json:"var_name,omitempty"`

	// var scope
	// Pattern: ^[^\s]+$
	VarScope string `json:"var_scope,omitempty"`

	// via socks4
	ViaSocks4 bool `json:"via_socks4,omitempty"`
}

// Validate validates this tcp check
func (m *TCPCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlpn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOkStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToutStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tcpCheckTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["comment","connect","expect","send","send-lf","send-binary","send-binary-lf","set-var","set-var-fmt","unset-var"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpCheckTypeActionPropEnum = append(tcpCheckTypeActionPropEnum, v)
	}
}

const (

	// TCPCheckActionComment captures enum value "comment"
	TCPCheckActionComment string = "comment"

	// TCPCheckActionConnect captures enum value "connect"
	TCPCheckActionConnect string = "connect"

	// TCPCheckActionExpect captures enum value "expect"
	TCPCheckActionExpect string = "expect"

	// TCPCheckActionSend captures enum value "send"
	TCPCheckActionSend string = "send"

	// TCPCheckActionSendLf captures enum value "send-lf"
	TCPCheckActionSendLf string = "send-lf"

	// TCPCheckActionSendBinary captures enum value "send-binary"
	TCPCheckActionSendBinary string = "send-binary"

	// TCPCheckActionSendBinaryLf captures enum value "send-binary-lf"
	TCPCheckActionSendBinaryLf string = "send-binary-lf"

	// TCPCheckActionSetVar captures enum value "set-var"
	TCPCheckActionSetVar string = "set-var"

	// TCPCheckActionSetVarFmt captures enum value "set-var-fmt"
	TCPCheckActionSetVarFmt string = "set-var-fmt"

	// TCPCheckActionUnsetVar captures enum value "unset-var"
	TCPCheckActionUnsetVar string = "unset-var"
)

// prop value enum
func (m *TCPCheck) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tcpCheckTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TCPCheck) validateAction(formats strfmt.Registry) error {

	if err := validate.RequiredString("action", "body", string(m.Action)); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *TCPCheck) validateAddr(formats strfmt.Registry) error {

	if swag.IsZero(m.Addr) { // not required
		return nil
	}

	if err := validate.Pattern("addr", "body", string(m.Addr), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *TCPCheck) validateAlpn(formats strfmt.Registry) error {

	if swag.IsZero(m.Alpn) { // not required
		return nil
	}

	if err := validate.Pattern("alpn", "body", string(m.Alpn), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var tcpCheckTypeErrorStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L7OKC","L7RSP","L7STS","L6RSP","L4CON"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpCheckTypeErrorStatusPropEnum = append(tcpCheckTypeErrorStatusPropEnum, v)
	}
}

const (

	// TCPCheckErrorStatusL7OKC captures enum value "L7OKC"
	TCPCheckErrorStatusL7OKC string = "L7OKC"

	// TCPCheckErrorStatusL7RSP captures enum value "L7RSP"
	TCPCheckErrorStatusL7RSP string = "L7RSP"

	// TCPCheckErrorStatusL7STS captures enum value "L7STS"
	TCPCheckErrorStatusL7STS string = "L7STS"

	// TCPCheckErrorStatusL6RSP captures enum value "L6RSP"
	TCPCheckErrorStatusL6RSP string = "L6RSP"

	// TCPCheckErrorStatusL4CON captures enum value "L4CON"
	TCPCheckErrorStatusL4CON string = "L4CON"
)

// prop value enum
func (m *TCPCheck) validateErrorStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tcpCheckTypeErrorStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TCPCheck) validateErrorStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateErrorStatusEnum("error_status", "body", m.ErrorStatus); err != nil {
		return err
	}

	return nil
}

func (m *TCPCheck) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var tcpCheckTypeMatchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","rstring","string-lf","binary","rbinary","binary-lf"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpCheckTypeMatchPropEnum = append(tcpCheckTypeMatchPropEnum, v)
	}
}

const (

	// TCPCheckMatchString captures enum value "string"
	TCPCheckMatchString string = "string"

	// TCPCheckMatchRstring captures enum value "rstring"
	TCPCheckMatchRstring string = "rstring"

	// TCPCheckMatchStringLf captures enum value "string-lf"
	TCPCheckMatchStringLf string = "string-lf"

	// TCPCheckMatchBinary captures enum value "binary"
	TCPCheckMatchBinary string = "binary"

	// TCPCheckMatchRbinary captures enum value "rbinary"
	TCPCheckMatchRbinary string = "rbinary"

	// TCPCheckMatchBinaryLf captures enum value "binary-lf"
	TCPCheckMatchBinaryLf string = "binary-lf"
)

// prop value enum
func (m *TCPCheck) validateMatchEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tcpCheckTypeMatchPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TCPCheck) validateMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if err := validate.Pattern("match", "body", string(m.Match), `^[^\s]+$`); err != nil {
		return err
	}

	// value enum
	if err := m.validateMatchEnum("match", "body", m.Match); err != nil {
		return err
	}

	return nil
}

var tcpCheckTypeOkStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L7OK","L7OKC","L6OK","L4OK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpCheckTypeOkStatusPropEnum = append(tcpCheckTypeOkStatusPropEnum, v)
	}
}

const (

	// TCPCheckOkStatusL7OK captures enum value "L7OK"
	TCPCheckOkStatusL7OK string = "L7OK"

	// TCPCheckOkStatusL7OKC captures enum value "L7OKC"
	TCPCheckOkStatusL7OKC string = "L7OKC"

	// TCPCheckOkStatusL6OK captures enum value "L6OK"
	TCPCheckOkStatusL6OK string = "L6OK"

	// TCPCheckOkStatusL4OK captures enum value "L4OK"
	TCPCheckOkStatusL4OK string = "L4OK"
)

// prop value enum
func (m *TCPCheck) validateOkStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tcpCheckTypeOkStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TCPCheck) validateOkStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.OkStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateOkStatusEnum("ok_status", "body", m.OkStatus); err != nil {
		return err
	}

	return nil
}

func (m *TCPCheck) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var tcpCheckTypeToutStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["L7TOUT","L6TOUT","L4TOUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tcpCheckTypeToutStatusPropEnum = append(tcpCheckTypeToutStatusPropEnum, v)
	}
}

const (

	// TCPCheckToutStatusL7TOUT captures enum value "L7TOUT"
	TCPCheckToutStatusL7TOUT string = "L7TOUT"

	// TCPCheckToutStatusL6TOUT captures enum value "L6TOUT"
	TCPCheckToutStatusL6TOUT string = "L6TOUT"

	// TCPCheckToutStatusL4TOUT captures enum value "L4TOUT"
	TCPCheckToutStatusL4TOUT string = "L4TOUT"
)

// prop value enum
func (m *TCPCheck) validateToutStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tcpCheckTypeToutStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TCPCheck) validateToutStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ToutStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateToutStatusEnum("tout_status", "body", m.ToutStatus); err != nil {
		return err
	}

	return nil
}

func (m *TCPCheck) validateVarName(formats strfmt.Registry) error {

	if swag.IsZero(m.VarName) { // not required
		return nil
	}

	if err := validate.Pattern("var_name", "body", string(m.VarName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *TCPCheck) validateVarScope(formats strfmt.Registry) error {

	if swag.IsZero(m.VarScope) { // not required
		return nil
	}

	if err := validate.Pattern("var_scope", "body", string(m.VarScope), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TCPCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TCPCheck) UnmarshalBinary(b []byte) error {
	var res TCPCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
