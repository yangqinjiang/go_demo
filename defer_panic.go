package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}
func A() {
	fmt.Println("call Func A")
}
func B() {
	defer func() { //匿名函数
		if err := recover(); err != nil {
			fmt.Println("Recover in Func B")
		}
	}()
	//panic必须放在defer之后,否则defer的func执行不了
	panic("Panic in Func B")
}
func C() {
	fmt.Println("call Func C")
}
