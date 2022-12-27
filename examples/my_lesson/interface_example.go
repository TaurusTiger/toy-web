package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
这里实现了io.Writer这个接口（实现了Write这个方法），所以UpperWriter这个结构类型可以被Fprintln当做第一个参数使用
*/
type UpperWriter struct {
	io.Writer
}

func (w *UpperWriter) Write(data []byte) (n int, err error) {
	return w.Writer.Write(bytes.ToUpper(data))
}

/*
任意隐式满足 fmt.Stringer 接口的对象都可以打印，不满足 fmt.Stringer 接口的依然可以通过反射的技术打印
对于每个要打印的对象，如果满足了 fmt.Stringer 接口，则默认使用对象的 String 方法返回的结果打印
*/
type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")

	fmt.Fprintln(os.Stdout, UpperString("hello, tiger"))
}
