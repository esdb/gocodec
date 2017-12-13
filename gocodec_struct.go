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

type structDecoder struct {
	BaseCodec
	fields []structFieldDecoder
}

type structFieldDecoder struct {
	offset  uintptr
	decoder ValDecoder
}

func (decoder *structDecoder) Decode(iter *Iterator) {
	baseCursor := iter.cursor
	for _, field := range decoder.fields {
		iter.cursor = baseCursor[field.offset:]
		field.decoder.Decode(iter)
	}
}
