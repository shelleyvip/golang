package main

import "fmt"

//定义参数时在函数名后面的（）里面定义的参数叫形参
//有参无返回值的调用 参数是int整型，所以就要给它传 个整型
func shelley (a int){
	fmt.Println("a=",a)
}
func main(){
	//有参数无返回值的调用：函数名shelley（加入所需参数）
	shelley(888)
}

