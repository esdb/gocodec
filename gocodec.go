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
	NewIterator(buf []byte) *Iterator
	NewStream(buf []byte) *Stream
}

type ValEncoder interface {
	Encode(ptr unsafe.Pointer, stream *Stream)
	EncodePointers(ptr unsafe.Pointer, stream *Stream)
}

type ValDecoder interface {
	Decode(ptr unsafe.Pointer, iter *Iterator)
	DecodePointers(ptr unsafe.Pointer, iter *Iterator)
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
	stream := cfg.NewStream(nil)
	stream.EncodeVal(obj)
	return stream.Buffer(), stream.Error
}

func (cfg *frozenConfig) Unmarshal(buf []byte, objPtr interface{}) error {
	iter := cfg.NewIterator(buf)
	iter.DecodeVal(objPtr)
	return iter.Error
}