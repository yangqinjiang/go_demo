package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if nil != err {
		fmt.Println("err = ", err)
	}
	defer listener.Close()

	//接收多个用户
	for {
		conn, err := listener.Accept()
		if nil != err {
			fmt.Println("err =", err)
			return
		}
		//接收用户的请求,新建一个协程
		go HandleConn(conn)
	}

}

//接收用户的请求
func HandleConn(conn net.Conn) {
	defer conn.Close()                     //关闭当前用户的连接
	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String();
	fmt.Println(addr," connect successful ",addr)
	//读取用户数据
	buf := make([]byte, 1024) //1024大小 的缓冲区

	for{
		n, err1 := conn.Read(buf)
		if nil != err1 {
			fmt.Println("err1 =", err1)
			return
		}
		txt := strings.TrimSpace(string(buf[:n]))

		fmt.Printf("[%s] :%s\n",addr ,txt ) //切片
		//转成小写,并删除换行符号
		if "exit" == strings.Trim(strings.ToLower(txt),"\r\n"){
			fmt.Println(addr," exit....")
			return
		}
		//把数据转换为大写,再给用户发送
		_,err2 :=conn.Write([]byte(strings.ToUpper(txt)))
		if nil != err2{
			fmt.Println("请求网络错误~")
			return
		}

	}
}
