package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//demo1()
	//demo2()
	//compute()
	//	compute2()
	UseWaitGroup()
}

//使用sync.WaitGrounp来实现并发的等待
func UseWaitGroup() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go3(&wg, i)
	}
	wg.Wait()
}

//wg 传地址
func Go3(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}

//使用chan的缓冲池,来解决comput()方法的bug
func compute2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//chan缓冲池
	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go Go2(c, i)
	}
	//读取c
	for i := 0; i < 10; i++ {
		<-c
	}
}
func Go2(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	//将true写入到c
	c <- true
}

//下面的代码是有bug的,因为是并发的.
//在Go方法中,很有可能会先执行index=9的情况
//而之后的index不会执行了
func compute() {
	//设置最大的CPU核数
	//实现并发执行
	//如果不设置,则是顺序执行
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)

	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	<-c
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	//当执行到最后的一个for index里,则将true写入到c
	//这里不能这样子写
	if index == 9 {
		c <- true
	}
}
func demo2() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!!!!")
		c <- true //将true保存到c 中
		close(c)  //关闭c
	}()
	//通过迭代c,读取c的值
	for v := range c {
		fmt.Println(v)
	}
}

//channel的简单使用
func demo1() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true //将true写入c中
	}()
	<-c //将c的值读出来,如果c没有值,则阻塞
}
