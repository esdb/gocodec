package gocodec

import (
	"fmt"
	"reflect"
	"unsafe"
	"hash/crc32"
)

type Stream struct {
	cfg       *frozenConfig
	// there are two pointers being written to
	// buf + ptrOffset => the place where the pointer will be updated
	// buf + len(buf) => the place where actual content of pointer will be appended to
	buf       []byte
	ptrOffset uintptr
	Error     error
}

func (cfg *frozenConfig) NewStream(buf []byte) *Stream {
	return &Stream{cfg: cfg, buf: buf}
}

func (stream *Stream) Reset(buf []byte) {
	stream.buf = buf
	stream.ptrOffset = 0
}

// buf + ptrOffset
func (stream *Stream) ptr() unsafe.Pointer {
	buf := stream.buf[stream.ptrOffset:]
	return unsafe.Pointer(&buf[0])
}

func (stream *Stream) EncodeVal(val interface{}) {
	typ := reflect.TypeOf(val)
	encoder, err := encoderOfType(stream.cfg, typ)
	if err != nil {
		stream.ReportError("EncodeVal", err)
		return
	}
	stream.buf = append(stream.buf, []byte{
		0, 0, 0, 0, // size
		0, 0, 0, 0, // crc32
		}...)
	beforeEncodeOffset := len(stream.buf)
	encoder.Encode(ptrOfEmptyInterface(val), stream)
	if stream.Error != nil {
		return
	}
	encoded := stream.buf[beforeEncodeOffset:]
	pSizeBuf := stream.buf[beforeEncodeOffset - 8:]
	pSize := unsafe.Pointer(&pSizeBuf[0])
	*(*uint32)(pSize) = uint32(len(encoded)) + 8
	pCrcBuf := stream.buf[beforeEncodeOffset - 4:]
	pCrc := unsafe.Pointer(&pCrcBuf[0])
	crc := crc32.NewIEEE()
	crc.Write(encoded)
	*(*uint32)(pCrc) = crc.Sum32()
}

func (stream *Stream) Buffer() []byte {
	return stream.buf
}

func (stream *Stream) ReportError(operation string, err error) {
	if stream.Error != nil {
		return
	}
	stream.Error = fmt.Errorf("%s: %s", operation, err)
}

func (stream *Stream) EncodeInt(val int) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeInt8(val int8) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[1]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeInt16(val int16) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[2]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeInt32(val int32) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[4]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeInt64(val int64) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeUint(val uint) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeUint8(val uint8) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[1]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeUint16(val uint16) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[2]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeUint32(val uint32) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[4]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeUint64(val uint64) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeUintptr(val uintptr) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeFloat32(val float32) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[4]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}

func (stream *Stream) EncodeFloat64(val float64) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	stream.buf = append(stream.buf, (*typedPtr)[:]...)
}