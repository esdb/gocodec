package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_uint32(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uint32(100))
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0}, encoded)
	var val uint32
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(uint32(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	encoder.EncodeUint32(200)
	should.Nil(encoder.Error)
	encoded = encoder.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uint32(100), decoder.DecodeUint32())
	should.Equal(uint32(200), decoder.DecodeUint32())
	should.Nil(decoder.Error)
}