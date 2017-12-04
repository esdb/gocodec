package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_uintptr(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uintptr(100))
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0}, encoded)
	var val uintptr
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(uintptr(100), val)
	encoder := gocodec.DefaultConfig.NewGocEncoder(encoded)
	encoder.EncodeUintptr(200)
	should.Nil(encoder.Error)
	encoded = encoder.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uintptr(100), decoder.DecodeUintptr())
	should.Equal(uintptr(200), decoder.DecodeUintptr())
	should.Nil(decoder.Error)
}