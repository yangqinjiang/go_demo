package main

import (
	"fmt"
	"reflect"
)

/*
通过反射来获取struct的信息
*/
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello,world.")
}
func main() {
	u := User{1, "OK", 12}
	Info(u)
}
func Info(o interface{}) { //以空接口为参数,接收任何类型的数据
	t := reflect.TypeOf(o) //返回o的类型
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o) //返回o的值
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v =%v\n", f.Name, f.Type, val)
	}
}
