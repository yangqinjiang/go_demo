package main

import (
	"fmt"
)

//使用select和chan来实现[生产者][消费者]模式
func main() {
	c := make(chan int)
	//消费者
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	//生产者
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
		
	}
}
