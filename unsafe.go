package gocodec

import (
	"unsafe"
	"reflect"
)

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

func ptrAsBytes(valType reflect.Type, ptr uintptr) []byte {
	valAsSlice := *(*[]byte)((unsafe.Pointer)(&sliceHeader{
		Data: ptr, Len: int(valType.Size()), Cap: int(valType.Size())}))
	return valAsSlice
}

func (header *sliceHeader) DataPtr() unsafe.Pointer {
	return unsafe.Pointer(header.Data)
}

func (header *stringHeader) DataPtr() unsafe.Pointer {
	return unsafe.Pointer(header.Data)
}
