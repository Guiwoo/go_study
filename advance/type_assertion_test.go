package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestTypeAssertion01(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter)

	w = new(bytes.Buffer)
	rw = w.(io.ReadWriter)

	fmt.Println(rw.Write([]byte("things")))
}
