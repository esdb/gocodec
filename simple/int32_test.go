package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_int32(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(int32(100))
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0}, encoded)
	var val int32
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(int32(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	stream.EncodeInt32(-1)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	should.Equal([]byte{100, 0, 0, 0, 0xff, 0xff, 0xff, 0xff}, encoded)
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(int32(100), iter.DecodeInt32())
	should.Equal(int32(-1), iter.DecodeInt32())
	should.Nil(iter.Error)
}