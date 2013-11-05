package main

import (
	"fmt"
)

type TZ int //
//TZ 的自定义func
func (a *TZ) Print() {
	fmt.Println("TZ Print")

}
func main() {
	var a TZ
	a.Print()       //这样调用Print函数
	(*TZ).Print(&a) //或者这样调用Print函数
}
