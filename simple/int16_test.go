package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_int16(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(int16(100))
	should.Nil(err)
	should.Equal([]byte{100, 0}, encoded[8:])
	var val int16
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal(int16(100), val)
	stream := gocodec.DefaultConfig.NewStream(encoded)
	stream.EncodeInt16(-1)
	should.Nil(stream.Error)
	encoded = stream.Buffer()
	should.Equal([]byte{100, 0, 0xff, 0xff}, encoded[8:])
	iter := gocodec.DefaultConfig.NewIterator(encoded[8:])
	should.Equal(int16(100), iter.DecodeInt16())
	should.Equal(int16(-1), iter.DecodeInt16())
	should.Nil(iter.Error)
}