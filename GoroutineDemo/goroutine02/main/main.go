package main

import (
	"fmt"
	"sync"
	"time"
)

//需求：现在需要计算1-200各个数的阶乘，并且把各个数的阶乘放入到map中。
//最后显示出来。要求时使用goroutine完成

//思路
// 1：编写一个函数来计算各个数的阶层，并放入到map中
// 2：我们启动多个协程，统一将结果放入map中
// 3：所以map应该是全局的

var(
	mymap = make(map[int]int,10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁
	//sync 是包：全称：synchronized 同步的意思
	//mutex ;是互斥
	lock sync.Mutex
)

//test函数就是计算n 让将这个结果放入到mymap
func test(n int)  {
	res := 1
	for i:=1;i<n;i++{
		res *= i
	}
	//将这个结果放入到mymap中
	//加锁
	lock.Lock()
	mymap[n] = res //concurrent map writes?
	//解锁
	lock.Unlock()
}

func main()  {

	//这里开启多个协程完成任务
	for i :=1;i<200;i++{
		go test(i)
	}
	time.Sleep(time.Second*10)

	//这里我们输出结果
	lock.Lock()
	for i,v := range mymap{
		fmt.Printf("map=[%d]%d\n",i,v)
	}
	lock.Unlock()
}

//多资源竞争问题 //也可以 go build -race也可以看到他的多资源竞争的问题
