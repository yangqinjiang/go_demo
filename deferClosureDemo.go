package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		//i = 0,1,2,3
		defer fmt.Println("defer i= ", i)
		//for完成后,在闭包函数里,i的值是4
		defer func() { fmt.Println("defer_closure i = ", i) }() //()表示执行这个闭包函数
		//这一句,不是defer要执行的函数,首先执行,
		//在for最后,i的值是4,
		fmt.Println("%p\n", &i) //i的内存地址不变
		fs[i] = func() { fmt.Println("closure i =", i) }
	}

	//
	for _, f := range fs {
		f()
	}
	testDefer()
}

//output:~
/*
	closure i = 4
	closure i = 4
	closure i = 4
	closure i = 4
	defer_closure i =  4
	defer i= 3
	defer_closure i =  4
	defer i= 2
	defer_closure i =  4
	defer i= 1
	defer_closure i =  4
	defer i= 0


*/

func testDefer() {
	for i := 0; i < 4; i++ {
		defer func() {
			fmt.Println(i) //i的内存地址不变
		}()
	}
}

/*output:
4
4
4
4


*/
