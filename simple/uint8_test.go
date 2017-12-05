package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_uint8(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uint8(100))
	should.Nil(err)
	should.Equal([]byte{100}, encoded)
	var val uint8
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(uint8(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	stream.EncodeUint8(200)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uint8(100), iter.DecodeUint8())
	should.Equal(uint8(200), iter.DecodeUint8())
	should.Nil(iter.Error)
}