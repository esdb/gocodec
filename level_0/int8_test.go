package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_int8(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(int8(100))
	should.Nil(err)
	should.Equal([]byte{100}, encoded[8:])
	decoded, err := gocodec.Unmarshal(encoded, (*int8)(nil))
	should.Nil(err)
	should.Equal(int8(100), reflect.ValueOf(decoded).Elem().Interface())
}