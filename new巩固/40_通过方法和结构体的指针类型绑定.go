package main

import "fmt"

type Circle1 struct {
	radius float64
}

//指针方法 高效
//c 是指针类型
func (c *Circle1)area2()float64  {
	fmt.Printf("c 指向的地址是=%p\n",c) //因为c本身就是指针所以不用&符号
  c.radius = 10
  return 3.14 *  c.radius * c.radius
}

func main()  {
	var c Circle1
	fmt.Printf("main:的地址是%p\n",&c)
	//编译器底层做了优化自动加上了&c
	c.radius = 7.0
	resuilt := c.area2()
	fmt.Println("面积是:",resuilt)

}