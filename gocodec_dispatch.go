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
	case reflect.Uint:
		return &uintCodec{}, nil
	case reflect.Uint8:
		return &uint8Codec{}, nil
	case reflect.Uint16:
		return &uint16Codec{}, nil
	case reflect.Uint32:
		return &uint32Codec{}, nil
	case reflect.Uint64:
		return &uint64Codec{}, nil
	case reflect.Uintptr:
		return &uintptrCodec{}, nil
	case reflect.Float32:
		return &float32Codec{}, nil
	case reflect.Float64:
		return &float64Codec{}, nil
	case reflect.String:
		return &stringCodec{}, nil
	case reflect.Struct:
		fields := make([]structFieldEncoder, typ.NumField())
		var err error
		for i := 0; i < typ.NumField(); i++ {
			fields[i].offset = typ.Field(i).Offset
			fields[i].encoder, err = createEncoderOfType(cfg, typ.Field(i).Type)
			if err != nil {
				return nil, err
			}
		}
		return &structEncoder{structSize: int(typ.Size()), fields: fields}, nil
	case reflect.Slice:
		elemEncoder, err := createEncoderOfType(cfg, typ.Elem())
		if err != nil {
			return nil, err
		}
		return &sliceEncoder{elemSize: int(typ.Elem().Size()), elemEncoder: elemEncoder}, nil
	case reflect.Ptr:
		elemEncoder, err := createEncoderOfType(cfg, typ.Elem())
		if err != nil {
			return nil, err
		}
		encoder := &pointerEncoder{elemEncoder: elemEncoder}
		switch typ.Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64:
			return &singlePointerFix{encoder:encoder}, nil
		}
		return encoder, nil
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
	case reflect.Uint:
		return &uintCodec{}, nil
	case reflect.Uint8:
		return &uint8Codec{}, nil
	case reflect.Uint16:
		return &uint16Codec{}, nil
	case reflect.Uint32:
		return &uint32Codec{}, nil
	case reflect.Uint64:
		return &uint64Codec{}, nil
	case reflect.Uintptr:
		return &uintptrCodec{}, nil
	case reflect.Float32:
		return &float32Codec{}, nil
	case reflect.Float64:
		return &float64Codec{}, nil
	case reflect.String:
		return &stringCodec{}, nil
	case reflect.Struct:
		fields := make([]structFieldDecoder, typ.NumField())
		var err error
		for i := 0; i < typ.NumField(); i++ {
			fields[i].offset = typ.Field(i).Offset
			fields[i].decoder, err = createDecoderOfType(cfg, typ.Field(i).Type)
			if err != nil {
				return nil, err
			}
		}
		return &structDecoder{structSize: int(typ.Size()), fields: fields}, nil
	case reflect.Slice:
		elemDecoder, err := createDecoderOfType(cfg, typ.Elem())
		if err != nil {
			return nil, err
		}
		return &sliceDecoder{elemSize: int(typ.Elem().Size()), elemDecoder: elemDecoder}, nil
	case reflect.Ptr:
		elemDecoder, err := createDecoderOfType(cfg, typ.Elem())
		if err != nil {
			return nil, err
		}
		return &pointerDecoder{elemType: typ.Elem(), elemDecoder: elemDecoder}, nil
	}
	return nil, fmt.Errorf("unsupported type %s", typ.String())
}
