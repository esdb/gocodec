package gocodec

import (
	"unsafe"
	"reflect"
)

type pointerEncoder struct {
	elemEncoder ValEncoder
}

func (encoder *pointerEncoder) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.buf = append(stream.buf, 8, 0, 0, 0, 0, 0, 0, 0)
	encoder.EncodePointers(ptr, stream)
}

func (encoder *pointerEncoder) EncodePointers(ptr unsafe.Pointer, stream *Stream) {
	*(*uintptr)(stream.ptr()) = uintptr(len(stream.buf)) - stream.ptrOffset
	encoder.elemEncoder.Encode(*(*unsafe.Pointer)(ptr), stream)
}

type pointerDecoder struct {
	elemType    reflect.Type
	elemDecoder ValDecoder
}

func (decoder *pointerDecoder) Decode(ptr unsafe.Pointer, iter *Iterator) {
	typedPtr := (*[8]byte)(ptr)
	copy(typedPtr[:], iter.buf)
	iter.ptrBuf = iter.buf[8:]
	decoder.DecodePointers(ptr, iter)
}

func (decoder *pointerDecoder) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
	ptrBufPtr := unsafe.Pointer(&iter.ptrBuf)
	*(*unsafe.Pointer)(ptr) = ptrOfSlice(ptrBufPtr)
	decoder.elemDecoder.Decode(ptrBufPtr, iter)
}