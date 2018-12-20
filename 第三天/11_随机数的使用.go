package main

import (
	"fmt"
	"math/rand" //种子函数
	"time"
)

func main()  {
 rand.Seed(time.Now().UnixNano()) //调用当前系统时间
 for i:=0; i < 5; i++ {           //（5）行
//产生随机数的，限制在10以内方法如下//               //最大值用Int
 	fmt.Println("rand =",rand.Intn(10))// 或用Intn是限制在100内的数字
 }

}
