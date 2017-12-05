package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_uint(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uint(100))
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0}, encoded)
	var val uint
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(uint(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	stream.EncodeUint(200)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uint(100), iter.DecodeUint())
	should.Equal(uint(200), iter.DecodeUint())
	should.Nil(iter.Error)
}