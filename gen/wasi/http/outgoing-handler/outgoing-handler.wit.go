// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package outgoinghandler represents the imported interface "wasi:http/outgoing-handler@0.2.0".
//
// This interface defines a handler of outgoing HTTP Requests. It should be
// imported by components which wish to make HTTP Requests.
package outgoinghandler

import (
	"github.com/LiamRandall/wanban/gen/wasi/http/types"
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// OutgoingRequest represents the imported type alias "wasi:http/outgoing-handler@0.2.0#outgoing-request".
//
// See [types.OutgoingRequest] for more information.
type OutgoingRequest = types.OutgoingRequest

// RequestOptions represents the imported type alias "wasi:http/outgoing-handler@0.2.0#request-options".
//
// See [types.RequestOptions] for more information.
type RequestOptions = types.RequestOptions

// FutureIncomingResponse represents the imported type alias "wasi:http/outgoing-handler@0.2.0#future-incoming-response".
//
// See [types.FutureIncomingResponse] for more information.
type FutureIncomingResponse = types.FutureIncomingResponse

// ErrorCode represents the type alias "wasi:http/outgoing-handler@0.2.0#error-code".
//
// See [types.ErrorCode] for more information.
type ErrorCode = types.ErrorCode

// Handle represents the imported function "handle".
//
// This function is invoked with an outgoing HTTP Request, and it returns
// a resource `future-incoming-response` which represents an HTTP Response
// which may arrive in the future.
//
// The `options` argument accepts optional parameters for the HTTP
// protocol's transport layer.
//
// This function may return an error if the `outgoing-request` is invalid
// or not allowed to be made. Otherwise, protocol errors are reported
// through the `future-incoming-response`.
//
//	handle: func(request: outgoing-request, options: option<request-options>) -> result<future-incoming-response,
//	error-code>
//
//go:nosplit
func Handle(request OutgoingRequest, options cm.Option[RequestOptions]) (result cm.Result[ErrorCodeShape, FutureIncomingResponse, ErrorCode]) {
	request0 := cm.Reinterpret[uint32](request)
	options0, options1 := lower_OptionRequestOptions(options)
	wasmimport_Handle((uint32)(request0), (uint32)(options0), (uint32)(options1), &result)
	return
}
