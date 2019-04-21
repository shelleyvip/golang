package main

import "fmt"

type Student struct {
	Id int
	Name string
}
type Stu Student

func main()  {
    var shu1 Student
    var shu2 Stu
    //shu2 = shu1 强转是不可以的 可以修改以下方式
    shu2 = Stu(shu1)
    fmt.Println(shu1,shu2)
}
