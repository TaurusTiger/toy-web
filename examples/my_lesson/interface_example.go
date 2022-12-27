package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
这里使用了
*/
type UpperWriter struct {
	io.Writer
}

func (w *UpperWriter) Write(data []byte) (n int, err error) {
	return w.Writer.Write(bytes.ToUpper(data))
}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")

	fmt.Fprintln(os.Stdout, UpperString("hello, tiger"))
}
