package main

import (
	"fmt"
	"math/rand"
	"time"
)
func CreatNum(p *int)  {
	rand.Seed(time.Now().UnixNano())//系统当前时间作为种子

	var num int
	for{
		num = rand.Intn(10000)
		if num >= 1000{
			break
		}
	}
	*p = num
}

func OnGame(randSilic []int)  {
	var num int
	for{
		fmt.Printf("请输入一个4位的数字:")
		fmt.Scanf(&num)
		if num > 999 && num < 10000{
			break
		}
		fmt.Println("num",num)
	}
	
}
//n1 := 1234/1000  //取商
//n2 := 1234%1000/100 //(1234%1000)取余数位234  234/100取商得到2
func Getnum(s []int,num int)  {
	s[0] = num /1000///取千位
	s[1] = num %1000/100//取百位
	s[2] = num %100/10//取百位     （去除随机四位的个数）
	s[3] = num %10//取十位

}


func main() {
	var randnum int
	CreatNum(&randnum)
	fmt.Println("randnum:",randnum)

	randSilic := make([]int,4)

	Getnum(randSilic,randnum)
	fmt.Println("randSilic=",randSilic)

	OnGame(randSilic)
}
