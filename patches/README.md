# Patches

## Binary Native Endian

By default upstream netstack pulls in the hostarch package for determining endianess. This will panic on ARM64 with non 4K pages (eg. mac osx). This patch changes the default to use Go's bultin `binary.NativeEndian` constant instead.