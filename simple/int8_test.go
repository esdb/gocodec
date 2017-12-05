package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_int8(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(int8(100))
	should.Nil(err)
	should.Equal([]byte{100}, encoded)
	var val int8
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(int8(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	stream.EncodeInt8(-1)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	should.Equal([]byte{100, 0xff}, encoded)
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(int8(100), iter.DecodeInt8())
	should.Equal(int8(-1), iter.DecodeInt8())
	should.Nil(iter.Error)
}