package main

import "fmt"

//这个可以接收 任意类型，任意个数的数值
func xxx(args ...interface{}){

}

func main(){
	//在这里空接口是万能类型，它可以保存任意类型的值
	 var i interface{}= 1
	 fmt.Println("i=",i)

	 i = "abc"
	 fmt.Println("i=",i)
}
