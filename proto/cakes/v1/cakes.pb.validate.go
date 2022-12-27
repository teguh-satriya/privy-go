// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cakes/v1/cakes.proto

package cakesv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateCakeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateCakeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCakeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCakeRequestMultiError, or nil if none found.
func (m *CreateCakeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCakeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTitle()); l < 3 || l > 150 {
		err := CreateCakeRequestValidationError{
			field:  "Title",
			reason: "value length must be between 3 and 150 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	if uri, err := url.Parse(m.GetImage()); err != nil {
		err = CreateCakeRequestValidationError{
			field:  "Image",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := CreateCakeRequestValidationError{
			field:  "Image",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateCakeRequestMultiError(errors)
	}

	return nil
}

// CreateCakeRequestMultiError is an error wrapping multiple validation errors
// returned by CreateCakeRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateCakeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCakeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCakeRequestMultiError) AllErrors() []error { return m }

// CreateCakeRequestValidationError is the validation error returned by
// CreateCakeRequest.Validate if the designated constraints aren't met.
type CreateCakeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCakeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCakeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCakeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCakeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCakeRequestValidationError) ErrorName() string {
	return "CreateCakeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCakeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCakeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCakeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCakeRequestValidationError{}

// Validate checks the field values on CreateCakeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCakeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCakeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCakeResponseMultiError, or nil if none found.
func (m *CreateCakeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCakeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCakeResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCakeResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCakeResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateCakeResponseMultiError(errors)
	}

	return nil
}

// CreateCakeResponseMultiError is an error wrapping multiple validation errors
// returned by CreateCakeResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateCakeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCakeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCakeResponseMultiError) AllErrors() []error { return m }

// CreateCakeResponseValidationError is the validation error returned by
// CreateCakeResponse.Validate if the designated constraints aren't met.
type CreateCakeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCakeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCakeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCakeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCakeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCakeResponseValidationError) ErrorName() string {
	return "CreateCakeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCakeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCakeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCakeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCakeResponseValidationError{}

// Validate checks the field values on GetCakeRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCakeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCakeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCakeRequestMultiError,
// or nil if none found.
func (m *GetCakeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCakeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetCakeRequestMultiError(errors)
	}

	return nil
}

// GetCakeRequestMultiError is an error wrapping multiple validation errors
// returned by GetCakeRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCakeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCakeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCakeRequestMultiError) AllErrors() []error { return m }

// GetCakeRequestValidationError is the validation error returned by
// GetCakeRequest.Validate if the designated constraints aren't met.
type GetCakeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCakeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCakeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCakeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCakeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCakeRequestValidationError) ErrorName() string { return "GetCakeRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetCakeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCakeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCakeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCakeRequestValidationError{}

// Validate checks the field values on GetCakeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCakeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCakeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCakeResponseMultiError, or nil if none found.
func (m *GetCakeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCakeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetCakeResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetCakeResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCakeResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetCakeResponseMultiError(errors)
	}

	return nil
}

// GetCakeResponseMultiError is an error wrapping multiple validation errors
// returned by GetCakeResponse.ValidateAll() if the designated constraints
// aren't met.
type GetCakeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCakeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCakeResponseMultiError) AllErrors() []error { return m }

// GetCakeResponseValidationError is the validation error returned by
// GetCakeResponse.Validate if the designated constraints aren't met.
type GetCakeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCakeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCakeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCakeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCakeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCakeResponseValidationError) ErrorName() string { return "GetCakeResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetCakeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCakeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCakeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCakeResponseValidationError{}

// Validate checks the field values on ListCakesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCakesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCakesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCakesRequestMultiError, or nil if none found.
func (m *ListCakesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCakesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListCakesRequestMultiError(errors)
	}

	return nil
}

// ListCakesRequestMultiError is an error wrapping multiple validation errors
// returned by ListCakesRequest.ValidateAll() if the designated constraints
// aren't met.
type ListCakesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCakesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCakesRequestMultiError) AllErrors() []error { return m }

// ListCakesRequestValidationError is the validation error returned by
// ListCakesRequest.Validate if the designated constraints aren't met.
type ListCakesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCakesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCakesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCakesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCakesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCakesRequestValidationError) ErrorName() string { return "ListCakesRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListCakesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCakesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCakesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCakesRequestValidationError{}

// Validate checks the field values on ListCakesResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCakesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCakesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCakesResponseMultiError, or nil if none found.
func (m *ListCakesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCakesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCakesResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCakesResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCakesResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListCakesResponseMultiError(errors)
	}

	return nil
}

