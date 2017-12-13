package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_uint16(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uint16(100))
	should.Nil(err)
	should.Equal([]byte{100, 0}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (*uint16)(nil))
	should.Nil(err)
	should.Equal(uint16(100), reflect.ValueOf(decoded).Elem().Interface())
}