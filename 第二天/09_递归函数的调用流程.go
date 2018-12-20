package main

import "fmt"

func zhou (a int) {
	if a == 1 {                   //2：if(如果) a 不等于 1 继续往下走 zhou（a-1）
		fmt.Println("a=",a)   //3:回来 a还是不满足条件（等于a ==1 的条件）继续调用一下zhou(a-1)
		                       ///4:if a == 1 满足条件开始打印：fmt("a=",a) 结果是：a=1
		return                 /// 5：中止函数的调用并返回 下列该执行的任务
		 }
	zhou(a-1)                 ///——2:zhou(a-1) a = 2
	fmt.Println("a=",a)    ///——3:zhou(a-1)  a = 1

}
func main(){
	zhou(3)        ////1：先从main函数开始执行 调用zhou 给它赋个值（3）开始执行zhou
	fmt.Println("main")
}
          ////函数调用流程：先调用后返回，先进后出
          ////函数递归，函数调用本身，利用此特点
