package main

import (
	"fmt"
)

var (
	a     [30000]byte
	prog  = "++++++++++[>++++++++++<-]>++++.+."
	p, pc int
)

func loop(inc int) {
	for i := inc; i != 0; pc += inc {
		switch prog[pc+inc] {
		case '[':
			i++
		case ']':
			i--
		}
	}
}
func main() {
	for {
		switch prog[pc] {
		case '>': //下一个单元作为当前数据单元
			p++
		case '<': //上一个单元作为当前数据单元
			p--
		case '+': //使当前数据单元的值增1
			a[p]++
		case '-': //使当前数据单元的值减1
			a[p]--
		case '.': //把当前数据单元的值作为字符输出
			fmt.Print(string(a[p]))
		case '[': //如果当前数据单元的值为0,下一条指令在对应的]后
			if a[p] == 0 {
				loop(1)
			}
		case ']': //如果当前数据单元的值不为0,下一条指令在对应的[后
			if a[p] != 0 {
				loop(-1)
			}
		default:
			fmt.Println("Illegal instruction")
		}
		pc++
		if pc == len(prog) {
			return
		}
	}
}
