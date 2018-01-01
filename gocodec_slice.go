package gocodec

import (
	"unsafe"
)

type sliceEncoder struct {
	BaseCodec
	elemSize    int
	elemEncoder ValEncoder
}

func (encoder *sliceEncoder) Encode(stream *Stream) {
	pSlice := unsafe.Pointer(&stream.buf[stream.cursor])
	header := (*sliceHeader)(pSlice)
	if header.Len == 0 {
		return
	}
	header.Cap = header.Len
	byteSlice := ptrAsBytes(encoder.elemSize*header.Len, header.Data)
	// replace actual pointer with relative offset
	header.Data = uintptr(len(stream.buf))
	stream.cursor = uintptr(len(stream.buf)) // start of the bytes
	stream.buf = append(stream.buf, byteSlice...)
	if encoder.elemEncoder != nil {
		endCursor := uintptr(len(stream.buf)) // end of the bytes
		cursor := stream.cursor
		for ; cursor < endCursor; cursor += uintptr(encoder.elemSize) {
			stream.cursor = cursor
			encoder.elemEncoder.Encode(stream)
		}
	}
}

type sliceDecoder struct {
	BaseCodec
	elemSize    int
	elemDecoder ValDecoder
}

func (decoder *sliceDecoder) Decode(iter *Iterator) {
	pSlice := unsafe.Pointer(&iter.cursor[0])
	header := (*sliceHeader)(pSlice)
	if header.Len == 0 {
		return
	}
	relOffset := header.Data
	pCursor := uintptr(unsafe.Pointer(&iter.cursor[0]))
	offset := relOffset - (pCursor - iter.baseOffset) - iter.oldBaseOffset
	header.Data = uintptr(unsafe.Pointer(&iter.cursor[offset]))
	if decoder.elemDecoder != nil {
		cursor := iter.cursor[offset:]
		for i := 0; i < header.Len; i++ {
			if i > 0 {
				cursor = cursor[decoder.elemSize:]
			}
			iter.cursor = cursor
			decoder.elemDecoder.Decode(iter)
		}
	}
}
