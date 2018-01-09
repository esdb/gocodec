package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
)

func Test_string(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal("hello")
	should.Nil(err)
	should.Equal([]byte{0x28, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 'h', 'e', 'l', 'l', 'o'}, encoded[24:])
	decoded, err := gocodec.ReadonlyConfig.Unmarshal(encoded, (*string)(nil))
	should.Nil(err)
	should.Equal("hello", *decoded.(*string))
	decoded, err = gocodec.Unmarshal(encoded, (*string)(nil))
	should.Nil(err)
	should.Equal("hello", *decoded.(*string))
}