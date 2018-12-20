package main

import "fmt"


func test99(){
	fmt.Println("uuuuuuu")
}

func test88(){
	fmt.Println("777777")
	//第一种：☝️显示调用panic 函数 导致程序中断
	panic("this is panic test")
}

func test77(w int){
	//第二种:✌️数组越界导致panic
	var a[10] int
	a[w] = 111//当w为20 的时候，导致数组越界，产生一个panic,导致程序崩溃
}

func test66(){
	////第三种：☝️✌️//设置recover
	defer func() {   //设置recover
		if err := recover();err !=nil{
			fmt.Println("err=",err)
		}
	}()   //别忘了这后面的（）是调用匿名函数

}
func main() {
	test99()
	test88()
	test77(20)
	test66()

}