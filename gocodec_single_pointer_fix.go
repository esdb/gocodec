package gocodec

import "unsafe"

type singlePointerFix struct {
	ValEncoder
}

func (encoder *singlePointerFix) EncodeEmptyInterface(ptr unsafe.Pointer, subEncoder ValEncoder, stream *Stream) {
	encoder.ValEncoder.EncodeEmptyInterface(unsafe.Pointer(&ptr), subEncoder, stream)
}
