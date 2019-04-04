package main

import (
	"fmt"
	"golang/工程管理不同目录调用/calc"
)

func init()  {
	fmt.Println("this is main init")

}

func main()  {
   a := calc.Add(10,20)
   fmt.Println("a=",a)
}
