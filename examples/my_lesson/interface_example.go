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
// 又重新的封装了一下io.Writer，自定义的Write收取参数，然后再调用默认的io.Writer实现的对象的Wirte方法
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

/*
interface的断言
*/
type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")

	fmt.Fprintln(os.Stdout, UpperString("hello, tiger"))

	s := S{}
	var i I //声明 i
	i = &s  //赋值 s 到 i
	f(i)

	if t, ok := i.(*S); ok {
		fmt.Println("s.implement I ", t, ok)
	}

	switch i.(type) {
	case *S:
		fmt.Println("ok")
	}

}
