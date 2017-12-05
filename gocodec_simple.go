package gocodec

import (
	"unsafe"
)

type simpleCodec struct {
}

func (codec *simpleCodec) EncodePointers(ptr unsafe.Pointer, stream *Stream) {
}

func (codec *simpleCodec) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
}

type intCodec struct {
	simpleCodec
}

func (codec *intCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *intCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*int)(ptr) = iter.DecodeInt()
}

type int8Codec struct {
	simpleCodec
}

func (codec *int8Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[1]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *int8Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*int8)(ptr) = iter.DecodeInt8()
}

type int16Codec struct {
	simpleCodec
}

func (codec *int16Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[2]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *int16Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*int16)(ptr) = iter.DecodeInt16()
}

type int32Codec struct {
	simpleCodec
}

func (codec *int32Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[4]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *int32Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*int32)(ptr) = iter.DecodeInt32()
}

type int64Codec struct {
	simpleCodec
}

func (codec *int64Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *int64Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*int64)(ptr) = iter.DecodeInt64()
}

type uintCodec struct {
	simpleCodec
}

func (codec *uintCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *uintCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*uint)(ptr) = iter.DecodeUint()
}

type uint8Codec struct {
	simpleCodec
}

func (codec *uint8Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[1]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *uint8Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*uint8)(ptr) = iter.DecodeUint8()
}

type uint16Codec struct {
	simpleCodec
}

func (codec *uint16Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[2]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *uint16Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*uint16)(ptr) = iter.DecodeUint16()
}

type uint32Codec struct {
	simpleCodec
}

func (codec *uint32Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[4]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *uint32Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*uint32)(ptr) = iter.DecodeUint32()
}

type uint64Codec struct {
	simpleCodec
}

func (codec *uint64Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *uint64Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*uint64)(ptr) = iter.DecodeUint64()
}

type uintptrCodec struct {
	simpleCodec
}

func (codec *uintptrCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *uintptrCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*uintptr)(ptr) = iter.DecodeUintptr()
}

type float32Codec struct {
	simpleCodec
}

func (codec *float32Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[4]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *float32Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*float32)(ptr) = iter.DecodeFloat32()
}

type float64Codec struct {
	simpleCodec
}

func (codec *float64Codec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (codec *float64Codec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*(*float64)(ptr) = iter.DecodeFloat64()
}
