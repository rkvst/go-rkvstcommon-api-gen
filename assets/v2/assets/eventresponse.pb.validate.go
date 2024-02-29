// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: datatrails-common-api/assets/v2/assets/eventresponse.proto

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

// Validate checks the field values on EventResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EventResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EventResponseMultiError, or
// nil if none found.
func (m *EventResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EventResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Identity

	// no validation rules for AssetIdentity

	{
		sorted_keys := make([]string, len(m.GetEventAttributes()))
		i := 0
		for key := range m.GetEventAttributes() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetEventAttributes()[key]
			_ = val

			// no validation rules for EventAttributes[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, EventResponseValidationError{
							field:  fmt.Sprintf("EventAttributes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, EventResponseValidationError{
							field:  fmt.Sprintf("EventAttributes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return EventResponseValidationError{
						field:  fmt.Sprintf("EventAttributes[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	{
		sorted_keys := make([]string, len(m.GetAssetAttributes()))
		i := 0
		for key := range m.GetAssetAttributes() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetAssetAttributes()[key]
			_ = val

			// no validation rules for AssetAttributes[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, EventResponseValidationError{
							field:  fmt.Sprintf("AssetAttributes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, EventResponseValidationError{
							field:  fmt.Sprintf("AssetAttributes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return EventResponseValidationError{
						field:  fmt.Sprintf("AssetAttributes[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for Operation

	// no validation rules for Behaviour

	if all {
		switch v := interface{}(m.GetTimestampDeclared()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "TimestampDeclared",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "TimestampDeclared",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampDeclared()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventResponseValidationError{
				field:  "TimestampDeclared",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampAccepted()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "TimestampAccepted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "TimestampAccepted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampAccepted()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventResponseValidationError{
				field:  "TimestampAccepted",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampCommitted()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "TimestampCommitted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "TimestampCommitted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampCommitted()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventResponseValidationError{
				field:  "TimestampCommitted",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPrincipalDeclared()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "PrincipalDeclared",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "PrincipalDeclared",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrincipalDeclared()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventResponseValidationError{
				field:  "PrincipalDeclared",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPrincipalAccepted()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "PrincipalAccepted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "PrincipalAccepted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrincipalAccepted()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventResponseValidationError{
				field:  "PrincipalAccepted",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ConfirmationStatus

	// no validation rules for TransactionId

	// no validation rules for BlockNumber

	// no validation rules for TransactionIndex

	// no validation rules for From

	// no validation rules for TenantIdentity

	if all {
		switch v := interface{}(m.GetMerklelogEntry()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "MerklelogEntry",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventResponseValidationError{
					field:  "MerklelogEntry",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMerklelogEntry()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventResponseValidationError{
				field:  "MerklelogEntry",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EventResponseMultiError(errors)
	}

	return nil
}

// EventResponseMultiError is an error wrapping multiple validation errors
// returned by EventResponse.ValidateAll() if the designated constraints
// aren't met.
type EventResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventResponseMultiError) AllErrors() []error { return m }

// EventResponseValidationError is the validation error returned by
// EventResponse.Validate if the designated constraints aren't met.
type EventResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventResponseValidationError) ErrorName() string { return "EventResponseValidationError" }

// Error satisfies the builtin error interface
func (e EventResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventResponseValidationError{}
