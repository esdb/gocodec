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
	dataPtr := *(*unsafe.Pointer)(ptr)
	if dataPtr != nil {
		*(*uintptr)(stream.ptr()) = uintptr(len(stream.buf)) - stream.ptrOffset
		encoder.elemEncoder.Encode(dataPtr, stream)
	}
}

type pointerDecoder struct {
	elemType    reflect.Type
	elemDecoder ValDecoder
}

func (decoder *pointerDecoder) Decode(ptr unsafe.Pointer, iter *Iterator) {
	typedPtr := (*[8]byte)(ptr)
	copy(typedPtr[:], iter.buf)
	iter.ptrBuf = iter.buf
	decoder.DecodePointers(ptr, iter)
}

func (decoder *pointerDecoder) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
	ptrBufPtr := unsafe.Pointer(&iter.ptrBuf)
	offset := *(*uintptr)(ptrOfSlice(ptrBufPtr))
	if offset == 0 {
		return
	}
	iter.ptrBuf = iter.ptrBuf[offset:]
	ptrBufPtr = unsafe.Pointer(&iter.ptrBuf)
	*(*unsafe.Pointer)(ptr) = ptrOfSlice(ptrBufPtr)
	oldBuf := iter.buf
	iter.buf = iter.ptrBuf
	decoder.elemDecoder.Decode(ptrOfSlice(ptrBufPtr), iter)
	iter.buf = oldBuf
}
