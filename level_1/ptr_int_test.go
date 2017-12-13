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
	should.Equal([]byte{0x20, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 100, 0, 0, 0, 0, 0, 0, 0}, encoded[24:])
	decoded, err := gocodec.Unmarshal(encoded, (**int)(nil))
	should.Nil(err)
	should.Equal(100, **decoded.(**int))
	gocodec.UpdateChecksum(encoded)
	decoded, err = gocodec.Unmarshal(encoded, (**int)(nil))
	should.Nil(err)
	should.Equal(100, **decoded.(**int))
}
