package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_int16(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(int16(100))
	should.Nil(err)
	should.Equal([]byte{100, 0}, encoded)
	var val int16
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(int16(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	encoder.EncodeInt16(-1)
	should.Nil(encoder.Error)
	encoded = encoder.Buffer()
	should.Equal([]byte{100, 0, 0xff, 0xff}, encoded)
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(int16(100), decoder.DecodeInt16())
	should.Equal(int16(-1), decoder.DecodeInt16())
	should.Nil(decoder.Error)
}