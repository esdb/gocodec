package gocodec

import (
	"unsafe"
)

type stringCodec struct {
	BaseCodec
}

func (codec *stringCodec) Encode(stream *Stream) {
	pStr := unsafe.Pointer(&stream.buf[stream.cursor])
	str := *(*string)(pStr)
	offset := uintptr(len(stream.buf))
	header := (*stringHeader)(pStr)
	header.Data = offset
	stream.buf = append(stream.buf, str...)
}

func (codec *stringCodec) Decode(iter *Iterator) {
	pStr := unsafe.Pointer(&iter.cursor[0])
	header := (*stringHeader)(pStr)
	relOffset := header.Data
	pCursor := uintptr(unsafe.Pointer(&iter.cursor[0]))
	offset := relOffset - (pCursor - iter.baseOffset) - iter.oldBaseOffset
	pStr = unsafe.Pointer(&iter.self[0])
	header = (*stringHeader)(pStr)
	header.Data = uintptr(unsafe.Pointer(&iter.cursor[offset]))
}

func (codec *stringCodec) HasPointer() bool {
	return true
}
