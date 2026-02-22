// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/Midbound/cloud-sdk-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type IdentityEnriched string         // Always "identity.enriched"
type IdentityQualified string        // Always "identity.qualified"
type IdentityResolved string         // Always "identity.resolved"
type IdentitySessionFinalized string // Always "identity.session.finalized"
type IdentityValidated string        // Always "identity.validated"

func (c IdentityEnriched) Default() IdentityEnriched   { return "identity.enriched" }
func (c IdentityQualified) Default() IdentityQualified { return "identity.qualified" }
func (c IdentityResolved) Default() IdentityResolved   { return "identity.resolved" }
func (c IdentitySessionFinalized) Default() IdentitySessionFinalized {
	return "identity.session.finalized"
}
func (c IdentityValidated) Default() IdentityValidated { return "identity.validated" }

func (c IdentityEnriched) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c IdentityQualified) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c IdentityResolved) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c IdentitySessionFinalized) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c IdentityValidated) MarshalJSON() ([]byte, error)        { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
