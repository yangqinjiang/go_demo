package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器,设置时间为2s,2s后,向time通道写内容(当前时间)
	timer := time.NewTimer(2*time.Second)
	fmt.Println("当前时间: ",time.Now())

	//2s后,往timer.c写数据,有数据后,就可以读取
	// channel没有数据前,会阻塞
	t := <- timer.C
	fmt.Println("t = ",t)


	//time.NewTimer只会响应一次
	timer2 := time.NewTimer(1*time.Second)
	for{
		//for第二次时,
		//fatal error: all goroutines are asleep - deadlock!
		<-timer2.C
		fmt.Println("时间到")
	}
}
