package tenancies

import (
	"strings"
)

const (
	// This is incorrect - should be 'tenancies/'
	identityPrefix = "tenant/"
)

// Handle UUID <-> Identity conversion.
// Ideally uuid and intentity would be different types so we don't use the wrong
// one in the wrong place, but protobuf does not seem to support "typedef" of one
// type to another

// IdentityFromUuid derives RRN identity for tenancy from a bare uuid id
func IdentityFromUuid(uuid string) string {
	return identityPrefix + uuid
}

// UuidFromIdentity derives tenancy base uuid id from RRN identity
// When the prefix does not match this will return the original identity unchanged
func UuidFromIdentity(identity string) string {
	return strings.TrimPrefix(identity, identityPrefix)
}
