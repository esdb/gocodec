package gocodec

import (
	"unsafe"
)

type pointerEncoder struct {
	BaseCodec
	elemEncoder ValEncoder
}

func (encoder *pointerEncoder) Encode(stream *Stream) {
	pPtr := unsafe.Pointer(&stream.buf[stream.cursor])
	ptr := *(*uintptr)(pPtr)
	if ptr != 0 {
		valAsBytes := ptrAsBytes(int(encoder.elemEncoder.Type().Size()), ptr)
		*(*uintptr)(pPtr) = uintptr(len(stream.buf))
		stream.cursor = uintptr(len(stream.buf))
		stream.buf = append(stream.buf, valAsBytes...)
		encoder.elemEncoder.Encode(stream)
	}
}

type pointerDecoder struct {
	BaseCodec
	elemDecoder ValDecoder
}

func (decoder *pointerDecoder) Decode(iter *Iterator) {
	pPtr := unsafe.Pointer(&iter.cursor[0])
	relOffset := *(*uintptr)(pPtr)
	if relOffset == 0 {
		return
	}
	pCursor := uintptr(pPtr)
	offset := relOffset - (pCursor - iter.baseOffset) - iter.oldBaseOffset
	iter.cursor = iter.cursor[offset:]
	*(*uintptr)(pPtr) = uintptr(unsafe.Pointer(&iter.cursor[0]))
	decoder.elemDecoder.Decode(iter)
}
