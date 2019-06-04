package main

import (
	"fmt"
	"time"
)

func main() {
	<-time.After(2*time.Second)//定时2s,阻塞2s后,产生一个事件,往channel写内容
	fmt.Println("时间到")

	t := time.AfterFunc(2*time.Second, func() {
		fmt.Println("又一个时间到")
	})
	time.Sleep(3*time.Second)
	t.Stop()


}
