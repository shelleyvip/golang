package main

import (
	"fmt"
	"math/rand"    //设置种子
	"time"
)

func Shelley( zhou *int){
	////设置种子，使用系统时间，格式为下列
	rand.Seed(time.Now().UnixNano()) ///系统时间
	xueli := rand.Intn(10000) //设置在10000以内的四位数随机数
	fmt.Printf("xueli=%d",xueli)
}

func main(){
	var zhou int
	Shelley(&zhou)

}