package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(100)
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0}, encoded)
	var val int
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(100, val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	stream.EncodeInt(200)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(100, iter.DecodeInt())
	should.Equal(200, iter.DecodeInt())
	should.Nil(iter.Error)
}