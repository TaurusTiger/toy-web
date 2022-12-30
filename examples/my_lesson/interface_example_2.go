package main

/*
定义函数类型，让相同签名的函数自动实现某个接口。
*/
import "fmt"

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	var t fmt.Stringer = FuncString(func() string { // 转换类型，使其实现Stringer接口
		return "hello, interface!"
	})

	fmt.Println(t)
}
