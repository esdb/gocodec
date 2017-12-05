package gocodec

import (
	"reflect"
	"fmt"
	"unsafe"
)

type Iterator struct {
	cfg   *frozenConfig
	buf   []byte
	ptrBuf []byte
	Error error
}

func (cfg *frozenConfig) NewIterator(buf []byte) *Iterator {
	return &Iterator{cfg: cfg, buf: buf}
}

func (iter *Iterator) Reset(buf []byte) {
	iter.buf = buf
	iter.ptrBuf = nil
}

func (iter *Iterator) DecodeVal(objPtr interface{}) {
	typ := reflect.TypeOf(objPtr)
	decoder, err := decoderOfType(iter.cfg, typ.Elem())
	if err != nil {
		iter.ReportError("DecodeVal", err)
		return
	}
	decoder.Decode(ptrOfEmptyInterface(objPtr), iter)
}

func (iter *Iterator) ReportError(operation string, err error) {
	if iter.Error != nil {
		return
	}
	iter.Error = fmt.Errorf("%s: %s", operation, err)
}

func (iter *Iterator) DecodeInt() int {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*int)(bufPtr)
	iter.buf = iter.buf[8:]
	return val
}

func (iter *Iterator) DecodeInt8() int8 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*int8)(bufPtr)
	iter.buf = iter.buf[1:]
	return val
}

func (iter *Iterator) DecodeInt16() int16 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*int16)(bufPtr)
	iter.buf = iter.buf[2:]
	return val
}

func (iter *Iterator) DecodeInt32() int32 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*int32)(bufPtr)
	iter.buf = iter.buf[4:]
	return val
}

func (iter *Iterator) DecodeInt64() int64 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*int64)(bufPtr)
	iter.buf = iter.buf[8:]
	return val
}

func (iter *Iterator) DecodeUint() uint {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*uint)(bufPtr)
	iter.buf = iter.buf[8:]
	return val
}

func (iter *Iterator) DecodeUint8() uint8 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*uint8)(bufPtr)
	iter.buf = iter.buf[1:]
	return val
}

func (iter *Iterator) DecodeUint16() uint16 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*uint16)(bufPtr)
	iter.buf = iter.buf[2:]
	return val
}

func (iter *Iterator) DecodeUint32() uint32 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*uint32)(bufPtr)
	iter.buf = iter.buf[4:]
	return val
}

func (iter *Iterator) DecodeUint64() uint64 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*uint64)(bufPtr)
	iter.buf = iter.buf[8:]
	return val
}

func (iter *Iterator) DecodeUintptr() uintptr {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*uintptr)(bufPtr)
	iter.buf = iter.buf[8:]
	return val
}

func (iter *Iterator) DecodeFloat32() float32 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*float32)(bufPtr)
	iter.buf = iter.buf[4:]
	return val
}

func (iter *Iterator) DecodeFloat64() float64 {
	bufPtr := ptrOfSlice(unsafe.Pointer(&iter.buf))
	val := *(*float64)(bufPtr)
	iter.buf = iter.buf[8:]
	return val
}