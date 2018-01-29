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
	valAsSlice := ptrAsBytes(int(encoder.valType.Size()), ptr)
	stream.buf = append(stream.buf, valAsSlice...)
	encoder.encoder.Encode(ptr, stream)
}

func (encoder *rootEncoder) Signature() uint32 {
	return encoder.signature
}

func (encoder *rootEncoder) Type() reflect.Type {
	return encoder.valType
}

type rootDecoderWithCopy struct {
	valType   reflect.Type
	signature uint32
	decoder   ValDecoder
}

func (decoder *rootDecoderWithCopy) Signature() uint32 {
	return decoder.signature
}

func (decoder *rootDecoderWithCopy) Type() reflect.Type {
	return decoder.valType
}

func (decoder *rootDecoderWithCopy) DecodeEmptyInterface(ptr *emptyInterface, iter *Iterator) {
	iter.self = iter.allocator.Allocate(iter.objectSeq, iter.buf[8:8+decoder.Type().Size()])
	ptr.word = unsafe.Pointer(&iter.self[0])
	iter.cursor = iter.buf[8:]
	decoder.decoder.Decode(iter)
}

type rootDecoderWithoutCopy struct {
	valType   reflect.Type
	signature uint32
	decoder   ValDecoder
}

func (decoder *rootDecoderWithoutCopy) Signature() uint32 {
	return decoder.signature
}

func (decoder *rootDecoderWithoutCopy) Type() reflect.Type {
	return decoder.valType
}

func (decoder *rootDecoderWithoutCopy) DecodeEmptyInterface(ptr *emptyInterface, iter *Iterator) {
	ptr.word = unsafe.Pointer(&iter.buf[8])
	iter.self = iter.buf[8:]
	iter.cursor = iter.buf[8:]
	decoder.decoder.Decode(iter)
}
