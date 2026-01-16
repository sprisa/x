package str

import "unsafe"

// Turns a byte string into a string without the extra alloc.
// Not safe to hold unto this string as underlying slice may change.
func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// Turns a string into a byte string without the extra alloc.
// The returned byte slice MUST NOT be modified.
// Go 1.22+ will attempt to prevent allocs if the compiler can guarantee,
// so only use this to force when the compiler doesn't know.
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
