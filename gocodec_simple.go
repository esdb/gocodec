package gocodec

import "unsafe"

type intCodec struct {
}

func (codec *intCodec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *intCodec) EncodePointers(ptr unsafe.Pointer, ptrOffset int, encoder *GocEncoder) {
}

func (codec *intCodec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	*(*int)(ptr) = *(*int)(bufPtr)
}

func (codec *intCodec) DecodePointers(ptr unsafe.Pointer, decoder *GocDecoder) {
}