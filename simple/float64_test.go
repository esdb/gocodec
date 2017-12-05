package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_float64(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(float64(100))
	should.Nil(err)
	should.Equal([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x59, 0x40}, encoded[8:])
	var val float64
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(float64(100), val)
	stream := gocodec.DefaultConfig.NewStream(encoded)
	stream.EncodeFloat64(-1)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	should.Equal([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x59, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf0, 0xbf}, encoded[8:])
	iter := gocodec.DefaultConfig.NewIterator(encoded[8:])
	should.Equal(float64(100), iter.DecodeFloat64())
	should.Equal(float64(-1), iter.DecodeFloat64())
	should.Nil(iter.Error)
}