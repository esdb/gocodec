package gocodec

import (
	"reflect"
)

type BaseCodec struct {
	valType reflect.Type
	signature uint32
}

func NewBaseCodec(valType reflect.Type, signature uint32) *BaseCodec {
	return &BaseCodec{valType: valType, signature: signature}
}

func (codec *BaseCodec) Encode(stream *Stream) {
}

func (codec *BaseCodec) EncodeEmptyInterface(ptr uintptr, encoder ValEncoder, stream *Stream) {
	stream.cursor = uintptr(len(stream.buf))
	valAsSlice := ptrAsBytes(int(codec.valType.Size()), ptr)
	stream.buf = append(stream.buf, valAsSlice...)
	encoder.Encode(stream)
}

func (codec *BaseCodec) Decode(iter *Iterator) {
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

type NoopCodec struct {
	BaseCodec
}

func (codec *NoopCodec) IsNoop() bool {
	return true
}