// automatically generated by stateify.

package network

import (
	"context"

	"github.com/noisysockets/netstack/pkg/state"
)

func (e *Endpoint) StateTypeName() string {
	return "pkg/tcpip/transport/internal/network.Endpoint"
}

func (e *Endpoint) StateFields() []string {
	return []string{
		"ops",
		"netProto",
		"transProto",
		"waiterQueue",
		"wasBound",
		"owner",
		"writeShutdown",
		"effectiveNetProto",
		"multicastMemberships",
		"ipv4TTL",
		"ipv6HopLimit",
		"multicastTTL",
		"multicastAddr",
		"multicastNICID",
		"ipv4TOS",
		"ipv6TClass",
		"info",
		"state",
	}
}

func (e *Endpoint) beforeSave() {}

// +checklocksignore
func (e *Endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.ops)
	stateSinkObject.Save(1, &e.netProto)
	stateSinkObject.Save(2, &e.transProto)
	stateSinkObject.Save(3, &e.waiterQueue)
	stateSinkObject.Save(4, &e.wasBound)
	stateSinkObject.Save(5, &e.owner)
	stateSinkObject.Save(6, &e.writeShutdown)
	stateSinkObject.Save(7, &e.effectiveNetProto)
	stateSinkObject.Save(8, &e.multicastMemberships)
	stateSinkObject.Save(9, &e.ipv4TTL)
	stateSinkObject.Save(10, &e.ipv6HopLimit)
	stateSinkObject.Save(11, &e.multicastTTL)
	stateSinkObject.Save(12, &e.multicastAddr)
	stateSinkObject.Save(13, &e.multicastNICID)
	stateSinkObject.Save(14, &e.ipv4TOS)
	stateSinkObject.Save(15, &e.ipv6TClass)
	stateSinkObject.Save(16, &e.info)
	stateSinkObject.Save(17, &e.state)
}

func (e *Endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *Endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.ops)
	stateSourceObject.Load(1, &e.netProto)
	stateSourceObject.Load(2, &e.transProto)
	stateSourceObject.Load(3, &e.waiterQueue)
	stateSourceObject.Load(4, &e.wasBound)
	stateSourceObject.Load(5, &e.owner)
	stateSourceObject.Load(6, &e.writeShutdown)
	stateSourceObject.Load(7, &e.effectiveNetProto)
	stateSourceObject.Load(8, &e.multicastMemberships)
	stateSourceObject.Load(9, &e.ipv4TTL)
	stateSourceObject.Load(10, &e.ipv6HopLimit)
	stateSourceObject.Load(11, &e.multicastTTL)
	stateSourceObject.Load(12, &e.multicastAddr)
	stateSourceObject.Load(13, &e.multicastNICID)
	stateSourceObject.Load(14, &e.ipv4TOS)
	stateSourceObject.Load(15, &e.ipv6TClass)
	stateSourceObject.Load(16, &e.info)
	stateSourceObject.Load(17, &e.state)
}

func (m *multicastMembership) StateTypeName() string {
	return "pkg/tcpip/transport/internal/network.multicastMembership"
}

func (m *multicastMembership) StateFields() []string {
	return []string{
		"nicID",
		"multicastAddr",
	}
}

func (m *multicastMembership) beforeSave() {}

// +checklocksignore
func (m *multicastMembership) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.nicID)
	stateSinkObject.Save(1, &m.multicastAddr)
}

func (m *multicastMembership) afterLoad(context.Context) {}

// +checklocksignore
func (m *multicastMembership) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.nicID)
	stateSourceObject.Load(1, &m.multicastAddr)
}

func init() {
	state.Register((*Endpoint)(nil))
	state.Register((*multicastMembership)(nil))
}
