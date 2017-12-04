package gocodec

import (
	"unsafe"
)

type sliceEncoder struct {
	elemSize    int
	elemEncoder ValEncoder
}

func (valEncoder *sliceEncoder) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*sliceHeader)(ptr)
	encoder.buf = append(encoder.buf, 24, 0, 0, 0, 0, 0, 0, 0)
	buf := [8]byte{}
	*(*int)(unsafe.Pointer(&buf)) = typedPtr.Len
	encoder.buf = append(encoder.buf, buf[:]...) // Len
	encoder.buf = append(encoder.buf, buf[:]...) // Cap
	valEncoder.EncodePointers(ptr, encoder)
}

func (valEncoder *sliceEncoder) EncodePointers(ptr unsafe.Pointer, encoder *GocEncoder) {
	// TODO: write offset to buffer
	typedPtr := (*sliceHeader)(ptr)
	bytesCount := valEncoder.elemSize * typedPtr.Len
	byteSliceHeader := sliceHeader{
		Data: typedPtr.Data,
		Len:  bytesCount,
		Cap:  bytesCount,
	}
	byteSlice := (*[]byte)(unsafe.Pointer(&byteSliceHeader))
	encoder.ptrOffset = uintptr(len(encoder.buf)) // start of the bytes
	encoder.buf = append(encoder.buf, *(byteSlice)...)
	endPtrOffset := uintptr(len(encoder.buf)) // end of the bytes
	for ; encoder.ptrOffset < endPtrOffset; encoder.ptrOffset += uintptr(valEncoder.elemSize) {
		// encoder.buf will be changed in the loop, so the pointer will need to updated every time
		elemPtr := uintptr(ptrOfSlice(unsafe.Pointer(&encoder.buf))) + encoder.ptrOffset
		valEncoder.elemEncoder.EncodePointers(unsafe.Pointer(elemPtr), encoder)
	}
}

type sliceDecoder struct {
	elemSize    int
	elemDecoder ValDecoder
}

func (valDecoder *sliceDecoder) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*[24]byte)(ptr)
	copy(typedPtr[:], decoder.buf)
	decoder.ptrBuf = decoder.buf
	valDecoder.DecodePointers(ptr, decoder)
	decoder.buf = decoder.buf[24:]
}

func (valDecoder *sliceDecoder) DecodePointers(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*sliceHeader)(ptr)
	sliceDataBuf := decoder.ptrBuf[uintptr(typedPtr.Data):]
	typedPtr.Data = uintptr(ptrOfSlice(unsafe.Pointer(&sliceDataBuf)))
	decoder.ptrBuf = sliceDataBuf
	for i := 0; i < typedPtr.Len; i++ {
		valDecoder.elemDecoder.DecodePointers(ptrOfSlice(unsafe.Pointer(&decoder.ptrBuf)), decoder)
		decoder.ptrBuf = decoder.ptrBuf[valDecoder.elemSize:]
	}
}
