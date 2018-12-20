package main

import "fmt"
//无参 有返回值只有一个返回值
func Shelley ()(zhou int){         //go推荐写法：给返回值取一个名字（xxx int）
	zhou = 666
	return          //有返回值的函数 需要return中断函数通过return返回
}
func main(){                 ///无参有返回值的调用
	c := Shelley()   //自动推导 c:= 的方式打印出 调用Shelley()的值
	fmt.Println("c=",c)
}
