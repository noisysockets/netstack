# Patches

## Binary Native Endian

By default upstream netstack pulls in the hostarch package for determining endianess. This will panic on ARM64 with non 4K pages (eg. mac osx). This patch changes the default to use Go's bultin `binary.NativeEndian` constant instead.

## View With Borrowed Data

This patch adds a `NewViewWithBorrowedData` method to the `buffer.View` type. This method allows a `View` to be created from a slice of bytes without copying the data. The `View` will borrow the data from the slice. This allows for zero-copy interop with other libraries that use slices of bytes.