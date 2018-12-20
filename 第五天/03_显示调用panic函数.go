package main

import "fmt"

func testa(){
	fmt.Println("aaaaaaaaaaaaa")
}
func testb(){
	fmt.Println("bbbbbbbbbbbbb")
	//显示调用panic 函数 导致程序中断
	panic("this is a panic test")
}
func testc(){
	fmt.Println("cccccccccccccc")
	//panic("this is a panic test")

}
func main(){
	testa()
	testb()
	testc()
}