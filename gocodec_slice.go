package gocodec

import (
	"unsafe"
)

type sliceEncoder struct {
	elemSize    int
	elemEncoder ValEncoder
}

func (encoder *sliceEncoder) Encode(ptr unsafe.Pointer, stream *Stream) {
	typedPtr := (*sliceHeader)(ptr)
	stream.ptrOffset = uintptr(len(stream.buf))
	stream.buf = append(stream.buf, 24, 0, 0, 0, 0, 0, 0, 0)
	buf := [8]byte{}
	*(*int)(unsafe.Pointer(&buf)) = typedPtr.Len
	stream.buf = append(stream.buf, buf[:]...) // Len
	stream.buf = append(stream.buf, buf[:]...) // Cap
	encoder.EncodePointers(ptr, stream)
}

func (encoder *sliceEncoder) EncodePointers(ptr unsafe.Pointer, stream *Stream) {
	// TODO: write offset to buffer
	typedPtr := (*sliceHeader)(ptr)
	bytesCount := encoder.elemSize * typedPtr.Len
	byteSliceHeader := sliceHeader{
		Data: typedPtr.Data,
		Len:  bytesCount,
		Cap:  bytesCount,
	}
	byteSlice := (*[]byte)(unsafe.Pointer(&byteSliceHeader))
	*(*uintptr)(stream.ptr()) = uintptr(len(stream.buf)) - stream.ptrOffset
	stream.ptrOffset = uintptr(len(stream.buf)) // start of the bytes
	stream.buf = append(stream.buf, *(byteSlice)...)
	endPtrOffset := uintptr(len(stream.buf)) // end of the bytes
	for ; stream.ptrOffset < endPtrOffset; stream.ptrOffset += uintptr(encoder.elemSize) {
		// stream.buf will be changed in the loop, so the pointer will need to updated every time
		elemPtr := uintptr(ptrOfSlice(unsafe.Pointer(&stream.buf))) + stream.ptrOffset
		encoder.elemEncoder.EncodePointers(unsafe.Pointer(elemPtr), stream)
	}
}

type sliceDecoder struct {
	elemSize    int
	elemDecoder ValDecoder
}

func (decoder *sliceDecoder) Decode(ptr unsafe.Pointer, iter *Iterator) {
	typedPtr := (*[24]byte)(ptr)
	copy(typedPtr[:], iter.buf)
	iter.ptrBuf = iter.buf
	decoder.DecodePointers(ptr, iter)
}

func (decoder *sliceDecoder) DecodePointers(ptr unsafe.Pointer, iter *Iterator) {
	typedPtr := (*sliceHeader)(ptr)
	sliceDataBuf := iter.ptrBuf[uintptr(typedPtr.Data):]
	typedPtr.Data = uintptr(ptrOfSlice(unsafe.Pointer(&sliceDataBuf)))
	iter.ptrBuf = sliceDataBuf
	for i := 0; i < typedPtr.Len; i++ {
		if i > 0 {
			iter.ptrBuf = iter.ptrBuf[decoder.elemSize:]
		}
		decoder.elemDecoder.DecodePointers(ptrOfSlice(unsafe.Pointer(&iter.ptrBuf)), iter)
	}
}
