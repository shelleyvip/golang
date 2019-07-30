package main

import (
	"fmt"
	"time"
)

//向intChan里写入1-8000个数
func putNum(intChan chan int){
	for i :=2;i<=80;i++{
		intChan <- i
	}
	//关闭管道intChan
	close(intChan)
}

func PrimeNum(intChan chan int,primeChan chan int,exitChan chan bool)  {
	var flag bool //判断是否是素数
	for{
		time.Sleep(time.Millisecond)
		num,ok := <-intChan
		if !ok{  //如果!ok 说明intChan取不到东西了
			break    //就退出来
		}
		flag = true
			//判断一个数是不是素数
			for i :=2;i<num;i++{
				if num %i == 0{ //说明该num不是素数
                  flag = false  //如果不是素数 就为假
                  break  //不是素数就退出
				}
			}
			if flag{
				//将这个数放入到primeChan
				primeChan <- num
			}
		}
	fmt.Println("有一个PrimeNum 协程 因为取不到数据 退出")
	//这里我们还不能关闭PrimeChan
	//向exitChan写入true
	exitChan <- true
	}


func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //用于存放结果的
	//标识退出的管道
	exitChan := make(chan bool, 4) //4个

	start := time.Now().Unix()


	//开启一个协程，向intChan放入1-8000个数
	go putNum(intChan)

	//开启4个协程 从intChan取出数据，并判断是否为素数 如果是 就
	// 结果放入到primeChan
	for i := 0; i < 4; i++ {
		go PrimeNum(intChan, primeChan, exitChan)
	}

	//这里我们主线程进行处理
	//为了避免跟主线程发生关系 于是用一个匿名函数把她跑起来
	//也就是起了一个协程让他去做
	//让他不要阻在这个地方
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan //如果成功取到了4个数据 直接丢掉
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=",end-start)
		//那么就可以放心的关闭primeChan了
		close(primeChan)
	}()


	//最后把结果取出来  遍历我们的结果
	for {
		res,ok := <-primeChan
		if !ok {
			break

		}
		//将结果输出
		fmt.Printf("素数有=%d\n", res)
	}
	fmt.Println("main线程退出")

}