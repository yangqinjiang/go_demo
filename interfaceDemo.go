package main

import (
	"fmt"
)

/*PhoneConnecter类型实例了USB的接口*/

//USB接口
type USB interface {
	Name() string
	Connect()
}
type PhoneConnecter struct {
	name string //字段
}

//鸭子类型
//PhoneConnecter结构实现Name方法
func (pc PhoneConnecter) Name() string {
	return pc.name
} //PhoneConnecter实现Connect方法
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB
	//a 的类型是USB接口,能被PhoneConnecter实例赋值
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	fmt.Println("call a.Name()", a.Name())
	Disconnect(a)
	receiverAnyType(a)
}

//能接收USB接口的实例
func Disconnect(usb USB) {
	fmt.Println("Disconnect.")
}

//以空接口为参数,能接收所有类型
func receiverAnyType(o interface{}) {
	fmt.Println("receiverAnyType interface{}")
}
