package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
)

func Test_single_ptr_in_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 *uint8
	}
	one := uint8(1)
	obj := TestObject{&one}
	encoded, err := gocodec.Marshal(obj)
	should.Nil(err)
	should.Equal([]byte{
		0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x1,
	}, encoded[8:])
	var decoded TestObject
	should.Nil(gocodec.Unmarshal(encoded, &decoded))
	should.Equal(obj, decoded)
}

func Test_two_ptrs_in_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 *uint8
		Field2 *uint8
	}
	one := uint8(1)
	obj := TestObject{&one, &one}
	encoded, err := gocodec.Marshal(obj)
	should.Nil(err)
	should.Equal([]byte{
		16, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		1,
		1,
	}, encoded[8:])
	var decoded TestObject
	should.Nil(gocodec.Unmarshal(encoded, &decoded))
	should.Equal(obj, decoded)
}
