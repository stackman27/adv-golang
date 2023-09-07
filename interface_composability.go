package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

// INTERFACE COMPOSABILITY

// Interface Composability pattern is used to create a new interface (HashReader) that extends the
// functionality of an existing interface (io.Reader). This pattern allows you to combine multiple interfaces and their
// implementations to achieve a more specific and composable behavior.
type HashReader interface {
	io.Reader
	hash() string
}

func Incmain() {
	payload := []byte("Hello high value software engineer")
	hashedPayload := NewHashReader(payload)
	 
	hashAndBroadcast(hashedPayload)
}

// we are composing bytes reader which already implements io reader into our hashReader 
type hashReader struct { 
	*bytes.Reader 
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf: bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string { 
	return hex.EncodeToString(h.buf.Bytes())
}


func hashAndBroadcast(hr HashReader) error {
	hash := hr.hash()
	fmt.Println("hash: ", hash)
	
	return broadcast(hr) 
}


func broadcast(r HashReader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err 
	}

	fmt.Println("string of the bytes: ", string(b))
	return nil 
}
 