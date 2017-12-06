package gocodec

import "reflect"

type NoopCodec struct {
	valType reflect.Type
	signature uint32
}

func NewNoopCodec(valType reflect.Type) *NoopCodec {
	return &NoopCodec{valType: valType, signature: 0}
}

func (codec *NoopCodec) Encode(stream *Stream) {
}

func (codec *NoopCodec) Decode(iter *Iterator) {
}

func (codec *NoopCodec) Type() reflect.Type {
	return codec.valType
}

func (codec *NoopCodec) IsNoop() bool {
	return true
}

func (codec *NoopCodec) Signature() uint32 {
	return codec.signature
}
