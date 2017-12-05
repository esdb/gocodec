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
	encoded, err  :=gocodec.Marshal(obj)
	should.Nil(err)
	should.Equal([]byte{}, encoded)
}