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
	var val int
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(100, val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	encoder.EncodeUint(200)
	should.Nil(encoder.Error)
	encoded = encoder.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uint(100), decoder.DecodeUint())
	should.Equal(uint(200), decoder.DecodeUint())
	should.Nil(decoder.Error)
}