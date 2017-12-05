package gocodec

import "unsafe"

// if only one pointer and embed in interface{}
// either directly or indirectly via array or struct
// the pointer itself will be merged as emptyInterface.Data
// the solution is to add one more direction before encoding
type singlePointerFix struct {
	simpleCodec
	encoder ValEncoder
}

func (encoder *singlePointerFix) Encode(ptr unsafe.Pointer, stream *Stream) {
	encoder.encoder.Encode(unsafe.Pointer(&ptr), stream)
}


