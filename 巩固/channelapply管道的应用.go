package main

import "fmt"

//Write Date
func writeData(intChan chan  int)  {
	for i := 0;i<=50;i++{
		intChan <- i //放入数据
		fmt.Println("i=",i)
	}
	close(intChan)
}
//Read Data
func readData(intChan chan  int,exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到的数据=%v\n", v)
	}
	//readdata读取后，即完成任务
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)
	for{
		_,ok :=  <-exitChan
		if !ok{
			break
		}
	}
}
