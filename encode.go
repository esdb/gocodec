package gocodec

import (
	"fmt"
	"reflect"
	"unsafe"
)

type GocEncoder struct {
	cfg       *frozenConfig
	buf       []byte
	Error     error
}

func (cfg *frozenConfig) NewGocEncoder(buf []byte) *GocEncoder {
	return &GocEncoder{cfg: cfg, buf: buf}
}

func (encoder *GocEncoder) EncodeVal(val interface{}) {
	typ := reflect.TypeOf(val)
	valEncoder, err := encoderOfType(encoder.cfg, typ)
	if err != nil {
		encoder.ReportError("EncodeVal", err)
		return
	}
	valEncoder.Encode(ptrOfEmptyInterface(val), encoder)
}

func (encoder *GocEncoder) Buffer() []byte {
	return encoder.buf
}

func (encoder *GocEncoder) ReportError(operation string, err error) {
	if encoder.Error != nil {
		return
	}
	encoder.Error = fmt.Errorf("%s: %s", operation, err)
}

func (encoder *GocEncoder) EncodeInt(val int) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeInt8(val int8) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[1]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeInt16(val int16) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[2]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeInt32(val int32) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[4]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeInt64(val int64) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint(val uint) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint8(val uint8) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[1]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint16(val uint16) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[2]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint32(val uint32) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[4]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}

func (encoder *GocEncoder) EncodeUint64(val uint64) {
	ptr := unsafe.Pointer(&val)
	typedPtr := (*[8]byte)(ptr)
	encoder.buf = append(encoder.buf, (*typedPtr)[:]...)
}