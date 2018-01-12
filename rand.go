package myrand

import (
	"crypto/rand"
	"io"
)

// TODO: find a better prng
var RandReader = rand.Reader

// RandFail is a Reader that doesn't deliver any data
var RandFail = eofReader{}

type eofReader struct{}

func (e eofReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}
