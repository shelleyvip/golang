package main

import "fmt"

func main(){
	 var d int       //变量var声明 a 的类型 int
	fmt.Println("a=",d)

	s := 10       ///给变量赋值10 := 是自动推导类型
	fmt.Printf("a type is %T\n",s) ////fmt.Printf 是格式化输出打印
	////////////////type is 是类型 %T 是自动推倒类型 以上（"a type is %T",a）（自动推导类型并格式化输出打印）

	var i int = 10 //初始化变量期间同时赋值
	fmt.Println("c=",i)

	//声明多个变量
	var a,b,c int
	a = 10
	b = 20
	c = 30
	fmt.Println("a=",a ,"b=",b,"c=",c )

}     /////注意：同一个任务栏不可以有两个相同的变量名
