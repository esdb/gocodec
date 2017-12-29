package gocodec

type arrayEncoder struct {
	BaseCodec
	arrayLength int
	elementSize uintptr
	elemEncoder ValEncoder
}

func (encoder *arrayEncoder) Encode(stream *Stream) {
	if encoder.IsNoop() {
		return
	}
	cursor := stream.cursor
	for i := 0; i < encoder.arrayLength; i++ {
		stream.cursor = cursor // stream.cursor will change in the elemEncoder
		encoder.elemEncoder.Encode(stream)
		cursor = cursor + encoder.elementSize
	}
}

func (encoder *arrayEncoder) IsNoop() bool {
	return encoder.elemEncoder == nil
}

type arrayDecoder struct {
	BaseCodec
	arrayLength int
	elementSize uintptr
	elemDecoder ValDecoder
}

func (decoder *arrayDecoder) Decode(iter *Iterator) {
	if decoder.IsNoop() {
		return
	}
	cursor := iter.cursor
	for i := 0; i < decoder.arrayLength; i++ {
		iter.cursor = cursor // iter.cursor will change in elemDecoder
		decoder.elemDecoder.Decode(iter)
		cursor = cursor[decoder.elementSize:]
	}
}

func (decoder *arrayDecoder) IsNoop() bool {
	return decoder.elemDecoder == nil
}