// ListCakesResponseMultiError is an error wrapping multiple validation errors
// returned by ListCakesResponse.ValidateAll() if the designated constraints
// aren't met.
type ListCakesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCakesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCakesResponseMultiError) AllErrors() []error { return m }

// ListCakesResponseValidationError is the validation error returned by
// ListCakesResponse.Validate if the designated constraints aren't met.
type ListCakesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCakesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCakesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCakesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCakesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCakesResponseValidationError) ErrorName() string {
	return "ListCakesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCakesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCakesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCakesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCakesResponseValidationError{}

// Validate checks the field values on UpdateCakeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateCakeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCakeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCakeRequestMultiError, or nil if none found.
func (m *UpdateCakeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCakeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) > 150 {
		err := UpdateCakeRequestValidationError{
			field:  "Title",
			reason: "value length must be at most 150 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	// no validation rules for Rating

	if m.GetImage() != "" {

		if uri, err := url.Parse(m.GetImage()); err != nil {
			err = UpdateCakeRequestValidationError{
				field:  "Image",
				reason: "value must be a valid URI",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else if !uri.IsAbs() {
			err := UpdateCakeRequestValidationError{
				field:  "Image",
				reason: "value must be absolute",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UpdateCakeRequestMultiError(errors)
	}

	return nil
}

// UpdateCakeRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateCakeRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateCakeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCakeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCakeRequestMultiError) AllErrors() []error { return m }

// UpdateCakeRequestValidationError is the validation error returned by
// UpdateCakeRequest.Validate if the designated constraints aren't met.
type UpdateCakeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCakeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCakeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCakeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCakeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCakeRequestValidationError) ErrorName() string {
	return "UpdateCakeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCakeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCakeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCakeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCakeRequestValidationError{}

// Validate checks the field values on UpdateCakeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCakeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCakeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCakeResponseMultiError, or nil if none found.
func (m *UpdateCakeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCakeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateCakeResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateCakeResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCakeResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateCakeResponseMultiError(errors)
	}

	return nil
}

// UpdateCakeResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateCakeResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateCakeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCakeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCakeResponseMultiError) AllErrors() []error { return m }

// UpdateCakeResponseValidationError is the validation error returned by
// UpdateCakeResponse.Validate if the designated constraints aren't met.
type UpdateCakeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCakeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCakeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCakeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCakeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCakeResponseValidationError) ErrorName() string {
	return "UpdateCakeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCakeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCakeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCakeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCakeResponseValidationError{}

// Validate checks the field values on DeleteCakeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteCakeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCakeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCakeRequestMultiError, or nil if none found.
func (m *DeleteCakeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCakeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteCakeRequestMultiError(errors)
	}

	return nil
}

// DeleteCakeRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteCakeRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteCakeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCakeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCakeRequestMultiError) AllErrors() []error { return m }

// DeleteCakeRequestValidationError is the validation error returned by
// DeleteCakeRequest.Validate if the designated constraints aren't met.
type DeleteCakeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCakeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCakeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCakeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCakeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCakeRequestValidationError) ErrorName() string {
	return "DeleteCakeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCakeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCakeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCakeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCakeRequestValidationError{}

// Validate checks the field values on DeleteCakeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCakeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCakeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCakeResponseMultiError, or nil if none found.
func (m *DeleteCakeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCakeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteCakeResponseMultiError(errors)
	}

	return nil
}

// DeleteCakeResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteCakeResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteCakeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCakeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCakeResponseMultiError) AllErrors() []error { return m }

// DeleteCakeResponseValidationError is the validation error returned by
// DeleteCakeResponse.Validate if the designated constraints aren't met.
type DeleteCakeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCakeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCakeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCakeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCakeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCakeResponseValidationError) ErrorName() string {
	return "DeleteCakeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCakeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCakeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCakeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCakeResponseValidationError{}

// Validate checks the field values on Cake with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Cake) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cake with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CakeMultiError, or nil if none found.
func (m *Cake) ValidateAll() error {
	return m.validate(true)
}

func (m *Cake) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for Rating

	// no validation rules for Image

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CakeValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CakeValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CakeValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CakeValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CakeValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CakeValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CakeMultiError(errors)
	}

	return nil
}

// CakeMultiError is an error wrapping multiple validation errors returned by
// Cake.ValidateAll() if the designated constraints aren't met.
type CakeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CakeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CakeMultiError) AllErrors() []error { return m }

// CakeValidationError is the validation error returned by Cake.Validate if the
// designated constraints aren't met.
type CakeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CakeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CakeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CakeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CakeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CakeValidationError) ErrorName() string { return "CakeValidationError" }

// Error satisfies the builtin error interface
func (e CakeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCake.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CakeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CakeValidationError{}