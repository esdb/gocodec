package gocodec

import "unsafe"

type sliceEncoder struct {
	elemSize int
}

func (valEncoder *sliceEncoder) Encode(ptr unsafe.Pointer, encoder *GocEncoder) {
	typedPtr := (*sliceHeader)(ptr)
	encoder.buf = append(encoder.buf, 24, 0, 0, 0, 0, 0, 0, 0)
	buf := [8]byte{}
	*(*int)(unsafe.Pointer(&buf)) = typedPtr.Len
	encoder.buf = append(encoder.buf, buf[:]...) // Len
	encoder.buf = append(encoder.buf, buf[:]...) // Cap
	valEncoder.EncodePointers(ptr, 0, encoder)
}

func (valEncoder *sliceEncoder) EncodePointers(ptr unsafe.Pointer, ptrOffset int, encoder *GocEncoder) {
	// TODO: write offset to buffer
	typedPtr := (*sliceHeader)(ptr)
	bytesCount := valEncoder.elemSize * typedPtr.Len
	byteSliceHeader := sliceHeader{
		Data: typedPtr.Data,
		Len: bytesCount,
		Cap: bytesCount,
	}
	byteSlice := (*[]byte)(unsafe.Pointer(&byteSliceHeader))
	encoder.buf = append(encoder.buf, *(byteSlice)...)
}

type sliceDecoder struct {
	elemSize int
}

func (valDecoder *sliceDecoder) Decode(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*[24]byte)(ptr)
	copy(typedPtr[:], decoder.buf)
	valDecoder.DecodePointers(ptr, decoder)
	decoder.buf = decoder.buf[24:]
}

func (valDecoder *sliceDecoder) DecodePointers(ptr unsafe.Pointer, decoder *GocDecoder) {
	typedPtr := (*sliceHeader)(ptr)
	offset := (uintptr)(typedPtr.Data)
	allocatedBytes := make([]byte, typedPtr.Len * valDecoder.elemSize)
	typedPtr.Data = ptrOfSlice(unsafe.Pointer(&allocatedBytes))
	copy(allocatedBytes, decoder.buf[offset:])
}