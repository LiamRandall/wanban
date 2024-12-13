// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package reveal represents the imported interface "wasmcloud:secrets/reveal@0.1.0-draft".
package reveal

import (
	"github.com/LiamRandall/wanban/gen/wasmcloud/secrets/store"
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Secret represents the imported type alias "wasmcloud:secrets/reveal@0.1.0-draft#secret".
//
// See [store.Secret] for more information.
type Secret = store.Secret

// SecretValue represents the type alias "wasmcloud:secrets/reveal@0.1.0-draft#secret-value".
//
// See [store.SecretValue] for more information.
type SecretValue = store.SecretValue

// Reveal represents the imported function "reveal".
//
//	reveal: func(s: borrow<secret>) -> secret-value
//
//go:nosplit
func Reveal(s Secret) (result SecretValue) {
	s0 := cm.Reinterpret[uint32](s)
	wasmimport_Reveal((uint32)(s0), &result)
	return
}
