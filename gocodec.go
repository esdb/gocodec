package gocodec

import (
	"unsafe"
	"sync/atomic"
	"reflect"
)

type Config struct {
	VerifyChecksum bool
	ReadonlyDecode bool
}

type API interface {
	Marshal(val interface{}) ([]byte, error)
	Unmarshal(buf []byte, candidatePointers ...interface{}) (interface{}, error)
	NewIterator(buf []byte) *Iterator
	NewStream(buf []byte) *Stream
}

type ValEncoder interface {
	Encode(stream *Stream)
	Type() reflect.Type
	IsNoop() bool
	Signature() uint32
}

type RootEncoder interface {
	Type() reflect.Type
	Signature() uint32
	EncodeEmptyInterface(ptr unsafe.Pointer, stream *Stream)
}

type ValDecoder interface {
	Decode(iter *Iterator)
	DecodeEmptyInterface(ptr unsafe.Pointer, iter *Iterator)
	Type() reflect.Type
	IsNoop() bool
	Signature() uint32
	HasPointer() bool
}

type frozenConfig struct {
	verifyChecksum bool
	readonlyDecode bool
	decoderCache   unsafe.Pointer
	encoderCache   unsafe.Pointer
}

func (cfg Config) Froze() API {
	api := &frozenConfig{verifyChecksum: cfg.VerifyChecksum, readonlyDecode: cfg.ReadonlyDecode}
	atomic.StorePointer(&api.decoderCache, unsafe.Pointer(&map[string]ValDecoder{}))
	atomic.StorePointer(&api.encoderCache, unsafe.Pointer(&map[string]ValEncoder{}))
	return api
}

var DefaultConfig = Config{VerifyChecksum: true}.Froze()
var ReadonlyConfig = Config{VerifyChecksum: true, ReadonlyDecode: true}.Froze()

func Marshal(obj interface{}) ([]byte, error) {
	return DefaultConfig.Marshal(obj)
}

func Unmarshal(buf []byte, candidatePointers ...interface{}) (interface{}, error) {
	return DefaultConfig.Unmarshal(buf, candidatePointers...)
}

func NewIterator(buf []byte) *Iterator {
	return DefaultConfig.NewIterator(buf)
}

func NewStream(buf []byte) *Stream {
	return DefaultConfig.NewStream(buf)
}

func (cfg *frozenConfig) Marshal(val interface{}) ([]byte, error) {
	stream := cfg.NewStream(nil)
	stream.Marshal(val)
	return stream.Buffer(), stream.Error
}

func (cfg *frozenConfig) Unmarshal(buf []byte, candidatePointers ...interface{}) (interface{}, error) {
	iter := cfg.NewIterator(buf)
	val := iter.Unmarshal(candidatePointers...)
	return val, iter.Error
}
