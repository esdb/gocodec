package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_uint16(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uint16(100))
	should.Nil(err)
	should.Equal([]byte{100, 0}, encoded[8:])
	var val uint16
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(uint16(100), val)
	stream := gocodec.DefaultConfig.NewStream(encoded)
	stream.EncodeUint16(200)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	iter := gocodec.DefaultConfig.NewIterator(encoded[8:])
	should.Equal(uint16(100), iter.DecodeUint16())
	should.Equal(uint16(200), iter.DecodeUint16())
	should.Nil(iter.Error)
}