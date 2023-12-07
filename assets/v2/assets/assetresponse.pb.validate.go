// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: datatrails-common-api/assets/v2/assets/assetresponse.proto

package assets

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

// Validate checks the field values on AssetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AssetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AssetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AssetResponseMultiError, or
// nil if none found.
func (m *AssetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AssetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Identity

	{
		sorted_keys := make([]string, len(m.GetAttributes()))
		i := 0
		for key := range m.GetAttributes() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetAttributes()[key]
			_ = val

			// no validation rules for Attributes[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, AssetResponseValidationError{
							field:  fmt.Sprintf("Attributes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, AssetResponseValidationError{
							field:  fmt.Sprintf("Attributes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return AssetResponseValidationError{
						field:  fmt.Sprintf("Attributes[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for ConfirmationStatus

	// no validation rules for Tracked

	{
		sorted_keys := make([]string, len(m.GetAccessPolicy()))
		i := 0
		for key := range m.GetAccessPolicy() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetAccessPolicy()[key]
			_ = val

			// no validation rules for AccessPolicy[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, AssetResponseValidationError{
							field:  fmt.Sprintf("AccessPolicy[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, AssetResponseValidationError{
							field:  fmt.Sprintf("AccessPolicy[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return AssetResponseValidationError{
						field:  fmt.Sprintf("AccessPolicy[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for Owner

	if all {
		switch v := interface{}(m.GetAtTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AssetResponseValidationError{
					field:  "AtTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AssetResponseValidationError{
					field:  "AtTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAtTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AssetResponseValidationError{
				field:  "AtTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ProofMechanism

	// no validation rules for ChainId

	// no validation rules for Public

	// no validation rules for TenantIdentity

	if len(errors) > 0 {
		return AssetResponseMultiError(errors)
	}

	return nil
}

// AssetResponseMultiError is an error wrapping multiple validation errors
// returned by AssetResponse.ValidateAll() if the designated constraints
// aren't met.
type AssetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AssetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AssetResponseMultiError) AllErrors() []error { return m }

// AssetResponseValidationError is the validation error returned by
// AssetResponse.Validate if the designated constraints aren't met.
type AssetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AssetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AssetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AssetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AssetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AssetResponseValidationError) ErrorName() string { return "AssetResponseValidationError" }

// Error satisfies the builtin error interface
func (e AssetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAssetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AssetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AssetResponseValidationError{}
