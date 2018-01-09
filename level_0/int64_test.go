package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_int64(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(int64(100))
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0}, encoded[16:])
	decoded, err := gocodec.Unmarshal(encoded, (*int64)(nil))
	should.Nil(err)
	should.Equal(int64(100), reflect.ValueOf(decoded).Elem().Interface())
}