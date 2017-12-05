package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_uint64(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uint64(100))
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0}, encoded)
	var val uint64
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(uint64(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	stream.EncodeUint64(200)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uint64(100), iter.DecodeUint64())
	should.Equal(uint64(200), iter.DecodeUint64())
	should.Nil(iter.Error)
}