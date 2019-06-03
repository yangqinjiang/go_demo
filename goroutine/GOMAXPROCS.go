package main

import (
	"fmt"
	"runtime"
)

func main() {
	//指定以x核运算
	x := 1
	n:=runtime.GOMAXPROCS(x)//CPU数量
	fmt.Println("n=",n)
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
