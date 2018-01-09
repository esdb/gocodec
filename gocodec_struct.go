package gocodec

type structEncoder struct {
	BaseCodec
	fields []structFieldEncoder
}

type structFieldEncoder struct {
	offset  uintptr
	encoder ValEncoder
}

func (encoder *structEncoder) Encode(stream *Stream) {
	baseCursor := stream.cursor
	for _, field := range encoder.fields {
		stream.cursor = baseCursor + field.offset
		field.encoder.Encode(stream)
	}
}

type structDecoderWithoutPointer struct {
	BaseCodec
	fields []structFieldDecoder
}

type structFieldDecoder struct {
	offset  uintptr
	decoder ValDecoder
}

func (decoder *structDecoderWithoutPointer) Decode(iter *Iterator) {
	baseCursor := iter.cursor
	for _, field := range decoder.fields {
		iter.cursor = baseCursor[field.offset:]
		field.decoder.Decode(iter)
	}
}

type structDecoderWithPointer struct {
	BaseCodec
	fields []structFieldDecoder
}

func (decoder *structDecoderWithPointer) Decode(iter *Iterator) {
	baseCursor := iter.cursor
	baseSelf := iter.self
	for _, field := range decoder.fields {
		iter.cursor = baseCursor[field.offset:]
		iter.self = baseSelf[field.offset:]
		field.decoder.Decode(iter)
	}
}

func (decoder *structDecoderWithPointer) HasPointer() bool {
	return true
}
