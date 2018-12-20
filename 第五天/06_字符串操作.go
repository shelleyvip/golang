package main

import (
	"fmt"
	"strings"
)
//Contains:包含 控制
//joins:连接 组合
//Index 指数 索引
//Repeat 重复
//Split 拆分 分开
//Trim 整理 装饰 修整
func main(){
	///Contains的作用是测你这个"hellogo"中是否包含"hello"，包含返回true 真 不包含返回false 假
	//"hellogo" 中是否包含"hello"，包含返回true 不包含返回false
	fmt.Println(strings.Contains("hellogo","hello"))
	fmt.Println(strings.Contains("hellogo","abc"))//另外看你是否包含这个abc，结果是不包含返回的false

	//joins 组合 连接
	s := []string{"hello","mike","student","result"}
	buf := strings.Join(s,"@")//也就是可以利用个@字符来把这些字符串以@的方式串联起来
	fmt.Println("buf=",buf)

	//index 查找子串的位置
	fmt.Println(strings.Index("hellogoto","goto"))
	fmt.Println(strings.Index("hellegoto","abc"))//不包含子串 返回—1

	//Repeat 重复
	tmp := strings.Repeat("go",5)
	fmt.Println("tmp=",tmp)

	//Split 以指定的分隔符拆分
	zhou := "hello@mike@student@result"
	s2 := strings.Split(zhou,"@")
	fmt.Println("s2=",s2)

	//Trim 去掉两头的字符 //也就是整理
	s3 := strings.Trim("         are you ok?          "," ")
	//fmt.Println("s3=",s3)  两种打印方式都可以
	fmt.Printf("s3=#%s#\n",s3)

	//Fields 去掉空格把元素放到切片中
	s4 := strings.Fields("         are you ok?          ")
	//fmt.Println("s4=",s4) 还有一种比较清晰的 打印方式 就是通过迭代👇
	for i,data := range s4{
		fmt.Println(i,",",data)
	}

}