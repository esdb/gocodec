package test

import (
	"testing"
	"github.com/esdb/gocodec"
	"github.com/stretchr/testify/require"
)

func Test_nil_struct_within_struct(t *testing.T) {
	should := require.New(t)
	type SubObject struct {
		length uint
		set    []uint64
	}
	type TestObject struct {
		f1 uint
		f2 uint
		f3 *SubObject
	}
	obj := TestObject{}
	encoded, err := gocodec.Marshal(obj)
	should.Nil(err)
	decoded, err := gocodec.Unmarshal(encoded, (*TestObject)(nil))
	should.Nil(err)
	should.Equal(obj, *decoded.(*TestObject))
}

func Test_struct_within_struct(t *testing.T) {
	should := require.New(t)
	type SubObject struct {
		length uint
		set    []uint64
	}
	type TestObject struct {
		f1 uint
		f2 uint
		f3 *SubObject
	}
	obj := TestObject{f1: 1, f2: 2, f3: &SubObject{length: 3, set: []uint64{100}}}
	encoded, err := gocodec.Marshal(obj)
	should.Nil(err)
	should.Equal([]byte{
		0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x18, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x64, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}, encoded[8:])
	decoded, err := gocodec.Unmarshal(encoded, (*TestObject)(nil))
	should.Nil(err)
	should.Equal(obj, *decoded.(*TestObject))
}