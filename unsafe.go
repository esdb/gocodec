package gocodec

import "unsafe"

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
	return (*emptyInterface)(unsafe.Pointer(&obj)).word
}

func ptrOfSlice(slicePtr unsafe.Pointer) unsafe.Pointer {
	return (*sliceHeader)(slicePtr).Data
}

