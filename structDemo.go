package main

import (
	"fmt"
)

/*struct与函数*/

//A struct
type A struct {
	Name string //公有变量,任何地方都可以访问
	age  int    //私有变量,只能在该包(当前是main)里能访问,
}

func (a *A) Print() {
	a.Name = "this is Name of struct A"
	fmt.Println("this is func Print of struct A")
}

//B struct
type B struct {
	Name string
}

func (b *B) Print() {
	b.Name = "this is Name of struct B"
	fmt.Println("this is func Print of struct B")
}

//call main 
func main() {
	a := A{}
	a.age = 20
	a.Print()
	fmt.Printf("Name:%s , age: %v\n", a.Name, a.age)

	b := B{}
	b.Print()
	fmt.Println(b.Name)
}
