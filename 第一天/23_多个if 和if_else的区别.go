package main

import "fmt"

func main() {
	a := 11
	if a == 10 { ///左大括号要和 if 同一行
		fmt.Println("a==10")
	} else if a < 10 { ////else 后面 继续跟大括号
		fmt.Println("a < 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	}
}