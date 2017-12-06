package gocodec

import (
	"unsafe"
	"sync/atomic"
	"reflect"
)

type Config struct {
}

type API interface {
	Marshal(val interface{}) ([]byte, error)
	Unmarshal(buf []byte, nilPtr interface{}) (interface{}, error)
	NewIterator(buf []byte) *Iterator
	NewStream(buf []byte) *Stream
}

type ValEncoder interface {
	Encode(stream *Stream)
	Type() reflect.Type
	IsNoop() bool
	Signature() uint32
}

type ValDecoder interface {
	Decode(iter *Iterator)
	Type() reflect.Type
	IsNoop() bool
	Signature() uint32
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

func Unmarshal(buf []byte, objPtr interface{}) (interface{}, error) {
	return DefaultConfig.Unmarshal(buf, objPtr)
}

func (cfg *frozenConfig) Marshal(val interface{}) ([]byte, error) {
	stream := cfg.NewStream(nil)
	stream.Marshal(val)
	return stream.Buffer(), stream.Error
}

func (cfg *frozenConfig) Unmarshal(buf []byte, nilPtr interface{}) (interface{}, error) {
	iter := cfg.NewIterator(buf)
	val := iter.Unmarshal(nilPtr)
	return val, iter.Error
}