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
	stream.EncodeUintptr(200)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	decoder := gocodec.DefaultConfig.NewGocDecoder(encoded)
	should.Equal(uintptr(100), iter.DecodeUintptr())
	should.Equal(uintptr(200), iter.DecodeUintptr())
	should.Nil(iter.Error)
}