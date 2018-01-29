package gocodec

import (
	"unsafe"
)

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func ptrOfEmptyInterface(obj interface{}) unsafe.Pointer {
	return unsafe.Pointer((*emptyInterface)(unsafe.Pointer(&obj)).word)
}

func ptrAsBytes(size int, ptr unsafe.Pointer) []byte {
	valAsSlice := *(*[]byte)((unsafe.Pointer)(&sliceHeader{
		Data: ptr, Len: size, Cap: size}))
	return valAsSlice
}
