package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if nil != err {
		fmt.Println("err = ", err)
	}
	defer listener.Close()

	fmt.Println("阻塞,等待用户连接...")
	//阻塞,等待用户连接

	conn, err := listener.Accept()
	if nil != err {
		fmt.Println("err =", err)
		return
	}
	//接收用户的请求
	buf := make([]byte, 1024) //1024大小 的缓冲区
	n, err1 := conn.Read(buf)
	if nil != err1 {
		fmt.Println("err1 =", err1)
		return
	}
	defer conn.Close()//关闭当前用户的连接
	fmt.Println("buf = ", string(buf[:n])) //切片

}
