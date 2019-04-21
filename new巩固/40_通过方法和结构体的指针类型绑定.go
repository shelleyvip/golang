package main

import "fmt"

type Circle1 struct {
	radius float64
}

func (c *Circle1)area2()float64  {
  c.radius = 10
  return 3.14 *  c.radius * c.radius
}

func main()  {
	var c Circle1
	c.radius = 7.0
	resuilt := c.area2()
	fmt.Println("面积是:",resuilt)
}