package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_uintptr(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(uintptr(100))
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0}, encoded[8:])
	decoded, err := gocodec.Unmarshal(encoded, (*uintptr)(nil))
	should.Nil(err)
	should.Equal(uintptr(100), reflect.ValueOf(decoded).Elem().Interface())
}