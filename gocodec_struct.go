package gocodec

import "unsafe"

type structEncoder struct {
	structSize int
	fields     []structFieldEncoder
}

func (encoder *structEncoder) Encode(ptr unsafe.Pointer, stream *Stream) {
	slice := sliceHeader{Data: uintptr(ptr), Len: encoder.structSize, Cap: encoder.structSize}
	slicePtr := unsafe.Pointer(&slice)
	stream.buf = append(stream.buf, *(*[]byte)(slicePtr)...)
}

func (encoder *structEncoder) EncodePointers(ptr unsafe.Pointer, stream *Stream) {

}

type structFieldEncoder struct {
	offset  uintptr
	encoder ValEncoder
}

func (encoder *structFieldEncoder) encode(ptr unsafe.Pointer, stream *Stream) {
}

type structDecoder struct {
	structSize int
	fields     []structFieldDecoder
}

func (decoder *structDecoder) Decode(ptr unsafe.Pointer, iter *Iterator) {
	slice := sliceHeader{Data: uintptr(ptr), Len: decoder.structSize, Cap: decoder.structSize}
	slicePtr := unsafe.Pointer(&slice)
	copy(*(*[]byte)(slicePtr), iter.buf)
	iter.buf = iter.buf[decoder.structSize:]
}

func (decoder *structDecoder) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
}

type structFieldDecoder struct {
	offset  uintptr
	decoder ValDecoder
}
