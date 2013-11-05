package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1: //如果c1已经关闭,则ok是false
				fmt.Printf("c1 ok =v ,%p\n", ok, &ok)
				if !ok { //如果ok=false,表示已经关闭,则跳出for循环
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2: //如果c2已经关闭,则ok是false
				fmt.Printf("c2 ok =v ,%p\n", ok, &ok)
				if !ok { //如果ok=false(ok不同于上面的ok),表示已经关闭,则跳出for循环
					o <- true
					break
				}
				fmt.Println("c1", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"
	close(c1) //实际上,关闭一个chan,就能退出for了
	close(c2)
	<-o
	//如果o是在缓冲区,则使用for循环,将o读出
}
