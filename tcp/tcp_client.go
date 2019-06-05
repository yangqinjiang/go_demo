package main

import (
	"fmt"
	"net"
)

func main() {
	conn,err := net.Dial("tcp","127.0.0.1:8000")
	if nil != err{
		fmt.Println("err = ",err)
		return
	}
	defer conn.Close()
	//发送数据
	conn.Write([]byte("are u ok?"))

}
