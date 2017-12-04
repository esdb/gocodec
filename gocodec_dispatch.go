package gocodec

import (
	"reflect"
	"sync/atomic"
	"unsafe"
	"fmt"
)

func (cfg *frozenConfig) addDecoderToCache(cacheKey reflect.Type, decoder ValDecoder) {
	done := false
	for !done {
		ptr := atomic.LoadPointer(&cfg.decoderCache)
		cache := *(*map[reflect.Type]ValDecoder)(ptr)
		copied := map[reflect.Type]ValDecoder{}
		for k, v := range cache {
			copied[k] = v
		}
		copied[cacheKey] = decoder
		done = atomic.CompareAndSwapPointer(&cfg.decoderCache, ptr, unsafe.Pointer(&copied))
	}
}

func (cfg *frozenConfig) addEncoderToCache(cacheKey reflect.Type, encoder ValEncoder) {
	done := false
	for !done {
		ptr := atomic.LoadPointer(&cfg.encoderCache)
		cache := *(*map[reflect.Type]ValEncoder)(ptr)
		copied := map[reflect.Type]ValEncoder{}
		for k, v := range cache {
			copied[k] = v
		}
		copied[cacheKey] = encoder
		done = atomic.CompareAndSwapPointer(&cfg.encoderCache, ptr, unsafe.Pointer(&copied))
	}
}

func (cfg *frozenConfig) getDecoderFromCache(cacheKey reflect.Type) ValDecoder {
	ptr := atomic.LoadPointer(&cfg.decoderCache)
	cache := *(*map[reflect.Type]ValDecoder)(ptr)
	return cache[cacheKey]
}

func (cfg *frozenConfig) getEncoderFromCache(cacheKey reflect.Type) ValEncoder {
	ptr := atomic.LoadPointer(&cfg.encoderCache)
	cache := *(*map[reflect.Type]ValEncoder)(ptr)
	return cache[cacheKey]
}

func encoderOfType(cfg *frozenConfig, typ reflect.Type) (ValEncoder, error) {
	cacheKey := typ
	encoder := cfg.getEncoderFromCache(cacheKey)
	if encoder != nil {
		return encoder, nil
	}
	encoder, err := createEncoderOfType(cfg, typ)
	if err != nil {
		return nil, err
	}
	cfg.addEncoderToCache(cacheKey, encoder)
	return encoder, err
}

func decoderOfType(cfg *frozenConfig, typ reflect.Type) (ValDecoder, error) {
	cacheKey := typ
	decoder := cfg.getDecoderFromCache(cacheKey)
	if decoder != nil {
		return decoder, nil
	}
	decoder, err := createDecoderOfType(cfg, typ)
	if err != nil {
		return nil, err
	}
	cfg.addDecoderToCache(cacheKey, decoder)
	return decoder, err
}

func createEncoderOfType(cfg *frozenConfig, typ reflect.Type) (ValEncoder, error) {
	switch typ.Kind() {
	case reflect.Int:
		return &intCodec{}, nil
	case reflect.Int8:
		return &int8Codec{}, nil
	case reflect.Int16:
		return &int16Codec{}, nil
	case reflect.Int32:
		return &int32Codec{}, nil
	case reflect.Int64:
		return &int64Codec{}, nil
	case reflect.String:
		return &stringCodec{}, nil
	case reflect.Ptr:
		elemEncoder, err := createEncoderOfType(cfg, typ.Elem())
		if err != nil {
			return nil, err
		}
		return &pointerEncoder{elemEncoder: elemEncoder}, nil
	}
	return nil, fmt.Errorf("unsupported type %s", typ.String())
}

func createDecoderOfType(cfg *frozenConfig, typ reflect.Type) (ValDecoder, error) {
	switch typ.Kind() {
	case reflect.Int:
		return &intCodec{}, nil
	case reflect.Int8:
		return &int8Codec{}, nil
	case reflect.Int16:
		return &int16Codec{}, nil
	case reflect.Int32:
		return &int32Codec{}, nil
	case reflect.Int64:
		return &int64Codec{}, nil
	case reflect.String:
		return &stringCodec{}, nil
	case reflect.Ptr:
		elemDecoder, err := createDecoderOfType(cfg, typ.Elem())
		if err != nil {
			return nil, err
		}
		return &pointerDecoder{elemType: typ.Elem(), elemDecoder: elemDecoder}, nil
	}
	return nil, fmt.Errorf("unsupported type %s", typ.String())
}
