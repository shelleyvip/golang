package main

import "fmt"

func testa(){
	fmt.Println("aaaaa")
}

func testb(x int){
	var a [10]int
	a[x] = 111//当x为20 的时候，导致数组越界，产生一个panic,导致程序崩溃

}
func testc(){
	fmt.Println("ccccc")
}

func main(){
	testa()
	testb(20)
	testc()
}
