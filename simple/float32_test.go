package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_float32(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(float32(100))
	should.Nil(err)
	should.Equal([]byte{0x0, 0x0, 0xc8, 0x42}, encoded)
	var val float32
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(float32(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	encoder.EncodeFloat32(-1)
	should.Nil(encoder.Error)
	encoded = encoder.Buffer()
	should.Equal([]byte{0x0, 0x0, 0xc8, 0x42, 0x0, 0x0, 0x80, 0xbf}, encoded)
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(float32(100), decoder.DecodeFloat32())
	should.Equal(float32(-1), decoder.DecodeFloat32())
	should.Nil(decoder.Error)
}