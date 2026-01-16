package str

import "unsafe"

// Turns a byte string into a string without the extra alloc
// Not safe to hold unto this string as underlying slice may change
func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
