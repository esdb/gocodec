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
		0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x1,
	}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (*TestObject)(nil))
	should.Nil(err)
	should.Equal(obj, *decoded.(*TestObject))
	gocodec.UpdateChecksum(encoded)
	decoded, err = gocodec.Unmarshal(encoded, (*TestObject)(nil))
	should.Nil(err)
	should.Equal(obj, *decoded.(*TestObject))
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
		0x28, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x29, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		1,
		1,
	}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (*TestObject)(nil))
	should.Nil(err)
	should.Equal(obj, *decoded.(*TestObject))
}

func Test_slice_in_struct(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 uint
		Field2 []uint
	}
	obj := TestObject{Field1: 1, Field2: []uint{2,3}}
	encoded, err := gocodec.Marshal(obj)
	should.Nil(err)
	should.Equal([]byte{
		0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x38, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (*TestObject)(nil))
	should.Nil(err)
	should.Equal(obj, *decoded.(*TestObject))
}

func Test_multiple_struct(t *testing.T) {
	should := require.New(t)
	type TestObj struct {
		Field []int
	}
	stream := gocodec.NewStream(nil)
	stream.Marshal(TestObj{[]int{1}})
	stream.Marshal(TestObj{[]int{2}})
	should.Nil(stream.Error)
	iter := gocodec.NewIterator(stream.Buffer())
	obj := iter.Unmarshal((*TestObj)(nil))
	should.Nil(iter.Error)
	should.Equal([]int{1}, obj.(*TestObj).Field)
	obj = iter.Unmarshal((*TestObj)(nil))
	should.Nil(iter.Error)
	should.Equal([]int{2}, obj.(*TestObj).Field)
}
