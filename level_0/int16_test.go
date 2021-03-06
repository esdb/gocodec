package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_int16(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(int16(100))
	should.Nil(err)
	should.Equal([]byte{100, 0}, encoded[8:])
	decoded, err := gocodec.Unmarshal(encoded, (*int16)(nil))
	should.Nil(err)
	should.Equal(int16(100), reflect.ValueOf(decoded).Elem().Interface())
}