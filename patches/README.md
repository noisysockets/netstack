# Patches

## Binary Native Endian

By default upstream netstack pulls in the hostarch package for determining endianess. This will panic on ARM64 with non 4K pages (eg. mac osx). This patch changes the default to use Go's bultin `binary.NativeEndian` constant instead.

## Expose Endpoints

Netstack's gonet abstraction does not provide a way to get the underlying endpoints. This patch adds a method to get the endpoints from a connection.

Remove once/if upstream [#10309](https://github.com/google/gvisor/pull/10309) is merged.

## User Cookie Option

This patch adds the ability to set the user cookie option on a socket.

Remove once/if upstream [#10308](https://github.com/google/gvisor/pull/10308) is merged.