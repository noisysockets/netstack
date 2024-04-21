# Patches

## Binary Native Endian

By default upstream netstack pulls in the hostarch package for determining endianess. This will panic on ARM64 with non 4K pages (eg. mac osx). This patch changes the default to use Go's bultin `binary.NativeEndian` constant instead.

## Expose Endpoints

Netstack's gonet abstraction does not provide a way to get the underlying endpoints. This patch adds a method to get the endpoints from a connection.

Remove once/if upstream [#10309](https://github.com/google/gvisor/pull/10309) is merged.

## Conn Mark

Add support for netfilter style connmarks so that packets can be associated with connections, and vice versa.

Remove once/if upstream [#10310](https://github.com/google/gvisor/pull/10310) is merged.