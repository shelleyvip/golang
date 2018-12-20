package main

import "fmt"

type shelley struct{
	name string
	age int
	sex string

}
type zhou struct{
	shelley
	id int
	addr string
}

func main(){//这里就是把zhou里面的成员都赋值给了s1（包括shelley的成员）
	s1 := zhou{shelley{"周",22,"女"},128,"北京"}
	//普通俄机构提成员的操作赋值
	//点操作.是s1.也就是s1里面指定的某某XX
	s1.name = "你好"
	s1.age = 25
	s1.sex = "男"
	s1.id = 346
	s1.addr = "深圳"
	fmt.Println(s1.name,s1.age,s1.sex,s1.id,s1.addr)


	//这里还可以操作一个成员当作整体赋值
	s1.shelley = shelley{"周晓莉",29,"nv"}
	//打印结果就是 改变过的//周晓莉 29 nv 346 深圳
	fmt.Println(s1.name,s1.age,s1.sex,s1.id,s1.addr)


}
