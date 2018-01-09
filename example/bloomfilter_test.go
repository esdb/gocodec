package test

import (
	"testing"
	"github.com/willf/bloom"
	"github.com/stretchr/testify/require"
	"github.com/esdb/gocodec"
	"io/ioutil"
	"github.com/edsrzf/mmap-go"
	"os"
	"github.com/json-iterator/go"
)

func Test_bloomfilter(t *testing.T) {
	should := require.New(t)
	f := bloom.New(1000*1024, 4)
	f.Add([]byte("hello"))
	f.Add([]byte("world"))
	should.True(f.Test([]byte("hello")))
	should.False(f.Test([]byte("hi")))
	encoded, err := gocodec.Marshal(*f)
	should.Nil(err)
	should.NotNil(encoded)
	ioutil.WriteFile("/tmp/bloomfilter.bin", encoded, 0666)
	decoded, err := gocodec.ReadonlyConfig.Unmarshal(encoded, (*bloom.BloomFilter)(nil))
	should.Nil(err)
	f2 := decoded.(*bloom.BloomFilter)
	should.True(f2.Test([]byte("hello")))
	should.False(f2.Test([]byte("hi")))
	decoded, err = gocodec.Unmarshal(encoded, (*bloom.BloomFilter)(nil))
	should.Nil(err)
}

func Test_mmap(t *testing.T) {
	should := require.New(t)
	f, err := os.Open("/tmp/bloomfilter.bin")
	should.Nil(err)
	mem, err := mmap.Map(f, mmap.COPY, 0)
	should.Nil(err)
	decoded, err := gocodec.Unmarshal(mem, (*bloom.BloomFilter)(nil))
	should.Nil(err)
	f2 := decoded.(*bloom.BloomFilter)
	should.True(f2.Test([]byte("hello")))
	should.False(f2.Test([]byte("hi")))
	mem.Unmap()
}

func Test_json(t *testing.T) {
	should := require.New(t)
	f := bloom.New(1000*1024, 4)
	f.Add([]byte("hello"))
	f.Add([]byte("world"))
	encoded, err := jsoniter.Marshal(f)
	should.Nil(err)
	ioutil.WriteFile("/tmp/bloomfilter.json", encoded, 0666)
}

func Benchmark(b *testing.B) {
	api := gocodec.Config{}.Froze()
	b.Run("gocodec", func(b *testing.B) {
		f, _ := os.OpenFile("/tmp/bloomfilter.bin", os.O_RDONLY, 0)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			mem, _ := mmap.Map(f, mmap.COPY, 0)
			_, err := api.Unmarshal(mem, (*bloom.BloomFilter)(nil))
			if err != nil {
				b.Error(err)
			}
			mem.Unmap()
		}
	})
	//b.Run("json", func(b *testing.B) {
	//	f, _ := os.Open("/tmp/bloomfilter.json")
	//	b.ReportAllocs()
	//	for i := 0; i < b.N; i++ {
	//		var f2 bloom.BloomFilter
	//		//bytes, _ := ioutil.ReadFile("/tmp/bloomfilter.bin")
	//		mem, _ := mmap.Map(f, mmap.COPY, 0)
	//		err := jsoniter.Unmarshal(mem, &f2)
	//		if err != nil {
	//			b.Error(err)
	//		}
	//		mem.Unmap()
	//	}
	//})
}
