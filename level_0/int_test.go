package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_int(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(100)
	should.Nil(err)
	should.Equal([]byte{100, 0, 0, 0, 0, 0, 0, 0}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (*int)(nil))
	should.Nil(err)
	should.Equal(int(100), reflect.ValueOf(decoded).Elem().Interface())
}