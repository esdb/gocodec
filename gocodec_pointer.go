package gocodec

import (
	"unsafe"
	"reflect"
)

type pointerEncoder struct {
	elemEncoder ValEncoder
}

func (valEncoder *pointerEncoder) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	encoder.buf = append(encoder.buf, 8, 0, 0, 0, 0, 0, 0, 0)
	valEncoder.EncodePointers(ptr, encoder)
}

func (valEncoder *pointerEncoder) EncodePointers(ptr unsafe.Pointer, encoder *GocEncoder) {
	// TODO: write offset to buffer
	valEncoder.elemEncoder.Encode(ptr, encoder)
}

type pointerDecoder struct {
	elemType    reflect.Type
	elemDecoder ValDecoder
}

func (valDecoder *pointerDecoder) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*[8]byte)(ptr)
	copy(typedPtr[:], decoder.buf)
	valDecoder.DecodePointers(ptr, decoder)
	decoder.buf = decoder.buf[8:]
}

func (valDecoder *pointerDecoder) DecodePointers(ptr unsafe.Pointer, decoder *GocDecoder) {
	offset := *(*int)(ptr)
	newVal := ptrOfEmptyInterface(reflect.New(valDecoder.elemType).Interface())
	*(*unsafe.Pointer)(ptr) = newVal
	oldBuf := decoder.buf
	decoder.buf = oldBuf[offset:]
	valDecoder.elemDecoder.Decode(newVal, decoder)
	decoder.buf = oldBuf
}