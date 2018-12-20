package main

import "fmt"

func main(){
    //defer是延时调用的作用，是在main函数结束前调用
	defer fmt.Println("aaaaaaaaaa")
	fmt.Println("cccccccccc")
}