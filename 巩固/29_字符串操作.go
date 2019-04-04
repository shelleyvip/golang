package main

import (
	"fmt"
	"strings"
)

func main() {
	//Contains 包含
	//看这个字符串helloeorld 是否包含"hello"  如果包含返回true 不包含返回false
	fmt.Println(strings.Contains("helloeorld","hello"))
	fmt.Println(strings.Contains("helloeorld","nihao"))

	//Joins链接
	s := []string{"hello","go","world"}
	buf := strings.Join(s,"-")//第二个参数是中间以什么来空隔 根据需要
	fmt.Println("buf=",buf)

	//Index索引 查找某个字符所在的位置
	fmt.Println(strings.Index("hellogo","go"))
	fmt.Println(strings.Index("helloworld","go"))//不包含子串返回-1

	//Repeat 重复
	//指定重复多少次
	buf = strings.Repeat("go",10) //第一个是元素 第二个是重复几次
	fmt.Println("buf=",buf) //gogogogogogogogogogo

	//Split 拆分
	//以指定的分隔符拆分
	buf = "hello&world&renzha"
	s2 := strings.Split(buf,"&")
	fmt.Println(s2)


	//Trim 去掉两头的字符
	buf = strings.Trim("    are you ok?   "," ") //去掉两头空隔
	fmt.Println(buf)

	//去掉两头空格把元素放入切片中
	s3 := strings.Fields("hello world go")
	for i,data := range s3{
		fmt.Println(i,"",data)
	}






}
