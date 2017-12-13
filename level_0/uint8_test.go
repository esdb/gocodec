package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_uint8(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uint8(100))
	should.Nil(err)
	should.Equal([]byte{100}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (*uint8)(nil))
	should.Nil(err)
	should.Equal(uint8(100), reflect.ValueOf(decoded).Elem().Interface())
}