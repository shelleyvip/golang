package main

import "fmt"

//开启一个协程向intChan 放入1-8000个数
func PutNum (intChan chan int)  {
		for i :=1;i<=1000;i++{
			intChan <- i
		}
		close(intChan)
}
//开启4个协程，从intChan取出数据，并判断是否为素数，如果是就放入到PrimeChan
func PrimeNum(intChan chan int,PrimeChan chan int,exitChan chan bool)  {
	//使用for循环
	var flag bool
	for{
		num,ok := <- intChan
		if !ok{
			break
		}
		flag = true //假设是素数
		//判断是不是素数
		for i := 2;i<num;i++{
			if num %i ==0 {
				flag = false
				break
			}
		}
		if flag{
			PrimeChan <- num
		}
	}
	fmt.Println("有一个PrimeNum 协程因为娶不到数据 退出")
	exitChan <- true
}
func main() {
	intChan := make(chan int,1000)
	PrimeChan := make(chan int,2000)//放入结果
	//标识退出管道
	exitChan := make(chan bool,4)

	//开启一个协程向intChan 放入1-8000个数
	go PutNum(intChan)


	//开启4个协程，从intChan取出数据，并判断是否为素数，如果是就放入到PrimeChan

	go func() {
		for i:=0;i<4;i++ {
			go PrimeNum(intChan,PrimeChan,exitChan)
		}

		for i :=0;i<4;i++{
			<- exitChan
		}
		//当我们从exitChan里买那u 出4个结果就可以放心的关闭 PrimeChan
		close(PrimeChan)
	}()
	//遍历我们的PrimeChan,把结果取出来
	for{
		result,ok := <- PrimeChan
		if !ok{
			break
		}
		fmt.Printf("素数=%d",result)
	}
	fmt.Println("main主线程退出")




}