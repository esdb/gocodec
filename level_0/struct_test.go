package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_simple_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 int
		Field2 int
	}
	obj := TestObject{1, 2}
	encoded, err := gocodec.Marshal(obj)
	should.Nil(err)
	should.Equal([]byte{
		0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (*TestObject)(nil))
	should.Nil(err)
	should.Equal(obj, *decoded.(*TestObject))
}

func Test_struct_signature(t *testing.T) {
	should := require.New(t)
	type TestVersion1 struct {
		Field1 int
	}
	type TestVersion2 struct {
		Field1 uint
		Field2 uint
	}
	obj := TestVersion1{1}
	encoded, err := gocodec.Marshal(obj)
	should.Nil(err)
	should.Equal([]byte{
		0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (*TestVersion2)(nil), (*TestVersion1)(nil))
	should.Nil(err)
	should.Equal(obj, *decoded.(*TestVersion1))
}
