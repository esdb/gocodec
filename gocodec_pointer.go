package gocodec

import (
	"unsafe"
	"github.com/v2pro/plz/countlog"
	"errors"
)

type pointerEncoder struct {
	BaseCodec
	elemEncoder ValEncoder
}

func (encoder *pointerEncoder) Encode(stream *Stream) {
	pPtr := unsafe.Pointer(&stream.buf[stream.cursor])
	ptr := *(*uintptr)(pPtr)
	if ptr != 0 {
		valAsBytes := ptrAsBytes(int(encoder.elemEncoder.Type().Size()), ptr)
		*(*uintptr)(pPtr) = uintptr(len(stream.buf))
		stream.cursor = uintptr(len(stream.buf))
		stream.buf = append(stream.buf, valAsBytes...)
		encoder.elemEncoder.Encode(stream)
	}
}

type pointerDecoderWithoutCopy struct {
	BaseCodec
	elemDecoder ValDecoder
}

func (decoder *pointerDecoderWithoutCopy) Decode(iter *Iterator) {
	if countlog.ShouldLog(countlog.LevelDebug) {
		defer func() {
			recovered := recover()
			countlog.LogPanic(recovered, "valueType", decoder.valType)
			if recovered != nil {
				iter.ReportError("pointerDecoder", errors.New(
					"decode failed at type: "+decoder.valType.String()))
			}
		}()
	}
	pPtr := unsafe.Pointer(&iter.cursor[0])
	relOffset := *(*uintptr)(pPtr)
	if relOffset == 0 {
		return
	}
	pCursor := uintptr(pPtr)
	offset := relOffset - (pCursor - iter.baseOffset) - iter.oldBaseOffset
	iter.cursor = iter.cursor[offset:]
	*(*uintptr)(unsafe.Pointer(&iter.self[0])) = uintptr(unsafe.Pointer(&iter.cursor[0]))
	iter.self = iter.cursor
	decoder.elemDecoder.Decode(iter)
}

func (decoder *pointerDecoderWithoutCopy) HasPointer() bool {
	return true
}

type pointerDecoderWithCopy struct {
	BaseCodec
	elemDecoder ValDecoder
}

func (decoder *pointerDecoderWithCopy) Decode(iter *Iterator) {
	if countlog.ShouldLog(countlog.LevelDebug) {
		defer func() {
			recovered := recover()
			countlog.LogPanic(recovered, "valueType", decoder.valType)
			if recovered != nil {
				iter.ReportError("pointerDecoder", errors.New(
					"decode failed at type: "+decoder.valType.String()))
			}
		}()
	}
	pPtr := unsafe.Pointer(&iter.cursor[0])
	relOffset := *(*uintptr)(pPtr)
	if relOffset == 0 {
		return
	}
	pCursor := uintptr(pPtr)
	offset := relOffset - (pCursor - iter.baseOffset) - iter.oldBaseOffset
	iter.cursor = iter.cursor[offset:]
	copied := append([]byte(nil), iter.cursor[:decoder.elemDecoder.Type().Size()]...)
	*(*uintptr)(unsafe.Pointer(&iter.self[0])) = uintptr(unsafe.Pointer(&copied[0]))
	iter.self = copied
	decoder.elemDecoder.Decode(iter)
}

func (decoder *pointerDecoderWithCopy) HasPointer() bool {
	return true
}
