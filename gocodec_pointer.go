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
	// TODO: write offset to buffer
	encoder.elemEncoder.Encode(ptr, stream)
}

type pointerDecoder struct {
	elemType    reflect.Type
	elemDecoder ValDecoder
}

func (decoder *pointerDecoder) Decode(ptr unsafe.Pointer, iter *Iterator) {
	typedPtr := (*[8]byte)(ptr)
	copy(typedPtr[:], iter.buf)
	decoder.DecodePointers(ptr, iter)
	iter.buf = iter.buf[8:]
}

func (decoder *pointerDecoder) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
	offset := *(*int)(ptr)
	newVal := ptrOfEmptyInterface(reflect.New(decoder.elemType).Interface())
	*(*unsafe.Pointer)(ptr) = newVal
	oldBuf := iter.buf
	iter.buf = oldBuf[offset:]
	decoder.elemDecoder.Decode(newVal, iter)
	iter.buf = oldBuf
}