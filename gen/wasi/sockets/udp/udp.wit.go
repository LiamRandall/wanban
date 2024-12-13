// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package udp represents the imported interface "wasi:sockets/udp@0.2.0".
package udp

import (
	"github.com/LiamRandall/wanban/gen/wasi/io/poll"
	"github.com/LiamRandall/wanban/gen/wasi/sockets/network"
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Pollable represents the imported type alias "wasi:sockets/udp@0.2.0#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// Network represents the imported type alias "wasi:sockets/udp@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// ErrorCode represents the type alias "wasi:sockets/udp@0.2.0#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IPSocketAddress represents the type alias "wasi:sockets/udp@0.2.0#ip-socket-address".
//
// See [network.IPSocketAddress] for more information.
type IPSocketAddress = network.IPSocketAddress

// IPAddressFamily represents the type alias "wasi:sockets/udp@0.2.0#ip-address-family".
//
// See [network.IPAddressFamily] for more information.
type IPAddressFamily = network.IPAddressFamily

// IncomingDatagram represents the record "wasi:sockets/udp@0.2.0#incoming-datagram".
//
//	record incoming-datagram {
//		data: list<u8>,
//		remote-address: ip-socket-address,
//	}
type IncomingDatagram struct {
	_             cm.HostLayout
	Data          cm.List[uint8]
	RemoteAddress IPSocketAddress
}

// OutgoingDatagram represents the record "wasi:sockets/udp@0.2.0#outgoing-datagram".
//
//	record outgoing-datagram {
//		data: list<u8>,
//		remote-address: option<ip-socket-address>,
//	}
type OutgoingDatagram struct {
	_             cm.HostLayout
	Data          cm.List[uint8]
	RemoteAddress cm.Option[IPSocketAddress]
}

// UDPSocket represents the imported resource "wasi:sockets/udp@0.2.0#udp-socket".
//
//	resource udp-socket
type UDPSocket cm.Resource

// ResourceDrop represents the imported resource-drop for resource "udp-socket".
//
// Drops a resource handle.
//
//go:nosplit
func (self UDPSocket) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_UDPSocketResourceDrop((uint32)(self0))
	return
}

// AddressFamily represents the imported method "address-family".
//
//	address-family: func() -> ip-address-family
//
//go:nosplit
func (self UDPSocket) AddressFamily() (result IPAddressFamily) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_UDPSocketAddressFamily((uint32)(self0))
	result = (network.IPAddressFamily)((uint32)(result0))
	return
}

// FinishBind represents the imported method "finish-bind".
//
//	finish-bind: func() -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) FinishBind() (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_UDPSocketFinishBind((uint32)(self0), &result)
	return
}

// LocalAddress represents the imported method "local-address".
//
//	local-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self UDPSocket) LocalAddress() (result cm.Result[IPSocketAddressShape, IPSocketAddress, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_UDPSocketLocalAddress((uint32)(self0), &result)
	return
}

// ReceiveBufferSize represents the imported method "receive-buffer-size".
//
//	receive-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self UDPSocket) ReceiveBufferSize() (result cm.Result[uint64, uint64, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_UDPSocketReceiveBufferSize((uint32)(self0), &result)
	return
}

// RemoteAddress represents the imported method "remote-address".
//
//	remote-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self UDPSocket) RemoteAddress() (result cm.Result[IPSocketAddressShape, IPSocketAddress, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_UDPSocketRemoteAddress((uint32)(self0), &result)
	return
}

// SendBufferSize represents the imported method "send-buffer-size".
//
//	send-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self UDPSocket) SendBufferSize() (result cm.Result[uint64, uint64, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_UDPSocketSendBufferSize((uint32)(self0), &result)
	return
}

// SetReceiveBufferSize represents the imported method "set-receive-buffer-size".
//
//	set-receive-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetReceiveBufferSize(value uint64) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint64)(value)
	wasmimport_UDPSocketSetReceiveBufferSize((uint32)(self0), (uint64)(value0), &result)
	return
}

// SetSendBufferSize represents the imported method "set-send-buffer-size".
//
//	set-send-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetSendBufferSize(value uint64) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint64)(value)
	wasmimport_UDPSocketSetSendBufferSize((uint32)(self0), (uint64)(value0), &result)
	return
}

// SetUnicastHopLimit represents the imported method "set-unicast-hop-limit".
//
//	set-unicast-hop-limit: func(value: u8) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetUnicastHopLimit(value uint8) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	value0 := (uint32)(value)
	wasmimport_UDPSocketSetUnicastHopLimit((uint32)(self0), (uint32)(value0), &result)
	return
}

