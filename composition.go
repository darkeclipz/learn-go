// https://www.youtube.com/watch?v=-gW7oSFxT2I&list=PL0xRBLFXXsP7-0IVCmoo2FEWBrQzfH2l8

package foohandler

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	*bytes.Reader
	buffer *bytes.Buffer
}

func main() {
	payload := []byte("hello high value software engineer")
	hashAndBroadcast(NewHashReader(payload))
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buffer: bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	ret := sha1.Sum(h.buffer.Bytes())
	return hex.EncodeToString(ret[:])
}

func hashAndBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println(hash)
	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of bytes:", string(b))
	return nil
}
