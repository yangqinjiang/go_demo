package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if nil != err {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	go func() {
		//从键盘输入内容,给服务器发送内容
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if nil != err {
				fmt.Println("os.Stdin.Read err=", err)
				return
			}
			//发送给服务器
			conn.Write(str[:n])
		}

	}()

	//接收服务器回复的数据,新任务
	//切片缓冲
	buf := make([]byte, 1024)
	for {
		//接收服务器的返回数据
		n, err := conn.Read(buf)
		if nil != err {
			fmt.Println("conn.Read err=", err)
			return
		}
		//打印接收的内容
		fmt.Println("REV: ",string(buf[:n]))
	}

}
