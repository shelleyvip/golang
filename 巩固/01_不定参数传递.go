package main

import "fmt"

func myfunc(tmp ...int)  {

	for _,data := range tmp{
		fmt.Println(data)
	}
}

func myfunc01(tmp ...int)  {
     for _,data := range tmp{
     	fmt.Println("data=",data)
	 }
	
}
func test(args ...int)  {
	//myfunc(args ...) //传递所有元素给myfunc
	myfunc01(args[3:]...) //只想吧后面的两个元素传递给myfunc01

}

//直接就一个rreturn返回一个返回值
func testa()int  {
	return 666

}
//只有一个返回值
func test02()(result int)  {
   result = 666
   return
}

func main() {
	test(1,2,3,4,5)

	a := test02()
	fmt.Println(a) //只有一个返回值

	b :=testa()
	fmt.Println(b) ////直接就是return一个返回值

}
