package main

import "fmt"

type student struct {
	id int
	name string
	age int
	addr string
}

func test(s student)  {
	s.age = 666
	fmt.Println("s=",s)
}

func test01(p *student)  {
	p.age = 666//这个是地址传递 也就是引用传递 因为传的是地址 所以形参可以改变实参

}

func main() {
	s :=student{1,"ss",22,"sssss"}

	test(s) //值传递 形参无法改变实参

	test01(&s)//这个是地址传递 也就是引用传递 因为传的是地址 所以形参可以改变实参)
	fmt.Println("s=",s)
}

