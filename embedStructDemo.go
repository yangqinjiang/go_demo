package main

import (
	"fmt"
)

type human struct {
	Name string
	Age  int
	Sex  int
}
type teacher struct {
	human           //匿名,嵌入结核
	JobTitle string //职称
}
type student struct {
	human
	class string //班级
}

//call main
func main() {
	t := teacher{JobTitle: "sir", human: human{Name: "teacher1", Age: 30, Sex: 0}}
	s := student{class: "class 5", human: human{Name: "student1", Age: 15, Sex: 1}}

	//修改t的值
	t.Name = "teacher2" //为了不产生歧义,应该像t.human.Name这样子
	t.human.Age = 25
	fmt.Println(t, s)
}
