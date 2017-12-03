package test

import (
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
	"testing"
)

func Test_ptr_int(t *testing.T) {
	should := require.New(t)
	val := 100
	encoded, err := gocodec.Marshal(&val)
	should.Nil(err)
	should.Equal([]byte{8, 0, 0, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0}, encoded)
	var pVal *int
	should.Nil(gocodec.Unmarshal(encoded, &pVal))
	should.Equal(100, *pVal)
}