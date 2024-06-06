// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: datatrails-common-api/assets/v2/assets/listevents.proto

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

// Validate checks the field values on ListEventsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListEventsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListEventsRequestMultiError, or nil if none found.
func (m *ListEventsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEventsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPageSize() < 0 {
		err := ListEventsRequestValidationError{
			field:  "PageSize",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PageToken

	// no validation rules for OrderBy

	if !_ListEventsRequest_Uuid_Pattern.MatchString(m.GetUuid()) {
		err := ListEventsRequestValidationError{
			field:  "Uuid",
			reason: "value does not match regex pattern \"^(-|[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ConfirmationStatus

	if all {
		switch v := interface{}(m.GetPrincipalDeclared()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "PrincipalDeclared",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "PrincipalDeclared",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrincipalDeclared()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
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
				errors = append(errors, ListEventsRequestValidationError{
					field:  "PrincipalAccepted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "PrincipalAccepted",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrincipalAccepted()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
				field:  "PrincipalAccepted",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampAcceptedSince()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampAcceptedSince",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampAcceptedSince",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampAcceptedSince()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
				field:  "TimestampAcceptedSince",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampAcceptedBefore()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampAcceptedBefore",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampAcceptedBefore",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampAcceptedBefore()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
				field:  "TimestampAcceptedBefore",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampCommittedSince()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampCommittedSince",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampCommittedSince",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampCommittedSince()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
				field:  "TimestampCommittedSince",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampCommittedBefore()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampCommittedBefore",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampCommittedBefore",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampCommittedBefore()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
				field:  "TimestampCommittedBefore",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampDeclaredSince()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampDeclaredSince",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampDeclaredSince",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampDeclaredSince()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
				field:  "TimestampDeclaredSince",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestampDeclaredBefore()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampDeclaredBefore",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "TimestampDeclaredBefore",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestampDeclaredBefore()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
				field:  "TimestampDeclaredBefore",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Operation

	// no validation rules for Behaviour

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

			if l := utf8.RuneCountInString(key); l < 1 || l > 1024 {
				err := ListEventsRequestValidationError{
					field:  fmt.Sprintf("EventAttributes[%v]", key),
					reason: "value length must be between 1 and 1024 runes, inclusive",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if !_ListEventsRequest_EventAttributes_Pattern.MatchString(key) {
				err := ListEventsRequestValidationError{
					field:  fmt.Sprintf("EventAttributes[%v]", key),
					reason: "value does not match regex pattern \"^[^[:cntrl:]]+$\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			// no validation rules for EventAttributes[key]
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

			if l := utf8.RuneCountInString(key); l < 1 || l > 1024 {
				err := ListEventsRequestValidationError{
					field:  fmt.Sprintf("AssetAttributes[%v]", key),
					reason: "value length must be between 1 and 1024 runes, inclusive",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if !_ListEventsRequest_AssetAttributes_Pattern.MatchString(key) {
				err := ListEventsRequestValidationError{
					field:  fmt.Sprintf("AssetAttributes[%v]", key),
					reason: "value does not match regex pattern \"^[^[:cntrl:]]+$\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			// no validation rules for AssetAttributes[key]
		}
	}

	if all {
		switch v := interface{}(m.GetAsset()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "Asset",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListEventsRequestValidationError{
					field:  "Asset",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAsset()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEventsRequestValidationError{
				field:  "Asset",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := ProofMechanism_name[int32(m.GetProofMechanism())]; !ok {
		err := ListEventsRequestValidationError{
			field:  "ProofMechanism",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListEventsRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListEventsRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListEventsRequestValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.MmrIndex != nil {

		if m.GetMmrIndex() < 0 {
			err := ListEventsRequestValidationError{
				field:  "MmrIndex",
				reason: "value must be greater than or equal to 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.From != nil {

		if !_ListEventsRequest_From_Pattern.MatchString(m.GetFrom()) {
			err := ListEventsRequestValidationError{
				field:  "From",
				reason: "value does not match regex pattern \"^0x[[:xdigit:]]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.MinimumTrust != nil {

		if _, ok := ConfirmationStatus_name[int32(m.GetMinimumTrust())]; !ok {
			err := ListEventsRequestValidationError{
				field:  "MinimumTrust",
				reason: "value must be one of the defined enum values",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.TransactionId != nil {
		// no validation rules for TransactionId
	}

	if len(errors) > 0 {
		return ListEventsRequestMultiError(errors)
	}

	return nil
}

// ListEventsRequestMultiError is an error wrapping multiple validation errors
// returned by ListEventsRequest.ValidateAll() if the designated constraints
// aren't met.
type ListEventsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEventsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEventsRequestMultiError) AllErrors() []error { return m }

// ListEventsRequestValidationError is the validation error returned by
// ListEventsRequest.Validate if the designated constraints aren't met.
type ListEventsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEventsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEventsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEventsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEventsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEventsRequestValidationError) ErrorName() string {
	return "ListEventsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListEventsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEventsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEventsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEventsRequestValidationError{}

var _ListEventsRequest_Uuid_Pattern = regexp.MustCompile("^(-|[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})$")

var _ListEventsRequest_EventAttributes_Pattern = regexp.MustCompile("^[^[:cntrl:]]+$")

var _ListEventsRequest_AssetAttributes_Pattern = regexp.MustCompile("^[^[:cntrl:]]+$")

var _ListEventsRequest_From_Pattern = regexp.MustCompile("^0x[[:xdigit:]]+$")

// Validate checks the field values on ListEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListEventsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListEventsResponseMultiError, or nil if none found.
func (m *ListEventsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEventsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListEventsResponseValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListEventsResponseValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListEventsResponseValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListEventsResponseMultiError(errors)
	}

	return nil
}

// ListEventsResponseMultiError is an error wrapping multiple validation errors
// returned by ListEventsResponse.ValidateAll() if the designated constraints
// aren't met.
type ListEventsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEventsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEventsResponseMultiError) AllErrors() []error { return m }

// ListEventsResponseValidationError is the validation error returned by
// ListEventsResponse.Validate if the designated constraints aren't met.
type ListEventsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEventsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEventsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEventsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEventsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEventsResponseValidationError) ErrorName() string {
	return "ListEventsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListEventsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEventsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEventsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEventsResponseValidationError{}

// Validate checks the field values on ListEventsRequest_AssetQuery with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListEventsRequest_AssetQuery) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEventsRequest_AssetQuery with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListEventsRequest_AssetQueryMultiError, or nil if none found.
func (m *ListEventsRequest_AssetQuery) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEventsRequest_AssetQuery) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

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

			if l := utf8.RuneCountInString(key); l < 1 || l > 1024 {
				err := ListEventsRequest_AssetQueryValidationError{
					field:  fmt.Sprintf("Attributes[%v]", key),
					reason: "value length must be between 1 and 1024 runes, inclusive",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if !_ListEventsRequest_AssetQuery_Attributes_Pattern.MatchString(key) {
				err := ListEventsRequest_AssetQueryValidationError{
					field:  fmt.Sprintf("Attributes[%v]", key),
					reason: "value does not match regex pattern \"^[^[:cntrl:]]+$\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			// no validation rules for Attributes[key]
		}
	}

	if len(errors) > 0 {
		return ListEventsRequest_AssetQueryMultiError(errors)
	}

	return nil
}

// ListEventsRequest_AssetQueryMultiError is an error wrapping multiple
// validation errors returned by ListEventsRequest_AssetQuery.ValidateAll() if
// the designated constraints aren't met.
type ListEventsRequest_AssetQueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEventsRequest_AssetQueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEventsRequest_AssetQueryMultiError) AllErrors() []error { return m }

// ListEventsRequest_AssetQueryValidationError is the validation error returned
// by ListEventsRequest_AssetQuery.Validate if the designated constraints
// aren't met.
type ListEventsRequest_AssetQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEventsRequest_AssetQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEventsRequest_AssetQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEventsRequest_AssetQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEventsRequest_AssetQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEventsRequest_AssetQueryValidationError) ErrorName() string {
	return "ListEventsRequest_AssetQueryValidationError"
}

// Error satisfies the builtin error interface
func (e ListEventsRequest_AssetQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEventsRequest_AssetQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEventsRequest_AssetQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEventsRequest_AssetQueryValidationError{}

var _ListEventsRequest_AssetQuery_Attributes_Pattern = regexp.MustCompile("^[^[:cntrl:]]+$")
