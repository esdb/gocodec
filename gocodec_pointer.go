package gocodec

import (
	"unsafe"
)

type pointerEncoder struct {
	BaseCodec
	elemEncoder ValEncoder
}

func (encoder *pointerEncoder) EncodeEmptyInterface(ptr uintptr, subEncoder ValEncoder, stream *Stream) {
	encoder.BaseCodec.EncodeEmptyInterface(uintptr(unsafe.Pointer(&ptr)), subEncoder, stream)
}

func (encoder *pointerEncoder) Encode(stream *Stream) {
	pPtr := unsafe.Pointer(&stream.buf[stream.cursor])
	ptr := *(*uintptr)(pPtr)
	if ptr != 0 {
		valAsBytes := ptrAsBytes(int(encoder.elemEncoder.Type().Size()), ptr)
		offset := uintptr(len(stream.buf)) - stream.cursor
		*(*uintptr)(pPtr) = offset
		stream.buf = append(stream.buf, valAsBytes...)
	}
}

type pointerDecoder struct {
	BaseCodec
	elemDecoder ValDecoder
}

func (decoder *pointerDecoder) Decode(iter *Iterator) {
	pPtr := unsafe.Pointer(&iter.cursor[0])
	offset := *(*uintptr)(pPtr)
	if offset == 0 {
		return
	}
	iter.cursor = iter.cursor[offset:]
	*(*uintptr)(pPtr) = uintptr(unsafe.Pointer(&iter.cursor[0]))
	decoder.elemDecoder.Decode(iter)
}
