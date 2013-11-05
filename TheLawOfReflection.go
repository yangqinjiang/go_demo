package main

import (
	"fmt"
	"reflect"
)

type MyInt int

func main() {
	var i int = 10
	var j MyInt = 10
	//invalid operation: i == j (mismatched types int and MyInt)
	//类型不匹配
	//fmt.Println(i == j)
	it := reflect.TypeOf(i) //int
	jt := reflect.TypeOf(j) //MyInt

	fmt.Println(it.Name(), "__", jt.Name())

	var x float64 = 3.14
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	v := reflect.ValueOf(x)
	//返回value的类型
	fmt.Println("type:", v.Type())
	//kind 返回值的类型
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//打印其值
	fmt.Println("value:", v.Float())
	//v.SetFloat(0.668)//Error: will panic.
	//fmt.Println("modify value:", v.Float())

	y := v.Interface().(float64) // y will have type float64.
	fmt.Println("y:", y)
	//The CanSet method of Value reports the settability of a Value; in our case,
	var z float64 = 3.4
	zv := reflect.ValueOf(z)
	fmt.Println("settability of v:", zv.CanSet())
	//zv.SetFloat(7.1)//Error: will panic.
	ModityValue()
	ModifyStructs()
	getAnonymousFiled()
	u := User{1, "hi", 20}
	Set(&u)
	fmt.Println(u)
	fmt.Println("Call Method...")
	callMethod()
}

func ModityValue() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) //Note:take the address of x
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet()) // false
	//获取p变量的指向,我们调用 p变量的Elem方法
	//间接指向x变量,并且保存结果在反射变量v
	v := p.Elem()
	fmt.Println("settability of p:", v.CanSet()) // true
	v.SetFloat(7.1)
	fmt.Println("通过反射来修改x的值 v= ", v.Interface())
	fmt.Println("修改后的x = ", x)

}

type T struct {
	A int
	B string
}

func (t T) Hello() {
	fmt.Println("call Hello func in T")
}
func ModifyStructs() {
	t := T{23, "skidoo"}
	st := reflect.TypeOf(t)
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	//打印字段值
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d:%s %s=%v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	fmt.Println("Method:")
	//st  =reflect.TypeOf(t),
	//不能这样s := reflect.ValueOf(&t).Elem()
	for i := 0; i < st.NumMethod(); i++ {
		m := st.Method(i)
		fmt.Println(m)
		//%6s 缩进
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
	//修改值
	//0 1 已经知道的值
	s.Field(0).SetInt(77)
	s.Field(1).SetString("sunset strip")
	fmt.Println("t is now: ", t)
}

type User struct {
	Id   int
	Name string
	Age  int
}

//不能这样:有指针*  func (u *User) SayHello(name string)
func (u User) SayHello(name string) {
	fmt.Println(name, "say :Hello,go-world")
}

type Manager struct {
	User
	title string
}

//获取struct里的匿名字段是否是 Anonymous
func getAnonymousFiled() {
	fmt.Println("getAnonymousFiled")
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	//Field 1  对应于 m User
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Println("Is Anonymous?", t.Field(0).Anonymous)

	fmt.Printf("%#v\n", t.Field(1))
	fmt.Println("Is Anonymous?", t.Field(1).Anonymous)

	//User Id
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}

//通过反射来设置新的值
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	fmt.Println("reflect.Ptr:", reflect.Ptr)
	//获取对象
	if v.Kind() == reflect.Ptr &&
		!v.Elem().CanSet() {
		fmt.Println("XXX")
	} else {
		v = v.Elem()
		fmt.Println("YYY")
	}

	//设置新的值
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("BYE")
	}
}
//通过反射可以"动态"调用方法
func callMethod() {
	u := User{1, "callMethod", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("SayHello")
	//准备参数
	args := []reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)
}
