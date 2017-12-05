package gocodec

import "unsafe"

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  uintptr
	word uintptr
}

type stringHeader struct {
	Data uintptr
	Len  int
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func ptrOfEmptyInterface(obj interface{}) unsafe.Pointer {
	return unsafe.Pointer((*emptyInterface)(unsafe.Pointer(&obj)).word)
}

func (header *sliceHeader) DataPtr() unsafe.Pointer {
	return unsafe.Pointer(header.Data)
}

func (header *stringHeader) DataPtr() unsafe.Pointer {
	return unsafe.Pointer(header.Data)
}
