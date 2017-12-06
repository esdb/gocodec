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

func emptyInterfaceAsBytes(typ reflect.Type, val interface{}) []byte {
	valAsSlice := *(*[]byte)((unsafe.Pointer)(&sliceHeader{
		Data: uintptr(ptrOfEmptyInterface(val)), Len: int(typ.Size()), Cap: int(typ.Size())}))
	return valAsSlice
}

func (header *sliceHeader) DataPtr() unsafe.Pointer {
	return unsafe.Pointer(header.Data)
}

func (header *stringHeader) DataPtr() unsafe.Pointer {
	return unsafe.Pointer(header.Data)
}
