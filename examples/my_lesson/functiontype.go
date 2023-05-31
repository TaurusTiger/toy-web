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
		这里要注意：这里不是接口，需要类型转换的，转换后才可以使用say这个方法
	*/
	fmt.Println(Greeting(english).say("interface"))

	/*
		这里g1是实例对象，say函数用的g(n)实际调用的是具体实例对象的方法体，比如english
	*/
	var g1 Greeting = english
	fmt.Println(g1.say("var interface"))

	g2 := Greeting(french)
	fmt.Println(g2("interface"))
}
