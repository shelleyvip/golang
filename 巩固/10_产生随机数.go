package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Creatnum(p *int)  {
	rand.Seed(time.Now().UnixNano())//当前系统时间
	var num int
	for{
		num = rand.Intn(10000)
		if num >= 1000{ //四位数以内
			break
		}
	}
	*p = num  //指针指向num的随机数内寸

}
func main() {
	var randNum int

	Creatnum(&randNum)

	fmt.Println("rendnum :",randNum)
}