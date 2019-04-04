package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int

}
type Student struct {
	Person
	id int
	addr string

}
func main()  {
	//顺序初始化
	s1 := Student{Person{"shelley",'m',26},2,"bj"}
	fmt.Println("s1=",s1)

	s2 := Student{id:1}
	fmt.Printf("s2=%+v\n",s2)

	s3 := Student{Person:Person{name:"zhou"},id:2}
	fmt.Printf("s3=%+v\n",s3)


}
