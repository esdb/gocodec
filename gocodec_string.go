package gocodec

import (
	"unsafe"
)

type stringCodec struct {
}

func (codec *stringCodec) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*[16]byte)(ptr)
	encoder.ptrOffset = uintptr(len(encoder.buf))
	encoder.buf = append(encoder.buf, typedPtr[:]...)
	codec.EncodePointers(encoder.ptr(), encoder)
}

func (codec *stringCodec) EncodePointers(ptr unsafe.Pointer, encoder *GocEncoder) {
	str := *(*string)(ptr)
	typedPtr := (*stringHeader)(ptr)
	// relative offset
	typedPtr.Data = uintptr(len(encoder.buf)) - encoder.ptrOffset
	encoder.buf = append(encoder.buf, str...)
}

func (codec *stringCodec) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*[16]byte)(ptr)
	copy(typedPtr[:], decoder.buf)
	decoder.ptrBuf = decoder.buf
	codec.DecodePointers(ptr, decoder)
	decoder.buf = decoder.buf[16:]
}

func (codec *stringCodec) DecodePointers(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*stringHeader)(ptr)
	strDataBuf := decoder.ptrBuf[uintptr(typedPtr.Data):]
	typedPtr.Data = uintptr(ptrOfSlice(unsafe.Pointer(&strDataBuf)))
}
