package gocodec

import "unsafe"

type simpleCodec struct {
}

func (codec *simpleCodec) EncodePointers(ptr unsafe.Pointer, ptrOffset int, encoder *GocEncoder) {
}

func (codec *simpleCodec) DecodePointers(ptr unsafe.Pointer, decoder *GocDecoder) {
}

type intCodec struct {
	simpleCodec
}

func (codec *intCodec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *intCodec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*int)(ptr) = decoder.DecodeInt()
}

type int8Codec struct {
	simpleCodec
}

func (codec *int8Codec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[1]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *int8Codec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*int8)(ptr) = decoder.DecodeInt8()
}

type int16Codec struct {
	simpleCodec
}

func (codec *int16Codec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[2]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *int16Codec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*int16)(ptr) = decoder.DecodeInt16()
}

type int32Codec struct {
	simpleCodec
}

func (codec *int32Codec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[4]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *int32Codec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*int32)(ptr) = decoder.DecodeInt32()
}

type int64Codec struct {
	simpleCodec
}

func (codec *int64Codec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *int64Codec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*int64)(ptr) = decoder.DecodeInt64()
}

type uintCodec struct {
	simpleCodec
}

func (codec *uintCodec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *uintCodec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*uint)(ptr) = decoder.DecodeUint()
}

type uint8Codec struct {
	simpleCodec
}

func (codec *uint8Codec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[1]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *uint8Codec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*uint8)(ptr) = decoder.DecodeUint8()
}

type uint16Codec struct {
	simpleCodec
}

func (codec *uint16Codec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[2]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *uint16Codec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*uint16)(ptr) = decoder.DecodeUint16()
}

type uint32Codec struct {
	simpleCodec
}

func (codec *uint32Codec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[4]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *uint32Codec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*uint32)(ptr) = decoder.DecodeUint32()
}

type uint64Codec struct {
	simpleCodec
}

func (codec *uint64Codec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *uint64Codec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	*(*uint64)(ptr) = decoder.DecodeUint64()
}