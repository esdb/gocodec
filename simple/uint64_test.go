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
	encoder.EncodeUint64(200)
	should.Nil(encoder.Error)
	encoded = encoder.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uint64(100), decoder.DecodeUint64())
	should.Equal(uint64(200), decoder.DecodeUint64())
	should.Nil(decoder.Error)
}