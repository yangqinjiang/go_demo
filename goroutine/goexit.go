package main

import (
	"fmt"
	"runtime"
	"time"
)

func test()  {
	defer fmt.Println("c")
	//return //终止此函数, 输出 acb
	runtime.Goexit()//终止所在的协程, 输出 ac
	fmt.Println("d")
}
func main() {
	go func(){
		fmt.Println("a")

		//调用了别的函数
		test()

		fmt.Println("b")
	}()


	//不退出
	for{
		time.Sleep(1*time.Second)
	}
}
