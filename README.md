# Netstack

An extract of the netstack, userspace TCP/IP stack, from the [gVisor](https://github.com/google/gvisor) project.

By default gVisor doesn't play nice with Go modules, and brings in a bunch of unnecessary stuff that limits cross platform compatibility. This package is an attempt to make it easier to use the netstack package in other projects.