package main

import "fmt"

type student struct {
	id int
	name string
	age int
	addr string

}

func main() {
	//1:指针有合法指向后才能操作成员
	//所以先定义一个普通结构体变量 接一下上面的结构体
	var a student

	//在定义一个指针变量保存s的地址
	var p1 *student
	p1 = &a


///	通过指针操作成员 使用.操作
	p1.age = 16
	p1.addr = "bj"
	p1.id = 1
	p1.name = "sss"

	fmt.Println("p1=",p1)

	//通过new来申请一个结构体
	p2 := new(student)
	p2.name = "sdsds"
	p2.id = 3
	p2.addr = "bj"
	p2.age = 29
	fmt.Println("p2=",p2)

	s1 := student{6,"shelley",26,"bj"}
	s2 := student{6,"shelley",26,"bj"} //数组的赋值和比较、
	s3 := student{2,"shelley",26,"bj"}

    fmt.Println("s1==s1",s1==s2)  //数组的赋值和比较、
	fmt.Println("s2==s3",s2==s3)

	var tmp student
	tmp = s3
	fmt.Println("tmp=",tmp)  //同类的两个结构体变量可以互相赋值

}