package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(path string){
	//打开文件,新建文件
	f,err := os.Create(path)
	if nil != err{
		fmt.Println("err = ",err)
		return
	}
	//使用完毕,需要关闭文件
	defer f.Close()

	var buf string
	for i:=0;i<10000;i++{
		//"i=%d\n" 这个字符串写入到buf中
		buf = fmt.Sprintf("i=%d\n",i)
		n ,err := f.WriteString(buf)
		if nil != err{
			fmt.Println("err = ",err," n=",n)
		}
		//fmt.Println("n=",n)
	}
}
func ReadFile(path string){
	//打开文件
	f,err := os.Open(path)
	if nil != err{
		fmt.Println("err =",err)
		return
	}
	//关闭文件
	defer f.Close()

	buf := make([]byte,2*1024)// 2k大小
	//n代表从文件读取内容的长度
	n,err1 := f.Read(buf)
	//读取出错,但是排除结束
	if nil != err1 && nil != io.EOF{
		fmt.Println("err1=",err1)
	}
	//只读取 buf 的容量大小的内容
	fmt.Println("buf len=", len(string(buf[:n])))
}
func main() {
	path := "./file/demo.txt"
	fmt.Println("WriteFile")
	WriteFile(path)
	fmt.Println("ReadFile")
	ReadFile(path)
}
