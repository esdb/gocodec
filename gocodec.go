package gocodec

import (
	"unsafe"
	"sync/atomic"
)

type Config struct {
}

type API interface {
	Marshal(obj interface{}) ([]byte, error)
	Unmarshal(buf []byte, objPtr interface{}) error
}

type ValEncoder interface {
	Encode(ptr unsafe.Pointer, encoder *GocEncoder)
	EncodePointers(ptr unsafe.Pointer, ptrOffset int, encoder *GocEncoder)
}

type ValDecoder interface {
	Decode(ptr unsafe.Pointer, decoder *GocDecoder)
	DecodePointers(ptr unsafe.Pointer, decoder *GocDecoder)
}

type frozenConfig struct {
	decoderCache                  unsafe.Pointer
	encoderCache                  unsafe.Pointer
}

func (cfg Config) Froze() API {
	api := &frozenConfig{}
	atomic.StorePointer(&api.decoderCache, unsafe.Pointer(&map[string]ValDecoder{}))
	atomic.StorePointer(&api.encoderCache, unsafe.Pointer(&map[string]ValEncoder{}))
	return api
}

var DefaultConfig = Config{}.Froze()

func Marshal(obj interface{}) ([]byte, error) {
	return DefaultConfig.Marshal(obj)
}

func Unmarshal(buf []byte, objPtr interface{}) error {
	return DefaultConfig.Unmarshal(buf, objPtr)
}

func (cfg *frozenConfig) Marshal(obj interface{}) ([]byte, error) {
	encoder := cfg.NewGocEncoder(nil)
	encoder.EncodeVal(obj)
	return encoder.Buffer(), encoder.Error
}

func (cfg *frozenConfig) Unmarshal(buf []byte, objPtr interface{}) error {
	decoder := cfg.NewGocDecoder(buf)
	decoder.DecodeVal(objPtr)
	return decoder.Error
}