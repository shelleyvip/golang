package main

import "fmt"

func zhou (x int){
	shelley := 100 / x
	fmt.Println("shelley=",shelley)
}
func main(){
	//如果函数调用没加defer，那么defer是最后执行的，如果出现多个defer执行是自下往上执行。
 defer fmt.Println("aaaaaaaa")
	defer fmt.Println("bbbbbbbb")
	 zhou(0)
	defer fmt.Println("cccccccc")
	 ////// defer执行顺序是 自下往上 执行
	 //哪怕函数某个环节 延时调用出错 这些调用依旧会被执行
}