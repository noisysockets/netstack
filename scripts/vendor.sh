#!/bin/bash
set -e

gvisor_dir=/workspace/gvisor

# List of subdirectories to copy and update
subdirs=(
    "atomicbitops"
    "bits"
    "buffer"
    "cleanup"
    "compressio"
    "context"
    "cpuid"
    "eventfd"
    "fd"
    "gohacks"
    "goid"
    "linewriter"
    "log"
    "memutil"
    "rand"
    "refs"
    "sleep"
    "sync"
    "state"
    "tcpip"
    "waiter"
    "xdp"
)

rmdirs=(
    "tcpip/link/tun"
)

# Copy directories.
mkdir -p pkg
for subdir in "${subdirs[@]}"; do
    cp -r "${gvisor_dir}/pkg/${subdir}" pkg/
done

# Remove blacklisted directories.
for subdir in "${rmdirs[@]}"; do
    rm -rf "pkg/$subdir"
done

# Rewrite all imports from 'gvisor.dev/gvisor/pkg/' to 'github.com/noisysockets/netstack/'.
for subdir in "${subdirs[@]}"; do
    find "pkg/$subdir" -type f -name '*.go' -exec sed -i 's|gvisor.dev/gvisor/|github.com/noisysockets/netstack/|g' {} +
done