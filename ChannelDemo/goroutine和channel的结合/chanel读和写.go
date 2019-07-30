package main

import (
	"fmt"
	"time"
)
//问题
//如果注销掉	go readData(intChan,exitChan)程序会怎样
//答：	如果只是向管道里写入数据而没有读取，就会出现阻塞deadlock（死锁），原因是intChan容量是10，
//而代码writeData会写入50个数据，因此会阻塞在writeData的 intChan <- i

//注：
//如果 编译器（运行）发现一个管道只有写，没有读 则该管道会阻塞
//第二种就是 只要有读，写管道和读管道的频率不一致 无所谓 不会阻塞 但是一定要让它动起来

//创建writeData
func writeData(intChan chan int)  {
     for i:=1;i<=50;i++{
     	//往管道里放入数据
     	intChan <- i
     	fmt.Println("writeData=",i)
	 }
     //
     //写完之后关闭管道 但不影响正常读取
     close(intChan)
}

//第一个参数是 要接收需要读取的管道 还得有一个退出的管道
func readData(intChan chan int, exitChan chan  bool)  {
   //上面关闭的好处 是直接可以用for循环
   for{
   	v,ok := <- intChan
   	if !ok {
   		break
	}
   	time.Sleep(time.Second)
   	fmt.Printf("readData 读到数据=%v\n",v)
   }
   //readData 读完数据后，即任务完成
   exitChan <- true
   close(exitChan )
}

func main() {
	//创建两个管道
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)


	go writeData(intChan)
	go readData(intChan,exitChan)

	for{

		_,ok := <- exitChan
		if !ok{
			break
		}
		fmt.Println("读取完毕！")
	}


	}
