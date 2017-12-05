package gocodec

import (
	"unsafe"
)

type stringCodec struct {
}

func (codec *stringCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*[16]byte)(ptr)
	stream.ptrOffset = uintptr(len(stream.buf))
	stream.buf = append(stream.buf, typedPtr[:]...)
	codec.EncodePointers(stream.ptr(), stream)
}

func (codec *stringCodec) EncodePointers(ptr unsafe.Pointer, stream *Stream) {
	str := *(*string)(ptr)
	typedPtr := (*stringHeader)(ptr)
	// relative offset
	typedPtr.Data = uintptr(len(stream.buf)) - stream.ptrOffset
	stream.buf = append(stream.buf, str...)
}

func (codec *stringCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	typedPtr := (*[16]byte)(ptr)
	copy(typedPtr[:], iter.buf)
	iter.ptrBuf = iter.buf
	codec.DecodePointers(ptr, iter)
	iter.buf = iter.buf[16:]
}

func (codec *stringCodec) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
	typedPtr := (*stringHeader)(ptr)
	strDataBuf := iter.ptrBuf[uintptr(typedPtr.Data):]
	typedPtr.Data = uintptr(unsafe.Pointer(&strDataBuf[0]))
}
