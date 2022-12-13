package main

import "fmt"

type Role struct {
	Name    string
	Ability string
	Level   int
	Kill    float64
}

func (r Role) Kungf() {
	fmt.Printf("我是:%s，我的武功:%s,已经练到%d级了，杀伤力%.1f\n", r.Name, r.Ability, r.Level, r.Kill)
}

type Haojiahuo int

func (h Haojiahuo) Add(num int) int {
	return int(h) + num
}
func (h Haojiahuo) Add2(num Haojiahuo) Haojiahuo {
	return h + num
}

func (h Haojiahuo) Clear() bool {
	h = 0
	return h == 0
}

//func main() {
//	rwx := &Role{"任我行", "吸星大法", 8, 10}
//	rwx.Kungf()
//
//	var h Haojiahuo
//	fmt.Println(h.Clear())
//	fmt.Println(h.Add(2))
//	fmt.Println(h.Clear())
//	fmt.Println(h.Add(6))
//	fmt.Println(h.Clear())
//	fmt.Println(h)
//
//}
