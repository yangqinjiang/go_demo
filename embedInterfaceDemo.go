package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter //Connecter作为嵌入接口
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

//不能是这样func(pc *PhoneConnecter) Name() string
func (pc PhoneConnecter) Name() string {
	return pc.name
}

//不能是这样func(pc *PhoneConnecter) Connect()
func (pc PhoneConnecter) Connect() {
	fmt.Printf("pc pointer: %p\n", &pc)
	fmt.Println("Connected:", pc.name)
}
func main() {
	//var a USB
	a := PhoneConnecter{"PhoneConnecter"}
	fmt.Printf(" a pointer: %p\n", &a)
	a.Connect() //调用Connecter接口的Connect
	Disconnect(a)
	//转成Connecter接口类型
	var c Connecter
	c = Connecter(a)
	fmt.Printf(" c pointer: %p\n", &c) //输出的指针与被转换的不一样,说明是值传递,copy
	c.Connect()
	Disconnect(c)
	Disconnect_2(c)

	a.name = "change this PhoneConnecter Name"
	fmt.Println("改变原来对象的name字段后,name的值是:", a.name)
	//不能改变copy版本的字段值
	//这个c对象与a不属于同一个的
	c.Connect() //打印原来对象的name字段
	checkNil()
}

//使用if语句
func Disconnect(usb interface{}) {
	//条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用
	//判断类型
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("--Disconnected:", pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}

//使用switch语句
func Disconnect_2(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("++Disconnected:", v.name)
	default:
		fmt.Println("++Unknown decive.")
	}
}
func checkNil() {
	//只有当接口存储的类型和对象都是为nil时,接口才是nil
	fmt.Println("只有当接口存储的类型和对象都是为nil时,接口才是nil")

	var a interface{}
	fmt.Println("接口为nil")
	fmt.Println(a == nil) //output:true

	var p *int = nil      //值为nil,p的值不为空
	fmt.Printf("%v\n", p) //&p是一个地址
	fmt.Printf("%v\n", a)
	fmt.Println("定义一个 p *int,初始为nil:并a=p")
	a = p
	fmt.Println(a == nil) //output:false

}

/*
func(pc *PhoneConnecter) Name() string
func(pc *PhoneConnecter) Connect()
如果这样写,会有报错
	.\embedInterfaceDemo.go:31: cannot convert a (type PhoneConnecter) to type Connecter:
		PhoneConnecter does not implement Connecter (Connect method requires pointer receiver)
*/
