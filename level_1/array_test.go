package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
)

func Test_single_ptr_int_array(t *testing.T) {
	should := require.New(t)
	type TestObject [1]*uint8
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

func Test_two_ptrs_in_array(t *testing.T) {
	should := require.New(t)
	type TestObject [2]*uint8
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