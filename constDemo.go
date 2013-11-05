package main

import (
	"fmt"
)

/*
使用[常量]来表示[计算机存储容量单位]的值
*/
const (
	//B          = 8                //byte
	_ = iota //reset init value to 0
	B
	KB float64 = 1 << (iota * 10) //kb =1024b
	MB                            //mb=1024kb
	GB                            //...
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
	fmt.Println("...........")
	fmt.Println(KB / B)  //1 KB =1024B
	fmt.Println(MB / KB) //1MB =1024KB
	//.......//其它的类似上两句代码的结果

	fmt.Println("打印[常量]的指针地址:")
	fmt.Printf("B: %p,KB %p, MB: %p", B, KB, MB) //打印出B,KB,MB的值,不是打印出它们的内存地址
}
