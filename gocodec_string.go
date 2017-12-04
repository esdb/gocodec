package gocodec

import (
	"unsafe"
)

type stringCodec struct {
}

func (codec *stringCodec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*stringHeader)(ptr)
	encoder.buf = append(encoder.buf, 16, 0, 0, 0, 0, 0, 0, 0)
	buf := [8]byte{}
	*(*int)(unsafe.Pointer(&buf)) = typedPtr.Len
	encoder.buf = append(encoder.buf, buf[:]...)
	codec.EncodePointers(ptr, 0, encoder)
}

func (codec *stringCodec) EncodePointers(ptr unsafe.Pointer, ptrOffset int, encoder *GocEncoder) {
	// TODO: write offset to buffer
	encoder.buf = append(encoder.buf, *(*string)(ptr)...)
}

func (codec *stringCodec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*[16]byte)(ptr)
	copy(typedPtr[:], decoder.buf)
	codec.DecodePointers(ptr, decoder)
	decoder.buf = decoder.buf[16:]
}

func (codec *stringCodec) DecodePointers(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*stringHeader)(ptr)
	offset := (uintptr)(typedPtr.Data)
	allocatedBytes := make([]byte, typedPtr.Len)
	typedPtr.Data = ptrOfSlice(unsafe.Pointer(&allocatedBytes))
	copy(allocatedBytes, decoder.buf[offset:])
}