package gocodec

import "unsafe"

type intCodec struct {
}

func (codec *intCodec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (codec *intCodec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*[8]byte)(ptr)
	copy((*typedPtr)[:], decoder.buf)
	decoder.buf = decoder.buf[8:]
}
