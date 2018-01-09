package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_float32(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(float32(100))
	should.Nil(err)
	should.Equal([]byte{0x0, 0x0, 0xc8, 0x42}, encoded[16:])
	val, err := gocodec.ReadonlyConfig.Unmarshal(encoded, (*float32)(nil))
	should.Nil(err)
	should.Equal(float32(100), *(val.(*float32)))
	val, err = gocodec.Unmarshal(encoded, (*float32)(nil))
	should.Nil(err)
	should.Equal(float32(100), *(val.(*float32)))
}
