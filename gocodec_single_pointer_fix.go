package gocodec

import "unsafe"

type singlePointerFix struct {
	ValEncoder
}

func (encoder *singlePointerFix) EncodeEmptyInterface(ptr uintptr, subEncoder ValEncoder, stream *Stream) {
	encoder.ValEncoder.EncodeEmptyInterface(uintptr(unsafe.Pointer(&ptr)), subEncoder, stream)
}
