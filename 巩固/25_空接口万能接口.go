package main

import "fmt"

func xxx(arg ...interface{})  {
	
}

func main() {
	//申请一个接口 空接口万能类型，保存任意类型的值
	var i interface{}

	i = 666
	fmt.Println("i=",i)     //可以是不同类型的变量

	i = "shelley"
	fmt.Println("i=",i)

}
