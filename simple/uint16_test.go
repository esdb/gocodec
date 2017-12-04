package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_uint16(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uint16(100))
	should.Nil(err)
	should.Equal([]byte{100, 0}, encoded)
	var val uint16
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(uint16(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	encoder.EncodeUint16(200)
	should.Nil(encoder.Error)
	encoded = encoder.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uint16(100), decoder.DecodeUint16())
	should.Equal(uint16(200), decoder.DecodeUint16())
	should.Nil(decoder.Error)
}