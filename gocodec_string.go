package gocodec

import (
	"unsafe"
)

type stringCodec struct {
	NoopCodec
}

func (codec *stringCodec) Encode(stream *Stream) {
	pStr := unsafe.Pointer(&stream.buf[stream.cursor])
	str := *(*string)(pStr)
	offset := uintptr(len(stream.buf)) - stream.cursor
	stream.buf = append(stream.buf, str...)
	header := (*stringHeader)(pStr)
	header.Data = offset
}

func (codec *stringCodec) Decode(iter *Iterator) {
	pStr := unsafe.Pointer(&iter.cursor[0])
	header := (*stringHeader)(pStr)
	offset := header.Data
	header.Data = uintptr(unsafe.Pointer(&iter.cursor[offset]))
}
