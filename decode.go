package gocodec

import (
	"reflect"
	"fmt"
	"unsafe"
	"hash/crc32"
	"errors"
)

type Iterator struct {
	cfg           *frozenConfig
	buf           []byte
	cursor        []byte
	baseOffset    uintptr
	oldBaseOffset uintptr
	Error         error
}

func (cfg *frozenConfig) NewIterator(buf []byte) *Iterator {
	return &Iterator{cfg: cfg, buf: buf, baseOffset: uintptr(unsafe.Pointer(&buf[0]))}
}

func (iter *Iterator) Reset(buf []byte) {
	iter.buf = buf
	iter.cursor = nil
	iter.baseOffset = uintptr(unsafe.Pointer(&buf[0]))
}

func (iter *Iterator) Unmarshal(candidatePointers ...interface{}) interface{} {
	size := *(*uint64)(unsafe.Pointer(&iter.buf[0]))
	encoded := iter.buf[12:size]
	nextBuf := iter.buf[size:]
	if iter.cfg.verifyChecksum {
		crcVal := *(*uint32)(unsafe.Pointer(&iter.buf[8]))
		crc := crc32.NewIEEE()
		crc.Write(encoded)
		if crc.Sum32() != crcVal {
			iter.ReportError("DecodeVal", errors.New("crc32 verification failed"))
			return nil
		}
	}
	sig := *(*uint32)(unsafe.Pointer(&iter.buf[12]))
	var decoder ValDecoder
	var val interface{}
	for _, candidatePointer := range candidatePointers {
		valType := reflect.TypeOf(candidatePointer).Elem()
		tryDecoder, err := decoderOfType(iter.cfg, valType)
		if err != nil {
			iter.ReportError("DecodeVal", err)
			return nil
		}
		if tryDecoder.Signature() == sig {
			decoder = tryDecoder
			val = candidatePointer
			break
		}
	}
	if decoder == nil {
		iter.ReportError("DecodeVal", errors.New("no decoder matches the signature"))
		return nil
	}
	oldBaseOffset := *(*uintptr)(unsafe.Pointer(&iter.buf[16]))
	*(*uintptr)(unsafe.Pointer(&iter.buf[16])) = iter.baseOffset
	iter.oldBaseOffset = oldBaseOffset
	(*emptyInterface)(unsafe.Pointer(&val)).word = uintptr(unsafe.Pointer(&iter.buf[24]))
	iter.cursor = iter.buf[24:]
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

func UpdateChecksum(buf []byte) {
	size := *(*uint64)(unsafe.Pointer(&buf[0]))
	encoded := buf[12:size]
	crc := crc32.NewIEEE()
	crc.Write(encoded)
	pCrc := unsafe.Pointer(&buf[8])
	*(*uint32)(pCrc) = crc.Sum32()
}
