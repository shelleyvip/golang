package main

import (
	"fmt"
	"regexp"
)

func main(){
	buf := "abc a2c auc a7c 888 tac"
	//解析规则，它会解析正则表达式，如果成功返回解析器

	//rge1 := regexp.MustCompile(`a.c`) //这个"."是只要a.c的都是在条件内的
	//rge2 := regexp.MustCompile(`a[0-9]c`)//第一种[0-9]是取：a和c之间0-9的数字
	rge2 := regexp.MustCompile(`a\dc`)//第二种 这个和上面0-9 是一样的效果
	if rge2 == nil{     //解析失败返回nil
		fmt.Println("regexp err")
		return
	}

	//根据规则提取关键信息
	result := rge2.FindAllStringSubmatch(buf,-1)//这里的-1 是匹配所有的
	fmt.Println("result=",result)



}