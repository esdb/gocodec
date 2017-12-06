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
	header.Cap = header.Len
	byteSlice := ptrAsBytes(encoder.elemSize * header.Len, header.Data)
	// replace actual pointer with relative offset
	header.Data = uintptr(len(stream.buf)) - stream.cursor
	stream.cursor = uintptr(len(stream.buf)) // start of the bytes
	stream.buf = append(stream.buf, byteSlice...)
	if encoder.elemEncoder != nil {
		endCursor := uintptr(len(stream.buf)) // end of the bytes
		for ; stream.cursor < endCursor; stream.cursor += uintptr(encoder.elemSize) {
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
	offset := header.Data
	header.Data = uintptr(unsafe.Pointer(&iter.cursor[offset]))
	if decoder.elemDecoder != nil {
		iter.cursor = iter.cursor[offset:]
		for i := 0; i < header.Len; i++ {
			if i > 0 {
				iter.cursor = iter.cursor[decoder.elemSize:]
			}
			decoder.elemDecoder.Decode(iter)
		}
	}
}
