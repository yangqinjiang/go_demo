package main

/*
import "fmt"
import "errors"
*/
import (
	"errors"
	"fmt"
)

func main() {
	array_example()
	complex_example()
	string_example()
	error_example()
	iota_example()
	fmt.Println("hello st3")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Print("...")
	}
	//Print
	fmt.Print("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
	var variableName int //=100
	variableName = 100
	fmt.Println("variableName:", variableName)
	var vname1, vname2 = 1, "string"
	vname3, vname4 := 1, "string"
	fmt.Println(vname1)
	fmt.Println(vname2)
	fmt.Println(vname3)
	fmt.Println(vname4)

	_, b := 2, 3
	fmt.Println(b)
}

//声明了全局变量,而不使用的话,编译会给通过的
var isActive bool                   //全局变量声明
var enabled, disables = true, false // 忽略类型的声明r
func test() {
	var available bool //一般声明
	valib := false     //简短声明
	available = true   //赋值操作
	fmt.Println(valib, available)
}

//常量
const PI = 3.1415926
const i = 10000
const MaxThread = 10
const prefix = "astaxie_"

func int_not_int32() {
	var a int8
	var b int
	a = 10
	b = 100
	a--
	b--
	//a+b的类型不同,不能相加
	//c := a + b //invalid operation: a + b (mismatched types int8 and int)
}
func complex_example() {
	//复数
	var c complex64 = 5 + 4i
	//output:(5+4i)
	fmt.Printf("Value is:%v\n", c)
}

var frenchHello string      // 声明变量为字符串的一般方法
var emptyString string = "" // 声明了一个字符串变量，初始化为空字符串

func string_example() {
	// 简短声明，同时声明多个变量
	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "Konichiwa"
	frenchHello = "Bonjour" //常规赋值
	fmt.Println(no, yes, maybe, japaneseHello)
	//修改字符串
	s := "hello"
	c := []byte(s) //将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s = string(c) //再转换回 string 类型
	fmt.Printf("%s\n", s)
	//修改字符串也可写为：
	ss := "hello"
	ss = "c" + ss[1:] //// 字符串虽不能更改，但可进行切片操作
	fmt.Printf("slice to  %s\n", ss)

	//Go中可以使用+操作符来连接两个字符串：
	s1 := "hello"
	m := " world"
	a := s1 + m
	fmt.Printf("%s\n", a)
	//如果要声明一个多行的字符串怎么办？可以通过`来声明：
	mulitString := `hello
				world,i am ` //按照这样的格式打印出来,包括换行和缩进
	//` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出
	fmt.Println(mulitString)
}

//错误类型
func error_example() {
	err := errors.New("emit macho dwarf:elf header corrupted.")
	if err != nil {
		fmt.Println(err)
	}
}

const (
	x = iota //x==0
	y = iota //y==1
	z = iota //z==2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)
const v = iota //每遇到一个const关键字，iota就会重置，此时v == 0

const (
	e, f, g = iota, iota, iota //e=0,f=0,g=0 iota在同一行值相同
)

//iota
func iota_example() {
	fmt.Printf("x=%d,y=%d,z=%d,w=%d,v=%d\n", x, y, z, w, v)
	fmt.Printf("e=%d,f=%d,g=%d\n", e, f, g)
}

//array
func array_example() {
	//由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度
	var arr [10]int // 声明了一个int类型的数组
	arr[0] = 42     // 数组下标是从0开始的
	arr[1] = 13     // 赋值操作
	// 获取数据，返回42
	fmt.Printf("The first element is %d\n", arr[0])
	//返回未赋值的最后一个元素，默认返回0

	fmt.Printf("The first element is %d\n", arr[9])

	var a [2]int
	var b [3]int
	//[2]int 与[3]int是不同的类型
	//fmt.Println(a == b)
	fmt.Println(a[0], b[0])

	//数组可以使用另一种:=来声明
	a1 := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b1 := [10]int{1, 2, 3}  // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c1 := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Println(":=", a1[0], b1[0], c1[0])

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Println(doubleArray[1][1])
	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(easyArray[1][3])
}
