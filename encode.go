package gocodec

import (
	"fmt"
	"reflect"
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
