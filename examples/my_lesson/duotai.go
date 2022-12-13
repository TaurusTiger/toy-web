package main

import "fmt"

//func main() {
//	h := &Haojiahuo2{name: "好家伙"}
//	fmt.Println(h.name)
//
//	l := Laolitou2{name: "老李头"}
//	fmt.Println(l.name)
//
//	testInterface(&l)
//
//	h.PlayGame()
//}

func testInterface(k Kongfu) {
	k.Toad()
	k.SixSwords()
}

type Kongfu interface {
	Toad()
	SixSwords()
}

type Haojiahuo2 struct {
	name string
}

func (h Haojiahuo2) Toad() {
	fmt.Println(h.name, "实现了蛤蟆功。。")
}

func (h Haojiahuo2) SixSwords() {
	fmt.Println(h.name, "实现了六脉神剑。。")
}

type Laolitou2 struct {
	name string
}

func (l *Laolitou2) Toad() {
	fmt.Println(l.name, "也实现了蛤蟆功。。")
}

func (l *Laolitou2) SixSwords() {
	fmt.Println(l.name, "也实现了六脉神剑。。")
}

func (h Haojiahuo2) PlayGame() {
	fmt.Println(h.name, "玩游戏。。")
}