// StartBind represents the imported method "start-bind".
//
//	start-bind: func(network: borrow<network>, local-address: ip-socket-address) ->
//	result<_, error-code>
//
//go:nosplit
func (self UDPSocket) StartBind(network_ Network, localAddress IPSocketAddress) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	network0 := cm.Reinterpret[uint32](network_)
	localAddress0, localAddress1, localAddress2, localAddress3, localAddress4, localAddress5, localAddress6, localAddress7, localAddress8, localAddress9, localAddress10, localAddress11 := lower_IPSocketAddress(localAddress)
	wasmimport_UDPSocketStartBind((uint32)(self0), (uint32)(network0), (uint32)(localAddress0), (uint32)(localAddress1), (uint32)(localAddress2), (uint32)(localAddress3), (uint32)(localAddress4), (uint32)(localAddress5), (uint32)(localAddress6), (uint32)(localAddress7), (uint32)(localAddress8), (uint32)(localAddress9), (uint32)(localAddress10), (uint32)(localAddress11), &result)
	return
}

// Stream represents the imported method "stream".
//
//	%stream: func(remote-address: option<ip-socket-address>) -> result<tuple<incoming-datagram-stream,
//	outgoing-datagram-stream>, error-code>
//
//go:nosplit
func (self UDPSocket) Stream(remoteAddress cm.Option[IPSocketAddress]) (result cm.Result[TupleIncomingDatagramStreamOutgoingDatagramStreamShape, cm.Tuple[IncomingDatagramStream, OutgoingDatagramStream], ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	remoteAddress0, remoteAddress1, remoteAddress2, remoteAddress3, remoteAddress4, remoteAddress5, remoteAddress6, remoteAddress7, remoteAddress8, remoteAddress9, remoteAddress10, remoteAddress11, remoteAddress12 := lower_OptionIPSocketAddress(remoteAddress)
	wasmimport_UDPSocketStream((uint32)(self0), (uint32)(remoteAddress0), (uint32)(remoteAddress1), (uint32)(remoteAddress2), (uint32)(remoteAddress3), (uint32)(remoteAddress4), (uint32)(remoteAddress5), (uint32)(remoteAddress6), (uint32)(remoteAddress7), (uint32)(remoteAddress8), (uint32)(remoteAddress9), (uint32)(remoteAddress10), (uint32)(remoteAddress11), (uint32)(remoteAddress12), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self UDPSocket) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_UDPSocketSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}

// UnicastHopLimit represents the imported method "unicast-hop-limit".
//
//	unicast-hop-limit: func() -> result<u8, error-code>
//
//go:nosplit
func (self UDPSocket) UnicastHopLimit() (result cm.Result[uint8, uint8, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_UDPSocketUnicastHopLimit((uint32)(self0), &result)
	return
}

// IncomingDatagramStream represents the imported resource "wasi:sockets/udp@0.2.0#incoming-datagram-stream".
//
//	resource incoming-datagram-stream
type IncomingDatagramStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "incoming-datagram-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self IncomingDatagramStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_IncomingDatagramStreamResourceDrop((uint32)(self0))
	return
}

// Receive represents the imported method "receive".
//
//	receive: func(max-results: u64) -> result<list<incoming-datagram>, error-code>
//
//go:nosplit
func (self IncomingDatagramStream) Receive(maxResults uint64) (result cm.Result[cm.List[IncomingDatagram], cm.List[IncomingDatagram], ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	maxResults0 := (uint64)(maxResults)
	wasmimport_IncomingDatagramStreamReceive((uint32)(self0), (uint64)(maxResults0), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self IncomingDatagramStream) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_IncomingDatagramStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}

// OutgoingDatagramStream represents the imported resource "wasi:sockets/udp@0.2.0#outgoing-datagram-stream".
//
//	resource outgoing-datagram-stream
type OutgoingDatagramStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "outgoing-datagram-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutgoingDatagramStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutgoingDatagramStreamResourceDrop((uint32)(self0))
	return
}

// CheckSend represents the imported method "check-send".
//
//	check-send: func() -> result<u64, error-code>
//
//go:nosplit
func (self OutgoingDatagramStream) CheckSend() (result cm.Result[uint64, uint64, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutgoingDatagramStreamCheckSend((uint32)(self0), &result)
	return
}

// Send represents the imported method "send".
//
//	send: func(datagrams: list<outgoing-datagram>) -> result<u64, error-code>
//
//go:nosplit
func (self OutgoingDatagramStream) Send(datagrams cm.List[OutgoingDatagram]) (result cm.Result[uint64, uint64, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	datagrams0, datagrams1 := cm.LowerList(datagrams)
	wasmimport_OutgoingDatagramStreamSend((uint32)(self0), (*OutgoingDatagram)(datagrams0), (uint32)(datagrams1), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self OutgoingDatagramStream) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_OutgoingDatagramStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}
