package gocodec

import (
	"reflect"
	"fmt"
)

type GocDecoder struct {
	cfg   *frozenConfig
	buf   []byte
	Error error
}

func (cfg *frozenConfig) NewGocDecoder(buf []byte) *GocDecoder {
	return &GocDecoder{cfg: cfg, buf: buf}
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
