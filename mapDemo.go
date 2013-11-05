package main

import (
	"fmt"
	"sort" //排序包
)

//map没有排序功能,使用slice对其Key排序
func main() {
	m := map[int]string{7: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		//s[i++]=k//不能将i++写在这里
		s[i] = k
		i++
	}
	sort.Ints(s) //对slice排序
	fmt.Println(m)
	fmt.Println(s)
	fmt.Println(m[s[0]]) //再使用s的值,作为m的key,取出相应的值
	makeMap()
	makeMap2()
	makeMapWithCond()
}

//初始化一个map数组
func makeMap() {
	sm := make([]map[int]string, 5) //cap 为 5
	for i := range sm {
		sm[i] = make(map[int]string, 1) //初始化数组元素
		sm[i][2] = "OK"                 //2是key,OK是value
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
func makeMap2() {
	sm := make([]map[int]string, 5)
	for _, v := range sm { //makeMap func的另一个版本
		v = make(map[int]string, 1)
		v[100] = "this is 1000"
		fmt.Println(v)
	}
	fmt.Println(sm)
}

//根据map的元素是否已经初始化的条件,来执行程序
func makeMapWithCond() {
	fmt.Println("call  Func makeMapWithCond")
	var m map[int]map[int]string
	m = make(map[int]map[int]string) //复杂的写法
	//if语句可以这样写
	if _, ok := m[1][2]; !ok {
		fmt.Println("m[1][2] not init,will be init.!!!")
	}

	a, ok := m[2][1] //判断是否初始化了
	if !ok {
		fmt.Println("m[2][1] not init,will be init.")
		m[2] = make(map[int]string) //如果没有初始化,则会执行make
	}
	m[2][1] = "have init value"
	a, ok = m[2][1]
	fmt.Println("m[2][1] have init,not to be init.")
	fmt.Println(a)
}
