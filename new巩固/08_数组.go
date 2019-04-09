package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(len(a))
	for i,v := range a{
		fmt.Printf("%d %d\n",i,v) //第一个是下标 第二个是元素的值
	}
	for _,v := range a{
		fmt.Printf("%d",v)   //忽略下标 只打印值
	}

	q1 :=[...]int{1,2,3,4} //...不定长度
	fmt.Println(q1)

	q2 := [...]int{4:2,10:-1}
	fmt.Println(q2)
	fmt.Printf("%d",len(q2))

	//数组变量赋值与比较
	a1 :=[3]int{1,2,3}
	var a2 [3]int
	a2 = a1
	fmt.Println(&a1[0],&a2[0])

	var n1,n2 int
	n2=n1
	fmt.Println(&n1,&n2)

}
