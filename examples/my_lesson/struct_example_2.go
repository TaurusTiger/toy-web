package main

import (
	"fmt"
	"reflect"
)

type node struct {
	_    int
	id   int
	next *node
}

// 字段标签
type user struct {
	name string `昵称`
	sex  byte   `性别`
}

type data2 struct {
	x int
}

func main() {
	n1 := node{id: 1}
	n2 := node{id: 2, next: &n1}

	fmt.Println(n1, &n1, n2)

	// 反射的案例
	u := user{"Tome", 1}
	v := reflect.ValueOf(u)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}

	// 空接口的使用，需要使用实例对象的类型来调用
	d := data2{100}
	var t2 interface{} = &d
	t2.(*data2).x = 200 // 此处的data是struct的自定义类型
	fmt.Println(t2.(*data2).x)
}
