package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
	"reflect"
)

func Test_float64(t *testing.T) {
	should := require.New(t)
	encoded, err := gocodec.Marshal(float64(100))
	should.Nil(err)
	should.Equal([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x59, 0x40}, encoded[8:])
	decoded, err := gocodec.Unmarshal(encoded, (*float64)(nil))
	should.Nil(err)
	should.Equal(float64(100), reflect.ValueOf(decoded).Elem().Interface())
}