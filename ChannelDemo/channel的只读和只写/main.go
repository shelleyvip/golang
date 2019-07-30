package main

import "fmt"

func main() {
	//管道可以声明只读或者只写
	//1，默认情况下 管道是双向的
	// var chan1 chan int //这种 的是可以读可以写的

	//2声明只写
	var chan2 chan<- int
	chan2 = make(chan int,3)
	chan2 <- 20
	fmt.Println("chan2=",chan2)
	//num := <-chan2  //error 声明的是只写 所以在这里读取的话会报错


	//3：声明只读
	var chan3 <-chan int
	num2 := <-chan3
	fmt.Println("num2=",num2)
	//chan3 <- //error

}
