// automatically generated by stateify.

package multicast

import (
	"context"

	"github.com/noisysockets/netstack/pkg/state"
)

func (r *RouteTable) StateTypeName() string {
	return "pkg/tcpip/network/internal/multicast.RouteTable"
}

func (r *RouteTable) StateFields() []string {
	return []string{
		"installedRoutes",
		"pendingRoutes",
		"cleanupPendingRoutesTimer",
		"isCleanupRoutineRunning",
		"config",
	}
}

func (r *RouteTable) beforeSave() {}

// +checklocksignore
func (r *RouteTable) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.installedRoutes)
	stateSinkObject.Save(1, &r.pendingRoutes)
	stateSinkObject.Save(2, &r.cleanupPendingRoutesTimer)
	stateSinkObject.Save(3, &r.isCleanupRoutineRunning)
	stateSinkObject.Save(4, &r.config)
}

func (r *RouteTable) afterLoad(context.Context) {}

// +checklocksignore
func (r *RouteTable) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.installedRoutes)
	stateSourceObject.Load(1, &r.pendingRoutes)
	stateSourceObject.Load(2, &r.cleanupPendingRoutesTimer)
	stateSourceObject.Load(3, &r.isCleanupRoutineRunning)
	stateSourceObject.Load(4, &r.config)
}

func (r *InstalledRoute) StateTypeName() string {
	return "pkg/tcpip/network/internal/multicast.InstalledRoute"
}

func (r *InstalledRoute) StateFields() []string {
	return []string{
		"MulticastRoute",
		"lastUsedTimestamp",
	}
}

func (r *InstalledRoute) beforeSave() {}

// +checklocksignore
func (r *InstalledRoute) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.MulticastRoute)
	stateSinkObject.Save(1, &r.lastUsedTimestamp)
}

func (r *InstalledRoute) afterLoad(context.Context) {}

// +checklocksignore
func (r *InstalledRoute) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.MulticastRoute)
	stateSourceObject.Load(1, &r.lastUsedTimestamp)
}

func (p *PendingRoute) StateTypeName() string {
	return "pkg/tcpip/network/internal/multicast.PendingRoute"
}

func (p *PendingRoute) StateFields() []string {
	return []string{
		"packets",
		"expiration",
	}
}

func (p *PendingRoute) beforeSave() {}

// +checklocksignore
func (p *PendingRoute) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.packets)
	stateSinkObject.Save(1, &p.expiration)
}

func (p *PendingRoute) afterLoad(context.Context) {}

// +checklocksignore
func (p *PendingRoute) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.packets)
	stateSourceObject.Load(1, &p.expiration)
}

func (c *Config) StateTypeName() string {
	return "pkg/tcpip/network/internal/multicast.Config"
}

func (c *Config) StateFields() []string {
	return []string{
		"MaxPendingQueueSize",
		"Clock",
	}
}

func (c *Config) beforeSave() {}

// +checklocksignore
func (c *Config) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.MaxPendingQueueSize)
	stateSinkObject.Save(1, &c.Clock)
}

func (c *Config) afterLoad(context.Context) {}

// +checklocksignore
func (c *Config) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.MaxPendingQueueSize)
	stateSourceObject.Load(1, &c.Clock)
}

func init() {
	state.Register((*RouteTable)(nil))
	state.Register((*InstalledRoute)(nil))
	state.Register((*PendingRoute)(nil))
	state.Register((*Config)(nil))
}
