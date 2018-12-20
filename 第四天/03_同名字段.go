package main

import "fmt"

type shelley struct{
	name string
	sex string
	age int
}
type zhou struct{
	//只有类型没有名字，继承了shelly的成员
	shelley //也就是说把shelley放在这里 把它的成员也一并带了过来
	addr string
	id int
	name string
}

func main(){
	//声明（定以）一个变量
	var s1 zhou //定以一个变量 把 zhou 赋值给s1
	//在这里操作的是zhou的name //还是shelley的name？//答案是zhou的name!
	s1.name = "楠楠"//因为默认规则是（就近原则）如果在本作用域能找到此成员就操作此成员
	s1.addr = "深圳"//如果找不到，就找//继承的字段//成员
	s1.id = 126 //
	s1.sex = "女"
	s1.age = 25

	//如果想调用上面👆Shelley的name 那就//显示调用 s1里面 带上指定字段 那么就可以
	s1.shelley.name = "周雪莉"

	fmt.Printf("s1=%+v\n",s1)

}
