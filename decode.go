package gocodec

import (
	"reflect"
	"fmt"
	"unsafe"
	"hash/crc32"
	"errors"
	"io"
	"github.com/v2pro/plz/countlog"
	"encoding/hex"
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
	baseOffset := uintptr(0)
	if buf != nil {
		baseOffset = uintptr(unsafe.Pointer(&buf[0]))
	}
	return &Iterator{cfg: cfg, buf: buf, baseOffset: baseOffset}
}

func (iter *Iterator) Reset(buf []byte) {
	iter.buf = buf
	iter.cursor = nil
	iter.baseOffset = uintptr(unsafe.Pointer(&buf[0]))
	iter.Error = nil
}

func (iter *Iterator) NextSize() uint64 {
	if len(iter.buf) < 24 {
		return 0
	}
	return *(*uint64)(unsafe.Pointer(&iter.buf[0]))
}

func (iter *Iterator) Skip() []byte {
	size := iter.NextSize()
	skipped := iter.buf[:size]
	iter.buf = iter.buf[size:]
	return skipped
}

func (iter *Iterator) Copy(candidatePointers ...interface{}) interface{} {
	size := iter.NextSize()
	if size == 0 {
		iter.Error = io.EOF
		return nil
	}
	copied := append([]byte(nil), iter.buf[:size]...)
	nextBuf := iter.buf[size:]
	iter.Reset(copied)
	result := iter.Unmarshal(candidatePointers...)
	iter.Reset(nextBuf)
	return result
}

func (iter *Iterator) Unmarshal(candidatePointers ...interface{}) interface{} {
	size := iter.NextSize()
	if size == 0 {
		iter.Error = io.EOF
		return nil
	}
	encoded := iter.buf[12:size]
	thisBuf := iter.buf[:size]
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!gocodec.failed to unmarshal",
				"err", recovered,
				"buf", hex.EncodeToString(thisBuf),
				"stacktrace", countlog.ProvideStacktrace)
			iter.ReportError("Unmarshal", fmt.Errorf("%v", recovered))
		}
	}()
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

func (iter *Iterator) Buffer() []byte {
	return iter.buf
}

func UpdateChecksum(buf []byte) {
	size := *(*uint64)(unsafe.Pointer(&buf[0]))
	encoded := buf[12:size]
	crc := crc32.NewIEEE()
	crc.Write(encoded)
	pCrc := unsafe.Pointer(&buf[8])
	*(*uint32)(pCrc) = crc.Sum32()
}
