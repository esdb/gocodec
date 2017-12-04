package gocodec

import (
	"reflect"
	"fmt"
	"unsafe"
)

type GocDecoder struct {
	cfg   *frozenConfig
	buf   []byte
	ptrBuf []byte
	Error error
}

func (cfg *frozenConfig) NewGocDecoder(buf []byte) *GocDecoder {
	return &GocDecoder{cfg: cfg, buf: buf}
}

func (decoder *GocDecoder) Reset(buf []byte) {
	decoder.buf = buf
	decoder.ptrBuf = nil
}

func (decoder *GocDecoder) DecodeVal(objPtr interface{}) {
	typ := reflect.TypeOf(objPtr)
	valDecoder, err := decoderOfType(decoder.cfg, typ.Elem())
	if err != nil {
		decoder.ReportError("DecodeVal", err)
		return
	}
	valDecoder.Decode(ptrOfEmptyInterface(objPtr), decoder)
}

func (decoder *GocDecoder) ReportError(operation string, err error) {
	if decoder.Error != nil {
		return
	}
	decoder.Error = fmt.Errorf("%s: %s", operation, err)
}

func (decoder *GocDecoder) DecodeInt() int {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*int)(bufPtr)
	decoder.buf = decoder.buf[8:]
	return val
}

func (decoder *GocDecoder) DecodeInt8() int8 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*int8)(bufPtr)
	decoder.buf = decoder.buf[1:]
	return val
}

func (decoder *GocDecoder) DecodeInt16() int16 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*int16)(bufPtr)
	decoder.buf = decoder.buf[2:]
	return val
}

func (decoder *GocDecoder) DecodeInt32() int32 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*int32)(bufPtr)
	decoder.buf = decoder.buf[4:]
	return val
}

func (decoder *GocDecoder) DecodeInt64() int64 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*int64)(bufPtr)
	decoder.buf = decoder.buf[8:]
	return val
}

func (decoder *GocDecoder) DecodeUint() uint {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*uint)(bufPtr)
	decoder.buf = decoder.buf[8:]
	return val
}

func (decoder *GocDecoder) DecodeUint8() uint8 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*uint8)(bufPtr)
	decoder.buf = decoder.buf[1:]
	return val
}

func (decoder *GocDecoder) DecodeUint16() uint16 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*uint16)(bufPtr)
	decoder.buf = decoder.buf[2:]
	return val
}

func (decoder *GocDecoder) DecodeUint32() uint32 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*uint32)(bufPtr)
	decoder.buf = decoder.buf[4:]
	return val
}

func (decoder *GocDecoder) DecodeUint64() uint64 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*uint64)(bufPtr)
	decoder.buf = decoder.buf[8:]
	return val
}

func (decoder *GocDecoder) DecodeUintptr() uintptr {
	bufPtr := ptrOfSlice(unsafe.Pointer(&decoder.buf))
	val := *(*uintptr)(bufPtr)
	decoder.buf = decoder.buf[8:]
	return val
}