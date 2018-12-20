package main

import "fmt"

func test(){
//	a := 10
{  // {}  声明的变量只作用与{}内部
	a := 9  //定义好的变量可以任意赋值
	b := a
	c := b
	d := c


	fmt.Println("d=",d)


}
}

func main(){
     test()
	a := 12
	fmt.Println("a=",a)
	{
		 a := 21
		fmt.Println("a=",a)
	}

}