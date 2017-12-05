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
	var decoded TestObject
	should.Nil(gocodec.Unmarshal(encoded, &decoded))
	should.Equal(obj, decoded)
}