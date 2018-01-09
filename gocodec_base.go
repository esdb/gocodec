package gocodec

import (
	"reflect"
	"unsafe"
)

type BaseCodec struct {
	valType   reflect.Type
	signature uint32
}

func newBaseCodec(valType reflect.Type, signature uint32) *BaseCodec {
	return &BaseCodec{valType: valType, signature: signature}
}

func (codec *BaseCodec) Encode(stream *Stream) {
	panic("not implemented")
}

func (codec *BaseCodec) EncodeEmptyInterface(ptr unsafe.Pointer, stream *Stream) {
	panic("not implemented")
}

func (codec *BaseCodec) Decode(iter *Iterator) {
	panic("not implemented")
}

func (codec *BaseCodec) DecodeEmptyInterface(ptr unsafe.Pointer, iter *Iterator) {
	panic("not implemented")
}

func (codec *BaseCodec) Type() reflect.Type {
	return codec.valType
}

func (codec *BaseCodec) IsNoop() bool {
	return false
}

func (codec *BaseCodec) Signature() uint32 {
	return codec.signature
}

func (codec *BaseCodec) HasPointer() bool {
	return false
}

type NoopCodec struct {
	BaseCodec
}

func (codec *NoopCodec) IsNoop() bool {
	return true
}

func (codec *NoopCodec) Decode(iter *Iterator) {
}

func (codec *NoopCodec) Encode(stream *Stream) {
}
