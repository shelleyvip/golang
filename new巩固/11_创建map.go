package main

import "fmt"

func main() {
	ages :=make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2
	fmt.Println(ages["a"])
	ages["a"] = ages["b"]+2

	//
	////第二种法式
	//ages1 := map[string]int{
	//	"a"  : 1,
	//	"b" : 2,
	//}
	//fmt.Println(ages1)

	c,ok :=ages["c"]  //查看这个c是否存在
	if ok{            //如果ok 打印出来
		fmt.Println(c)
	}else {        //否则就打印出下面一段话
		fmt.Println("not found")
	}
}
