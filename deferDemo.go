package main

import (
	"fmt"
)

func main() {
	r := noFunc()
	fmt.Println(r)
}

/*
 Deferred functions may read and
 assign to the returning function's named return values.
 ->http://blog.golang.org/defer-panic-and-recover

 defer 函数可以读取和指定(修改??)返回值
 因为i就是返回值,如果retunr (返回)1之后,defer函数再执行
 就会修改了返回值i 的值,结果在调用的地方,返回的是2
*/
func noFunc() (i int) { //i 是返回值
	j := 9
	defer func() { i++; j++ }()
	return 1
}
