package main

import (
	"fmt"
	"io"
	"os"
)
/**
copy_file_command src.txt dst.txt
 */
func main() {
	list := os.Args//获取命令行参数
	if len(list) != 3{
		fmt.Println("Usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName{
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	//只读方式打开源文件
	sF,err1 := os.Open(srcFileName)
	if err1 != nil{
		fmt.Println("open src file error=",err1)
		return
	}
	//新建目的文件
	dF,err2 := os.Create(dstFileName)
	if err2 != nil{
		fmt.Println("create dst file error=",err2)
		return
	}

	//操作完毕,需要关闭文件
	defer  sF.Close()
	defer  dF.Close()

	//核心处理,从源文件读取内容,往目的文件写
	//注意, 读多少, 则写多少
	buf := make([]byte,4*1024) // 4kd大小临时缓冲区
	for  {
		read_bytes_number, err := sF.Read(buf)//从源文件读取内容
		if nil != err{
			if io.EOF == err{//文件读取完毕
				fmt.Println("copy over")
				break
			}
			fmt.Println("copy error=",err)
		}
		//往目录文件写,读多少,写多少
		_,err = dF.Write(buf[:read_bytes_number])
		if nil!= err{
			fmt.Println("write error=",err)
		}
	}
}
