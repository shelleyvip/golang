package main

import (
	"fmt" //fmt提供打印输出
	"os" //os包提供一些函数和变量
)

func main()  {

zhouxueli := os.Args   //接受字符串的方法

	// 变量os.Args是一个字符串slice。可以理解它是一个动态容量的顺序数组s，可以通过s[i]来访问单个元素，
	//通过s[m:n]来访问一段连续子区间，数组长度用len(s)表示。
	//在Go中，所有的索引使用半开区间，即包含第一个索引，不包含最后一个索引。os.Args的第一个元素是os.Args[0],它是命令本身的名字；
	//另外的元素是程序开始执行时的参数。表达式s[m:n]表示一个从第m个到第n-1个元素的slice

n :=  len(zhouxueli)  //len()  内建函数主要获取函数的长度
f := cap(zhouxueli)    //cap() 内建函数主要获取函数的容量
fmt.Println("f=",f)
fmt.Println("n=",n)

for  i,v := range  zhouxueli {       //for i ,v := range name { }  遍历语句
	fmt.Println("zhouxueli[%d] %s",i,v)
}

}
