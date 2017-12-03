package gocodec

import (
	"unsafe"
	"errors"
	"reflect"
)

type pointerEncoder struct {
	elemEncoder ValEncoder
}

func (valEncoder *pointerEncoder) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	if encoder.jmpOffset == 0 {
		encoder.buf = append(encoder.buf, 8, 0, 0, 0, 0, 0, 0, 0)
		valEncoder.elemEncoder.Encode(ptr, encoder)
		return
	}
	encoder.ReportError("encode pointer", errors.New("encode pointer in struct is not supported"))
}

type pointerDecoder struct {
	elemType    reflect.Type
	elemDecoder ValDecoder
}

func (valDecoder *pointerDecoder) Decode(ptrToPtr unsafe.Pointer, decoder *GocDecoder) {
	ptr := *(*unsafe.Pointer)(ptrToPtr)
	if ptr == nil {
		ptr = ptrOfEmptyInterface(reflect.New(valDecoder.elemType).Interface())
		*(*unsafe.Pointer)(ptrToPtr) = ptr
	}
	b := decoder.buf
	offset := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	decoder.buf = b[offset:]
	valDecoder.elemDecoder.Decode(ptr, decoder)
	decoder.buf = b
}
