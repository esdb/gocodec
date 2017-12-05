package gocodec

import "unsafe"

type structEncoder struct {
	structSize uintptr
	fields []structFieldEncoder
}

func (encoder *structEncoder) Encode(ptr unsafe.Pointer, stream *Stream) {

}

func (encoder *structEncoder) EncodePointers(ptr unsafe.Pointer, stream *Stream) {

}

type structFieldEncoder struct {
	offset uintptr
	valEncoder ValEncoder
}

func (encoder *structFieldEncoder) encode(ptr unsafe.Pointer, stream *Stream) {
}

type structDecoder struct {
	fields []structFieldDecoder
}

type structFieldDecoder struct {
	offset uintptr
	valDecoder ValDecoder
}

func (decoder *structDecoder) Decode(ptr unsafe.Pointer, iter *Iterator) {
}
func (decoder *structDecoder) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
}
