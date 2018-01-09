package gocodec

import (
	"unsafe"
	"reflect"
)

type rootEncoder struct {
	valType   reflect.Type
	signature uint32
	encoder   ValEncoder
}

func (encoder *rootEncoder) EncodeEmptyInterface(ptr unsafe.Pointer, stream *Stream) {
	stream.cursor = uintptr(len(stream.buf))
	valAsSlice := ptrAsBytes(int(encoder.valType.Size()), uintptr(ptr))
	stream.buf = append(stream.buf, valAsSlice...)
	encoder.encoder.Encode(stream)
}

func (encoder *rootEncoder) Signature() uint32 {
	return encoder.signature
}

func (encoder *rootEncoder) Encode(stream *Stream) {
	panic("not implemented")
}

func (encoder *rootEncoder) Type() reflect.Type {
	return encoder.valType
}

func (encoder *rootEncoder) IsNoop() bool {
	panic("not implemented")
}

type rootDecoder struct {
}

type verifyChecksumRootDecoder struct {
}

type readonlyRootDecoder struct {
}

type readonlyVerifyChecksumRootDecoder struct {
}
