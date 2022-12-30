package main

import "fmt"

//import "fmt"
//
//// 这个只是一个“函数类型”
//type Greeing func(name string) string
//
//func say(g Greeing, n string) {
//	fmt.Println(g(n))
//}
//
//// 这个是具体的函数对象
//func english(name string) string {
//	return "hello, " + name
//}
//
//func main() {
//	say(english, "world")
//}

/*
函数类型是表示所有包含相同参数和返回类型的函数集合
*/
type Greeting func(name string) string

func (g Greeting) say(n string) string {
	return g(n)
}

func english(name string) string {
	return "hello, " + name
}

func french(name string) string {
	return "Bonjour, " + name
}

func main() {
	/*
		这里要注意：这里不是接口，需要类型转换的
	*/
	fmt.Println(Greeting(english).say("interface"))

	g := Greeting(french)
	fmt.Println(g("interface"))
}
