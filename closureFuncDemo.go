package main

import (
	"fmt"
)

/*
Go Func闭包的使用
*/
func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}
func closure(x int) func(int) int { //返回一个func  参数是int ,返回值也是int
	fmt.Printf("%v , %p\n", x, &x) //只执行一次
	return func(y int) int {
		fmt.Printf("x: %p,y:%p \n", &x, &y)
		return x + y
	}
}

/*
output:
	10 , 0xc080000038
	x: 0xc080000038,y:0xc080000058
	11
	x: 0xc080000038,y:0xc080000070
	12

*/
