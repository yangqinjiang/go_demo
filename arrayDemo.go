package main

import (
	"fmt"
)

/*
当len>cap里,cap会自增一倍于自身,并且引起copy操作,到新内存块
%v 打印值,%p打印指针
*/
func main() {
	s1 := make([]int, 3, 6) //3 len 6 cap
	fmt.Println("0.....")
	fmt.Println("len of s1", len(s1)) //s1的元素长度
	fmt.Println("cap of s1", cap(s1)) //s1的容量
	fmt.Printf("%p\n", s1)            //打印指针
	fmt.Println("1,append.....")
	s1 = append(s1, 1, 2, 3)          //再append三个元素
	fmt.Println("len of s1", len(s1)) //s1的元素长度
	fmt.Println("cap of s1", cap(s1)) //s1的容量
	fmt.Printf("%v %p\n", s1, s1)
	fmt.Println("2,append.....")
	s1 = append(s1, 1) //再append一个元素,len会大于cap,则引发copy到新的内存块
	fmt.Printf("%v %p\n", s1, s1)
	fmt.Println("len of s1", len(s1)) //s1的元素长度
	fmt.Println("cap of s1", cap(s1)) //s1的容量
	Sort()
	SliceTest()
}

/*冒泡排序法*/
func Sort() {
	a := [...]int{5, 2, 6, 3, 9}
	fmt.Println("冒泡排序前:", a)

	num := len(a) //获取a数组的长度
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i] //每次得到的temp的地址是不一样的
				fmt.Printf("%v\n", &temp)
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println("冒泡排序后:", a)
}
func SliceTest() {
	a := [10]int{}
	a[1] = 2 //修改
	fmt.Println(a)

	//通过p来指向a
	p := new([15]int)
	p[2] = 3
	fmt.Println(p)
	fmt.Println("len: ", len(p))
	fmt.Println("cap:  ", cap(p))

}
