package gocodec

import (
	"unsafe"
)

type structEncoder struct {
	structSize int
	fields     []structFieldEncoder
}

type structFieldEncoder struct {
	offset  uintptr
	encoder ValEncoder
}

func (encoder *structEncoder) Encode(ptr unsafe.Pointer, stream *Stream) {
	slice := sliceHeader{Data: uintptr(ptr), Len: encoder.structSize, Cap: encoder.structSize}
	slicePtr := unsafe.Pointer(&slice)
	stream.ptrOffset = uintptr(len(stream.buf))
	stream.buf = append(stream.buf, *(*[]byte)(slicePtr)...)
	encoder.EncodePointers(ptr, stream)
}

func (encoder *structEncoder) EncodePointers(ptr unsafe.Pointer, stream *Stream) {
	basePtrOffset := stream.ptrOffset
	for _, field := range encoder.fields {
		fieldPtr := uintptr(ptr) + field.offset
		stream.ptrOffset = basePtrOffset + field.offset
		field.encoder.EncodePointers(unsafe.Pointer(fieldPtr), stream)
	}
}

type structDecoder struct {
	structSize int
	fields     []structFieldDecoder
}

type structFieldDecoder struct {
	offset  uintptr
	decoder ValDecoder
}

func (decoder *structDecoder) Decode(ptr unsafe.Pointer, iter *Iterator) {
	slice := sliceHeader{Data: uintptr(ptr), Len: decoder.structSize, Cap: decoder.structSize}
	slicePtr := unsafe.Pointer(&slice)
	copy(*(*[]byte)(slicePtr), iter.buf)
	iter.ptrBuf = iter.buf
	decoder.DecodePointers(ptr, iter)
}

func (decoder *structDecoder) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
	basePtrBuf := iter.ptrBuf
	for _, field := range decoder.fields {
		fieldPtr := unsafe.Pointer(uintptr(ptr) + field.offset)
		iter.ptrBuf = basePtrBuf[field.offset:]
		field.decoder.DecodePointers(fieldPtr, iter)
	}
}
