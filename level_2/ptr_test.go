package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_ptr_of_string(t *testing.T) {
	should := require.New(t)
	obj := "hello"
	encoded, err := gocodec.Marshal(&obj)
	should.Nil(err)
	should.Equal([]byte{
		0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x68, 0x65, 0x6c, 0x6c, 0x6f,
	}, encoded[24:])
	decoded, err := gocodec.ReadonlyConfig.Unmarshal(encoded, (**string)(nil))
	should.Nil(err)
	should.Equal("hello", **decoded.(**string))
	decoded, err = gocodec.Unmarshal(encoded, (**string)(nil))
	should.Nil(err)
	should.Equal("hello", **decoded.(**string))
}

func Test_ptr_of_slice(t *testing.T) {
	should := require.New(t)
	obj := []byte("hello")
	encoded, err := gocodec.Marshal(&obj)
	should.Nil(err)
	should.Equal([]byte{
		0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x38, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x68, 0x65, 0x6c, 0x6c, 0x6f,
	}, encoded[24:])
	decoded, err := gocodec.ReadonlyConfig.Unmarshal(encoded, (**[]byte)(nil))
	should.Nil(err)
	should.Equal("hello", string(**decoded.(**[]byte)))
	decoded, err = gocodec.Unmarshal(encoded, (**[]byte)(nil))
	should.Nil(err)
	should.Equal("hello", string(**decoded.(**[]byte)))
}
