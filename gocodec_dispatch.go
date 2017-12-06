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

func encoderOfType(cfg *frozenConfig, valType reflect.Type) (ValEncoder, error) {
	cacheKey := valType
	encoder := cfg.getEncoderFromCache(cacheKey)
	if encoder != nil {
		return encoder, nil
	}
	encoder, err := createEncoderOfType(cfg, valType)
	if err != nil {
		return nil, err
	}
	cfg.addEncoderToCache(cacheKey, encoder)
	return encoder, err
}

func decoderOfType(cfg *frozenConfig, valType reflect.Type) (ValDecoder, error) {
	cacheKey := valType
	decoder := cfg.getDecoderFromCache(cacheKey)
	if decoder != nil {
		return decoder, nil
	}
	decoder, err := createDecoderOfType(cfg, valType)
	if err != nil {
		return nil, err
	}
	cfg.addDecoderToCache(cacheKey, decoder)
	return decoder, err
}

func createEncoderOfType(cfg *frozenConfig, valType reflect.Type) (ValEncoder, error) {
	switch valType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64:
		return &NoopCodec{BaseCodec: *NewBaseCodec(valType)}, nil
	case reflect.String:
		return &stringCodec{BaseCodec: *NewBaseCodec(valType)}, nil
	case reflect.Struct:
		fields := make([]structFieldEncoder, 0, valType.NumField())
		for i := 0; i < valType.NumField(); i++ {
			encoder, err := createEncoderOfType(cfg, valType.Field(i).Type)
			if err != nil {
				return nil, err
			}
			if !encoder.IsNoop() {
				fields = append(fields, structFieldEncoder{
					offset:  valType.Field(i).Offset,
					encoder: encoder,
				})
			}
		}
		encoder := &structEncoder{BaseCodec: *NewBaseCodec(valType), fields: fields}
		//if len(fields) == 1 && valType.Field(0).Type.Kind() == reflect.Ptr {
		//	return &singlePointerFix{encoder: encoder}, nil
		//}
		return encoder, nil
		//case reflect.Slice:
		//	elemEncoder, err := createEncoderOfType(cfg, CodecMeta.Elem())
		//	if err != nil {
		//		return nil, err
		//	}
		//	return &sliceEncoder{elemSize: int(CodecMeta.Elem().Size()), elemEncoder: elemEncoder}, nil
	case reflect.Ptr:
		elemEncoder, err := createEncoderOfType(cfg, valType.Elem())
		if err != nil {
			return nil, err
		}
		return &pointerEncoder{BaseCodec: *NewBaseCodec(valType), elemEncoder: elemEncoder}, nil
	}
	return nil, fmt.Errorf("unsupported type %s", valType.String())
}

func createDecoderOfType(cfg *frozenConfig, valType reflect.Type) (ValDecoder, error) {
	switch valType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64:
		return &NoopCodec{BaseCodec: *NewBaseCodec(valType)}, nil
	case reflect.String:
		return &stringCodec{BaseCodec: *NewBaseCodec(valType)}, nil
	case reflect.Struct:
		fields := make([]structFieldDecoder, 0, valType.NumField())
		for i := 0; i < valType.NumField(); i++ {
			decoder, err := createDecoderOfType(cfg, valType.Field(i).Type)
			if err != nil {
				return nil, err
			}
			if !decoder.IsNoop() {
				fields = append(fields, structFieldDecoder{
					offset:  valType.Field(i).Offset,
					decoder: decoder,
				})
			}
		}
		return &structDecoder{BaseCodec: *NewBaseCodec(valType), fields: fields}, nil
		//case reflect.Slice:
		//	elemDecoder, err := createDecoderOfType(cfg, CodecMeta.Elem())
		//	if err != nil {
		//		return nil, err
		//	}
		//	return &sliceDecoder{elemSize: int(CodecMeta.Elem().Size()), elemDecoder: elemDecoder}, nil
	case reflect.Ptr:
		elemDecoder, err := createDecoderOfType(cfg, valType.Elem())
		if err != nil {
			return nil, err
		}
		return &pointerDecoder{BaseCodec: *NewBaseCodec(valType), elemDecoder: elemDecoder}, nil
	}
	return nil, fmt.Errorf("unsupported type %s", valType.String())
}
