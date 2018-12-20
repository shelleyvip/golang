package main

import "fmt"

func li(c int){
	fmt.Println("c=",c)///4:到li的时候 遇到 fmt于是 执行打印
}                        //// 结果是：c =1

func xue(b int){
	li(b - 1)           //// 3：到xue的时候 它的内容是开始调用li 并执行(b -1)的任务继续往下跳转执行li
	fmt.Println("b=",b)////结果是：b = 2

}

func zhou (a int){
	xue(a - 1)           ////2：到zhou的时候 它的内容是开始调用xue并执行(a -1)的任务继续往下执行xue
	fmt.Println("a=",a)//// 结果是：a = 3
}

func main(){
	zhou(3)              //1：先从main函数开始执行 调用zhou 给它赋个值（3）开始执行zhou
	fmt.Println("main")
}
         ////函数调用流程：先调用后返回，先进后出
         ////函数递归，函数调用本身，利用此特点
