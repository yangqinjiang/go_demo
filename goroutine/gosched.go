package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(){//子协程
		for i:=0;i<50;i++{
			fmt.Println("go ",i)
		}
	}()

	for i:=0;i<2;i++{
		//让出时间片,先让别人执行,它执行完,
		//再回来执行此协程
		runtime.Gosched()
		fmt.Println("hello")

	}
}
