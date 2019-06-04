package main

import "fmt"
//通过select实现斐波那契数列
// fibonacci  1 1 2 3 5 8 13 21...
func main() {

	ch:= make(chan int)//数字通信
	quit := make(chan bool)//程序是否结束

	//消费者,从channel读取内容
	//新建协程
	go func() {
		for i:=0;i<10;i++{
			num := <-ch
			fmt.Print(num," ")
		}
		//可以停止
		quit <- true
	}()
	//生产者,产生数字,写入channel
	fibonacci(ch,quit)

}
//ch 只写, quit 只读
func fibonacci(ch chan<- int,quit <-chan bool)  {
	x,y := 1 ,1
	for{
		//监听channel数据的滚动
		select {
		case ch <- x:
			x,y = y ,x+y // 通过select实现斐波那契数列
		case flag :=<-quit:
				fmt.Println("flag= ",flag)
				return//停止,返回

		}
	}
}
