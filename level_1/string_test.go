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
	should.Equal([]byte{16, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 'h', 'e', 'l', 'l', 'o'}, encoded[8:])
	var val string
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal("hello", val)
}