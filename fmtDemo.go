package main

import (
	"fmt"
)

type Person struct {
	Name string
}

/*
	fmt包的使用
*/
func main() {
	var i int = 10
	fmt.Printf("value %v \n", i)  //value
	fmt.Printf("十六进制 %x \n", i)   //十六进制
	fmt.Printf("value %#v \n", i) //value
	fmt.Printf("Type is [%T] \n", i)
	fmt.Printf("%% \n", i) //将%转义
	fmt.Printf("%%t : %t \n", true)
	//
	fmt.Printf("%v base 2  : %b \n", i, i) //将i(integer)以二进制输出
	//%d	base 10
	fmt.Printf("%v base 10  : %d \n", i, i)
	//%o	base 8
	fmt.Printf("%v base 8  : %o \n", i, i)
	//%x	base 16, with lower-case letters for a-f

	fmt.Printf("%v base 16 ,lower-case :%x \n", i, i)
	//%X	base 16, with upper-case letters for A-F
	fmt.Printf("%v base 16 ,upper-case :%X \n", i, i)
	//%q	a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%%q :%q \n", i)
	//%U	Unicode format: U+1234; same as "U+%04X"
	fmt.Printf("%U \n", " U+1234")
	//%x	base 16, lower-case, two characters per byte
	fmt.Printf("%x \n", "HELLO,WORLD") //以小写输出byte
	//%X	base 16, upper-case, two characters per byte
	fmt.Printf("%X \n", "HELLO,WORLD") //以大写输出byte
	var p = []int{1, 2, 3, 4}
	//%p	base 16 notation, with leading 0x
	fmt.Printf("Pointer:  %p\n", p)
	var person = new(Person) //new函数的使用,使用new,才有pointer
	person.Name = "hi,i am ..."
	fmt.Printf("person pointer:  %p , name :%s\n", person, person.Name)
	fmt.Println(-100)
	fmt.Printf("%v\n", -100)
	fmt.Printf("%.4g\n", 123.45)
}
