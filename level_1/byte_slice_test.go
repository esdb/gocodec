package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
)

func Test_byte_slice(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal([]byte("hello"))
	should.Nil(err)
	should.Equal([]byte{24, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 'h', 'e', 'l', 'l', 'o'}, encoded)
	var val []byte
	should.Nil(gocodec.Unmarshal(encoded, &val))
	should.Equal([]byte("hello"), val)
}