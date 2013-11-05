package main
import "fmt"
func main() {
	fmt.Printf("hello,world!你好,世界!\n");
	var i int;
	i=10;
	i=i+i;
	var a int;
	var b int32;
	a=15;
	a=a;
	//b=a+a;
	b=b+5;//+a;
	a=a+9.0;
	/*
	第一个 iota 表示为 0，因此 a 等于 0，当 iota 再次在新的一行使用时，它的值增加
	了 1，因此 b 的值是 1
	*/
	const(
		c=iota
		d=iota
	)
	//如果需要，可以明确指定常量的类型：
	const(
		e=0
		f string ="this is f"
		)

	//字符串的修改
	s:="hello"
	cc:=[]rune(s)//转换 s 为 rune 数组
	cc[0]='c'
	s2:=string(cc)
	fmt.Printf("修改后的:%s\n",s2)
	s3:="Starting Part"+		//将Starting Part和Ending Part串联
		"Ending Part"
	s4:=`Starting part
			Ending part`		//按照这个格式输出,类似于html里的<pre>标签
	fmt.Printf("%s,\n%s",s3,s4);

	//复数
	var complex complex64 = 5+6i;
	fmt.Printf("\nValue is : %v",complex)

	ifelse:=3;
	if ifelse>2{	/*{} 是强制的  { 必须同 if 在同一行*/
		ifelse=4;
	}else{
		ifelse=0;
	}
	fmt.Printf("\n%v",ifelse);
	 myfunc();
/*
	if err:=Chmod(0664);err!=nil{// nil 与 C 的 NULL 类似
		fmt.Printf(err);	//err 的作用域被限定在 if 内
	}
*/
/*	//平行赋值 i,j:= 0,len(a)-1
	//reverse a
	for i,j:= 0,len(a)-1; i < j; i,j=i+1,j-1 {
		a[i],a[j]=a[j],a[i];
	}
*/

	breakByTag();
}
func myfunc() {
	i:=0
Here:
	println(i)
	i++
	if i<10{//判断条件
		goto Here
	}
}
func breakByTag() {
	J:for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			fmt.Printf("j=%v,i=%v",j,i);
			if i>5{
				break J;//现在终止的是 j 循环，而不是 i 的那个
			}
		}
	}
	
}
func breakN() {
	for i := 0; i < 10; i++ {
		if i>6{
			break  //终止这个循环,只打印 0 到6
		}
	}
}
func ContinueN() {
	for i := 0; i < 10; i++ {
		if i>5{
			continue;//让循环进入下一个迭代，而略过剩下的所有代码
		}
	}
	
}
func RangeN() {
	list:=[]string{"a","b","c","d","e","f"};
	for k,v:=range list {
		fmt.Printf("%s=%s",k,v);//range 是个迭代器
	}
	
}

func switchN() {
	switch i{
	case 0://空的case体
	case 1: //当i==0时,f不会被调用
		f()
	}

	//==============
	switch i{
	case 0:fallthrough  //空的case体
	case 1: //当i==0时,f会被调用
		f()
	}

		//==============
	switch i{
	case 0:
	case 1: 
		f()
	default:
		g()//当 i 不等于 0 或 1 时调用
	}
	
}
//分支可以使用逗号分隔的列表。
func shouldEscape(c by te) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+': // as ”or” 或
	return true
	}
	return  false
}
func f() {
	
}
func g() {
	
}