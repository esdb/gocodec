package gocodec

import (
	"reflect"
	"fmt"
	"unsafe"
	"hash/crc32"
	"errors"
)

type Iterator struct {
	cfg    *frozenConfig
	buf    []byte
	cursor []byte
	Error  error
}

func (cfg *frozenConfig) NewIterator(buf []byte) *Iterator {
	return &Iterator{cfg: cfg, buf: buf}
}

func (iter *Iterator) Reset(buf []byte) {
	iter.buf = buf
	iter.cursor = nil
}

func (iter *Iterator) Unmarshal(nilPtr interface{}) interface{} {
	valType := reflect.TypeOf(nilPtr).Elem()
	decoder, err := decoderOfType(iter.cfg, valType)
	if err != nil {
		iter.ReportError("DecodeVal", err)
		return nil
	}
	size := *(*uint32)(unsafe.Pointer(&iter.buf[0]))
	encoded := iter.buf[8:size]
	nextBuf := iter.buf[size:]
	iter.buf = iter.buf[4:]
	crcVal := *(*uint32)(unsafe.Pointer(&iter.buf[0]))
	crc := crc32.NewIEEE()
	crc.Write(encoded)
	if crc.Sum32() != crcVal {
		iter.ReportError("DecodeVal", errors.New("crc32 verification failed"))
		return nil
	}
	iter.buf = iter.buf[4:]
	val := nilPtr
	(*emptyInterface)(unsafe.Pointer(&val)).word = uintptr(unsafe.Pointer(&iter.buf[0]))
	iter.cursor = iter.buf
	decoder.Decode(iter)
	iter.buf = nextBuf
	return val
}

func (iter *Iterator) ReportError(operation string, err error) {
	if iter.Error != nil {
		return
	}
	iter.Error = fmt.Errorf("%s: %s", operation, err)
}