package main

import (
	"fmt"
	"regexp"  //
)

func main(){
	rege01 := "3.14 2.5 dfg 1. 23.6 3. wrr"
  //解释正则表达式 "+"匹配 前一个字符 的一次或多次
	su1 := regexp.MustCompile(`\d+\.\d+`)
	if su1 == nil{
		fmt.Println("MustCompile err")
		return
	}
    //提取关键信息
	result := su1.FindAllStringSubmatch(rege01,-1)
	fmt.Println("result=",result)
}