package main

import "fmt"

func func03(c int)  {
	fmt.Println("c=",c)

}

func func02(b int)  {
	func03(b -1)    //调用func03 条件是(b-1)
	fmt.Println("b=",b)

}

func func01(a int)  {
	func02(a-1)      //调用func02 条件是(a-1)
	fmt.Println("a=",a) //第一步给a传参为3
}

func main() {
	func01(3)
	fmt.Println("main")
}
