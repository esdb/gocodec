package gocodec

import (
	"fmt"
	"reflect"
	"unsafe"
	"hash/crc32"
)

type Stream struct {
	cfg *frozenConfig
	// there are two pointers
	// buf + cursor => the input of encoder
	// buf + len(buf) => the output of encoder
	buf    []byte
	cursor uintptr
	Error  error
}

func (cfg *frozenConfig) NewStream(buf []byte) *Stream {
	return &Stream{cfg: cfg, buf: buf}
}

func (stream *Stream) Reset(buf []byte) {
	stream.buf = buf
	stream.cursor = 0
}

func (stream *Stream) Marshal(val interface{}) uint64 {
	valType := reflect.TypeOf(val)
	encoder, err := encoderOfType(stream.cfg, valType)
	if err != nil {
		stream.ReportError("EncodeVal", err)
		return 0
	}
	baseCursor := len(stream.buf)
	stream.buf = append(stream.buf, []byte{
		0, 0, 0, 0, 0, 0, 0, 0, // size
		0, 0, 0, 0,             // crc32
		0, 0, 0, 0,             // signature
	}...)
	encoder.EncodeEmptyInterface(uintptr(ptrOfEmptyInterface(val)), encoder, stream)
	if stream.Error != nil {
		return 0
	}
	encoded := stream.buf[baseCursor+12:]
	pSize := unsafe.Pointer(&stream.buf[baseCursor])
	size := uint64(len(stream.buf) - baseCursor)
	*(*uint64)(pSize) = size
	pSig := unsafe.Pointer(&stream.buf[baseCursor+12])
	*(*uint32)(pSig) = encoder.Signature()
	pCrc := unsafe.Pointer(&stream.buf[baseCursor+8])
	crc := crc32.NewIEEE()
	crc.Write(encoded)
	*(*uint32)(pCrc) = crc.Sum32()
	return size
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
