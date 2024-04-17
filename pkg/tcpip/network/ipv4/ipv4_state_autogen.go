// automatically generated by stateify.

package ipv4

import (
	"context"

	"github.com/noisysockets/netstack/pkg/state"
)

func (i *icmpv4DestinationUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4DestinationUnreachableSockError"
}

func (i *icmpv4DestinationUnreachableSockError) StateFields() []string {
	return []string{}
}

func (i *icmpv4DestinationUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv4DestinationUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *icmpv4DestinationUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv4DestinationUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
}

func (i *icmpv4DestinationHostUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4DestinationHostUnreachableSockError"
}

func (i *icmpv4DestinationHostUnreachableSockError) StateFields() []string {
	return []string{
		"icmpv4DestinationUnreachableSockError",
	}
}

func (i *icmpv4DestinationHostUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv4DestinationHostUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationHostUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv4DestinationHostUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationNetUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4DestinationNetUnreachableSockError"
}

func (i *icmpv4DestinationNetUnreachableSockError) StateFields() []string {
	return []string{
		"icmpv4DestinationUnreachableSockError",
	}
}

func (i *icmpv4DestinationNetUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv4DestinationNetUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationNetUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv4DestinationNetUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationPortUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4DestinationPortUnreachableSockError"
}

func (i *icmpv4DestinationPortUnreachableSockError) StateFields() []string {
	return []string{
		"icmpv4DestinationUnreachableSockError",
	}
}

func (i *icmpv4DestinationPortUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv4DestinationPortUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationPortUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv4DestinationPortUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationProtoUnreachableSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4DestinationProtoUnreachableSockError"
}

func (i *icmpv4DestinationProtoUnreachableSockError) StateFields() []string {
	return []string{
		"icmpv4DestinationUnreachableSockError",
	}
}

func (i *icmpv4DestinationProtoUnreachableSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv4DestinationProtoUnreachableSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationProtoUnreachableSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv4DestinationProtoUnreachableSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4SourceRouteFailedSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4SourceRouteFailedSockError"
}

func (i *icmpv4SourceRouteFailedSockError) StateFields() []string {
	return []string{
		"icmpv4DestinationUnreachableSockError",
	}
}

func (i *icmpv4SourceRouteFailedSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv4SourceRouteFailedSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4SourceRouteFailedSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv4SourceRouteFailedSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4SourceHostIsolatedSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4SourceHostIsolatedSockError"
}

func (i *icmpv4SourceHostIsolatedSockError) StateFields() []string {
	return []string{
		"icmpv4DestinationUnreachableSockError",
	}
}

func (i *icmpv4SourceHostIsolatedSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv4SourceHostIsolatedSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4SourceHostIsolatedSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv4SourceHostIsolatedSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationHostUnknownSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4DestinationHostUnknownSockError"
}

func (i *icmpv4DestinationHostUnknownSockError) StateFields() []string {
	return []string{
		"icmpv4DestinationUnreachableSockError",
	}
}

func (i *icmpv4DestinationHostUnknownSockError) beforeSave() {}

// +checklocksignore
func (i *icmpv4DestinationHostUnknownSockError) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.icmpv4DestinationUnreachableSockError)
}

func (i *icmpv4DestinationHostUnknownSockError) afterLoad(context.Context) {}

// +checklocksignore
func (i *icmpv4DestinationHostUnknownSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.icmpv4DestinationUnreachableSockError)
}

func (e *icmpv4FragmentationNeededSockError) StateTypeName() string {
	return "pkg/tcpip/network/ipv4.icmpv4FragmentationNeededSockError"
}

func (e *icmpv4FragmentationNeededSockError) StateFields() []string {
	return []string{
		"icmpv4DestinationUnreachableSockError",
		"mtu",
	}
}

func (e *icmpv4FragmentationNeededSockError) beforeSave() {}

// +checklocksignore
func (e *icmpv4FragmentationNeededSockError) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.icmpv4DestinationUnreachableSockError)
	stateSinkObject.Save(1, &e.mtu)
}

func (e *icmpv4FragmentationNeededSockError) afterLoad(context.Context) {}

// +checklocksignore
func (e *icmpv4FragmentationNeededSockError) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.icmpv4DestinationUnreachableSockError)
	stateSourceObject.Load(1, &e.mtu)
}

func init() {
	state.Register((*icmpv4DestinationUnreachableSockError)(nil))
	state.Register((*icmpv4DestinationHostUnreachableSockError)(nil))
	state.Register((*icmpv4DestinationNetUnreachableSockError)(nil))
	state.Register((*icmpv4DestinationPortUnreachableSockError)(nil))
	state.Register((*icmpv4DestinationProtoUnreachableSockError)(nil))
	state.Register((*icmpv4SourceRouteFailedSockError)(nil))
	state.Register((*icmpv4SourceHostIsolatedSockError)(nil))
	state.Register((*icmpv4DestinationHostUnknownSockError)(nil))
	state.Register((*icmpv4FragmentationNeededSockError)(nil))
}