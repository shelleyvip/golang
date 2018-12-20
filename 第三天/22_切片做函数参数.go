package main

import (
	"fmt"
	"math/rand"
	"time"
)

func data(s []int)  {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	for i :=0;i < len(s); i++{
		s[i] = rand.Intn(31)// 获取100以内的随机数
	}
}

func main()  {
	//a := 7
	s := make([]int,7,9) //创建一个切片
	data(s)
fmt.Printf("排序前s=%d\n",s)
}
