package gocodec

import "unsafe"

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func ptrOfEmptyInterface(obj interface{}) unsafe.Pointer {
	return (*emptyInterface)(unsafe.Pointer(&obj)).word
}