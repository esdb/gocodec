package test

import (
	"testing"
	"github.com/willf/bloom"
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
)

func Test_bloomfilter(t *testing.T) {
	should := require.New(t)
	f := bloom.New(1000, 4)
	f.Add([]byte("hello"))
	f.Add([]byte("world"))
	should.True(f.Test([]byte("hello")))
	should.False(f.Test([]byte("hi")))
	encoded, err := gocodec.Marshal(*f)
	should.Nil(err)
	should.NotNil(encoded)
	var f2 bloom.BloomFilter
	should.Nil(gocodec.Unmarshal(encoded, &f2))
	should.True(f2.Test([]byte("hello")))
	should.False(f2.Test([]byte("hi")))
}
