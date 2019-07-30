package channeldemo01

import "fmt"

func main() {
	//演示一下管道的使用
	//1：创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int ,3)

	//2:看看intChan是什么
	fmt.Printf("intChan的值=%v\n intChan本身的地址=%v\n",intChan,&intChan)
	 //结果：intChan的值=0xc000078000是个地址
	 // intChan本身的地址=0xc00007e018

	 //3:向管道写入数据
	 intChan <- 10
	 num := 211
	 intChan <- num

	 intChan <- 20

	 //intChan <- 98   //报错 原因fatal error: all goroutines are asleep - deadlock!
	 //当我们给管道写入数据时，不能超出其容量

	 //4:看看管道的长度和容量
	 fmt.Printf("intChan长度是%v 容量是%v\n",len(intChan),cap(intChan)) //长度3 //容量3


	 //5:从管道里去数据
	 var num2 int
	 num2 = <-intChan
	 fmt.Println("num2=",num2)
	fmt.Printf("intChan长度是%v 容量是%v\n",len(intChan),cap(intChan)) //长度2 容量3

	//6：在没有使用协程的情况下，如果我们管道的数据已经全部取出 在取就会报错deadlock
	num3 := <-intChan
	num4 := <-intChan
	//已经取完 在取就会deadlock
	//num5 := <-intChan

	fmt.Println("num3=",num3,"num4=",num4)//,"num5=",num5)
	//fatal error: all goroutines are asleep - deadlock!




}
