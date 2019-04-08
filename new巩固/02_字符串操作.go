package main

import "fmt"

func main() {
	s1 := "hello"+"world"   //相加
	s2 := "helloworld"

	if s1 == s2{
		fmt.Println("equal")
	}
	var c1 byte
	c1 =s1[0]     //取字符
	fmt.Println(s1,s2,c1)

	s3 := s1[3:4]  //切片  左臂右开 右边的是不包括在里面的 从0开始计算
	fmt.Println(s3)

	//字符串本身不能修改，通过跟【】byte相互转换来修改
	array := []byte(s1)  //把字符串s1强制转换 成字符  字符就是
	fmt.Println(array)
	array[0] = 'H'    //指定某一个字符修改
	s6 := string(array) //在把它转回来
	fmt.Println(s6)

	var b bool
	b = true
	b = false
	b = ("hello" == "world")
	if b{

	}


}